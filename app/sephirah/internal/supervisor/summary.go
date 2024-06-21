package supervisor

import "github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"

func (s *Supervisor) GetFeatureSummary() *modeltiphereth.ServerFeatureSummary {
	s.muFeatureSummary.RLock()
	defer s.muFeatureSummary.RUnlock()
	return s.featureSummary
}

func (s *Supervisor) updateFeatureSummary() {
	s.muFeatureSummary.Lock()
	defer s.muFeatureSummary.Unlock()
	features := make([]*modeltiphereth.PorterFeatureSummary, 0, len(s.aliveInstances))
	for _, instance := range s.aliveInstances {
		if s.knownInstances[instance.Address] != nil &&
			s.knownInstances[instance.Address].Status == modeltiphereth.PorterInstanceStatusActive {
			features = append(features, instance.FeatureSummary)
		}
	}
	s.featureSummary = s.summarize(features)
}

func (s *Supervisor) summarize( //nolint:gocognit // how?
	features []*modeltiphereth.PorterFeatureSummary,
) *modeltiphereth.ServerFeatureSummary {
	res := new(modeltiphereth.ServerFeatureSummary)
	supportedAccountPlatforms := make(map[string]bool)
	supportedAppSources := make(map[string]bool)
	supportedFeedSources := make(map[string]bool)
	supportedNotifyDestinations := make(map[string]bool)
	for _, feat := range features {
		if feat == nil {
			continue
		}
		for _, account := range feat.AccountPlatforms {
			if supportedAccountPlatforms[account.ID] {
				continue
			}
			res.AccountPlatforms = append(res.AccountPlatforms, account)
			supportedAccountPlatforms[account.ID] = true
		}
		for _, appSource := range feat.AppInfoSources {
			if supportedAppSources[appSource.ID] {
				continue
			}
			res.AppInfoSources = append(res.AppInfoSources, appSource)
			supportedAppSources[appSource.ID] = true
		}
		for _, feedSource := range feat.FeedSources {
			if supportedFeedSources[feedSource.ID] {
				continue
			}
			res.FeedSources = append(res.FeedSources, feedSource)
			supportedFeedSources[feedSource.ID] = true
		}
		for _, notifyDestination := range feat.NotifyDestinations {
			if supportedNotifyDestinations[notifyDestination.ID] {
				continue
			}
			res.NotifyDestinations = append(res.NotifyDestinations, notifyDestination)
			supportedNotifyDestinations[notifyDestination.ID] = true
		}
	}
	return res
}
