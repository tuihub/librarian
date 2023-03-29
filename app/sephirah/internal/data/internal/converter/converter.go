package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter --packageName converter --output ./generated.go ./

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
func ToEntFeedConfigSourceList(a []modelyesod.FeedConfigSource) []feedconfig.Source {
	return toEnt.ToEntFeedConfigSourceList(a)
}
func ToEntFeedConfigStatusList(a []modelyesod.FeedConfigStatus) []feedconfig.Status {
	return toEnt.ToEntFeedConfigStatusList(a)
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
