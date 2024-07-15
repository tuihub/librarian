package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
)

func (s *Supervisor) CheckAccountPlatform(platform string) bool {
	for _, p := range s.featureSummary.AccountPlatforms {
		if p.ID == platform {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallAccountPlatform(ctx context.Context, platform string) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.AccountPlatforms {
			if a.ID == platform {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckAppInfoSource(source string) bool {
	for _, p := range s.featureSummary.AppInfoSources {
		if p.ID == source {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallAppInfoSource(ctx context.Context, source string) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.AppInfoSources {
			if a.ID == source {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckFeedSource(source *modeltiphereth.FeatureRequest) bool {
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

func (s *Supervisor) CallFeedSource(ctx context.Context, source *modeltiphereth.FeatureRequest) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.FeedSources {
			if a.ID == source.ID {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckNotifyDestination(destination *modeltiphereth.FeatureRequest) bool {
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

func (s *Supervisor) CallNotifyDestination(ctx context.Context, destination *modeltiphereth.FeatureRequest) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.NotifyDestinations {
			if a.ID == destination.ID {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckFeedItemAction(request *modeltiphereth.FeatureRequest) bool {
	for _, p := range s.featureSummary.FeedItemActions {
		if p.Match(request) {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallFeedItemAction(ctx context.Context, request *modeltiphereth.FeatureRequest) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.FeedItemActions {
			if a.Match(request) {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}
