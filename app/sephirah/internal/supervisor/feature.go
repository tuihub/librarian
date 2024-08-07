package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
)

func (s *Supervisor) HasAccountPlatform(platform string) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	for _, p := range s.featureSummary.AccountPlatforms {
		if p.ID == platform {
			return true
		}
	}
	return false
}

func (s *Supervisor) WithAccountPlatform(ctx context.Context, platform string) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if platforms, ok := s.featureSummaryMap.AccountPlatforms.Load(platform); ok {
		return client.WithPorterAddress(ctx, platforms)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasAppInfoSource(source string) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	for _, p := range s.featureSummary.AppInfoSources {
		if p.ID == source {
			return true
		}
	}
	return false
}

func (s *Supervisor) WithAppInfoSource(ctx context.Context, source string) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if sources, ok := s.featureSummaryMap.AppInfoSources.Load(source); ok {
		return client.WithPorterAddress(ctx, sources)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasFeedSource(source *modelsupervisor.FeatureRequest) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
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

func (s *Supervisor) WithFeedSource(ctx context.Context, source *modelsupervisor.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if sources, ok := s.featureSummaryMap.FeedSources.Load(source.ID); ok {
		return client.WithPorterAddress(ctx, sources)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasNotifyDestination(destination *modelsupervisor.FeatureRequest) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
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

func (s *Supervisor) WithNotifyDestination(ctx context.Context, destination *modelsupervisor.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if destinations, ok := s.featureSummaryMap.NotifyDestinations.Load(destination.ID); ok {
		return client.WithPorterAddress(ctx, destinations)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasFeedItemAction(request *modelsupervisor.FeatureRequest) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	for _, p := range s.featureSummary.FeedItemActions {
		if p.Match(request) {
			return true
		}
	}
	return false
}

func (s *Supervisor) WithFeedItemAction(ctx context.Context, request *modelsupervisor.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if actions, ok := s.featureSummaryMap.FeedItemActions.Load(request.ID); ok {
		return client.WithPorterAddress(ctx, actions)
	}
	return client.WithPorterFastFail(ctx)
}
