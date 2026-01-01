//go:build !goverter

package converter

import (
	"strings"

	"github.com/tuihub/librarian/internal/data/orm/model"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelyesod"
)

func ToBizNotifyFlowExtend(a *model.NotifyFlow) *modelnetzach.NotifyFlow {
	res := ToBizNotifyFlow(a)
	if res == nil {
		return res
	}
	if len(a.NotifyFlowSources) > 0 {
		res.Sources = make([]*modelnetzach.NotifyFlowSource, 0, len(a.NotifyFlowSources))
		for _, source := range a.NotifyFlowSources {
			res.Sources = append(res.Sources, &modelnetzach.NotifyFlowSource{
				SourceID: source.NotifySourceID,
				Filter: &modelnetzach.NotifyFilter{
					ExcludeKeywords: source.FilterExcludeKeywords,
					IncludeKeywords: source.FilterIncludeKeywords,
				},
			})
		}
	}
	if len(a.NotifyFlowTargets) > 0 {
		targets := make([]*modelnetzach.NotifyFlowTarget, 0, len(a.NotifyFlowTargets))
		for _, target := range a.NotifyFlowTargets {
			targets = append(targets, &modelnetzach.NotifyFlowTarget{
				TargetID: target.NotifyTargetID,
				Filter: &modelnetzach.NotifyFilter{
					ExcludeKeywords: target.FilterExcludeKeywords,
					IncludeKeywords: target.FilterIncludeKeywords,
				},
			})
		}
		res.Targets = targets
	}
	return res
}

func ToBizFeedItemDigest(a *model.FeedItem) *modelyesod.FeedItemDigest {
	if a == nil {
		return nil
	}
	digest := new(modelyesod.FeedItemDigest)
	digest.FeedID = a.FeedID
	digest.ItemID = a.ID
	digest.PublishedParsedTime = a.PublishedParsed
	digest.Title = a.Title
	digest.PublishPlatform = a.PublishPlatform
	digest.ShortDescription = a.DigestDescription
	digest.ReadCount = a.ReadCount
	if a.Image != nil {
		digest.AvatarURL = a.Image.URL
	}
	if len(a.Authors) > 0 {
		digest.Authors = ""
		for _, author := range a.Authors {
			digest.Authors = strings.Join([]string{digest.Authors, author.Name}, ", ")
		}
		digest.Authors = strings.TrimPrefix(digest.Authors, ", ")
	}
	for _, img := range a.DigestImages {
		digest.ImageUrls = append(digest.ImageUrls, img.URL)
	}
	if a.Feed != nil {
		if a.Feed.Image != nil {
			digest.FeedAvatarURL = a.Feed.Image.URL
		}
		if a.Feed.Config != nil {
			digest.FeedConfigName = a.Feed.Config.Name
		}
	}
	return digest
}

func ToBizAppCategoryExtend(ac *model.AppCategory) *modelgebura.AppCategory {
	res := ToBizAppCategory(ac)
	// var res *modelgebura.AppCategory
	if res == nil {
		return res
	}
	if len(ac.AppAppCategories) > 0 {
		res.AppIDs = make([]libmodel.InternalID, 0, len(ac.AppAppCategories))
		for _, aac := range ac.AppAppCategories {
			res.AppIDs = append(res.AppIDs, aac.AppID)
		}
	}
	return res
}
