package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
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

func (s *Supervisor) CheckFeedSource(source string) bool {
	for _, p := range s.featureSummary.FeedSources {
		if p.ID == source {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallFeedSource(ctx context.Context, source string) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.FeedSources {
			if a.ID == source {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckNotifyDestination(destination string) bool {
	for _, p := range s.featureSummary.NotifyDestinations {
		if p.ID == destination {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallNotifyDestination(ctx context.Context, destination string) context.Context {
	for _, i := range s.aliveInstances {
		for _, a := range i.FeatureSummary.NotifyDestinations {
			if a.ID == destination {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}
