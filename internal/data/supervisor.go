package data

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"entgo.io/ent/dialect/sql"
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
	var config model.PullAccountInfoConfig
	if err := json.Unmarshal([]byte(request.ConfigJSON), &config); err != nil {
		return client.WithPorterFastFail(ctx)
	}
	if platforms, ok := s.featureMap.AccountPlatforms.Load(config.Platform); ok {
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
	res, err := s.data.db.PorterInstance.Query().Where(porterinstance.AddressEQ(instance.Address)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorter(res), nil
}

func (s *SupervisorRepo) UpsertPorters(ctx context.Context, il []*modelsupervisor.PorterInstance) error {
	instances := make([]*ent.PorterInstanceCreate, len(il))
	for i, instance := range il {
		if instance.BinarySummary == nil {
			instance.BinarySummary = new(modelsupervisor.PorterBinarySummary)
		}
		instances[i] = s.data.db.PorterInstance.Create().
			SetID(instance.ID).
			SetName(instance.BinarySummary.Name).
			SetVersion(instance.BinarySummary.Version).
			SetDescription(instance.BinarySummary.Description).
			SetSourceCodeAddress(instance.BinarySummary.SourceCodeAddress).
			SetBuildVersion(instance.BinarySummary.BuildVersion).
			SetBuildDate(instance.BinarySummary.BuildDate).
			SetGlobalName(instance.GlobalName).
			SetRegion(instance.Region).
			SetAddress(instance.Address).
			SetStatus(converter.ToEntPorterInstanceStatus(instance.Status)).
			SetFeatureSummary(instance.FeatureSummary).
			SetContextJSONSchema(instance.ContextJSONSchema).
			SetConnectionStatus(converter.ToEntPorterConnectionStatus(instance.ConnectionStatus)).
			SetConnectionStatusMessage(instance.ConnectionStatusMessage)
	}
	return s.data.db.PorterInstance.
		CreateBulk(instances...).
		OnConflict(
			sql.ConflictColumns(porterinstance.FieldAddress),
			resolveWithIgnores([]string{
				porterinstance.FieldID,
				porterinstance.FieldStatus,
			}),
		).
		Exec(ctx)
}

func (s *SupervisorRepo) FetchPorterByAddress(
	ctx context.Context,
	address string,
) (*modelsupervisor.PorterInstance, error) {
	p, err := s.data.db.PorterInstance.Query().Where(
		porterinstance.AddressEQ(address),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorter(p), nil
}

func (s *SupervisorRepo) UpdatePorterContext(
	ctx context.Context,
	pc *modelsupervisor.PorterContext,
) (*modelsupervisor.PorterContext, error) {
	err := s.data.db.PorterContext.UpdateOneID(pc.ID).
		SetHandleStatus(converter.ToEntPorterContextHandleStatus(pc.HandleStatus)).
		SetHandleStatusMessage(pc.HandleStatusMessage).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.data.db.PorterContext.Get(ctx, pc.ID)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorterContext(res), nil
}
