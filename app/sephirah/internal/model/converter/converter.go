//go:build !goverter

package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter@v1.4.0 gen -g ignoreUnexported .

var toPB = &toPBConverterImpl{}   //nolint:gochecknoglobals // checked
var toBiz = &toBizConverterImpl{} //nolint:gochecknoglobals // checked

func ToPBServerFeatureSummary(a *modeltiphereth.ServerFeatureSummary) *pb.ServerFeatureSummary {
	return toPB.ToPBServerFeatureSummary(a)
}
func ToPBDeviceInfo(a *modeltiphereth.DeviceInfo) *pb.DeviceInfo {
	return toPB.ToPBDeviceInfo(a)
}
func ToPBDeviceInfoList(a []*modeltiphereth.DeviceInfo) []*pb.DeviceInfo {
	return toPB.ToPBDeviceInfoList(a)
}
func ToPBUserSessionList(a []*modeltiphereth.UserSession) []*pb.UserSession {
	return toPB.ToPBUserSessionList(a)
}
func ToPBUser(a *modeltiphereth.User) *pb.User {
	return toPB.ToPBUser(a)
}
func ToPBUserList(a []*modeltiphereth.User) []*pb.User {
	return toPB.ToPBUserList(a)
}
func ToPBAccount(a *modeltiphereth.Account) *librarian.Account {
	return toPB.ToPBAccount(a)
}
func ToPBAccountList(a []*modeltiphereth.Account) []*librarian.Account {
	return toPB.ToPBAccountList(a)
}
func ToPBPorterList(a []*modeltiphereth.PorterInstance) []*pb.Porter {
	return toPB.ToPBPorterList(a)
}
func ToPBAppInfo(a *modelgebura.AppInfo) *librarian.AppInfo {
	return toPB.ToPBAppInfo(a)
}
func ToPBAppInfoList(a []*modelgebura.AppInfo) []*librarian.AppInfo {
	return toPB.ToPBAppInfoList(a)
}
func ToPBAppInfoMixed(a *modelgebura.AppInfoMixed) *librarian.AppInfoMixed {
	return toPB.ToPBAppInfoMixed(a)
}
func ToPBAppInfoMixedList(a []*modelgebura.AppInfoMixed) []*librarian.AppInfoMixed {
	return toPB.ToPBAppInfoMixedList(a)
}
func ToPBApp(a *modelgebura.App) *pb.App {
	return toPB.ToPBApp(a)
}
func ToPBAppList(a []*modelgebura.App) []*pb.App {
	return toPB.ToPBAppList(a)
}
func ToPBAppBinary(a *modelgebura.AppBinary) *pb.AppBinary {
	return toPB.ToPBAppBinary(a)
}
func ToPBAppInstList(a []*modelgebura.AppInst) []*pb.AppInst {
	return toPB.ToPBAppInstList(a)
}
func ToPBFeed(a *modelfeed.Feed) *librarian.Feed {
	return toPB.ToPBFeed(a)
}
func ToPBFeedItem(a *modelfeed.Item) *librarian.FeedItem {
	return toPB.ToPBFeedItem(a)
}
func ToPBFeedItemList(a []*modelfeed.Item) []*librarian.FeedItem {
	return toPB.ToPBFeedItemList(a)
}
func ToPBFeedImage(a *modelfeed.Image) *librarian.FeedImage {
	return toPB.ToPBFeedImage(a)
}
func ToPBEnclosure(a *modelfeed.Enclosure) *librarian.FeedEnclosure {
	return toPB.ToPBEnclosure(a)
}
func ToPBFeedConfig(a *modelyesod.FeedConfig) *pb.FeedConfig {
	return toPB.ToPBFeedConfig(a)
}
func ToPBFeedWithConfig(a *modelyesod.FeedWithConfig) *pb.ListFeedConfigsResponse_FeedWithConfig {
	return toPB.ToPBFeedWithConfig(a)
}
func ToPBFeedWithConfigList(a []*modelyesod.FeedWithConfig) []*pb.ListFeedConfigsResponse_FeedWithConfig {
	return toPB.ToPBFeedWithConfigList(a)
}
func ToPBItemIDWithFeedID(a *modelyesod.FeedItemDigest) *pb.FeedItemDigest {
	return toPB.ToPBFeedItemDigest(a)
}
func ToPBFeedItemDigestList(a []*modelyesod.FeedItemDigest) []*pb.FeedItemDigest {
	return toPB.ToPBFeedItemDigestList(a)
}
func ToPBFeedItemCollectionList(a []*modelyesod.FeedItemCollection) []*pb.FeedItemCollection {
	return toPB.ToPBFeedItemCollectionList(a)
}
func ToPBTimeRange(a *model.TimeRange) *librarian.TimeRange {
	return toPB.ToPBTimeRange(a)
}
func ToPBInternalIDList(a []model.InternalID) []*librarian.InternalID {
	return toPB.ToPBInternalIDList(a)
}
func ToPBNotifyTarget(a *modelnetzach.NotifyTarget) *pb.NotifyTarget {
	return toPB.ToPBNotifyTarget(a)
}
func ToPBNotifyTargetList(a []*modelnetzach.NotifyTarget) []*pb.NotifyTarget {
	return toPB.ToPBNotifyTargetList(a)
}
func ToPBNotifyFlow(a *modelnetzach.NotifyFlow) *pb.NotifyFlow {
	res := toPB.ToPBNotifyFlow(a)
	for i := range a.Sources {
		res.Sources[i].Source = &pb.NotifyFlowSource_FeedConfigId{
			FeedConfigId: ToPBInternalID(a.Sources[i].FeedConfigID),
		}
	}
	return res
}
func ToPBNotifyFlowList(a []*modelnetzach.NotifyFlow) []*pb.NotifyFlow {
	res := make([]*pb.NotifyFlow, 0, len(a))
	for i := range a {
		res[i] = ToPBNotifyFlow(a[i])
	}
	return res
}
func ToPBSystemNotificationList(a []*modelnetzach.SystemNotification) []*pb.SystemNotification {
	return toPB.ToPBSystemNotificationList(a)
}
func ToBizPorterFeatureSummary(a *porter.PorterFeatureSummary) *modeltiphereth.PorterFeatureSummary {
	return toBiz.ToBizPorterFeatureSummary(a)
}
func ToBizInternalIDList(a []*librarian.InternalID) []model.InternalID {
	return toBiz.ToBizInternalIDList(a)
}
func ToBizUser(a *pb.User) *modeltiphereth.User {
	return toBiz.ToBizUser(a)
}
func ToLibAuthUserTypeList(a []pb.UserType) []libauth.UserType {
	return toBiz.ToLibAuthUserTypeList(a)
}
func ToBizUserStatusList(a []pb.UserStatus) []modeltiphereth.UserStatus {
	return toBiz.ToBizUserStatusList(a)
}
func ToBizPorterPrivilege(a *pb.PorterPrivilege) *modeltiphereth.PorterInstancePrivilege {
	return toBiz.ToBizPorterPrivilege(a)
}
func ToBizDeviceInfo(a *pb.DeviceInfo) *modeltiphereth.DeviceInfo {
	return toBiz.ToBizDeviceInfo(a)
}
func ToBizAppInfo(a *librarian.AppInfo) *modelgebura.AppInfo {
	return toBiz.ToBizAppInfo(a)
}
func ToBizAppInfoList(a []*librarian.AppInfo) []*modelgebura.AppInfo {
	return toBiz.ToBizAppInfoList(a)
}
func ToBizAppInfoID(a *librarian.AppInfoID) *modelgebura.AppInfoID {
	return toBiz.ToBizAppInfoID(a)
}
func ToBizAppInfoIDList(a []*librarian.AppInfoID) []*modelgebura.AppInfoID {
	return toBiz.ToBizAppInfoIDList(a)
}
func ToBizAppTypeList(a []librarian.AppType) []modelgebura.AppType {
	return toBiz.ToBizAppTypeList(a)
}
func ToBizApp(a *pb.App) *modelgebura.App {
	return toBiz.ToBizApp(a)
}
func ToBizAppBinaryList(a []*pb.AppBinary) []*modelgebura.AppBinary {
	return toBiz.ToBizAppBinaryList(a)
}
func ToBizAppInst(a *pb.AppInst) *modelgebura.AppInst {
	return toBiz.ToBizAppInst(a)
}
func ToBizFeedConfig(a *pb.FeedConfig) *modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfig(a)
}
func ToBizFeedConfigStatusList(a []pb.FeedConfigStatus) []modelyesod.FeedConfigStatus {
	return toBiz.ToBizFeedConfigStatusList(a)
}
func ToBizFeedItemCollection(a *pb.FeedItemCollection) *modelyesod.FeedItemCollection {
	return toBiz.ToBizFeedItemCollection(a)
}
func ToBizTimeRange(a *librarian.TimeRange) *model.TimeRange {
	return toBiz.ToBizTimeRange(a)
}
func ToBizNotifyTarget(a *pb.NotifyTarget) *modelnetzach.NotifyTarget {
	return toBiz.ToBizNotifyTarget(a)
}
func ToBizNotifyTargetStatusList(a []pb.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus {
	return toBiz.ToBizNotifyTargetStatusList(a)
}
func ToBizNotifyFlow(a *pb.NotifyFlow) *modelnetzach.NotifyFlow {
	res := toBiz.ToBizNotifyFlow(a)
	for i := range a.GetSources() {
		res.Sources[i].FeedConfigID = ToBizInternalID(a.GetSources()[i].GetFeedConfigId())
	}
	return res
}
func ToBizFileMetadata(a *pb.FileMetadata) *modelbinah.FileMetadata {
	return toBiz.ToBizFileMetadata(a)
}
func ToBizSystemNotificationTypeList(a []pb.SystemNotificationType) []modelnetzach.SystemNotificationType {
	return toBiz.ToBizSystemNotificationTypeList(a)
}
func ToBizSystemNotificationLevelList(a []pb.SystemNotificationLevel) []modelnetzach.SystemNotificationLevel {
	return toBiz.ToBizSystemNotificationLevelList(a)
}
func ToBizSystemNotificationStatusList(a []pb.SystemNotificationStatus) []modelnetzach.SystemNotificationStatus {
	return toBiz.ToBizSystemNotificationStatusList(a)
}
