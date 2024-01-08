package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
)

func (s *Supervisor) CheckAccountPlatform(platform string) bool {
	for _, p := range s.featureSummary.SupportedAccountPlatforms {
		if p == platform {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallAccountPlatform(ctx context.Context, platform string) context.Context {
	for _, i := range s.instances {
		for _, a := range i.FeatureSummary.SupportedAccounts {
			if a.Platform == platform {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckAppSource(source string) bool {
	for _, p := range s.featureSummary.SupportedAppSources {
		if p == source {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallAppSource(ctx context.Context, source string) context.Context {
	for _, i := range s.instances {
		for _, a := range i.FeatureSummary.SupportedAppSources {
			if a == source {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckFeedSource(source string) bool {
	for _, p := range s.featureSummary.SupportedFeedSources {
		if p == source {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallFeedSource(ctx context.Context, source string) context.Context {
	for _, i := range s.instances {
		for _, a := range i.FeatureSummary.SupportedFeedSources {
			if a == source {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}

func (s *Supervisor) CheckNotifyDestination(destination string) bool {
	for _, p := range s.featureSummary.SupportedNotifyDestinations {
		if p == destination {
			return true
		}
	}
	return false
}

func (s *Supervisor) CallNotifyDestination(ctx context.Context, destination string) context.Context {
	for _, i := range s.instances {
		for _, a := range i.FeatureSummary.SupportedNotifyDestinations {
			if a == destination {
				return client.WithPorterName(ctx, i.GlobalName)
			}
		}
	}
	return client.WithPorterFastFail(ctx)
}
