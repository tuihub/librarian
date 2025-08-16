package bizsupervisor

import (
	"context"
	"math"
	"slices"
	"sync"
	"time"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

const (
	// Default downgrade timeout.
	defaultContextDowngradeTimeout = 3 * time.Minute
	// Default re-enable intervals.
	defaultContextEnableInitialInterval = 5 * time.Second
	defaultContextEnableMaxInterval     = 10 * time.Minute
	defaultContextEnableBackoffFactor   = 2
)

type PorterContextController struct {
	modelsupervisor.PorterContext

	s                    *Supervisor
	heartbeatCount       int       // Positive for successful heartbeats, negative for failed heartbeats
	lastHeartbeat        time.Time // Last heartbeat time
	lastSuccessHeartbeat time.Time // Last successful heartbeat time
	handlerAddress       string
	mu                   sync.Mutex
}

func (s *Supervisor) NewPorterContextController(
	c *modelsupervisor.PorterContext,
) *PorterContextController {
	return &PorterContextController{
		PorterContext:        *c,
		s:                    s,
		heartbeatCount:       0,
		lastHeartbeat:        time.Time{},
		lastSuccessHeartbeat: time.Time{},
		handlerAddress:       "",
		mu:                   sync.Mutex{},
	}
}

func (c *PorterContextController) GetName() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.PorterContext.Name
}

func (c *PorterContextController) GetGlobalName() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.PorterContext.GlobalName
}

func (c *PorterContextController) GetRegion() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.PorterContext.Region
}

func (c *PorterContextController) UpdateStatus( //nolint:gocognit // TODO
	ctx context.Context,
	availableAddresses []string,
) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Check heartbeat interval
	if c.heartbeatCount < 0 {
		interval := math.Min(
			float64(
				defaultContextEnableInitialInterval,
			)*math.Pow(
				defaultContextEnableBackoffFactor,
				float64(-c.heartbeatCount-1),
			),
			float64(defaultContextEnableMaxInterval),
		)
		if time.Since(c.lastHeartbeat) < time.Duration(interval) {
			return c.handlerAddress
		}
	}
	defer func() {
		c.lastHeartbeat = time.Now()
		if c.heartbeatCount > 0 {
			c.lastSuccessHeartbeat = c.lastHeartbeat
		}
		newContext, err := c.s.repo.UpdatePorterContext(ctx, &c.PorterContext)
		if err != nil {
			logger.Errorf("failed to update porter context: %v", err)
		} else {
			c.PorterContext = *newContext
		}
	}()
	if c.Status == modelsupervisor.PorterContextStatusDisabled {
		err := c.disable(ctx)
		if err != nil {
			return ""
		}
	}
	switch c.HandleStatus {
	case modelsupervisor.PorterContextHandleStatusUnspecified:
		// Initial state, set to queueing
		c.HandleStatus = modelsupervisor.PorterContextHandleStatusQueueing
	case modelsupervisor.PorterContextHandleStatusActive:
		if slices.Contains(availableAddresses, c.handlerAddress) {
			// Handler address still available, no need to change
			c.heartbeatCount = max(1, c.heartbeatCount+1)
		} else {
			// Handler address not available, downgrade
			c.HandleStatus = modelsupervisor.PorterContextHandleStatusDowngraded
			c.HandleStatusMessage = "handler address not available, downgraded"
		}
	case modelsupervisor.PorterContextHandleStatusDowngraded:
		if slices.Contains(availableAddresses, c.handlerAddress) {
			// Handler address available again, re-activate
			c.HandleStatus = modelsupervisor.PorterContextHandleStatusActive
			c.HandleStatusMessage = "handler address available again, re-activated"
			c.heartbeatCount = max(1, c.heartbeatCount+1)
		} else if time.Since(c.lastSuccessHeartbeat) >= defaultContextDowngradeTimeout {
			// Downgrade timeout reached, re-queue
			c.HandleStatus = modelsupervisor.PorterContextHandleStatusQueueing
			c.HandleStatusMessage = "previous handler address not available, re-queueing"
			c.handlerAddress = ""
			c.heartbeatCount = 0
		}
		// Still in downgrade period, keep the same handler address
	case modelsupervisor.PorterContextHandleStatusQueueing:
		if len(availableAddresses) == 0 {
			// No available addresses, keep queueing
			c.HandleStatusMessage = "no available instance"
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
		} else {
			// Set to blocked status to start enabling context
			c.HandleStatus = modelsupervisor.PorterContextHandleStatusBlocked
		}
	case modelsupervisor.PorterContextHandleStatusBlocked:
		// Try to enable context
		for _, address := range availableAddresses {
			_, err := c.s.porter.EnableContext(
				client.WithPorterAddress(ctx, []string{address}),
				&porter.EnableContextRequest{
					ContextId:   converter.ToPBInternalID(c.ID),
					ContextJson: c.ContextJSON,
				},
			)
			if err == nil {
				c.HandleStatus = modelsupervisor.PorterContextHandleStatusActive
				c.HandleStatusMessage = ""
				c.handlerAddress = address
				c.heartbeatCount = 1
				break
			} else {
				c.HandleStatusMessage = err.Error()
				c.heartbeatCount = min(-1, c.heartbeatCount-1)
			}
		}
	}
	return c.handlerAddress
}

func (c *PorterContextController) disable(
	ctx context.Context,
) error {
	if c.HandleStatus == modelsupervisor.PorterContextHandleStatusActive {
		_, err := c.s.porter.DisableContext(
			client.WithPorterAddress(ctx, []string{c.handlerAddress}),
			&porter.DisableContextRequest{
				ContextId: converter.ToPBInternalID(c.ID),
			},
		)
		c.HandleStatus = modelsupervisor.PorterContextHandleStatusQueueing
		c.HandleStatusMessage = ""
		c.handlerAddress = ""
		c.heartbeatCount = 0
		if err != nil {
			return err
		}
	}
	return nil
}
