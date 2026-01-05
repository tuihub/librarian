package data

import (
	"context"
	"sync"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/data/internal/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"gorm.io/gorm/clause"
)

type SupervisorRepo struct {
	data           *Data
	featureSummary *modelsupervisor.ServerFeatureSummary
	featureMap     *modelsupervisor.ServerFeatureSummaryMap
	featureMu      sync.RWMutex
}

func NewSupervisorRepo(
	data *Data,
) *SupervisorRepo {
	return &SupervisorRepo{
		data:           data,
		featureSummary: new(modelsupervisor.ServerFeatureSummary),
		featureMap:     modelsupervisor.NewServerFeatureSummaryMap(),
		featureMu:      sync.RWMutex{},
	}
}

func (s *SupervisorRepo) SetFeatureSummary(summary *modelsupervisor.ServerFeatureSummary) {
	s.featureMu.Lock()
	defer s.featureMu.Unlock()
	s.featureSummary = summary
}
func (s *SupervisorRepo) GetFeatureSummary() *modelsupervisor.ServerFeatureSummary {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	return s.featureSummary
}
func (s *SupervisorRepo) SetFeatureMap(featureMap *modelsupervisor.ServerFeatureSummaryMap) {
	s.featureMu.Lock()
	defer s.featureMu.Unlock()
	s.featureMap = featureMap
}
func (s *SupervisorRepo) GetFeatureMap() *modelsupervisor.ServerFeatureSummaryMap {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	return s.featureMap
}
func (s *SupervisorRepo) HasAccountPlatform(request *libmodel.FeatureRequest) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	for _, p := range s.featureSummary.AccountPlatforms {
		if p.ID == request.ID {
			return true
		}
	}
	return false
}
func (s *SupervisorRepo) WithAccountPlatform(ctx context.Context, request *libmodel.FeatureRequest) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if platforms, ok := s.featureMap.AccountPlatforms.Load(request.ID); ok {
		return client.WithPorterAddress(ctx, platforms)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasAppInfoSource(source string) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	for _, p := range s.featureSummary.AppInfoSources {
		if p.ID == source {
			return true
		}
	}
	return false
}
func (s *SupervisorRepo) WithAppInfoSource(ctx context.Context, source string) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if sources, ok := s.featureMap.AppInfoSources.Load(source); ok {
		return client.WithPorterAddress(ctx, sources)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasFeedSource(source *libmodel.FeatureRequest) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if source == nil {
		return false
	}
	for _, p := range s.featureSummary.FeedSources {
		if p.ID == source.ID {
			return true
		}
	}
	return false
}
func (s *SupervisorRepo) WithFeedSource(ctx context.Context, source *libmodel.FeatureRequest) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if sources, ok := s.featureMap.FeedSources.Load(source.ID); ok {
		return client.WithPorterAddress(ctx, sources)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasNotifyDestination(destination *libmodel.FeatureRequest) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if destination == nil {
		return false
	}
	for _, p := range s.featureSummary.NotifyDestinations {
		if p.ID == destination.ID {
			return true
		}
	}
	return false
}
func (s *SupervisorRepo) WithNotifyDestination(
	ctx context.Context,
	destination *libmodel.FeatureRequest,
) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if destinations, ok := s.featureMap.NotifyDestinations.Load(destination.ID); ok {
		return client.WithPorterAddress(ctx, destinations)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasFeedItemAction(request *libmodel.FeatureRequest) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	for _, p := range s.featureSummary.FeedItemActions {
		if p.Match(request) {
			return true
		}
	}
	return false
}

func (s *SupervisorRepo) WithFeedItemAction(
	ctx context.Context,
	request *libmodel.FeatureRequest,
) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if actions, ok := s.featureMap.FeedItemActions.Load(request.ID); ok {
		return client.WithPorterAddress(ctx, actions)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *SupervisorRepo) UpsertPorter(
	ctx context.Context,
	instance *modelsupervisor.PorterInstance,
) (*modelsupervisor.PorterInstance, error) {
	err := s.UpsertPorters(ctx, []*modelsupervisor.PorterInstance{instance})
	if err != nil {
		return nil, err
	}

	q := query.Use(s.data.db).PorterInstance
	res, err := q.WithContext(ctx).Where(q.Address.Eq(instance.Address)).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SupervisorRepo) UpsertPorters(ctx context.Context, il []*modelsupervisor.PorterInstance) error {
	for _, instance := range il {
		if instance.BinarySummary == nil {
			instance.BinarySummary = new(modelsupervisor.PorterBinarySummary)
		}
		// BeforeSave hook will handle mapping BinarySummary to fields
	}

	return query.Use(s.data.db).PorterInstance.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "version", "description", "source_code_address",
			"build_version", "build_date", "global_name", "region",
			"feature_summary", "context_json_schema", "connection_status",
			"connection_status_message",
		}),
	}).Create(il...)
}

func (s *SupervisorRepo) FetchPorterByAddress(
	ctx context.Context,
	address string,
) (*modelsupervisor.PorterInstance, error) {
	q := query.Use(s.data.db).PorterInstance
	p, err := q.WithContext(ctx).Where(q.Address.Eq(address)).First()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *SupervisorRepo) UpdatePorterContext(
	ctx context.Context,
	pc *modelsupervisor.PorterContext,
) (*modelsupervisor.PorterContext, error) {
	q := query.Use(s.data.db).PorterContext
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(pc.ID))).
		Updates(&modelsupervisor.PorterContext{
			HandleStatus:        pc.HandleStatus,
			HandleStatusMessage: pc.HandleStatusMessage,
		})
	if err != nil {
		return nil, err
	}

	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(pc.ID))).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}
