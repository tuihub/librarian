package modelsupervisor

import "github.com/tuihub/librarian/model"

type PorterInstance struct {
	ID             model.InternalID
	Name           string
	Version        string
	GlobalName     string
	Address        string
	FeatureSummary *PorterFeatureSummary
}

type PorterFeatureSummary struct {
	SupportedAccounts []struct {
		Platform         string
		AppRelationTypes []model.AccountAppRelationType
	}
	SupportedAppSources         []string
	SupportedFeedSources        []string
	SupportedNotifyDestinations []string
}

type ServerFeatureSummary struct {
	SupportedAccountPlatforms   []string
	SupportedAppSources         []string
	SupportedFeedSources        []string
	SupportedNotifyDestinations []string
}

func Summarize(sums []*PorterFeatureSummary) *ServerFeatureSummary {
	if len(sums) == 0 {
		return nil
	}
	res := new(ServerFeatureSummary)
	for _, sum := range sums {
		if sum == nil {
			continue
		}
		for _, platform := range sum.SupportedAccounts {
			res.SupportedAccountPlatforms = append(res.SupportedAccountPlatforms, platform.Platform)
		}
		res.SupportedAppSources = append(res.SupportedAppSources, sum.SupportedAppSources...)
		res.SupportedFeedSources = append(res.SupportedFeedSources, sum.SupportedFeedSources...)
		res.SupportedNotifyDestinations = append(res.SupportedNotifyDestinations, sum.SupportedNotifyDestinations...)
	}
	return res
}
