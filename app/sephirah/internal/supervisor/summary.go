package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libtype"
)

func (s *Supervisor) GetFeatureSummary() *modelsupervisor.ServerFeatureSummary {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	featureSummary := new(modelsupervisor.ServerFeatureSummary)
	if s.featureSummary != nil {
		_ = libtype.DeepCopyStruct(s.featureSummary, &featureSummary)
	}
	return s.featureSummary
}

func (s *Supervisor) updateFeatureSummary(ctx context.Context) {
	var instances []*modelsupervisor.PorterInstance
	s.instanceController.Range(func(key string, controller modelsupervisor.PorterInstanceController) bool {
		ins, err := s.instanceCache.Get(ctx, key)
		if err == nil && ins != nil && ins.Status == modeltiphereth.UserStatusActive {
			instances = append(instances, ins)
		}
		return true
	})

	featureSummary, featureSummaryMap := summarize(instances)

	s.featureSummaryRWMu.Lock()
	defer s.featureSummaryRWMu.Unlock()
	s.featureSummary = featureSummary
	s.featureSummaryMap = featureSummaryMap
}

func summarize( //nolint:gocognit // how?
	instances []*modelsupervisor.PorterInstance,
) (*modelsupervisor.ServerFeatureSummary, *modelsupervisor.ServerFeatureSummaryMap) {
	res := new(modelsupervisor.ServerFeatureSummary)
	resMap := modelsupervisor.NewServerFeatureSummaryMap()

	supportedAccountPlatforms := make(map[string]bool)
	supportedAppSources := make(map[string]bool)
	supportedFeedSources := make(map[string]bool)
	supportedNotifyDestinations := make(map[string]bool)
	for _, ins := range instances {
		if ins == nil {
			continue
		}
		for _, account := range ins.FeatureSummary.AccountPlatforms {
			a, _ := resMap.AccountPlatforms.Load(ins.Address)
			if a == nil {
				a = []string{}
			}
			a = append(a, account.ID)
			resMap.AccountPlatforms.Store(ins.Address, a)
			if supportedAccountPlatforms[account.ID] {
				continue
			}
			res.AccountPlatforms = append(res.AccountPlatforms, account)
			supportedAccountPlatforms[account.ID] = true
		}
		for _, appSource := range ins.FeatureSummary.AppInfoSources {
			a, _ := resMap.AppInfoSources.Load(ins.Address)
			if a == nil {
				a = []string{}
			}
			a = append(a, appSource.ID)
			resMap.AppInfoSources.Store(ins.Address, a)
			if supportedAppSources[appSource.ID] {
				continue
			}
			res.AppInfoSources = append(res.AppInfoSources, appSource)
			supportedAppSources[appSource.ID] = true
		}
		for _, feedSource := range ins.FeatureSummary.FeedSources {
			a, _ := resMap.FeedSources.Load(ins.Address)
			if a == nil {
				a = []string{}
			}
			a = append(a, feedSource.ID)
			resMap.FeedSources.Store(ins.Address, a)
			if supportedFeedSources[feedSource.ID] {
				continue
			}
			res.FeedSources = append(res.FeedSources, feedSource)
			supportedFeedSources[feedSource.ID] = true
		}
		for _, notifyDestination := range ins.FeatureSummary.NotifyDestinations {
			a, _ := resMap.NotifyDestinations.Load(ins.Address)
			if a == nil {
				a = []string{}
			}
			a = append(a, notifyDestination.ID)
			resMap.NotifyDestinations.Store(ins.Address, a)
			if supportedNotifyDestinations[notifyDestination.ID] {
				continue
			}
			res.NotifyDestinations = append(res.NotifyDestinations, notifyDestination)
			supportedNotifyDestinations[notifyDestination.ID] = true
		}
	}
	return res, resMap
}
