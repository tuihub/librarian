// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/schema"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[6].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[7].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescUpdatedAt is the schema descriptor for updated_at field.
	appDescUpdatedAt := appFields[13].Descriptor()
	// app.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	app.DefaultUpdatedAt = appDescUpdatedAt.Default.(func() time.Time)
	// app.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	app.UpdateDefaultUpdatedAt = appDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appDescCreatedAt is the schema descriptor for created_at field.
	appDescCreatedAt := appFields[14].Descriptor()
	// app.DefaultCreatedAt holds the default value on creation for the created_at field.
	app.DefaultCreatedAt = appDescCreatedAt.Default.(func() time.Time)
	apppackageFields := schema.AppPackage{}.Fields()
	_ = apppackageFields
	// apppackageDescUpdatedAt is the schema descriptor for updated_at field.
	apppackageDescUpdatedAt := apppackageFields[10].Descriptor()
	// apppackage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apppackage.DefaultUpdatedAt = apppackageDescUpdatedAt.Default.(func() time.Time)
	// apppackage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apppackage.UpdateDefaultUpdatedAt = apppackageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apppackageDescCreatedAt is the schema descriptor for created_at field.
	apppackageDescCreatedAt := apppackageFields[11].Descriptor()
	// apppackage.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppackage.DefaultCreatedAt = apppackageDescCreatedAt.Default.(func() time.Time)
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
	feedconfigFields := schema.FeedConfig{}.Fields()
	_ = feedconfigFields
	// feedconfigDescLatestPullAt is the schema descriptor for latest_pull_at field.
	feedconfigDescLatestPullAt := feedconfigFields[7].Descriptor()
	// feedconfig.DefaultLatestPullAt holds the default value on creation for the latest_pull_at field.
	feedconfig.DefaultLatestPullAt = feedconfigDescLatestPullAt.Default.(time.Time)
	// feedconfigDescNextPullBeginAt is the schema descriptor for next_pull_begin_at field.
	feedconfigDescNextPullBeginAt := feedconfigFields[8].Descriptor()
	// feedconfig.DefaultNextPullBeginAt holds the default value on creation for the next_pull_begin_at field.
	feedconfig.DefaultNextPullBeginAt = feedconfigDescNextPullBeginAt.Default.(time.Time)
	// feedconfigDescUpdatedAt is the schema descriptor for updated_at field.
	feedconfigDescUpdatedAt := feedconfigFields[9].Descriptor()
	// feedconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedconfig.DefaultUpdatedAt = feedconfigDescUpdatedAt.Default.(func() time.Time)
	// feedconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedconfig.UpdateDefaultUpdatedAt = feedconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedconfigDescCreatedAt is the schema descriptor for created_at field.
	feedconfigDescCreatedAt := feedconfigFields[10].Descriptor()
	// feedconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	feedconfig.DefaultCreatedAt = feedconfigDescCreatedAt.Default.(func() time.Time)
	feeditemFields := schema.FeedItem{}.Fields()
	_ = feeditemFields
	// feeditemDescUpdatedAt is the schema descriptor for updated_at field.
	feeditemDescUpdatedAt := feeditemFields[15].Descriptor()
	// feeditem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feeditem.DefaultUpdatedAt = feeditemDescUpdatedAt.Default.(func() time.Time)
	// feeditem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feeditem.UpdateDefaultUpdatedAt = feeditemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feeditemDescCreatedAt is the schema descriptor for created_at field.
	feeditemDescCreatedAt := feeditemFields[16].Descriptor()
	// feeditem.DefaultCreatedAt holds the default value on creation for the created_at field.
	feeditem.DefaultCreatedAt = feeditemDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}