SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="accounts"
CREATE TABLE `accounts` (`id` bigint,`platform` text,`platform_account_id` text,`bound_user_id` bigint,`name` text,`profile_url` text,`avatar_url` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_account` FOREIGN KEY (`bound_user_id`) REFERENCES `users`(`id`))
CREATE INDEX `idx_accounts_bound_user_id` ON `accounts`(`bound_user_id`)
CREATE INDEX `idx_account_platform_id` ON `accounts`(`platform`,`platform_account_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="devices"
CREATE TABLE `devices` (`id` bigint,`device_name` text,`system_type` text,`system_version` text,`client_name` text,`client_source_code_address` text,`client_version` text,`client_local_id` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="apps"
CREATE TABLE `apps` (`id` bigint,`version_number` bigint,`version_date` timestamp with time zone,`user_id` bigint,`creator_device_id` bigint,`app_sources` text,`public` boolean,`bound_store_app_id` bigint,`stop_store_manage` boolean,`name` text,`type` text,`short_description` text,`description` text,`icon_image_url` text,`icon_image_id` bigint,`background_image_url` text,`background_image_id` bigint,`cover_image_url` text,`cover_image_id` bigint,`release_date` text,`developer` text,`publisher` text,`tags` text,`alternative_names` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_app` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),CONSTRAINT `fk_devices_app` FOREIGN KEY (`creator_device_id`) REFERENCES `devices`(`id`))
CREATE INDEX `idx_apps_user_id` ON `apps`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_categories"
CREATE TABLE `app_categories` (`id` bigint,`user_id` bigint,`version_number` bigint,`version_date` timestamp with time zone,`name` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_app_categories_user_id` ON `app_categories`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_app_categories"
CREATE TABLE `app_app_categories` (`app_category_id` bigint,`app_id` bigint,PRIMARY KEY (`app_category_id`,`app_id`),CONSTRAINT `fk_app_app_categories_app_category` FOREIGN KEY (`app_category_id`) REFERENCES `app_categories`(`id`),CONSTRAINT `fk_app_app_categories_app` FOREIGN KEY (`app_id`) REFERENCES `apps`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_categories"
CREATE TABLE `app_categories` (`id` bigint,`user_id` bigint,`version_number` bigint,`version_date` timestamp with time zone,`name` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_app_categories_user_id` ON `app_categories`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="devices"
CREATE TABLE `devices` (`id` bigint,`device_name` text,`system_type` text,`system_version` text,`client_name` text,`client_source_code_address` text,`client_version` text,`client_local_id` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="apps"
CREATE TABLE `apps` (`id` bigint,`version_number` bigint,`version_date` timestamp with time zone,`user_id` bigint,`creator_device_id` bigint,`app_sources` text,`public` boolean,`bound_store_app_id` bigint,`stop_store_manage` boolean,`name` text,`type` text,`short_description` text,`description` text,`icon_image_url` text,`icon_image_id` bigint,`background_image_url` text,`background_image_id` bigint,`cover_image_url` text,`cover_image_id` bigint,`release_date` text,`developer` text,`publisher` text,`tags` text,`alternative_names` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_devices_app` FOREIGN KEY (`creator_device_id`) REFERENCES `devices`(`id`),CONSTRAINT `fk_users_app` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`))
CREATE INDEX `idx_apps_user_id` ON `apps`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_app_categories"
CREATE TABLE `app_app_categories` (`app_id` bigint,`app_category_id` bigint,PRIMARY KEY (`app_id`,`app_category_id`),CONSTRAINT `fk_app_app_categories_app` FOREIGN KEY (`app_id`) REFERENCES `apps`(`id`),CONSTRAINT `fk_app_app_categories_app_category` FOREIGN KEY (`app_category_id`) REFERENCES `app_categories`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_app_categories"
CREATE TABLE `app_app_categories` (`app_category_id` bigint,`app_id` bigint,PRIMARY KEY (`app_category_id`,`app_id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="devices"
CREATE TABLE `devices` (`id` bigint,`device_name` text,`system_type` text,`system_version` text,`client_name` text,`client_source_code_address` text,`client_version` text,`client_local_id` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_action_sets"
CREATE TABLE `feed_action_sets` (`id` bigint,`user_feed_action_set` bigint,`name` text,`description` text,`actions` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_action_sets_owner` FOREIGN KEY (`user_feed_action_set`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_config_actions"
CREATE TABLE `feed_config_actions` (`feed_action_set_id` bigint,`feed_config_id` bigint,PRIMARY KEY (`feed_action_set_id`,`feed_config_id`),CONSTRAINT `fk_feed_config_actions_feed_action_set` FOREIGN KEY (`feed_action_set_id`) REFERENCES `feed_action_sets`(`id`),CONSTRAINT `fk_feed_config_actions_feed_config` FOREIGN KEY (`feed_config_id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_items"
CREATE TABLE `feed_items` (`id` bigint,`feed_id` bigint,`title` text,`authors` text,`description` text,`content` text,`guid` text,`link` text,`image` text,`published` text,`published_parsed` timestamp with time zone,`updated` text,`updated_parsed` timestamp with time zone,`enclosures` text,`publish_platform` text,`read_count` bigint,`digest_description` text,`digest_images` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_item` FOREIGN KEY (`feed_id`) REFERENCES `feeds`(`id`))
CREATE INDEX `idx_feed_items_publish_platform` ON `feed_items`(`publish_platform`)
CREATE INDEX `idx_feed_item_feed_id_guid` ON `feed_items`(`feed_id`,`guid`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collections"
CREATE TABLE `feed_item_collections` (`id` bigint,`user_feed_item_collection` bigint,`name` text,`description` text,`category` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_item_collections_owner` FOREIGN KEY (`user_feed_item_collection`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_item_collections_category` ON `feed_item_collections`(`category`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collection_feed_items"
CREATE TABLE `feed_item_collection_feed_items` (`feed_item_collection_id` bigint,`feed_item_id` bigint,PRIMARY KEY (`feed_item_collection_id`,`feed_item_id`),CONSTRAINT `fk_feed_item_collection_feed_items_feed_item_collection` FOREIGN KEY (`feed_item_collection_id`) REFERENCES `feed_item_collections`(`id`),CONSTRAINT `fk_feed_item_collection_feed_items_feed_item` FOREIGN KEY (`feed_item_id`) REFERENCES `feed_items`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_action_sets"
CREATE TABLE `feed_action_sets` (`id` bigint,`user_feed_action_set` bigint,`name` text,`description` text,`actions` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_action_sets_owner` FOREIGN KEY (`user_feed_action_set`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_config_actions"
CREATE TABLE `feed_config_actions` (`feed_config_id` bigint,`feed_action_set_id` bigint,PRIMARY KEY (`feed_config_id`,`feed_action_set_id`),CONSTRAINT `fk_feed_config_actions_feed_config` FOREIGN KEY (`feed_config_id`) REFERENCES `feed_configs`(`id`),CONSTRAINT `fk_feed_config_actions_feed_action_set` FOREIGN KEY (`feed_action_set_id`) REFERENCES `feed_action_sets`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_config_actions"
CREATE TABLE `feed_config_actions` (`feed_config_id` bigint,`feed_action_set_id` bigint,`index` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`feed_config_id`,`feed_action_set_id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collections"
CREATE TABLE `feed_item_collections` (`id` bigint,`user_feed_item_collection` bigint,`name` text,`description` text,`category` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_item_collections_owner` FOREIGN KEY (`user_feed_item_collection`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_item_collections_category` ON `feed_item_collections`(`category`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_items"
CREATE TABLE `feed_items` (`id` bigint,`feed_id` bigint,`title` text,`authors` text,`description` text,`content` text,`guid` text,`link` text,`image` text,`published` text,`published_parsed` timestamp with time zone,`updated` text,`updated_parsed` timestamp with time zone,`enclosures` text,`publish_platform` text,`read_count` bigint,`digest_description` text,`digest_images` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_item` FOREIGN KEY (`feed_id`) REFERENCES `feeds`(`id`))
CREATE INDEX `idx_feed_items_publish_platform` ON `feed_items`(`publish_platform`)
CREATE INDEX `idx_feed_item_feed_id_guid` ON `feed_items`(`feed_id`,`guid`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collection_feed_items"
CREATE TABLE `feed_item_collection_feed_items` (`feed_item_id` bigint,`feed_item_collection_id` bigint,PRIMARY KEY (`feed_item_id`,`feed_item_collection_id`),CONSTRAINT `fk_feed_item_collection_feed_items_feed_item` FOREIGN KEY (`feed_item_id`) REFERENCES `feed_items`(`id`),CONSTRAINT `fk_feed_item_collection_feed_items_feed_item_collection` FOREIGN KEY (`feed_item_collection_id`) REFERENCES `feed_item_collections`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collections"
CREATE TABLE `feed_item_collections` (`id` bigint,`user_feed_item_collection` bigint,`name` text,`description` text,`category` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_item_collections_owner` FOREIGN KEY (`user_feed_item_collection`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_item_collections_category` ON `feed_item_collections`(`category`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_sources"
CREATE TABLE `notify_sources` (`id` bigint,`user_notify_source` bigint,`feed_config_id` bigint,`feed_item_collection_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_notify_sources_feed_config` FOREIGN KEY (`feed_config_id`) REFERENCES `feed_configs`(`id`),CONSTRAINT `fk_notify_sources_feed_item_collection` FOREIGN KEY (`feed_item_collection_id`) REFERENCES `feed_item_collections`(`id`),CONSTRAINT `fk_feed_item_collections_notify_source` FOREIGN KEY (`id`) REFERENCES `feed_item_collections`(`id`),CONSTRAINT `fk_feed_configs_notify_source` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`),CONSTRAINT `fk_users_notify_source` FOREIGN KEY (`user_notify_source`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flows"
CREATE TABLE `notify_flows` (`id` bigint,`user_notify_flow` bigint,`name` text,`description` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_notify_flow` FOREIGN KEY (`user_notify_flow`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_sources"
CREATE TABLE `notify_flow_sources` (`notify_flow_id` bigint,`notify_source_id` bigint,PRIMARY KEY (`notify_flow_id`,`notify_source_id`),CONSTRAINT `fk_notify_flow_sources_notify_flow` FOREIGN KEY (`notify_flow_id`) REFERENCES `notify_flows`(`id`),CONSTRAINT `fk_notify_flow_sources_notify_source` FOREIGN KEY (`notify_source_id`) REFERENCES `notify_sources`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_targets"
CREATE TABLE `notify_targets` (`id` bigint,`user_notify_target` bigint,`name` text,`description` text,`destination` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_notify_target` FOREIGN KEY (`user_notify_target`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flows"
CREATE TABLE `notify_flows` (`id` bigint,`user_notify_flow` bigint,`name` text,`description` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_notify_flow` FOREIGN KEY (`user_notify_flow`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_targets"
CREATE TABLE `notify_flow_targets` (`notify_flow_id` bigint,`notify_target_id` bigint,PRIMARY KEY (`notify_flow_id`,`notify_target_id`),CONSTRAINT `fk_notify_flow_targets_notify_flow` FOREIGN KEY (`notify_flow_id`) REFERENCES `notify_flows`(`id`),CONSTRAINT `fk_notify_flow_targets_notify_target` FOREIGN KEY (`notify_target_id`) REFERENCES `notify_targets`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flows"
CREATE TABLE `notify_flows` (`id` bigint,`user_notify_flow` bigint,`name` text,`description` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_notify_flow` FOREIGN KEY (`user_notify_flow`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_targets"
CREATE TABLE `notify_targets` (`id` bigint,`user_notify_target` bigint,`name` text,`description` text,`destination` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_notify_target` FOREIGN KEY (`user_notify_target`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_targets"
CREATE TABLE `notify_flow_targets` (`notify_target_id` bigint,`notify_flow_id` bigint,PRIMARY KEY (`notify_target_id`,`notify_flow_id`),CONSTRAINT `fk_notify_flow_targets_notify_target` FOREIGN KEY (`notify_target_id`) REFERENCES `notify_targets`(`id`),CONSTRAINT `fk_notify_flow_targets_notify_flow` FOREIGN KEY (`notify_flow_id`) REFERENCES `notify_flows`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_item_collections"
CREATE TABLE `feed_item_collections` (`id` bigint,`user_feed_item_collection` bigint,`name` text,`description` text,`category` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_item_collections_owner` FOREIGN KEY (`user_feed_item_collection`) REFERENCES `users`(`id`))
CREATE INDEX `idx_feed_item_collections_category` ON `feed_item_collections`(`category`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feeds"
CREATE TABLE `feeds` (`id` bigint,`title` text,`link` text,`description` text,`language` text,`authors` text,`image` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_configs_feed` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="feed_configs"
CREATE TABLE `feed_configs` (`id` bigint,`user_feed_config` bigint,`name` text,`description` text,`source` text,`status` text,`category` text,`pull_interval` bigint,`hide_items` boolean,`latest_pull_at` timestamp with time zone,`latest_pull_status` text,`latest_pull_message` text,`next_pull_begin_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_feed_config` FOREIGN KEY (`user_feed_config`) REFERENCES `users`(`id`),CONSTRAINT `fk_feeds_config` FOREIGN KEY (`id`) REFERENCES `feeds`(`id`))
CREATE INDEX `idx_feed_configs_category` ON `feed_configs`(`category`)
CREATE INDEX `idx_feed_configs_user_feed_config` ON `feed_configs`(`user_feed_config`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_sources"
CREATE TABLE `notify_sources` (`id` bigint,`user_notify_source` bigint,`feed_config_id` bigint,`feed_item_collection_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_feed_item_collections_notify_source` FOREIGN KEY (`id`) REFERENCES `feed_item_collections`(`id`),CONSTRAINT `fk_feed_configs_notify_source` FOREIGN KEY (`id`) REFERENCES `feed_configs`(`id`),CONSTRAINT `fk_users_notify_source` FOREIGN KEY (`user_notify_source`) REFERENCES `users`(`id`),CONSTRAINT `fk_notify_sources_feed_config` FOREIGN KEY (`feed_config_id`) REFERENCES `feed_configs`(`id`),CONSTRAINT `fk_notify_sources_feed_item_collection` FOREIGN KEY (`feed_item_collection_id`) REFERENCES `feed_item_collections`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_sources"
CREATE TABLE `notify_flow_sources` (`notify_source_id` bigint,`notify_flow_id` bigint,PRIMARY KEY (`notify_source_id`,`notify_flow_id`),CONSTRAINT `fk_notify_flow_sources_notify_source` FOREIGN KEY (`notify_source_id`) REFERENCES `notify_sources`(`id`),CONSTRAINT `fk_notify_flow_sources_notify_flow` FOREIGN KEY (`notify_flow_id`) REFERENCES `notify_flows`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_sources"
CREATE TABLE `notify_flow_sources` (`notify_flow_id` bigint,`notify_source_id` bigint,`filter_include_keywords` text,`filter_exclude_keywords` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`notify_flow_id`,`notify_source_id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="notify_flow_targets"
CREATE TABLE `notify_flow_targets` (`notify_flow_id` bigint,`notify_target_id` bigint,`filter_include_keywords` text,`filter_exclude_keywords` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`notify_flow_id`,`notify_target_id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="devices"
CREATE TABLE `devices` (`id` bigint,`device_name` text,`system_type` text,`system_version` text,`client_name` text,`client_source_code_address` text,`client_version` text,`client_local_id` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sessions"
CREATE TABLE `sessions` (`id` bigint,`user_id` bigint,`device_id` bigint,`refresh_token` text,`expire_at` timestamp with time zone,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_devices_session` FOREIGN KEY (`device_id`) REFERENCES `devices`(`id`),CONSTRAINT `fk_users_session` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_sessions_refresh_token` ON `sessions`(`refresh_token`)
CREATE INDEX `idx_session_user_id_device_id` ON `sessions`(`user_id`,`device_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="files"
CREATE TABLE `files` (`id` bigint,`user_file` bigint,`name` text,`size` bigint,`type` text,`sha256` bytea,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_file` FOREIGN KEY (`user_file`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="files"
CREATE TABLE `files` (`id` bigint,`user_file` bigint,`name` text,`size` bigint,`type` text,`sha256` bytea,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_file` FOREIGN KEY (`user_file`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="images"
CREATE TABLE `images` (`id` bigint,`user_image` bigint,`file_image` bigint,`name` text,`description` text,`status` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_files_image` FOREIGN KEY (`file_image`) REFERENCES `files`(`id`),CONSTRAINT `fk_users_image` FOREIGN KEY (`user_image`) REFERENCES `users`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="tags"
CREATE TABLE `tags` (`id` bigint,`user_tag` bigint,`name` text,`description` text,`public` boolean,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_tag` FOREIGN KEY (`user_tag`) REFERENCES `users`(`id`))
CREATE INDEX `idx_tags_user_tag` ON `tags`(`user_tag`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="porter_contexts"
CREATE TABLE `porter_contexts` (`id` bigint,`user_porter_context` bigint,`global_name` text,`region` text,`context_json` text,`name` text,`description` text,`status` text,`handle_status` text,`handle_status_message` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_porter_context` FOREIGN KEY (`user_porter_context`) REFERENCES `users`(`id`))
CREATE INDEX `idx_porter_context_global_name_region` ON `porter_contexts`(`global_name`,`region`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="porter_instances"
CREATE TABLE `porter_instances` (`id` bigint,`name` text,`version` text,`description` text,`source_code_address` text,`build_version` text,`build_date` text,`global_name` text,`address` text,`region` text,`feature_summary` text,`context_json_schema` text,`status` text,`connection_status` text,`connection_status_message` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE UNIQUE INDEX `idx_porter_instances_address` ON `porter_instances`(`address`)
CREATE INDEX `idx_porter_instance_global_name_region` ON `porter_instances`(`global_name`,`region`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="kvs"
CREATE TABLE `kvs` (`bucket` text,`key` text,`value` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`bucket`,`key`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="system_notifications"
CREATE TABLE `system_notifications` (`id` bigint,`user_id` bigint,`type` text,`level` text,`status` text,`title` text,`content` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_system_notifications_user_id` ON `system_notifications`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinels"
CREATE TABLE `sentinels` (`id` bigint,`name` text,`description` text,`url` text,`alternative_urls` text,`get_token_path` text,`download_file_base_path` text,`creator_id` bigint,`library_report_sequence` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinel_app_binaries"
CREATE TABLE `sentinel_app_binaries` (`id` bigint,`union_id` text,`sentinel_id` bigint,`sentinel_library_reported_id` bigint,`library_snapshot` timestamp with time zone,`generated_id` text,`size_bytes` bigint,`need_token` boolean,`name` text,`version` text,`developer` text,`publisher` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_sentinel_app_binaries_union_id` ON `sentinel_app_binaries`(`union_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="store_apps"
CREATE TABLE `store_apps` (`id` bigint,`name` text,`description` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="store_app_binaries"
CREATE TABLE `store_app_binaries` (`store_app_id` bigint,`sentinel_app_binary_union_id` bigint,PRIMARY KEY (`store_app_id`,`sentinel_app_binary_union_id`),CONSTRAINT `fk_store_app_binaries_store_app` FOREIGN KEY (`store_app_id`) REFERENCES `store_apps`(`id`),CONSTRAINT `fk_store_app_binaries_sentinel_app_binary` FOREIGN KEY (`sentinel_app_binary_union_id`) REFERENCES `sentinel_app_binaries`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinel_app_binary_files"
CREATE TABLE `sentinel_app_binary_files` (`id` bigint,`sentinel_id` bigint,`sentinel_library_reported_id` bigint,`library_snapshot` timestamp with time zone,`sentinel_app_binary_generated_id` text,`name` text,`size_bytes` bigint,`sha256` bytea,`server_file_path` text,`chunks_info` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinels"
CREATE TABLE `sentinels` (`id` bigint,`name` text,`description` text,`url` text,`alternative_urls` text,`get_token_path` text,`download_file_base_path` text,`creator_id` bigint,`library_report_sequence` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinel_libraries"
CREATE TABLE `sentinel_libraries` (`id` bigint,`sentinel_id` bigint,`reported_id` bigint,`download_base_path` text,`active_snapshot` timestamp with time zone,`library_report_sequence` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_sentinels_sentinel_libraries` FOREIGN KEY (`sentinel_id`) REFERENCES `sentinels`(`id`))
CREATE INDEX `idx_sentinel_libraries_library_report_sequence` ON `sentinel_libraries`(`library_report_sequence`)
CREATE INDEX `idx_sentinel_library_sentinel_id_reported_id` ON `sentinel_libraries`(`sentinel_id`,`reported_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinels"
CREATE TABLE `sentinels` (`id` bigint,`name` text,`description` text,`url` text,`alternative_urls` text,`get_token_path` text,`download_file_base_path` text,`creator_id` bigint,`library_report_sequence` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinel_sessions"
CREATE TABLE `sentinel_sessions` (`id` bigint,`sentinel_id` bigint,`refresh_token` text,`expire_at` timestamp with time zone,`status` text,`creator_id` bigint,`last_used_at` timestamp with time zone,`last_refreshed_at` timestamp with time zone,`refresh_count` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_sentinels_sentinel_sessions` FOREIGN KEY (`sentinel_id`) REFERENCES `sentinels`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="store_apps"
CREATE TABLE `store_apps` (`id` bigint,`name` text,`description` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="sentinel_app_binaries"
CREATE TABLE `sentinel_app_binaries` (`id` bigint,`union_id` text,`sentinel_id` bigint,`sentinel_library_reported_id` bigint,`library_snapshot` timestamp with time zone,`generated_id` text,`size_bytes` bigint,`need_token` boolean,`name` text,`version` text,`developer` text,`publisher` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_sentinel_app_binaries_union_id` ON `sentinel_app_binaries`(`union_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="store_app_binaries"
CREATE TABLE `store_app_binaries` (`sentinel_app_binary_union_id` bigint,`store_app_id` bigint,PRIMARY KEY (`sentinel_app_binary_union_id`,`store_app_id`),CONSTRAINT `fk_store_app_binaries_sentinel_app_binary` FOREIGN KEY (`sentinel_app_binary_union_id`) REFERENCES `sentinel_app_binaries`(`id`),CONSTRAINT `fk_store_app_binaries_store_app` FOREIGN KEY (`store_app_id`) REFERENCES `store_apps`(`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="store_app_binaries"
CREATE TABLE `store_app_binaries` (`store_app_id` bigint,`sentinel_app_binary_union_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`store_app_id`,`sentinel_app_binary_union_id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="devices"
CREATE TABLE `devices` (`id` bigint,`device_name` text,`system_type` text,`system_version` text,`client_name` text,`client_source_code_address` text,`client_version` text,`client_local_id` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
CREATE TABLE `users` (`id` bigint,`username` text,`password` text,`status` text,`type` text,`creator_id` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_created_user` FOREIGN KEY (`creator_id`) REFERENCES `users`(`id`))
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="apps"
CREATE TABLE `apps` (`id` bigint,`version_number` bigint,`version_date` timestamp with time zone,`user_id` bigint,`creator_device_id` bigint,`app_sources` text,`public` boolean,`bound_store_app_id` bigint,`stop_store_manage` boolean,`name` text,`type` text,`short_description` text,`description` text,`icon_image_url` text,`icon_image_id` bigint,`background_image_url` text,`background_image_id` bigint,`cover_image_url` text,`cover_image_id` bigint,`release_date` text,`developer` text,`publisher` text,`tags` text,`alternative_names` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_users_app` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),CONSTRAINT `fk_devices_app` FOREIGN KEY (`creator_device_id`) REFERENCES `devices`(`id`))
CREATE INDEX `idx_apps_user_id` ON `apps`(`user_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_run_times"
CREATE TABLE `app_run_times` (`id` bigint,`user_id` bigint,`app_id` bigint,`device_id` bigint,`start_time` timestamp with time zone,`duration` bigint,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`),CONSTRAINT `fk_apps_app_run_time` FOREIGN KEY (`app_id`) REFERENCES `apps`(`id`))
CREATE UNIQUE INDEX `idx_app_run_time_start_time_duration` ON `app_run_times`(`start_time`,`duration`)
CREATE INDEX `idx_app_run_time_user_id_app_id` ON `app_run_times`(`user_id`,`app_id`)
SELECT count(*) FROM sqlite_master WHERE type='table' AND name="app_infos"
CREATE TABLE `app_infos` (`id` bigint,`source` text,`source_app_id` text,`source_url` text,`name` text,`type` text,`short_description` text,`description` text,`icon_image_url` text,`icon_image_id` bigint,`background_image_url` text,`background_image_id` bigint,`cover_image_url` text,`cover_image_id` bigint,`release_date` text,`developer` text,`publisher` text,`tags` text,`alternative_names` text,`raw_data` text,`updated_at` timestamp with time zone,`created_at` timestamp with time zone,PRIMARY KEY (`id`))
CREATE INDEX `idx_app_info_source_source_app_id` ON `app_infos`(`source`,`source_app_id`)
