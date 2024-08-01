package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
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
	if platforms := s.featureSummaryMap.AccountPlatforms.Load(platform); platforms != nil {
		return client.WithPorterAddress(ctx, *platforms)
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
	if sources := s.featureSummaryMap.AppInfoSources.Load(source); sources != nil {
		return client.WithPorterAddress(ctx, *sources)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasFeedSource(source *modeltiphereth.FeatureRequest) bool {
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

func (s *Supervisor) WithFeedSource(ctx context.Context, source *modeltiphereth.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if sources := s.featureSummaryMap.FeedSources.Load(source.ID); sources != nil {
		return client.WithPorterAddress(ctx, *sources)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasNotifyDestination(destination *modeltiphereth.FeatureRequest) bool {
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

func (s *Supervisor) WithNotifyDestination(ctx context.Context, destination *modeltiphereth.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if destinations := s.featureSummaryMap.NotifyDestinations.Load(destination.ID); destinations != nil {
		return client.WithPorterAddress(ctx, *destinations)
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) HasFeedItemAction(request *modeltiphereth.FeatureRequest) bool {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	for _, p := range s.featureSummary.FeedItemActions {
		if p.Match(request) {
			return true
		}
	}
	return false
}

func (s *Supervisor) WithFeedItemAction(ctx context.Context, request *modeltiphereth.FeatureRequest) context.Context {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	if actions := s.featureSummaryMap.FeedItemActions.Load(request.ID); actions != nil {
		return client.WithPorterAddress(ctx, *actions)
	}
	return client.WithPorterFastFail(ctx)
}
