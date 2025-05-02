package data

import (
	"context"

	"github.com/tuihub/librarian/internal/model/modelangela"
)

const (
	angelaKVBucket           = "angela"
	serverInstanceSummaryKey = "server_instance_summary"
)

type AngelaRepo struct {
	data *Data
}

func NewAngelaRepo(data *Data) *AngelaRepo {
	return &AngelaRepo{
		data: data,
	}
}

func (a *AngelaRepo) SetServerInstanceSummary(
	ctx context.Context,
	summary *modelangela.ServerInstanceSummary,
) error {
	return a.data.kvSetJSON(ctx, angelaKVBucket, serverInstanceSummaryKey, summary)
}

func (a *AngelaRepo) GetServerInstanceSummary(
	ctx context.Context,
) (*modelangela.ServerInstanceSummary, error) {
	summary := new(modelangela.ServerInstanceSummary)
	if err := a.data.kvGetJSON(ctx, angelaKVBucket, serverInstanceSummaryKey, summary); err != nil {
		if exist, err2 := a.data.kvExists(ctx, angelaKVBucket, serverInstanceSummaryKey); !exist && err2 == nil {
			return summary, nil
		}
		return nil, err
	}
	return summary, nil
}
