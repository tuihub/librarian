// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/tuihub/librarian/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/internal/data/internal/ent/appcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/appruntime"
	"github.com/tuihub/librarian/internal/data/internal/ent/device"
	"github.com/tuihub/librarian/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedactionset"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedconfigaction"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/internal/data/internal/ent/file"
	"github.com/tuihub/librarian/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/internal/data/internal/ent/schema"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinaryfile"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinellibrary"
	"github.com/tuihub/librarian/internal/data/internal/ent/session"
	"github.com/tuihub/librarian/internal/data/internal/ent/storeapp"
	"github.com/tuihub/librarian/internal/data/internal/ent/storeappbinary"
	"github.com/tuihub/librarian/internal/data/internal/ent/systemnotification"
	"github.com/tuihub/librarian/internal/data/internal/ent/tag"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[7].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[8].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescUpdatedAt is the schema descriptor for updated_at field.
	appDescUpdatedAt := appFields[24].Descriptor()
	// app.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	app.DefaultUpdatedAt = appDescUpdatedAt.Default.(func() time.Time)
	// app.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	app.UpdateDefaultUpdatedAt = appDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appDescCreatedAt is the schema descriptor for created_at field.
	appDescCreatedAt := appFields[25].Descriptor()
	// app.DefaultCreatedAt holds the default value on creation for the created_at field.
	app.DefaultCreatedAt = appDescCreatedAt.Default.(func() time.Time)
	appcategoryFields := schema.AppCategory{}.Fields()
	_ = appcategoryFields
	// appcategoryDescUpdatedAt is the schema descriptor for updated_at field.
	appcategoryDescUpdatedAt := appcategoryFields[5].Descriptor()
	// appcategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appcategory.DefaultUpdatedAt = appcategoryDescUpdatedAt.Default.(func() time.Time)
	// appcategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appcategory.UpdateDefaultUpdatedAt = appcategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appcategoryDescCreatedAt is the schema descriptor for created_at field.
	appcategoryDescCreatedAt := appcategoryFields[6].Descriptor()
	// appcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	appcategory.DefaultCreatedAt = appcategoryDescCreatedAt.Default.(func() time.Time)
	appinfoFields := schema.AppInfo{}.Fields()
	_ = appinfoFields
	// appinfoDescUpdatedAt is the schema descriptor for updated_at field.
	appinfoDescUpdatedAt := appinfoFields[20].Descriptor()
	// appinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appinfo.DefaultUpdatedAt = appinfoDescUpdatedAt.Default.(func() time.Time)
	// appinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appinfo.UpdateDefaultUpdatedAt = appinfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appinfoDescCreatedAt is the schema descriptor for created_at field.
	appinfoDescCreatedAt := appinfoFields[21].Descriptor()
	// appinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	appinfo.DefaultCreatedAt = appinfoDescCreatedAt.Default.(func() time.Time)
	appruntimeFields := schema.AppRunTime{}.Fields()
	_ = appruntimeFields
	// appruntimeDescUpdatedAt is the schema descriptor for updated_at field.
	appruntimeDescUpdatedAt := appruntimeFields[6].Descriptor()
	// appruntime.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appruntime.DefaultUpdatedAt = appruntimeDescUpdatedAt.Default.(func() time.Time)
	// appruntime.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appruntime.UpdateDefaultUpdatedAt = appruntimeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appruntimeDescCreatedAt is the schema descriptor for created_at field.
	appruntimeDescCreatedAt := appruntimeFields[7].Descriptor()
	// appruntime.DefaultCreatedAt holds the default value on creation for the created_at field.
	appruntime.DefaultCreatedAt = appruntimeDescCreatedAt.Default.(func() time.Time)
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescUpdatedAt is the schema descriptor for updated_at field.
	deviceDescUpdatedAt := deviceFields[8].Descriptor()
	// device.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	device.DefaultUpdatedAt = deviceDescUpdatedAt.Default.(func() time.Time)
	// device.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	device.UpdateDefaultUpdatedAt = deviceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deviceDescCreatedAt is the schema descriptor for created_at field.
	deviceDescCreatedAt := deviceFields[9].Descriptor()
	// device.DefaultCreatedAt holds the default value on creation for the created_at field.
	device.DefaultCreatedAt = deviceDescCreatedAt.Default.(func() time.Time)
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedFields[7].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedFields[8].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	feedactionsetFields := schema.FeedActionSet{}.Fields()
	_ = feedactionsetFields
	// feedactionsetDescUpdatedAt is the schema descriptor for updated_at field.
	feedactionsetDescUpdatedAt := feedactionsetFields[4].Descriptor()
	// feedactionset.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedactionset.DefaultUpdatedAt = feedactionsetDescUpdatedAt.Default.(func() time.Time)
	// feedactionset.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedactionset.UpdateDefaultUpdatedAt = feedactionsetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedactionsetDescCreatedAt is the schema descriptor for created_at field.
	feedactionsetDescCreatedAt := feedactionsetFields[5].Descriptor()
	// feedactionset.DefaultCreatedAt holds the default value on creation for the created_at field.
	feedactionset.DefaultCreatedAt = feedactionsetDescCreatedAt.Default.(func() time.Time)
	feedconfigFields := schema.FeedConfig{}.Fields()
	_ = feedconfigFields
	// feedconfigDescHideItems is the schema descriptor for hide_items field.
	feedconfigDescHideItems := feedconfigFields[8].Descriptor()
	// feedconfig.DefaultHideItems holds the default value on creation for the hide_items field.
	feedconfig.DefaultHideItems = feedconfigDescHideItems.Default.(bool)
	// feedconfigDescLatestPullAt is the schema descriptor for latest_pull_at field.
	feedconfigDescLatestPullAt := feedconfigFields[9].Descriptor()
	// feedconfig.DefaultLatestPullAt holds the default value on creation for the latest_pull_at field.
	feedconfig.DefaultLatestPullAt = feedconfigDescLatestPullAt.Default.(time.Time)
	// feedconfigDescNextPullBeginAt is the schema descriptor for next_pull_begin_at field.
	feedconfigDescNextPullBeginAt := feedconfigFields[12].Descriptor()
	// feedconfig.DefaultNextPullBeginAt holds the default value on creation for the next_pull_begin_at field.
	feedconfig.DefaultNextPullBeginAt = feedconfigDescNextPullBeginAt.Default.(time.Time)
	// feedconfigDescUpdatedAt is the schema descriptor for updated_at field.
	feedconfigDescUpdatedAt := feedconfigFields[13].Descriptor()
	// feedconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedconfig.DefaultUpdatedAt = feedconfigDescUpdatedAt.Default.(func() time.Time)
	// feedconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedconfig.UpdateDefaultUpdatedAt = feedconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedconfigDescCreatedAt is the schema descriptor for created_at field.
	feedconfigDescCreatedAt := feedconfigFields[14].Descriptor()
	// feedconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	feedconfig.DefaultCreatedAt = feedconfigDescCreatedAt.Default.(func() time.Time)
	feedconfigactionFields := schema.FeedConfigAction{}.Fields()
	_ = feedconfigactionFields
	// feedconfigactionDescUpdatedAt is the schema descriptor for updated_at field.
	feedconfigactionDescUpdatedAt := feedconfigactionFields[3].Descriptor()
	// feedconfigaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedconfigaction.DefaultUpdatedAt = feedconfigactionDescUpdatedAt.Default.(func() time.Time)
	// feedconfigaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedconfigaction.UpdateDefaultUpdatedAt = feedconfigactionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedconfigactionDescCreatedAt is the schema descriptor for created_at field.
	feedconfigactionDescCreatedAt := feedconfigactionFields[4].Descriptor()
	// feedconfigaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	feedconfigaction.DefaultCreatedAt = feedconfigactionDescCreatedAt.Default.(func() time.Time)
	feeditemFields := schema.FeedItem{}.Fields()
	_ = feeditemFields
	// feeditemDescReadCount is the schema descriptor for read_count field.
	feeditemDescReadCount := feeditemFields[15].Descriptor()
	// feeditem.DefaultReadCount holds the default value on creation for the read_count field.
	feeditem.DefaultReadCount = feeditemDescReadCount.Default.(int64)
	// feeditemDescUpdatedAt is the schema descriptor for updated_at field.
	feeditemDescUpdatedAt := feeditemFields[18].Descriptor()
	// feeditem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feeditem.DefaultUpdatedAt = feeditemDescUpdatedAt.Default.(func() time.Time)
	// feeditem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feeditem.UpdateDefaultUpdatedAt = feeditemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feeditemDescCreatedAt is the schema descriptor for created_at field.
	feeditemDescCreatedAt := feeditemFields[19].Descriptor()
	// feeditem.DefaultCreatedAt holds the default value on creation for the created_at field.
	feeditem.DefaultCreatedAt = feeditemDescCreatedAt.Default.(func() time.Time)
	feeditemcollectionFields := schema.FeedItemCollection{}.Fields()
	_ = feeditemcollectionFields
	// feeditemcollectionDescUpdatedAt is the schema descriptor for updated_at field.
	feeditemcollectionDescUpdatedAt := feeditemcollectionFields[4].Descriptor()
	// feeditemcollection.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feeditemcollection.DefaultUpdatedAt = feeditemcollectionDescUpdatedAt.Default.(func() time.Time)
	// feeditemcollection.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feeditemcollection.UpdateDefaultUpdatedAt = feeditemcollectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feeditemcollectionDescCreatedAt is the schema descriptor for created_at field.
	feeditemcollectionDescCreatedAt := feeditemcollectionFields[5].Descriptor()
	// feeditemcollection.DefaultCreatedAt holds the default value on creation for the created_at field.
	feeditemcollection.DefaultCreatedAt = feeditemcollectionDescCreatedAt.Default.(func() time.Time)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileFields[5].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[6].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescUpdatedAt is the schema descriptor for updated_at field.
	imageDescUpdatedAt := imageFields[4].Descriptor()
	// image.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	image.DefaultUpdatedAt = imageDescUpdatedAt.Default.(func() time.Time)
	// image.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	image.UpdateDefaultUpdatedAt = imageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageFields[5].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
	notifyflowFields := schema.NotifyFlow{}.Fields()
	_ = notifyflowFields
	// notifyflowDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowDescUpdatedAt := notifyflowFields[4].Descriptor()
	// notifyflow.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflow.DefaultUpdatedAt = notifyflowDescUpdatedAt.Default.(func() time.Time)
	// notifyflow.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflow.UpdateDefaultUpdatedAt = notifyflowDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowDescCreatedAt is the schema descriptor for created_at field.
	notifyflowDescCreatedAt := notifyflowFields[5].Descriptor()
	// notifyflow.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflow.DefaultCreatedAt = notifyflowDescCreatedAt.Default.(func() time.Time)
	notifyflowsourceFields := schema.NotifyFlowSource{}.Fields()
	_ = notifyflowsourceFields
	// notifyflowsourceDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowsourceDescUpdatedAt := notifyflowsourceFields[4].Descriptor()
	// notifyflowsource.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflowsource.DefaultUpdatedAt = notifyflowsourceDescUpdatedAt.Default.(func() time.Time)
	// notifyflowsource.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflowsource.UpdateDefaultUpdatedAt = notifyflowsourceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowsourceDescCreatedAt is the schema descriptor for created_at field.
	notifyflowsourceDescCreatedAt := notifyflowsourceFields[5].Descriptor()
	// notifyflowsource.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflowsource.DefaultCreatedAt = notifyflowsourceDescCreatedAt.Default.(func() time.Time)
	notifyflowtargetFields := schema.NotifyFlowTarget{}.Fields()
	_ = notifyflowtargetFields
	// notifyflowtargetDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowtargetDescUpdatedAt := notifyflowtargetFields[4].Descriptor()
	// notifyflowtarget.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflowtarget.DefaultUpdatedAt = notifyflowtargetDescUpdatedAt.Default.(func() time.Time)
	// notifyflowtarget.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflowtarget.UpdateDefaultUpdatedAt = notifyflowtargetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowtargetDescCreatedAt is the schema descriptor for created_at field.
	notifyflowtargetDescCreatedAt := notifyflowtargetFields[5].Descriptor()
	// notifyflowtarget.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflowtarget.DefaultCreatedAt = notifyflowtargetDescCreatedAt.Default.(func() time.Time)
	notifysourceFields := schema.NotifySource{}.Fields()
	_ = notifysourceFields
	// notifysourceDescUpdatedAt is the schema descriptor for updated_at field.
	notifysourceDescUpdatedAt := notifysourceFields[3].Descriptor()
	// notifysource.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifysource.DefaultUpdatedAt = notifysourceDescUpdatedAt.Default.(func() time.Time)
	// notifysource.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifysource.UpdateDefaultUpdatedAt = notifysourceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifysourceDescCreatedAt is the schema descriptor for created_at field.
	notifysourceDescCreatedAt := notifysourceFields[4].Descriptor()
	// notifysource.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifysource.DefaultCreatedAt = notifysourceDescCreatedAt.Default.(func() time.Time)
	notifytargetFields := schema.NotifyTarget{}.Fields()
	_ = notifytargetFields
	// notifytargetDescUpdatedAt is the schema descriptor for updated_at field.
	notifytargetDescUpdatedAt := notifytargetFields[5].Descriptor()
	// notifytarget.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifytarget.DefaultUpdatedAt = notifytargetDescUpdatedAt.Default.(func() time.Time)
	// notifytarget.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifytarget.UpdateDefaultUpdatedAt = notifytargetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifytargetDescCreatedAt is the schema descriptor for created_at field.
	notifytargetDescCreatedAt := notifytargetFields[6].Descriptor()
	// notifytarget.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifytarget.DefaultCreatedAt = notifytargetDescCreatedAt.Default.(func() time.Time)
	portercontextFields := schema.PorterContext{}.Fields()
	_ = portercontextFields
	// portercontextDescUpdatedAt is the schema descriptor for updated_at field.
	portercontextDescUpdatedAt := portercontextFields[7].Descriptor()
	// portercontext.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	portercontext.DefaultUpdatedAt = portercontextDescUpdatedAt.Default.(func() time.Time)
	// portercontext.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	portercontext.UpdateDefaultUpdatedAt = portercontextDescUpdatedAt.UpdateDefault.(func() time.Time)
	// portercontextDescCreatedAt is the schema descriptor for created_at field.
	portercontextDescCreatedAt := portercontextFields[8].Descriptor()
	// portercontext.DefaultCreatedAt holds the default value on creation for the created_at field.
	portercontext.DefaultCreatedAt = portercontextDescCreatedAt.Default.(func() time.Time)
	porterinstanceFields := schema.PorterInstance{}.Fields()
	_ = porterinstanceFields
	// porterinstanceDescUpdatedAt is the schema descriptor for updated_at field.
	porterinstanceDescUpdatedAt := porterinstanceFields[13].Descriptor()
	// porterinstance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	porterinstance.DefaultUpdatedAt = porterinstanceDescUpdatedAt.Default.(func() time.Time)
	// porterinstance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	porterinstance.UpdateDefaultUpdatedAt = porterinstanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// porterinstanceDescCreatedAt is the schema descriptor for created_at field.
	porterinstanceDescCreatedAt := porterinstanceFields[14].Descriptor()
	// porterinstance.DefaultCreatedAt holds the default value on creation for the created_at field.
	porterinstance.DefaultCreatedAt = porterinstanceDescCreatedAt.Default.(func() time.Time)
	sentinelappbinaryFields := schema.SentinelAppBinary{}.Fields()
	_ = sentinelappbinaryFields
	// sentinelappbinaryDescUpdatedAt is the schema descriptor for updated_at field.
	sentinelappbinaryDescUpdatedAt := sentinelappbinaryFields[9].Descriptor()
	// sentinelappbinary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sentinelappbinary.DefaultUpdatedAt = sentinelappbinaryDescUpdatedAt.Default.(func() time.Time)
	// sentinelappbinary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sentinelappbinary.UpdateDefaultUpdatedAt = sentinelappbinaryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sentinelappbinaryDescCreatedAt is the schema descriptor for created_at field.
	sentinelappbinaryDescCreatedAt := sentinelappbinaryFields[10].Descriptor()
	// sentinelappbinary.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentinelappbinary.DefaultCreatedAt = sentinelappbinaryDescCreatedAt.Default.(func() time.Time)
	sentinelappbinaryfileFields := schema.SentinelAppBinaryFile{}.Fields()
	_ = sentinelappbinaryfileFields
	// sentinelappbinaryfileDescUpdatedAt is the schema descriptor for updated_at field.
	sentinelappbinaryfileDescUpdatedAt := sentinelappbinaryfileFields[8].Descriptor()
	// sentinelappbinaryfile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sentinelappbinaryfile.DefaultUpdatedAt = sentinelappbinaryfileDescUpdatedAt.Default.(func() time.Time)
	// sentinelappbinaryfile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sentinelappbinaryfile.UpdateDefaultUpdatedAt = sentinelappbinaryfileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sentinelappbinaryfileDescCreatedAt is the schema descriptor for created_at field.
	sentinelappbinaryfileDescCreatedAt := sentinelappbinaryfileFields[9].Descriptor()
	// sentinelappbinaryfile.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentinelappbinaryfile.DefaultCreatedAt = sentinelappbinaryfileDescCreatedAt.Default.(func() time.Time)
	sentinelinfoFields := schema.SentinelInfo{}.Fields()
	_ = sentinelinfoFields
	// sentinelinfoDescUpdatedAt is the schema descriptor for updated_at field.
	sentinelinfoDescUpdatedAt := sentinelinfoFields[5].Descriptor()
	// sentinelinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sentinelinfo.DefaultUpdatedAt = sentinelinfoDescUpdatedAt.Default.(func() time.Time)
	// sentinelinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sentinelinfo.UpdateDefaultUpdatedAt = sentinelinfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sentinelinfoDescCreatedAt is the schema descriptor for created_at field.
	sentinelinfoDescCreatedAt := sentinelinfoFields[6].Descriptor()
	// sentinelinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentinelinfo.DefaultCreatedAt = sentinelinfoDescCreatedAt.Default.(func() time.Time)
	sentinellibraryFields := schema.SentinelLibrary{}.Fields()
	_ = sentinellibraryFields
	// sentinellibraryDescUpdatedAt is the schema descriptor for updated_at field.
	sentinellibraryDescUpdatedAt := sentinellibraryFields[3].Descriptor()
	// sentinellibrary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sentinellibrary.DefaultUpdatedAt = sentinellibraryDescUpdatedAt.Default.(func() time.Time)
	// sentinellibrary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sentinellibrary.UpdateDefaultUpdatedAt = sentinellibraryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sentinellibraryDescCreatedAt is the schema descriptor for created_at field.
	sentinellibraryDescCreatedAt := sentinellibraryFields[4].Descriptor()
	// sentinellibrary.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentinellibrary.DefaultCreatedAt = sentinellibraryDescCreatedAt.Default.(func() time.Time)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[5].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[6].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	storeappFields := schema.StoreApp{}.Fields()
	_ = storeappFields
	// storeappDescUpdatedAt is the schema descriptor for updated_at field.
	storeappDescUpdatedAt := storeappFields[2].Descriptor()
	// storeapp.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	storeapp.DefaultUpdatedAt = storeappDescUpdatedAt.Default.(func() time.Time)
	// storeapp.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	storeapp.UpdateDefaultUpdatedAt = storeappDescUpdatedAt.UpdateDefault.(func() time.Time)
	// storeappDescCreatedAt is the schema descriptor for created_at field.
	storeappDescCreatedAt := storeappFields[3].Descriptor()
	// storeapp.DefaultCreatedAt holds the default value on creation for the created_at field.
	storeapp.DefaultCreatedAt = storeappDescCreatedAt.Default.(func() time.Time)
	storeappbinaryFields := schema.StoreAppBinary{}.Fields()
	_ = storeappbinaryFields
	// storeappbinaryDescUpdatedAt is the schema descriptor for updated_at field.
	storeappbinaryDescUpdatedAt := storeappbinaryFields[5].Descriptor()
	// storeappbinary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	storeappbinary.DefaultUpdatedAt = storeappbinaryDescUpdatedAt.Default.(func() time.Time)
	// storeappbinary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	storeappbinary.UpdateDefaultUpdatedAt = storeappbinaryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// storeappbinaryDescCreatedAt is the schema descriptor for created_at field.
	storeappbinaryDescCreatedAt := storeappbinaryFields[6].Descriptor()
	// storeappbinary.DefaultCreatedAt holds the default value on creation for the created_at field.
	storeappbinary.DefaultCreatedAt = storeappbinaryDescCreatedAt.Default.(func() time.Time)
	systemnotificationFields := schema.SystemNotification{}.Fields()
	_ = systemnotificationFields
	// systemnotificationDescUpdatedAt is the schema descriptor for updated_at field.
	systemnotificationDescUpdatedAt := systemnotificationFields[7].Descriptor()
	// systemnotification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	systemnotification.DefaultUpdatedAt = systemnotificationDescUpdatedAt.Default.(func() time.Time)
	// systemnotification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	systemnotification.UpdateDefaultUpdatedAt = systemnotificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// systemnotificationDescCreatedAt is the schema descriptor for created_at field.
	systemnotificationDescCreatedAt := systemnotificationFields[8].Descriptor()
	// systemnotification.DefaultCreatedAt holds the default value on creation for the created_at field.
	systemnotification.DefaultCreatedAt = systemnotificationDescCreatedAt.Default.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescPublic is the schema descriptor for public field.
	tagDescPublic := tagFields[4].Descriptor()
	// tag.DefaultPublic holds the default value on creation for the public field.
	tag.DefaultPublic = tagDescPublic.Default.(bool)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagFields[5].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[6].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
