package supervisor

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/internal/client/client"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model/converter"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func (s *Supervisor) QueuePorterContext(
	ctx context.Context,
	porterContext modelsupervisor.PorterContext,
) error {
	if ic, ok := s.instanceContextController.Load(porterContext.ID); ok && ic.HandleStatus == modelsupervisor.PorterContextHandleStatusActive {
		return nil
	}
	return s.enableContextTopic.Publish(ctx, porterContext)
}

func newEnablePorterContextTopic(
	s *Supervisor,
) *libmq.Topic[modelsupervisor.PorterContext] {
	return libmq.NewTopic[modelsupervisor.PorterContext](
		"EnablePorterContext",
		func(ctx context.Context, c *modelsupervisor.PorterContext) error {
			ic, ok := s.instanceContextController.Load(c.ID)
			if ok && ic.HandleStatus == modelsupervisor.PorterContextHandleStatusActive {
				return nil
			}
			if !ok {
				ic = modelsupervisor.PorterContextController{
					PorterContext:       *c,
					HandleStatus:        modelsupervisor.PorterContextHandleStatusQueueing,
					HandleStatusMessage: "",
					HandlerAddress:      "",
				}
			}
			var available []string
			s.instanceController.Range(func(key string, controller modelsupervisor.PorterInstanceController) bool {
				if controller.ConnectionStatus == modelsupervisor.PorterConnectionStatusActive &&
					controller.GlobalName == c.GlobalName &&
					controller.Region == c.Region {
					available = append(available, controller.PorterInstance.Address)
				}
				return true
			})
			if len(available) == 0 {
				ic.HandleStatus = modelsupervisor.PorterContextHandleStatusBlocked
				ic.HandleStatusMessage = "no available instance"
				s.instanceContextController.Store(c.ID, ic)
				return nil
			}
			libtype.ShuffleSlice(available)
			var err error
			ic.HandleStatus = modelsupervisor.PorterContextHandleStatusBlocked
			ic.HandleStatusMessage = ""
			for _, address := range available {
				_, err = s.porter.EnableContext(
					client.WithPorterAddress(ctx, []string{address}),
					&porter.EnableContextRequest{
						ContextId:   converter.ToPBInternalID(c.ID),
						ContextJson: c.ContextJSON,
					},
				)
				if err == nil {
					ic.HandleStatus = modelsupervisor.PorterContextHandleStatusActive
					ic.HandlerAddress = address
					break
				}
			}
			if err != nil {
				ic.HandleStatusMessage = err.Error()
			}
			s.instanceContextController.Store(c.ID, ic)
			return err
		},
	)
}

func (s *Supervisor) DisablePorterContext(
	ctx context.Context,
	porterContext *modelsupervisor.PorterContext,
) error {
	ic, ok := s.instanceContextController.Load(porterContext.ID)
	if !ok || len(ic.HandlerAddress) == 0 {
		return nil
	}
	_, ok = s.instanceController.Load(ic.HandlerAddress)
	if !ok {
		return fmt.Errorf("instance not found: %s", ic.HandlerAddress)
	}
	_, err := s.porter.DisableContext(
		client.WithPorterAddress(ctx, []string{ic.HandlerAddress}),
		&porter.DisableContextRequest{
			ContextId: converter.ToPBInternalID(porterContext.ID),
		},
	)
	if err != nil {
		return err
	}
	ic.HandleStatus = modelsupervisor.PorterContextHandleStatusQueueing
	ic.HandleStatusMessage = ""
	ic.HandlerAddress = ""
	s.instanceContextController.Store(porterContext.ID, ic)
	return nil
}
