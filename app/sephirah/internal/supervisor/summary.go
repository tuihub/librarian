package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/lib/libtype"
)

func (s *Supervisor) GetFeatureSummary() *modelsupervisor.ServerFeatureSummary {
	s.featureSummaryRWMu.RLock()
	defer s.featureSummaryRWMu.RUnlock()
	featureSummary := new(modelsupervisor.ServerFeatureSummary)
	if s.featureSummary != nil {
		_ = libtype.DeepCopyStruct(s.featureSummary, featureSummary)
	}
	return featureSummary
}

func (s *Supervisor) updateFeatureSummary(_ context.Context) {
	s.featureSummaryRWMu.Lock()
	defer s.featureSummaryRWMu.Unlock()

	var instances []*modelsupervisor.PorterInstance
	s.instanceController.Range(func(key string, controller modelsupervisor.PorterInstanceController) bool {
		if controller.ConnectionStatus == modelsupervisor.PorterConnectionStatusActive {
			instances = append(instances, &controller.PorterInstance)
		}
		return true
	})

	featureSummary, featureSummaryMap := summarize(instances)
	s.featureSummary = featureSummary
	s.featureSummaryMap = featureSummaryMap
}

func summarize(
	instances []*modelsupervisor.PorterInstance,
) (*modelsupervisor.ServerFeatureSummary, *modelsupervisor.ServerFeatureSummaryMap) {
	res := new(modelsupervisor.ServerFeatureSummary)
	resMap := modelsupervisor.NewServerFeatureSummaryMap()

	for _, ins := range instances {
		if ins == nil {
			continue
		}
		do := func(flags []*modelsupervisor.FeatureFlag, resMap *libtype.SyncMap[string, []string], res []*modelsupervisor.FeatureFlag) {
			markMap := make(map[string]bool)
			for _, flag := range flags {
				a, _ := resMap.Load(flag.ID)
				if a == nil {
					a = []string{}
				}
				a = append(a, ins.Address)
				resMap.Store(flag.ID, a)
				if markMap[flag.ID] {
					continue
				}
				res = append(res, flag)
				markMap[flag.ID] = true
			}
		}
		do(ins.FeatureSummary.AccountPlatforms, resMap.AccountPlatforms, res.AccountPlatforms)
		do(ins.FeatureSummary.AppInfoSources, resMap.AppInfoSources, res.AppInfoSources)
		do(ins.FeatureSummary.FeedSources, resMap.FeedSources, res.FeedSources)
		do(ins.FeatureSummary.NotifyDestinations, resMap.NotifyDestinations, res.NotifyDestinations)
		do(ins.FeatureSummary.FeedItemActions, resMap.FeedItemActions, res.FeedItemActions)
	}
	return res, resMap
}
