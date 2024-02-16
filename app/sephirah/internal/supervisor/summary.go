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
		for _, account := range feat.SupportedAccounts {
			if supportedAccountPlatforms[account.Platform] {
				continue
			}
			res.SupportedAccountPlatforms = append(res.SupportedAccountPlatforms, account.Platform)
			supportedAccountPlatforms[account.Platform] = true
		}
		for _, appSource := range feat.SupportedAppInfoSources {
			if supportedAppSources[appSource] {
				continue
			}
			res.SupportedAppInfoSources = append(res.SupportedAppInfoSources, appSource)
			supportedAppSources[appSource] = true
		}
		for _, feedSource := range feat.SupportedFeedSources {
			if supportedFeedSources[feedSource] {
				continue
			}
			res.SupportedFeedSources = append(res.SupportedFeedSources, feedSource)
			supportedFeedSources[feedSource] = true
		}
		for _, notifyDestination := range feat.SupportedNotifyDestinations {
			if supportedNotifyDestinations[notifyDestination] {
				continue
			}
			res.SupportedNotifyDestinations = append(res.SupportedNotifyDestinations, notifyDestination)
			supportedNotifyDestinations[notifyDestination] = true
		}
	}
	return res
}
