//go:build !goverter

package converter

import (
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

func ToBizNotifyFlowExtend(a *ent.NotifyFlow) *modelnetzach.NotifyFlow {
	res := ToBizNotifyFlow(a)
	if res == nil {
		return res
	}
	if len(a.Edges.NotifyFlowSource) > 0 {
		res.Sources = make([]*modelnetzach.NotifyFlowSource, 0, len(a.Edges.NotifyFlowSource))
		for _, source := range a.Edges.NotifyFlowSource {
			res.Sources = append(res.Sources, &modelnetzach.NotifyFlowSource{
				SourceID: source.NotifySourceID,
				Filter: &modelnetzach.NotifyFilter{
					ExcludeKeywords: source.FilterExcludeKeywords,
					IncludeKeywords: source.FilterIncludeKeywords,
				},
			})
		}
	}
	if len(a.Edges.NotifyFlowTarget) > 0 {
		targets := make([]*modelnetzach.NotifyFlowTarget, 0, len(a.Edges.NotifyFlowTarget))
		for _, target := range a.Edges.NotifyFlowTarget {
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
func ToBizFeedItemDigest(a *ent.FeedItem) *modelyesod.FeedItemDigest {
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
	if a.Edges.Feed != nil {
		if a.Edges.Feed.Image != nil {
			digest.FeedAvatarURL = a.Edges.Feed.Image.URL
		}
		if a.Edges.Feed.Edges.Config != nil {
			digest.FeedConfigName = a.Edges.Feed.Edges.Config.Name
		}
	}
	// TODO incomplete
	return digest
}
