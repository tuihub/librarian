-- V1 SQLite Fresh Install Schema
-- This creates all tables from scratch for a new installation

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    type TEXT NOT NULL,
    status TEXT NOT NULL,
    creator_id INTEGER,
    created_at DATETIME,
    updated_at DATETIME
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER PRIMARY KEY,
    platform TEXT NOT NULL,
    platform_account_id TEXT NOT NULL,
    name TEXT,
    profile_url TEXT,
    avatar_url TEXT,
    latest_update_time DATETIME,
    bound_user_id INTEGER,
    created_at DATETIME,
    updated_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_accounts_bound_user_id ON accounts(bound_user_id);
CREATE INDEX IF NOT EXISTS idx_account_platform_id ON accounts(platform, platform_account_id);

-- Devices table
CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY,
    device_name TEXT,
    system_type TEXT,
    system_version TEXT,
    client_name TEXT,
    client_source_code_address TEXT,
    client_version TEXT,
    client_local_id TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Sessions table
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    device_id INTEGER NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at DATETIME,
    expire_at DATETIME,
    updated_at DATETIME
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_sessions_refresh_token ON sessions(refresh_token);
CREATE INDEX IF NOT EXISTS idx_session_user_id_device_id ON sessions(user_id, device_id);

-- Tags table
CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY,
    user_tag INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    public NUMERIC,
    updated_at DATETIME,
    created_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_tags_user_tag ON tags(user_tag);

-- KVs table
CREATE TABLE IF NOT EXISTS kvs (
    bucket TEXT NOT NULL,
    key TEXT NOT NULL,
    value TEXT,
    updated_at DATETIME,
    created_at DATETIME,
    PRIMARY KEY (bucket, key)
);

-- Apps table
CREATE TABLE IF NOT EXISTS apps (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    version_number INTEGER,
    version_date DATETIME,
    creator_device_id INTEGER,
    app_sources TEXT,
    public NUMERIC,
    bound_store_app_id INTEGER,
    stop_store_manage NUMERIC,
    name TEXT NOT NULL,
    type TEXT,
    short_description TEXT,
    description TEXT,
    icon_image_url TEXT,
    icon_image_id INTEGER,
    background_image_url TEXT,
    background_image_id INTEGER,
    cover_image_url TEXT,
    cover_image_id INTEGER,
    release_date TEXT,
    developer TEXT,
    publisher TEXT,
    tags TEXT,
    alternative_names TEXT,
    updated_at DATETIME,
    created_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_apps_user_id ON apps(user_id);

-- App categories table
CREATE TABLE IF NOT EXISTS app_categories (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    version_number INTEGER,
    version_date DATETIME,
    name TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_app_categories_user_id ON app_categories(user_id);

-- App-AppCategories junction table
CREATE TABLE IF NOT EXISTS app_app_categories (
    app_id INTEGER NOT NULL,
    app_category_id INTEGER NOT NULL,
    PRIMARY KEY (app_id, app_category_id)
);

-- App infos table
CREATE TABLE IF NOT EXISTS app_infos (
    id INTEGER PRIMARY KEY,
    source TEXT NOT NULL,
    source_app_id TEXT NOT NULL,
    source_url TEXT,
    name TEXT,
    type TEXT,
    short_description TEXT,
    description TEXT,
    icon_image_url TEXT,
    icon_image_id INTEGER,
    background_image_url TEXT,
    background_image_id INTEGER,
    cover_image_url TEXT,
    cover_image_id INTEGER,
    release_date TEXT,
    developer TEXT,
    publisher TEXT,
    tags TEXT,
    alternative_names TEXT,
    raw_data TEXT,
    updated_at DATETIME,
    created_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_app_info_source_source_app_id ON app_infos(source, source_app_id);

-- App run times table
CREATE TABLE IF NOT EXISTS app_run_times (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    app_id INTEGER NOT NULL,
    device_id INTEGER,
    start_time DATETIME,
    duration INTEGER,
    created_at DATETIME,
    updated_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_app_run_times_user_id ON app_run_times(user_id);

-- Sentinels table
CREATE TABLE IF NOT EXISTS sentinels (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL,
    alternative_urls TEXT,
    get_token_path TEXT,
    download_file_base_path TEXT,
    creator_id INTEGER,
    library_report_sequence INTEGER,
    updated_at DATETIME,
    created_at DATETIME
);

-- Sentinel libraries table
CREATE TABLE IF NOT EXISTS sentinel_libraries (
    id INTEGER PRIMARY KEY,
    sentinel_id INTEGER NOT NULL,
    reported_id INTEGER,
    download_base_path TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Sentinel app binaries table
CREATE TABLE IF NOT EXISTS sentinel_app_binaries (
    id INTEGER PRIMARY KEY,
    union_id TEXT,
    sentinel_library_id INTEGER NOT NULL,
    generated_id TEXT,
    size_bytes INTEGER,
    need_token NUMERIC,
    name TEXT,
    version TEXT,
    developer TEXT,
    publisher TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Sentinel app binary files table
CREATE TABLE IF NOT EXISTS sentinel_app_binary_files (
    id INTEGER PRIMARY KEY,
    sentinel_app_binary_id INTEGER NOT NULL,
    name TEXT,
    size_bytes INTEGER,
    sha256 BLOB,
    server_file_path TEXT,
    chunks_info TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Sentinel sessions table
CREATE TABLE IF NOT EXISTS sentinel_sessions (
    id INTEGER PRIMARY KEY,
    sentinel_id INTEGER NOT NULL,
    refresh_token TEXT NOT NULL,
    status TEXT,
    creator_id INTEGER,
    expire_at DATETIME,
    last_used_at DATETIME,
    last_refreshed_at DATETIME,
    refresh_count INTEGER,
    created_at DATETIME,
    updated_at DATETIME
);

-- Store apps table
CREATE TABLE IF NOT EXISTS store_apps (
    id INTEGER PRIMARY KEY,
    source TEXT NOT NULL,
    source_app_id TEXT NOT NULL,
    name TEXT,
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Store app binaries table
CREATE TABLE IF NOT EXISTS store_app_binaries (
    id INTEGER PRIMARY KEY,
    app_id INTEGER NOT NULL,
    union_id TEXT,
    size_bytes INTEGER,
    need_token NUMERIC,
    name TEXT,
    version TEXT,
    developer TEXT,
    publisher TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- Feeds table
CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    link TEXT,
    authors TEXT,
    language TEXT,
    image TEXT,
    updated_at DATETIME,
    created_at DATETIME
);

-- Feed items table
CREATE TABLE IF NOT EXISTS feed_items (
    id INTEGER PRIMARY KEY,
    feed_id INTEGER NOT NULL,
    title TEXT,
    description TEXT,
    content TEXT,
    link TEXT,
    updated TEXT,
    updated_parsed DATETIME,
    published TEXT,
    published_parsed DATETIME,
    authors TEXT,
    guid TEXT,
    image TEXT,
    enclosures TEXT,
    publish_platform TEXT,
    read_count INTEGER,
    digest_description TEXT,
    digest_images TEXT,
    updated_at DATETIME,
    created_at DATETIME
);
