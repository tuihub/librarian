package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	converter2 "github.com/tuihub/librarian/internal/service/sephirah/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahServiceService) CreateNotifyTarget(ctx context.Context, req *pb.CreateNotifyTargetRequest) (
	*pb.CreateNotifyTargetResponse, error) {
	id, err := s.n.CreateNotifyTarget(ctx, converter2.ToBizNotifyTarget(req.GetTarget()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotifyTargetResponse{
		Id: converter2.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateNotifyTarget(ctx context.Context, req *pb.UpdateNotifyTargetRequest) (
	*pb.UpdateNotifyTargetResponse, error) {
	err := s.n.UpdateNotifyTarget(ctx, converter2.ToBizNotifyTarget(req.GetTarget()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateNotifyTargetResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListNotifyTargets(ctx context.Context, req *pb.ListNotifyTargetsRequest) (
	*pb.ListNotifyTargetsResponse, error) {
	t, total, err := s.n.ListNotifyTargets(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter2.ToBizInternalIDList(req.GetIdFilter()),
		converter2.ToBizNotifyTargetStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListNotifyTargetsResponse{
		Paging:  &librarian.PagingResponse{TotalSize: total},
		Targets: converter2.ToPBNotifyTargetList(t),
	}, nil
}
func (s *LibrarianSephirahServiceService) CreateNotifyFlow(ctx context.Context, req *pb.CreateNotifyFlowRequest) (
	*pb.CreateNotifyFlowResponse, error) {
	id, err := s.n.CreateNotifyFlow(ctx, converter2.ToBizNotifyFlow(req.GetFlow()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotifyFlowResponse{
		Id: converter2.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateNotifyFlow(ctx context.Context, req *pb.UpdateNotifyFlowRequest) (
	*pb.UpdateNotifyFlowResponse, error) {
	err := s.n.UpdateNotifyFlow(ctx, converter2.ToBizNotifyFlow(req.GetFlow()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateNotifyFlowResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListNotifyFlows(ctx context.Context, req *pb.ListNotifyFlowsRequest) (
	*pb.ListNotifyFlowsResponse, error) {
	res, total, err := s.n.ListNotifyFlows(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter2.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListNotifyFlowsResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Flows:  converter2.ToPBNotifyFlowList(res),
	}, nil
}
func (s *LibrarianSephirahServiceService) ListSystemNotifications(ctx context.Context, req *pb.ListSystemNotificationsRequest) (
	*pb.ListSystemNotificationsResponse, error) {
	res, total, err := s.n.ListSystemNotifications(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter2.ToBizSystemNotificationTypeList(req.GetTypeFilter()),
		converter2.ToBizSystemNotificationLevelList(req.GetLevelFilter()),
		converter2.ToBizSystemNotificationStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListSystemNotificationsResponse{
		Paging:        &librarian.PagingResponse{TotalSize: total},
		Notifications: converter2.ToPBSystemNotificationList(res),
	}, nil
}
