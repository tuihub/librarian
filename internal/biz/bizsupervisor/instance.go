package bizsupervisor

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/samber/lo"
)

const (
	// Default heartbeat intervals.
	defaultInstanceHeartbeatInitialInterval = 5 * time.Second
	defaultInstanceHeartbeatMaxInterval     = 5 * time.Minute
	defaultInstanceHeartbeatBackoffFactor   = 1.5
	// Default downgrade timeout.
	defaultInstanceDowngradeTimeout = 3 * time.Minute
	// Default re-connect intervals.
	defaultInstanceConnectInitialInterval = 5 * time.Second
	defaultInstanceConnectMaxInterval     = 10 * time.Minute
	defaultInstanceConnectBackoffFactor   = 2
)

type PorterInstanceController struct {
	modelsupervisor.PorterInstance
	s                    *Supervisor
	address              string
	heartbeatCount       int       // Positive for successful heartbeats, negative for failed heartbeats
	lastSuccessHeartbeat time.Time // Last successful heartbeat time
	lastHeartbeat        time.Time
	lastEnabledContext   []model.InternalID
	mu                   sync.Mutex
}

func (s *Supervisor) NewPorterInstanceController(
	address string,
) *PorterInstanceController {
	return &PorterInstanceController{
		PorterInstance: modelsupervisor.PorterInstance{
			ID:                      0,
			BinarySummary:           nil,
			GlobalName:              "",
			Address:                 "",
			Region:                  "",
			FeatureSummary:          nil,
			Status:                  model.UserStatusUnspecified,
			ContextJSONSchema:       "",
			ConnectionStatus:        modelsupervisor.PorterConnectionStatusUnspecified,
			ConnectionStatusMessage: "",
		},
		s:                    s,
		address:              address,
		heartbeatCount:       0,
		lastSuccessHeartbeat: time.Time{},
		lastHeartbeat:        time.Time{},
		lastEnabledContext:   nil,
		mu:                   sync.Mutex{},
	}
}

func (c *PorterInstanceController) UpdateStatus(ctx context.Context) *modelsupervisor.PorterInstance {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Check heartbeat interval
	if c.heartbeatCount > 0 {
		interval := math.Min(
			float64(
				defaultInstanceHeartbeatInitialInterval,
			)*math.Pow(
				defaultInstanceHeartbeatBackoffFactor,
				float64(c.heartbeatCount-1),
			),
			float64(defaultInstanceHeartbeatMaxInterval),
		)
		if time.Since(c.lastHeartbeat) < time.Duration(interval) {
			return &c.PorterInstance
		}
	}
	if c.heartbeatCount < 0 {
		interval := math.Min(
			float64(
				defaultInstanceConnectInitialInterval,
			)*math.Pow(
				defaultInstanceConnectBackoffFactor,
				float64(-c.heartbeatCount-1),
			),
			float64(defaultInstanceConnectMaxInterval),
		)
		if time.Since(c.lastHeartbeat) < time.Duration(interval) {
			return &c.PorterInstance
		}
	}
	// State machine for connection status
	defer func() {
		c.lastHeartbeat = time.Now()
		if c.heartbeatCount > 0 {
			c.lastSuccessHeartbeat = c.lastHeartbeat
		}
	}()
	switch c.ConnectionStatus {
	case modelsupervisor.PorterConnectionStatusUnspecified:
		c.updateStatusUnspecified(ctx)
	case modelsupervisor.PorterConnectionStatusQueueing:
		c.updateStatusQueueing(ctx)
	case modelsupervisor.PorterConnectionStatusConnected:
		c.updateStatusConnected(ctx)
	case modelsupervisor.PorterConnectionStatusActivationFailed:
		c.updateStatusActivationFailed(ctx)
	case modelsupervisor.PorterConnectionStatusActive:
		c.updateStatusActive(ctx)
	case modelsupervisor.PorterConnectionStatusDowngraded:
		c.updateStatusDowngraded(ctx)
	case modelsupervisor.PorterConnectionStatusDisconnected:
		c.updateStatusDisconnected(ctx)
	}
	// Save instance info
	if c.PorterInstance.ID != 0 {
		newInstance, err := c.s.repo.UpsertPorter(ctx, &c.PorterInstance)
		if err != nil {
			logger.Errorf("Failed to upsert porter instance: %v", err)
		} else {
			c.PorterInstance = *newInstance
		}
	}
	return &c.PorterInstance
}

