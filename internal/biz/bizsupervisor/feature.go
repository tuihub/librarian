package bizsupervisor

import (
	"fmt"
	"sync"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
)

type PorterFeatureController struct {
	repo              *data.SupervisorRepo
	nextVersion       int
	newFeatureSummary *modelsupervisor.ServerFeatureSummary
	newFeatureMap     *modelsupervisor.ServerFeatureSummaryMap
	contextInstances  map[string]map[string]int // map[Instance.GlobalName:Instance.Region]map[Instance.Address]version
	mu                sync.Mutex
}

func NewPorterFeatureController(
	sv *Supervisor,
) *PorterFeatureController {
	return &PorterFeatureController{
		repo:              sv.repo,
		nextVersion:       1,
		newFeatureSummary: new(modelsupervisor.ServerFeatureSummary),
		newFeatureMap:     modelsupervisor.NewServerFeatureSummaryMap(),
		contextInstances:  make(map[string]map[string]int),
		mu:                sync.Mutex{},
	}
}

func (c *PorterFeatureController) GetFeatureSummary() *modelsupervisor.ServerFeatureSummary {
	return c.repo.GetFeatureSummary()
}

func (c *PorterFeatureController) GetFeatureMap() *modelsupervisor.ServerFeatureSummaryMap {
	return c.repo.GetFeatureMap()
}

func (c *PorterFeatureController) Update(instance *modelsupervisor.PorterInstance) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if instance == nil {
		return
	}

	do := func(flags []*modelsupervisor.FeatureFlag, resMap *libtype.SyncMap[string, []string], res []*modelsupervisor.FeatureFlag) []*modelsupervisor.FeatureFlag {
		markMap := make(map[string]bool)
		for _, flag := range flags {
			a, _ := resMap.Load(flag.ID)
			if a == nil {
				a = []string{}
			}
			a = append(a, instance.Address)
			resMap.Store(flag.ID, a)
			key := fmt.Sprintf("%s:%s", instance.GlobalName, instance.Region)
			if _, ok := c.contextInstances[key]; !ok {
				c.contextInstances[key] = make(map[string]int)
			}
			c.contextInstances[key][instance.Address] = c.nextVersion
			if markMap[flag.ID] {
				continue
			}
			res = append(res, flag)
			markMap[flag.ID] = true
		}
		return res
	}
	c.newFeatureSummary.AccountPlatforms = do(
		instance.FeatureSummary.AccountPlatforms,
		c.newFeatureMap.AccountPlatforms,
		c.newFeatureSummary.AccountPlatforms,
	)
	c.newFeatureSummary.AppInfoSources = do(
		instance.FeatureSummary.AppInfoSources,
		c.newFeatureMap.AppInfoSources,
		c.newFeatureSummary.AppInfoSources,
	)
	c.newFeatureSummary.FeedSources = do(
		instance.FeatureSummary.FeedSources,
		c.newFeatureMap.FeedSources,
		c.newFeatureSummary.FeedSources,
	)
	c.newFeatureSummary.NotifyDestinations = do(
		instance.FeatureSummary.NotifyDestinations,
		c.newFeatureMap.NotifyDestinations,
		c.newFeatureSummary.NotifyDestinations,
	)
	c.newFeatureSummary.FeedItemActions = do(
		instance.FeatureSummary.FeedItemActions,
		c.newFeatureMap.FeedItemActions,
		c.newFeatureSummary.FeedItemActions,
	)
}

func (c *PorterFeatureController) Commit() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.newFeatureSummary == nil || c.newFeatureMap == nil {
		return
	}

	c.repo.SetFeatureSummary(c.newFeatureSummary)
	c.repo.SetFeatureMap(c.newFeatureMap)
	c.newFeatureSummary = new(modelsupervisor.ServerFeatureSummary)
	c.newFeatureMap = modelsupervisor.NewServerFeatureSummaryMap()
	for key, instances := range c.contextInstances {
		for address, version := range instances {
			if version != c.nextVersion {
				delete(c.contextInstances, address)
			}
		}
		if len(instances) == 0 {
			delete(c.contextInstances, key)
			continue
		}
		delete(c.contextInstances, key)
	}
	c.nextVersion++
}
