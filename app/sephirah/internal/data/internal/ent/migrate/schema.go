// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "platform", Type: field.TypeString},
		{Name: "platform_account_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "profile_url", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_bind_account", Type: field.TypeInt64, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_bind_account",
				Columns:    []*schema.Column{AccountsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "account_platform_platform_account_id",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[1], AccountsColumns[2]},
			},
		},
	}
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "public", Type: field.TypeBool},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_info_app", Type: field.TypeInt64, Nullable: true},
		{Name: "user_app", Type: field.TypeInt64},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "apps_app_infos_app",
				Columns:    []*schema.Column{AppsColumns[6]},
				RefColumns: []*schema.Column{AppInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "apps_users_app",
				Columns:    []*schema.Column{AppsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppBinariesColumns holds the columns for the "app_binaries" table.
	AppBinariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "size_bytes", Type: field.TypeInt64, Nullable: true},
		{Name: "public_url", Type: field.TypeString, Nullable: true},
		{Name: "sha256", Type: field.TypeBytes, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_info_app_binary", Type: field.TypeInt64, Nullable: true},
	}
	// AppBinariesTable holds the schema information for the "app_binaries" table.
	AppBinariesTable = &schema.Table{
		Name:       "app_binaries",
		Columns:    AppBinariesColumns,
		PrimaryKey: []*schema.Column{AppBinariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_binaries_app_infos_app_binary",
				Columns:    []*schema.Column{AppBinariesColumns[7]},
				RefColumns: []*schema.Column{AppInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "appbinary_sha256",
				Unique:  true,
				Columns: []*schema.Column{AppBinariesColumns[4]},
			},
		},
	}
	// AppInfosColumns holds the columns for the "app_infos" table.
	AppInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "internal", Type: field.TypeBool},
		{Name: "source", Type: field.TypeString},
		{Name: "source_app_id", Type: field.TypeString},
		{Name: "source_url", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"unknown", "game"}},
		{Name: "short_description", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "icon_image_url", Type: field.TypeString, Nullable: true},
		{Name: "background_image_url", Type: field.TypeString, Nullable: true},
		{Name: "cover_image_url", Type: field.TypeString, Nullable: true},
		{Name: "release_date", Type: field.TypeString, Nullable: true},
		{Name: "developer", Type: field.TypeString, Nullable: true},
		{Name: "publisher", Type: field.TypeString, Nullable: true},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_info_bind_external", Type: field.TypeInt64, Nullable: true},
	}
	// AppInfosTable holds the schema information for the "app_infos" table.
	AppInfosTable = &schema.Table{
		Name:       "app_infos",
		Columns:    AppInfosColumns,
		PrimaryKey: []*schema.Column{AppInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_infos_app_infos_bind_external",
				Columns:    []*schema.Column{AppInfosColumns[18]},
				RefColumns: []*schema.Column{AppInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "appinfo_source_source_app_id",
				Unique:  true,
				Columns: []*schema.Column{AppInfosColumns[2], AppInfosColumns[3]},
			},
		},
	}
	// AppInstsColumns holds the columns for the "app_insts" table.
	AppInstsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "device_id", Type: field.TypeInt64},
		{Name: "app_id", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_app_inst", Type: field.TypeInt64},
	}
	// AppInstsTable holds the schema information for the "app_insts" table.
	AppInstsTable = &schema.Table{
		Name:       "app_insts",
		Columns:    AppInstsColumns,
		PrimaryKey: []*schema.Column{AppInstsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_insts_users_app_inst",
				Columns:    []*schema.Column{AppInstsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppInstRunTimesColumns holds the columns for the "app_inst_run_times" table.
	AppInstRunTimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "app_inst_id", Type: field.TypeInt64},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "run_duration", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AppInstRunTimesTable holds the schema information for the "app_inst_run_times" table.
	AppInstRunTimesTable = &schema.Table{
		Name:       "app_inst_run_times",
		Columns:    AppInstRunTimesColumns,
		PrimaryKey: []*schema.Column{AppInstRunTimesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appinstruntime_user_id_app_inst_id",
				Unique:  false,
				Columns: []*schema.Column{AppInstRunTimesColumns[1], AppInstRunTimesColumns[2]},
			},
		},
	}
	// DeviceInfosColumns holds the columns for the "device_infos" table.
	DeviceInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "device_name", Type: field.TypeString},
		{Name: "system_type", Type: field.TypeEnum, Enums: []string{"ios", "android", "web", "windows", "macos", "linux", "unknown"}},
		{Name: "system_version", Type: field.TypeString},
		{Name: "client_name", Type: field.TypeString},
		{Name: "client_source_code_address", Type: field.TypeString},
		{Name: "client_version", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// DeviceInfosTable holds the schema information for the "device_infos" table.
	DeviceInfosTable = &schema.Table{
		Name:       "device_infos",
		Columns:    DeviceInfosColumns,
		PrimaryKey: []*schema.Column{DeviceInfosColumns[0]},
	}
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "language", Type: field.TypeString, Nullable: true},
		{Name: "authors", Type: field.TypeJSON, Nullable: true},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_config_feed", Type: field.TypeInt64, Unique: true},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:       "feeds",
		Columns:    FeedsColumns,
		PrimaryKey: []*schema.Column{FeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feeds_feed_configs_feed",
				Columns:    []*schema.Column{FeedsColumns[9]},
				RefColumns: []*schema.Column{FeedConfigsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeedActionSetsColumns holds the columns for the "feed_action_sets" table.
	FeedActionSetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "actions", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_feed_action_set", Type: field.TypeInt64},
	}
	// FeedActionSetsTable holds the schema information for the "feed_action_sets" table.
	FeedActionSetsTable = &schema.Table{
		Name:       "feed_action_sets",
		Columns:    FeedActionSetsColumns,
		PrimaryKey: []*schema.Column{FeedActionSetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_action_sets_users_feed_action_set",
				Columns:    []*schema.Column{FeedActionSetsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeedConfigsColumns holds the columns for the "feed_configs" table.
	FeedConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "source", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "category", Type: field.TypeString},
		{Name: "pull_interval", Type: field.TypeInt64},
		{Name: "hide_items", Type: field.TypeBool, Default: false},
		{Name: "latest_pull_at", Type: field.TypeTime},
		{Name: "latest_pull_status", Type: field.TypeEnum, Enums: []string{"processing", "success", "failed"}},
		{Name: "latest_pull_message", Type: field.TypeString},
		{Name: "next_pull_begin_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_feed_config", Type: field.TypeInt64},
	}
	// FeedConfigsTable holds the schema information for the "feed_configs" table.
	FeedConfigsTable = &schema.Table{
		Name:       "feed_configs",
		Columns:    FeedConfigsColumns,
		PrimaryKey: []*schema.Column{FeedConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_configs_users_feed_config",
				Columns:    []*schema.Column{FeedConfigsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feedconfig_category",
				Unique:  false,
				Columns: []*schema.Column{FeedConfigsColumns[5]},
			},
		},
	}
	// FeedConfigActionsColumns holds the columns for the "feed_config_actions" table.
	FeedConfigActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "index", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_config_id", Type: field.TypeInt64},
		{Name: "feed_action_set_id", Type: field.TypeInt64},
	}
	// FeedConfigActionsTable holds the schema information for the "feed_config_actions" table.
	FeedConfigActionsTable = &schema.Table{
		Name:       "feed_config_actions",
		Columns:    FeedConfigActionsColumns,
		PrimaryKey: []*schema.Column{FeedConfigActionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_config_actions_feed_configs_feed_config",
				Columns:    []*schema.Column{FeedConfigActionsColumns[4]},
				RefColumns: []*schema.Column{FeedConfigsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "feed_config_actions_feed_action_sets_feed_action_set",
				Columns:    []*schema.Column{FeedConfigActionsColumns[5]},
				RefColumns: []*schema.Column{FeedActionSetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feedconfigaction_feed_config_id_feed_action_set_id",
				Unique:  true,
				Columns: []*schema.Column{FeedConfigActionsColumns[4], FeedConfigActionsColumns[5]},
			},
		},
	}
	// FeedItemsColumns holds the columns for the "feed_items" table.
	FeedItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "authors", Type: field.TypeJSON, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "guid", Type: field.TypeString},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
		{Name: "published", Type: field.TypeString, Nullable: true},
		{Name: "published_parsed", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeString, Nullable: true},
		{Name: "updated_parsed", Type: field.TypeTime, Nullable: true},
		{Name: "enclosures", Type: field.TypeJSON, Nullable: true},
		{Name: "publish_platform", Type: field.TypeString, Nullable: true},
		{Name: "read_count", Type: field.TypeInt64, Default: 0},
		{Name: "digest_description", Type: field.TypeString, Nullable: true},
		{Name: "digest_images", Type: field.TypeJSON, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_id", Type: field.TypeInt64},
	}
	// FeedItemsTable holds the schema information for the "feed_items" table.
	FeedItemsTable = &schema.Table{
		Name:       "feed_items",
		Columns:    FeedItemsColumns,
		PrimaryKey: []*schema.Column{FeedItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_items_feeds_item",
				Columns:    []*schema.Column{FeedItemsColumns[19]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feeditem_feed_id_guid",
				Unique:  true,
				Columns: []*schema.Column{FeedItemsColumns[19], FeedItemsColumns[5]},
			},
			{
				Name:    "feeditem_publish_platform",
				Unique:  false,
				Columns: []*schema.Column{FeedItemsColumns[13]},
			},
		},
	}
	// FeedItemCollectionsColumns holds the columns for the "feed_item_collections" table.
	FeedItemCollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_feed_item_collection", Type: field.TypeInt64},
	}
	// FeedItemCollectionsTable holds the schema information for the "feed_item_collections" table.
	FeedItemCollectionsTable = &schema.Table{
		Name:       "feed_item_collections",
		Columns:    FeedItemCollectionsColumns,
		PrimaryKey: []*schema.Column{FeedItemCollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_item_collections_users_feed_item_collection",
				Columns:    []*schema.Column{FeedItemCollectionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feeditemcollection_category",
				Unique:  false,
				Columns: []*schema.Column{FeedItemCollectionsColumns[3]},
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"gebura_save", "chesed_image"}},
		{Name: "sha256", Type: field.TypeBytes},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_file", Type: field.TypeInt64, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_users_file",
				Columns:    []*schema.Column{FilesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"uploaded", "scanned"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "file_image", Type: field.TypeInt64, Unique: true, Nullable: true},
		{Name: "user_image", Type: field.TypeInt64},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_files_image",
				Columns:    []*schema.Column{ImagesColumns[6]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_users_image",
				Columns:    []*schema.Column{ImagesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotifyFlowsColumns holds the columns for the "notify_flows" table.
	NotifyFlowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_notify_flow", Type: field.TypeInt64},
	}
	// NotifyFlowsTable holds the schema information for the "notify_flows" table.
	NotifyFlowsTable = &schema.Table{
		Name:       "notify_flows",
		Columns:    NotifyFlowsColumns,
		PrimaryKey: []*schema.Column{NotifyFlowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flows_users_notify_flow",
				Columns:    []*schema.Column{NotifyFlowsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotifyFlowSourcesColumns holds the columns for the "notify_flow_sources" table.
	NotifyFlowSourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "filter_include_keywords", Type: field.TypeJSON},
		{Name: "filter_exclude_keywords", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "notify_flow_id", Type: field.TypeInt64},
		{Name: "notify_source_id", Type: field.TypeInt64},
	}
	// NotifyFlowSourcesTable holds the schema information for the "notify_flow_sources" table.
	NotifyFlowSourcesTable = &schema.Table{
		Name:       "notify_flow_sources",
		Columns:    NotifyFlowSourcesColumns,
		PrimaryKey: []*schema.Column{NotifyFlowSourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flow_sources_notify_flows_notify_flow",
				Columns:    []*schema.Column{NotifyFlowSourcesColumns[5]},
				RefColumns: []*schema.Column{NotifyFlowsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "notify_flow_sources_notify_sources_notify_source",
				Columns:    []*schema.Column{NotifyFlowSourcesColumns[6]},
				RefColumns: []*schema.Column{NotifySourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notifyflowsource_notify_flow_id_notify_source_id",
				Unique:  true,
				Columns: []*schema.Column{NotifyFlowSourcesColumns[5], NotifyFlowSourcesColumns[6]},
			},
		},
	}
	// NotifyFlowTargetsColumns holds the columns for the "notify_flow_targets" table.
	NotifyFlowTargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "filter_include_keywords", Type: field.TypeJSON},
		{Name: "filter_exclude_keywords", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "notify_flow_id", Type: field.TypeInt64},
		{Name: "notify_target_id", Type: field.TypeInt64},
	}
	// NotifyFlowTargetsTable holds the schema information for the "notify_flow_targets" table.
	NotifyFlowTargetsTable = &schema.Table{
		Name:       "notify_flow_targets",
		Columns:    NotifyFlowTargetsColumns,
		PrimaryKey: []*schema.Column{NotifyFlowTargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flow_targets_notify_flows_notify_flow",
				Columns:    []*schema.Column{NotifyFlowTargetsColumns[5]},
				RefColumns: []*schema.Column{NotifyFlowsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "notify_flow_targets_notify_targets_notify_target",
				Columns:    []*schema.Column{NotifyFlowTargetsColumns[6]},
				RefColumns: []*schema.Column{NotifyTargetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notifyflowtarget_notify_flow_id_notify_target_id",
				Unique:  true,
				Columns: []*schema.Column{NotifyFlowTargetsColumns[5], NotifyFlowTargetsColumns[6]},
			},
		},
	}
	// NotifySourcesColumns holds the columns for the "notify_sources" table.
	NotifySourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_config_id", Type: field.TypeInt64, Nullable: true},
		{Name: "feed_item_collection_id", Type: field.TypeInt64, Nullable: true},
		{Name: "user_notify_source", Type: field.TypeInt64},
	}
	// NotifySourcesTable holds the schema information for the "notify_sources" table.
	NotifySourcesTable = &schema.Table{
		Name:       "notify_sources",
		Columns:    NotifySourcesColumns,
		PrimaryKey: []*schema.Column{NotifySourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_sources_feed_configs_notify_source",
				Columns:    []*schema.Column{NotifySourcesColumns[3]},
				RefColumns: []*schema.Column{FeedConfigsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "notify_sources_feed_item_collections_notify_source",
				Columns:    []*schema.Column{NotifySourcesColumns[4]},
				RefColumns: []*schema.Column{FeedItemCollectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "notify_sources_users_notify_source",
				Columns:    []*schema.Column{NotifySourcesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotifyTargetsColumns holds the columns for the "notify_targets" table.
	NotifyTargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "destination", Type: field.TypeJSON},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_notify_target", Type: field.TypeInt64},
	}
	// NotifyTargetsTable holds the schema information for the "notify_targets" table.
	NotifyTargetsTable = &schema.Table{
		Name:       "notify_targets",
		Columns:    NotifyTargetsColumns,
		PrimaryKey: []*schema.Column{NotifyTargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_targets_users_notify_target",
				Columns:    []*schema.Column{NotifyTargetsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PorterContextsColumns holds the columns for the "porter_contexts" table.
	PorterContextsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "porter_id", Type: field.TypeInt64},
		{Name: "context", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PorterContextsTable holds the schema information for the "porter_contexts" table.
	PorterContextsTable = &schema.Table{
		Name:       "porter_contexts",
		Columns:    PorterContextsColumns,
		PrimaryKey: []*schema.Column{PorterContextsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "portercontext_user_id_porter_id",
				Unique:  true,
				Columns: []*schema.Column{PorterContextsColumns[1], PorterContextsColumns[2]},
			},
		},
	}
	// PorterInstancesColumns holds the columns for the "porter_instances" table.
	PorterInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "global_name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "region", Type: field.TypeString},
		{Name: "feature_summary", Type: field.TypeJSON},
		{Name: "context_json_schema", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "blocked"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// PorterInstancesTable holds the schema information for the "porter_instances" table.
	PorterInstancesTable = &schema.Table{
		Name:       "porter_instances",
		Columns:    PorterInstancesColumns,
		PrimaryKey: []*schema.Column{PorterInstancesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "porterinstance_address",
				Unique:  true,
				Columns: []*schema.Column{PorterInstancesColumns[4]},
			},
		},
	}
	// SystemNotificationsColumns holds the columns for the "system_notifications" table.
	SystemNotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"system", "user"}},
		{Name: "level", Type: field.TypeEnum, Enums: []string{"info", "warn", "error", "ongoing"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unread", "read", "dismissed"}},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// SystemNotificationsTable holds the schema information for the "system_notifications" table.
	SystemNotificationsTable = &schema.Table{
		Name:       "system_notifications",
		Columns:    SystemNotificationsColumns,
		PrimaryKey: []*schema.Column{SystemNotificationsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "public", Type: field.TypeBool, Default: false},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_tag", Type: field.TypeInt64},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_users_tag",
				Columns:    []*schema.Column{TagsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "blocked"}},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"admin", "normal", "sentinel"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_created_user", Type: field.TypeInt64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_created_user",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserDevicesColumns holds the columns for the "user_devices" table.
	UserDevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "device_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// UserDevicesTable holds the schema information for the "user_devices" table.
	UserDevicesTable = &schema.Table{
		Name:       "user_devices",
		Columns:    UserDevicesColumns,
		PrimaryKey: []*schema.Column{UserDevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_devices_device_infos_device_info",
				Columns:    []*schema.Column{UserDevicesColumns[3]},
				RefColumns: []*schema.Column{DeviceInfosColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_devices_users_user",
				Columns:    []*schema.Column{UserDevicesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userdevice_user_id_device_id",
				Unique:  true,
				Columns: []*schema.Column{UserDevicesColumns[4], UserDevicesColumns[3]},
			},
		},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "expire_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "device_info_user_session", Type: field.TypeInt64, Nullable: true},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_device_infos_user_session",
				Columns:    []*schema.Column{UserSessionsColumns[6]},
				RefColumns: []*schema.Column{DeviceInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersession_refresh_token",
				Unique:  true,
				Columns: []*schema.Column{UserSessionsColumns[2]},
			},
		},
	}
	// AccountPurchasedAppColumns holds the columns for the "account_purchased_app" table.
	AccountPurchasedAppColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeInt64},
		{Name: "app_info_id", Type: field.TypeInt64},
	}
	// AccountPurchasedAppTable holds the schema information for the "account_purchased_app" table.
	AccountPurchasedAppTable = &schema.Table{
		Name:       "account_purchased_app",
		Columns:    AccountPurchasedAppColumns,
		PrimaryKey: []*schema.Column{AccountPurchasedAppColumns[0], AccountPurchasedAppColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_purchased_app_account_id",
				Columns:    []*schema.Column{AccountPurchasedAppColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_purchased_app_app_info_id",
				Columns:    []*schema.Column{AccountPurchasedAppColumns[1]},
				RefColumns: []*schema.Column{AppInfosColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FeedItemFeedItemCollectionColumns holds the columns for the "feed_item_feed_item_collection" table.
	FeedItemFeedItemCollectionColumns = []*schema.Column{
		{Name: "feed_item_id", Type: field.TypeInt64},
		{Name: "feed_item_collection_id", Type: field.TypeInt64},
	}
	// FeedItemFeedItemCollectionTable holds the schema information for the "feed_item_feed_item_collection" table.
	FeedItemFeedItemCollectionTable = &schema.Table{
		Name:       "feed_item_feed_item_collection",
		Columns:    FeedItemFeedItemCollectionColumns,
		PrimaryKey: []*schema.Column{FeedItemFeedItemCollectionColumns[0], FeedItemFeedItemCollectionColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_item_feed_item_collection_feed_item_id",
				Columns:    []*schema.Column{FeedItemFeedItemCollectionColumns[0]},
				RefColumns: []*schema.Column{FeedItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "feed_item_feed_item_collection_feed_item_collection_id",
				Columns:    []*schema.Column{FeedItemFeedItemCollectionColumns[1]},
				RefColumns: []*schema.Column{FeedItemCollectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPurchasedAppColumns holds the columns for the "user_purchased_app" table.
	UserPurchasedAppColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "app_info_id", Type: field.TypeInt64},
	}
	// UserPurchasedAppTable holds the schema information for the "user_purchased_app" table.
	UserPurchasedAppTable = &schema.Table{
		Name:       "user_purchased_app",
		Columns:    UserPurchasedAppColumns,
		PrimaryKey: []*schema.Column{UserPurchasedAppColumns[0], UserPurchasedAppColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_purchased_app_user_id",
				Columns:    []*schema.Column{UserPurchasedAppColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_purchased_app_app_info_id",
				Columns:    []*schema.Column{UserPurchasedAppColumns[1]},
				RefColumns: []*schema.Column{AppInfosColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AppsTable,
		AppBinariesTable,
		AppInfosTable,
		AppInstsTable,
		AppInstRunTimesTable,
		DeviceInfosTable,
		FeedsTable,
		FeedActionSetsTable,
		FeedConfigsTable,
		FeedConfigActionsTable,
		FeedItemsTable,
		FeedItemCollectionsTable,
		FilesTable,
		ImagesTable,
		NotifyFlowsTable,
		NotifyFlowSourcesTable,
		NotifyFlowTargetsTable,
		NotifySourcesTable,
		NotifyTargetsTable,
		PorterContextsTable,
		PorterInstancesTable,
		SystemNotificationsTable,
		TagsTable,
		UsersTable,
		UserDevicesTable,
		UserSessionsTable,
		AccountPurchasedAppTable,
		FeedItemFeedItemCollectionTable,
		UserPurchasedAppTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AppsTable.ForeignKeys[0].RefTable = AppInfosTable
	AppsTable.ForeignKeys[1].RefTable = UsersTable
	AppBinariesTable.ForeignKeys[0].RefTable = AppInfosTable
	AppInfosTable.ForeignKeys[0].RefTable = AppInfosTable
	AppInstsTable.ForeignKeys[0].RefTable = UsersTable
	FeedsTable.ForeignKeys[0].RefTable = FeedConfigsTable
	FeedActionSetsTable.ForeignKeys[0].RefTable = UsersTable
	FeedConfigsTable.ForeignKeys[0].RefTable = UsersTable
	FeedConfigActionsTable.ForeignKeys[0].RefTable = FeedConfigsTable
	FeedConfigActionsTable.ForeignKeys[1].RefTable = FeedActionSetsTable
	FeedItemsTable.ForeignKeys[0].RefTable = FeedsTable
	FeedItemCollectionsTable.ForeignKeys[0].RefTable = UsersTable
	FilesTable.ForeignKeys[0].RefTable = UsersTable
	ImagesTable.ForeignKeys[0].RefTable = FilesTable
	ImagesTable.ForeignKeys[1].RefTable = UsersTable
	NotifyFlowsTable.ForeignKeys[0].RefTable = UsersTable
	NotifyFlowSourcesTable.ForeignKeys[0].RefTable = NotifyFlowsTable
	NotifyFlowSourcesTable.ForeignKeys[1].RefTable = NotifySourcesTable
	NotifyFlowTargetsTable.ForeignKeys[0].RefTable = NotifyFlowsTable
	NotifyFlowTargetsTable.ForeignKeys[1].RefTable = NotifyTargetsTable
	NotifySourcesTable.ForeignKeys[0].RefTable = FeedConfigsTable
	NotifySourcesTable.ForeignKeys[1].RefTable = FeedItemCollectionsTable
	NotifySourcesTable.ForeignKeys[2].RefTable = UsersTable
	NotifyTargetsTable.ForeignKeys[0].RefTable = UsersTable
	TagsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UserDevicesTable.ForeignKeys[0].RefTable = DeviceInfosTable
	UserDevicesTable.ForeignKeys[1].RefTable = UsersTable
	UserSessionsTable.ForeignKeys[0].RefTable = DeviceInfosTable
	AccountPurchasedAppTable.ForeignKeys[0].RefTable = AccountsTable
	AccountPurchasedAppTable.ForeignKeys[1].RefTable = AppInfosTable
	FeedItemFeedItemCollectionTable.ForeignKeys[0].RefTable = FeedItemsTable
	FeedItemFeedItemCollectionTable.ForeignKeys[1].RefTable = FeedItemCollectionsTable
	UserPurchasedAppTable.ForeignKeys[0].RefTable = UsersTable
	UserPurchasedAppTable.ForeignKeys[1].RefTable = AppInfosTable
}
