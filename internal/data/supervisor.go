package data

import (
	"context"
	"sync"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/model"
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
func (s *SupervisorRepo) HasAccountPlatform(request *model.FeatureRequest) bool {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	for _, p := range s.featureSummary.AccountPlatforms {
		if p.ID == request.ID {
			return true
		}
	}
	return false
}
func (s *SupervisorRepo) WithAccountPlatform(ctx context.Context, request *model.FeatureRequest) context.Context {
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
func (s *SupervisorRepo) HasFeedSource(source *model.FeatureRequest) bool {
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
func (s *SupervisorRepo) WithFeedSource(ctx context.Context, source *model.FeatureRequest) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if sources, ok := s.featureMap.FeedSources.Load(source.ID); ok {
		return client.WithPorterAddress(ctx, sources)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasNotifyDestination(destination *model.FeatureRequest) bool {
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
	destination *model.FeatureRequest,
) context.Context {
	s.featureMu.RLock()
	defer s.featureMu.RUnlock()
	if destinations, ok := s.featureMap.NotifyDestinations.Load(destination.ID); ok {
		return client.WithPorterAddress(ctx, destinations)
	}
	return client.WithPorterFastFail(ctx)
}
func (s *SupervisorRepo) HasFeedItemAction(request *model.FeatureRequest) bool {
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
	request *model.FeatureRequest,
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
	var p gormschema.PorterInstance
	if err := s.data.db.WithContext(ctx).Where("address = ?", instance.Address).First(&p).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorter(&p), nil
}

func (s *SupervisorRepo) UpsertPorters(ctx context.Context, il []*modelsupervisor.PorterInstance) error {
	instances := make([]gormschema.PorterInstance, len(il))
	for i, instance := range il {
		if instance.BinarySummary == nil {
			instance.BinarySummary = new(modelsupervisor.PorterBinarySummary)
		}
		var featureSummaryVal *gormschema.PorterFeatureSummaryVal
		if instance.FeatureSummary != nil {
			v := gormschema.PorterFeatureSummaryVal(*instance.FeatureSummary)
			featureSummaryVal = &v
		}
		instances[i] = gormschema.PorterInstance{
			ID:                      instance.ID,
			Name:                    instance.BinarySummary.Name,
			Version:                 instance.BinarySummary.Version,
			Description:             instance.BinarySummary.Description,
			SourceCodeAddress:       instance.BinarySummary.SourceCodeAddress,
			BuildVersion:            instance.BinarySummary.BuildVersion,
			BuildDate:               instance.BinarySummary.BuildDate,
			GlobalName:              instance.GlobalName,
			Address:                 instance.Address,
			Region:                  instance.Region,
			FeatureSummary:          featureSummaryVal,
			ContextJSONSchema:       instance.ContextJSONSchema,
			Status:                  gormschema.ToSchemaUserStatus(instance.Status),
			ConnectionStatus:        gormschema.ToSchemaPorterConnectionStatus(instance.ConnectionStatus),
			ConnectionStatusMessage: instance.ConnectionStatusMessage,
		}
	}
	return s.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "version", "description", "source_code_address", "build_version", "build_date", "global_name", "region", "feature_summary", "context_json_schema", "connection_status", "connection_status_message", "updated_at"}),
	}).Create(&instances).Error
}

func (s *SupervisorRepo) FetchPorterByAddress(
	ctx context.Context,
	address string,
) (*modelsupervisor.PorterInstance, error) {
	var p gormschema.PorterInstance
	if err := s.data.db.WithContext(ctx).Where("address = ?", address).First(&p).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorter(&p), nil
}

func (s *SupervisorRepo) UpdatePorterContext(
	ctx context.Context,
	pc *modelsupervisor.PorterContext,
) (*modelsupervisor.PorterContext, error) {
	if err := s.data.db.WithContext(ctx).Model(&gormschema.PorterContext{}).
		Where("id = ?", pc.ID).
		Updates(map[string]any{
			"handle_status":         gormschema.ToSchemaPorterContextHandleStatus(pc.HandleStatus),
			"handle_status_message": pc.HandleStatusMessage,
		}).Error; err != nil {
		return nil, err
	}
	var res gormschema.PorterContext
	if err := s.data.db.WithContext(ctx).First(&res, pc.ID).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorterContext(&res), nil
}
