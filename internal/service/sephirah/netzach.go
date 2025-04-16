package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahService) CreateNotifyTarget(ctx context.Context, req *sephirah.CreateNotifyTargetRequest) (
	*sephirah.CreateNotifyTargetResponse, error) {
	id, err := s.n.CreateNotifyTarget(ctx, converter.ToBizNotifyTarget(req.GetTarget()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateNotifyTargetResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahService) UpdateNotifyTarget(ctx context.Context, req *sephirah.UpdateNotifyTargetRequest) (
	*sephirah.UpdateNotifyTargetResponse, error) {
	err := s.n.UpdateNotifyTarget(ctx, converter.ToBizNotifyTarget(req.GetTarget()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateNotifyTargetResponse{}, nil
}
func (s *LibrarianSephirahService) ListNotifyTargets(ctx context.Context, req *sephirah.ListNotifyTargetsRequest) (
	*sephirah.ListNotifyTargetsResponse, error) {
	t, total, err := s.n.ListNotifyTargets(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		converter.ToBizNotifyTargetStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListNotifyTargetsResponse{
		Paging:  &librarian.PagingResponse{TotalSize: total},
		Targets: converter.ToPBNotifyTargetList(t),
	}, nil
}
func (s *LibrarianSephirahService) CreateNotifyFlow(ctx context.Context, req *sephirah.CreateNotifyFlowRequest) (
	*sephirah.CreateNotifyFlowResponse, error) {
	id, err := s.n.CreateNotifyFlow(ctx, converter.ToBizNotifyFlow(req.GetFlow()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateNotifyFlowResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahService) UpdateNotifyFlow(ctx context.Context, req *sephirah.UpdateNotifyFlowRequest) (
	*sephirah.UpdateNotifyFlowResponse, error) {
	err := s.n.UpdateNotifyFlow(ctx, converter.ToBizNotifyFlow(req.GetFlow()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateNotifyFlowResponse{}, nil
}
func (s *LibrarianSephirahService) ListNotifyFlows(ctx context.Context, req *sephirah.ListNotifyFlowsRequest) (
	*sephirah.ListNotifyFlowsResponse, error) {
	res, total, err := s.n.ListNotifyFlows(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListNotifyFlowsResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Flows:  converter.ToPBNotifyFlowList(res),
	}, nil
}

func (s *LibrarianSephirahService) ListSystemNotifications(
	ctx context.Context,
	req *sephirah.ListSystemNotificationsRequest,
) (
	*sephirah.ListSystemNotificationsResponse, error) {
	res, total, err := s.n.ListSystemNotifications(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizSystemNotificationLevelList(req.GetLevelFilter()),
		converter.ToBizSystemNotificationStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListSystemNotificationsResponse{
		Paging:        &librarian.PagingResponse{TotalSize: total},
		Notifications: converter.ToPBSystemNotificationList(res),
	}, nil
}
