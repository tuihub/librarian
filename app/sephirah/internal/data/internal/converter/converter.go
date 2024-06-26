//go:build !goverter

package converter

import (
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/systemnotification"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter@v1.4.0 gen -g ignoreUnexported .

var toEnt = &toEntConverterImpl{} //nolint:gochecknoglobals // checked
var toBiz = &toBizConverterImpl{} //nolint:gochecknoglobals // checked

func ToEntUserTypeList(a []libauth.UserType) []user.Type {
	return toEnt.ToEntUserTypeList(a)
}
func ToEntUserStatusList(a []modeltiphereth.UserStatus) []user.Status {
	return toEnt.ToEntUserStatusList(a)
}
func ToEntAppInfo(a modelgebura.AppInfo) ent.AppInfo {
	return toEnt.ToEntAppInfo(a)
}
func ToEntFeedConfigStatusList(a []modelyesod.FeedConfigStatus) []feedconfig.Status {
	return toEnt.ToEntFeedConfigStatusList(a)
}
func ToEntNotifyTargetStatusList(a []modelnetzach.NotifyTargetStatus) []notifytarget.Status {
	return toEnt.ToEntNotifyTargetStatusList(a)
}
func ToEntSystemNotificationTypeList(a []modelnetzach.SystemNotificationType) []systemnotification.Type {
	return toEnt.ToEntSystemNotificationTypeList(a)
}
func ToEntSystemNotificationLevelList(a []modelnetzach.SystemNotificationLevel) []systemnotification.Level {
	return toEnt.ToEntSystemNotificationLevelList(a)
}
func ToEntSystemNotificationStatusList(a []modelnetzach.SystemNotificationStatus) []systemnotification.Status {
	return toEnt.ToEntSystemNotificationStatusList(a)
}

func ToBizUser(a *ent.User) *modeltiphereth.User {
	return toBiz.ToBizUser(a)
}
func ToBizUserList(a []*ent.User) []*modeltiphereth.User {
	return toBiz.ToBizUserList(a)
}
func ToBizUserSession(a *ent.UserSession) *modeltiphereth.UserSession {
	return toBiz.ToBizUserSession(a)
}
func ToBizUserSessionList(a []*ent.UserSession) []*modeltiphereth.UserSession {
	return toBiz.ToBizUserSessionList(a)
}
func ToBizDeviceInfo(a *ent.DeviceInfo) *modeltiphereth.DeviceInfo {
	return toBiz.ToBizDeviceInfo(a)
}
func ToBizDeviceInfoList(a []*ent.DeviceInfo) []*modeltiphereth.DeviceInfo {
	return toBiz.ToBizDeviceInfoList(a)
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
func ToBizAppInfo(a *ent.AppInfo) *modelgebura.AppInfo {
	return toBiz.ToBizAppInfo(a)
}
func ToBizAppInfoList(a []*ent.AppInfo) []*modelgebura.AppInfo {
	return toBiz.ToBizAppInfoList(a)
}
func ToBizApp(a *ent.App) *modelgebura.App {
	return toBiz.ToBizApp(a)
}
func ToBizAppBinary(a ent.AppBinary) modelgebura.AppBinary {
	return toBiz.ToBizAppBinary(a)
}
func ToBizAppList(a []*ent.App) []*modelgebura.App {
	return toBiz.ToBizAppList(a)
}
func ToBizAppInstList(a []*ent.AppInst) []*modelgebura.AppInst {
	return toBiz.ToBizAppInstList(a)
}
func ToBizFeedConfig(a *ent.FeedConfig) *modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfig(a)
}
func ToBizFeedConfigList(a []*ent.FeedConfig) []*modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfigList(a)
}
func ToBizFeedActionSetList(a []*ent.FeedActionSet) []*modelyesod.FeedActionSet {
	return toBiz.ToBizFeedActionSetList(a)
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

func ToBizFeedItemCollectionList(a []*ent.FeedItemCollection) []*modelyesod.FeedItemCollection {
	return toBiz.ToBizFeedItemCollectionList(a)
}

func ToBizImage(a *ent.Image) *modelchesed.Image {
	return toBiz.ToBizImage(a)
}

func ToBizImageList(a []*ent.Image) []*modelchesed.Image {
	return toBiz.ToBizImageList(a)
}

func ToBizSystemNotificationList(a []*ent.SystemNotification) []*modelnetzach.SystemNotification {
	return toBiz.ToBizSystemNotificationList(a)
}
