package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/sephirah/internal/model/converter --packageName converter --output ./generated.go ./

var toPB = &toPBConverterImpl{}   //nolint:gochecknoglobals // checked
var toBiz = &toBizConverterImpl{} //nolint:gochecknoglobals // checked

func PtrToString(u *string) string {
	if u == nil {
		return ""
	}
	return *u
}

func DurationPBToDuration(t *durationpb.Duration) time.Duration {
	if t == nil {
		return time.Duration(0)
	}
	return t.AsDuration()
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
func ToPBApp(a *modelgebura.App) *librarian.App {
	return toPB.ToPBApp(a)
}
func ToPBAppList(a []*modelgebura.App) []*librarian.App {
	return toPB.ToPBAppList(a)
}
func ToPBAppPackage(a *modelgebura.AppPackage) *librarian.AppPackage {
	return toPB.ToPBAppPackage(a)
}
func ToPBAppPackageBinary(a *modelgebura.AppPackageBinary) *librarian.AppPackageBinary {
	return toPB.ToPBAppPackageBinary(a)
}
func ToPBAppPackageList(a []*modelgebura.AppPackage) []*librarian.AppPackage {
	return toPB.ToPBAppPackageList(a)
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
func ToPBItemIDWithFeedID(a *modelyesod.FeedItemIDWithFeedID) *pb.FeedItemIDWithFeedID {
	return toPB.ToPBItemIDWithFeedID(a)
}
func ToPBItemIDWithFeedIDList(a []*modelyesod.FeedItemIDWithFeedID) []*pb.FeedItemIDWithFeedID {
	return toPB.ToPBItemIDWithFeedIDList(a)
}
func ToPBTimeRange(a *model.TimeRange) *librarian.TimeRange {
	return toPB.ToPBTimeRange(a)
}
func ToPBInternalIDList(a []model.InternalID) []*librarian.InternalID {
	return toPB.ToPBInternalIDList(a)
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
func ToBizApp(a *librarian.App) *modelgebura.App {
	return toBiz.ToBizApp(a)
}
func ToBizAppTypeList(a []librarian.AppType) []modelgebura.AppType {
	return toBiz.ToBizAppTypeList(a)
}
func ToBizAppSourceList(a []librarian.AppSource) []modelgebura.AppSource {
	return toBiz.ToBizAppSourceList(a)
}
func ToBizAppPackage(a *librarian.AppPackage) *modelgebura.AppPackage {
	return toBiz.ToBizAppPackage(a)
}
func ToBizAppPackageBinary(a *librarian.AppPackageBinary) *modelgebura.AppPackageBinary {
	return toBiz.ToBizAppPackageBinary(a)
}
func ToBizAppPackageSourceList(a []librarian.AppPackageSource) []modelgebura.AppPackageSource {
	return toBiz.ToBizAppPackageSourceList(a)
}
func ToBizFeedConfig(a *pb.FeedConfig) *modelyesod.FeedConfig {
	return toBiz.ToBizFeedConfig(a)
}
func ToBizFeedConfigSourceList(a []pb.FeedConfigSource) []modelyesod.FeedConfigSource {
	return toBiz.ToBizFeedConfigSourceList(a)
}
func ToBizFeedConfigStatusList(a []pb.FeedConfigStatus) []modelyesod.FeedConfigStatus {
	return toBiz.ToBizFeedConfigStatusList(a)
}
func ToBizTimeRange(a *librarian.TimeRange) *model.TimeRange {
	return toBiz.ToBizTimeRange(a)
}