// Initial state, load or create instance info.
func (c *PorterInstanceController) updateStatusUnspecified(ctx context.Context) {
	instance, err := c.s.repo.FetchPorterByAddress(ctx, c.address)
	if data.ErrorIsNotFound(err) {
		id, err1 := c.s.id.New()
		if err1 != nil {
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
		c.ID = id
		c.Address = c.address
		c.Status = model.UserStatusBlocked
		c.ConnectionStatusMessage = "First Initialization"
		_, err2 := c.s.repo.UpsertPorter(ctx, &c.PorterInstance)
		if err2 != nil {
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
	} else if err != nil {
		c.heartbeatCount = min(-1, c.heartbeatCount-1)
		return
	} else {
		c.PorterInstance = *instance
	}
	c.ConnectionStatus = modelsupervisor.PorterConnectionStatusQueueing
}

// Try to connect.
func (c *PorterInstanceController) updateStatusQueueing(ctx context.Context) {
	inst, err := c.getInstanceInfo(ctx)
	if err != nil {
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
		c.ConnectionStatusMessage = err.Error()
		c.heartbeatCount = -1
		return
	}
	c.PorterInstance = lo.FromPtr(inst)
	c.ConnectionStatus = modelsupervisor.PorterConnectionStatusConnected
	c.ConnectionStatusMessage = ""
	c.heartbeatCount = 0
}

// Connected but not active.
func (c *PorterInstanceController) updateStatusConnected(ctx context.Context) {
	if c.PorterInstance.Status == model.UserStatusActive { // Try enable
		err := c.enableInstance(ctx)
		if err != nil {
			c.ConnectionStatus = modelsupervisor.PorterConnectionStatusActivationFailed
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = -1
			return
		}
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusActive
		c.ConnectionStatusMessage = ""
		c.heartbeatCount = 1
	} else { // Keep heartbeat
		_, err := c.getInstanceInfo(ctx)
		if err != nil {
			c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
		c.heartbeatCount = max(1, c.heartbeatCount+1)
	}
}

// Last activation failed.
func (c *PorterInstanceController) updateStatusActivationFailed(ctx context.Context) {
	if c.PorterInstance.Status == model.UserStatusActive { // Try to enable again
		err := c.enableInstance(ctx)
		if err != nil {
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusActive
		c.ConnectionStatusMessage = ""
		c.heartbeatCount = 1
	} else { // Downgrade to connected
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusConnected
		c.heartbeatCount = 0
	}
}

// Activated.
func (c *PorterInstanceController) updateStatusActive(ctx context.Context) {
	if c.PorterInstance.Status == model.UserStatusActive { // Heartbeat
		err := c.enableInstance(ctx)
		if err != nil {
			c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDowngraded
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = -1
			return
		}
		c.heartbeatCount = max(1, c.heartbeatCount+1)
	} else {
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDowngraded
		c.ConnectionStatusMessage = "Disabling instance"
		c.heartbeatCount = 0
	}
}

// Last heartbeat failed.
func (c *PorterInstanceController) updateStatusDowngraded(ctx context.Context) {
	if time.Since(c.lastSuccessHeartbeat) > defaultInstanceDowngradeTimeout {
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
		return
	}
	if c.PorterInstance.Status == model.UserStatusActive { // Try to enable again
		err := c.enableInstance(ctx)
		if err != nil {
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusActive
		c.ConnectionStatusMessage = ""
		c.heartbeatCount = 1
	} else { // Downgrade to connected
		_, err := c.getInstanceInfo(ctx)
		if err != nil {
			c.ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
			c.ConnectionStatusMessage = err.Error()
			c.heartbeatCount = min(-1, c.heartbeatCount-1)
			return
		}
		c.ConnectionStatus = modelsupervisor.PorterConnectionStatusConnected
		c.ConnectionStatusMessage = ""
		c.heartbeatCount = 1
	}
}

// Try to connect.
func (c *PorterInstanceController) updateStatusDisconnected(ctx context.Context) {
	_, err := c.getInstanceInfo(ctx)
	if err != nil {
		c.ConnectionStatusMessage = err.Error()
		c.heartbeatCount = min(-1, c.heartbeatCount-1)
		return
	}
	c.ConnectionStatus = modelsupervisor.PorterConnectionStatusConnected
	c.ConnectionStatusMessage = ""
	c.heartbeatCount = 0
}

func (c *PorterInstanceController) getInstanceInfo(
	ctx context.Context,
) (*modelsupervisor.PorterInstance, error) {
	if c.address == "" {
		// bad address
		return nil, errors.New("address is empty")
	}
	info, err := c.s.porter.GetPorterInformation(
		client.WithPorterAddress(ctx, []string{c.address}),
		&porter.GetPorterInformationRequest{},
	)
	if err != nil {
		// bad instance
		logger.Infof("%s", err.Error())
		return nil, err
	}
	if info == nil || info.GetBinarySummary() == nil {
		// bad instance
		return nil, errors.New("bad instance info")
	}
	feature := converter.ToBizPorterFeatureSummary(info.GetFeatureSummary())
	id, err := c.s.id.New()
	if err != nil {
		return nil, err
	}
	return &modelsupervisor.PorterInstance{
		ID:                      id,
		BinarySummary:           converter.ToBizPorterBinarySummary(info.GetBinarySummary()),
		GlobalName:              info.GetGlobalName(),
		Address:                 c.address,
		Region:                  info.GetRegion(),
		FeatureSummary:          feature,
		Status:                  c.Status,
		ContextJSONSchema:       info.GetContextJsonSchema(),
		ConnectionStatus:        c.ConnectionStatus,
		ConnectionStatusMessage: c.ConnectionStatusMessage,
	}, nil
}

func (c *PorterInstanceController) enableInstance(
	ctx context.Context,
) error {
	resp, err := c.s.porter.EnablePorter(
		client.WithPorterAddress(ctx, []string{c.PorterInstance.Address}),
		&porter.EnablePorterRequest{
			SephirahId:   int64(c.s.app.InstanceUUID.ID()),
			RefreshToken: nil,
		},
	)
	if err != nil {
		return err
	}
	if resp.GetNeedRefreshToken() {
		var refreshToken string
		refreshToken, err = c.s.auth.GenerateToken(
			c.PorterInstance.ID,
			0,
			libauth.ClaimsTypeRefreshToken,
			model.UserTypePorter,
			libtime.Hour,
		)
		if err != nil {
			return err
		}
		_, err = c.s.porter.EnablePorter(
			client.WithPorterAddress(ctx, []string{c.PorterInstance.Address}),
			&porter.EnablePorterRequest{
				SephirahId:   int64(c.s.app.InstanceUUID.ID()),
				RefreshToken: &refreshToken,
			},
		)
	}
	return err
}
