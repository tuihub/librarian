package converter

import (
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen -g ignoreUnexported -g "output:package github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter:converter" -g "output:file generated.go" .

var toEnt = &toEntConverterImpl{} //nolint:gochecknoglobals // checked
var toBiz = &toBizConverterImpl{} //nolint:gochecknoglobals // checked

func ToEntUserTypeList(a []libauth.UserType) []user.Type {
	return toEnt.ToEntUserTypeList(a)
}
func ToEntUserStatusList(a []modeltiphereth.UserStatus) []user.Status {
	return toEnt.ToEntUserStatusList(a)
}
func ToEntApp(a modelgebura.App) ent.App {
	return toEnt.ToEntApp(a)
}
func ToEntAppPackageSourceList(a []modelgebura.AppPackageSource) []apppackage.Source {
	return toEnt.ToEntAppPackageSourceList(a)
}
func ToEntFeedConfigStatusList(a []modelyesod.FeedConfigStatus) []feedconfig.Status {
	return toEnt.ToEntFeedConfigStatusList(a)
}
func ToEntNotifyTargetStatusList(a []modelnetzach.NotifyTargetStatus) []notifytarget.Status {
	return toEnt.ToEntNotifyTargetStatusList(a)
}

func ToBizUser(a *ent.User) *modeltiphereth.User {
	return toBiz.ToBizUser(a)
}
func ToBizUserList(a []*ent.User) []*modeltiphereth.User {
	return toBiz.ToBizUserList(a)
}
func ToBizAccount(a *ent.Account) *modeltiphereth.Account {
	return toBiz.ToBizAccount(a)
}
func ToBizAccountList(a []*ent.Account) []*modeltiphereth.Account {
	return toBiz.ToBizAccountList(a)
}
func ToBizPorterList(a []*ent.PorterInstance) []*modeltiphereth.PorterInstance {
	return toBiz.ToBizPorterList(a)
}
func ToBizApp(a *ent.App) *modelgebura.App {
	return toBiz.ToBizApp(a)
}
func ToBizAppList(a []*ent.App) []*modelgebura.App {
	return toBiz.ToBizAppList(a)
}
func ToBizAppPackage(a *ent.AppPackage) *modelgebura.AppPackage {
	return toBiz.ToBizAppPackage(a)
}
func ToBizAppPackageBinary(a ent.AppPackage) modelgebura.AppPackageBinary {
	return toBiz.ToBizAppPackageBinary(a)
}
func ToBizAppPackageList(a []*ent.AppPackage) []*modelgebura.AppPackage {
	return toBiz.ToBizAppPackageList(a)
}
func ToBizFeedConfig(a *ent.FeedConfig) *modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfig(a)
}
func ToBizFeedConfigList(a []*ent.FeedConfig) []*modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfigList(a)
}
func ToBizFeed(a *ent.Feed) *modelfeed.Feed {
	return toBiz.ToBizFeed(a)
}
func ToBizFeedItem(a *ent.FeedItem) *modelfeed.Item {
	return toBiz.ToBizFeedItem(a)
}
func ToBizFeedItemList(a []*ent.FeedItem) []*modelfeed.Item {
	return toBiz.ToBizFeedItemList(a)
}
func ToBizNotifyTarget(a *ent.NotifyTarget) *modelnetzach.NotifyTarget {
	return toBiz.ToBizNotifyTarget(a)
}
func ToBizNotifyTargetList(a []*ent.NotifyTarget) []*modelnetzach.NotifyTarget {
	return toBiz.ToBizNotifyTargetList(a)
}
func ToBizNotifyFlow(a *ent.NotifyFlow) *modelnetzach.NotifyFlow {
	res := toBiz.ToBizNotifyFlow(a)
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
				ChannelID: target.ChannelID,
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

func ToBizImage(a *ent.Image) *modelchesed.Image {
	return toBiz.ToBizImage(a)
}

func ToBizImageList(a []*ent.Image) []*modelchesed.Image {
	return toBiz.ToBizImageList(a)
}
