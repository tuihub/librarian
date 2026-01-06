-- V1 PostgreSQL Fresh Install Schema
-- This creates all tables from scratch for a new installation

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    creator_id BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id BIGINT PRIMARY KEY,
    platform VARCHAR(100) NOT NULL,
    platform_account_id VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    profile_url TEXT,
    avatar_url TEXT,
    latest_update_time TIMESTAMP,
    bound_user_id BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_accounts_bound_user_id ON accounts(bound_user_id);
CREATE INDEX IF NOT EXISTS idx_account_platform_id ON accounts(platform, platform_account_id);

-- Devices table
CREATE TABLE IF NOT EXISTS devices (
    id BIGINT PRIMARY KEY,
    device_name VARCHAR(255),
    system_type VARCHAR(100),
    system_version VARCHAR(100),
    client_name VARCHAR(255),
    client_source_code_address TEXT,
    client_version VARCHAR(100),
    client_local_id VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Sessions table
CREATE TABLE IF NOT EXISTS sessions (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    device_id BIGINT NOT NULL,
    refresh_token VARCHAR(500) NOT NULL,
    created_at TIMESTAMP,
    expire_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_sessions_refresh_token ON sessions(refresh_token);
CREATE INDEX IF NOT EXISTS idx_session_user_id_device_id ON sessions(user_id, device_id);

-- Tags table
CREATE TABLE IF NOT EXISTS tags (
    id BIGINT PRIMARY KEY,
    user_tag BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    public BOOLEAN,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_tags_user_tag ON tags(user_tag);

-- KVs table
CREATE TABLE IF NOT EXISTS kvs (
    bucket VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    value TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    PRIMARY KEY (bucket, key)
);

-- Apps table
CREATE TABLE IF NOT EXISTS apps (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    version_number INTEGER,
    version_date TIMESTAMP,
    creator_device_id BIGINT,
    app_sources TEXT,
    public BOOLEAN,
    bound_store_app_id BIGINT,
    stop_store_manage BOOLEAN,
    name VARCHAR(500) NOT NULL,
    type VARCHAR(50),
    short_description TEXT,
    description TEXT,
    icon_image_url TEXT,
    icon_image_id BIGINT,
    background_image_url TEXT,
    background_image_id BIGINT,
    cover_image_url TEXT,
    cover_image_id BIGINT,
    release_date VARCHAR(50),
    developer VARCHAR(255),
    publisher VARCHAR(255),
    tags TEXT,
    alternative_names TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_apps_user_id ON apps(user_id);

-- App categories table
CREATE TABLE IF NOT EXISTS app_categories (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    version_number INTEGER,
    version_date TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_app_categories_user_id ON app_categories(user_id);

-- App-AppCategories junction table
CREATE TABLE IF NOT EXISTS app_app_categories (
    app_id BIGINT NOT NULL,
    app_category_id BIGINT NOT NULL,
    PRIMARY KEY (app_id, app_category_id)
);

-- App infos table
CREATE TABLE IF NOT EXISTS app_infos (
    id BIGINT PRIMARY KEY,
    source VARCHAR(100) NOT NULL,
    source_app_id VARCHAR(255) NOT NULL,
    source_url TEXT,
    name VARCHAR(500),
    type VARCHAR(50),
    short_description TEXT,
    description TEXT,
    icon_image_url TEXT,
    icon_image_id BIGINT,
    background_image_url TEXT,
    background_image_id BIGINT,
    cover_image_url TEXT,
    cover_image_id BIGINT,
    release_date VARCHAR(50),
    developer VARCHAR(255),
    publisher VARCHAR(255),
    tags TEXT,
    alternative_names TEXT,
    raw_data TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_app_info_source_source_app_id ON app_infos(source, source_app_id);

-- App run times table
CREATE TABLE IF NOT EXISTS app_run_times (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    app_id BIGINT NOT NULL,
    device_id BIGINT,
    start_time TIMESTAMP,
    duration BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_app_run_times_user_id ON app_run_times(user_id);

-- Sentinels table
CREATE TABLE IF NOT EXISTS sentinels (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    url TEXT NOT NULL,
    alternative_urls TEXT,
    get_token_path TEXT,
    download_file_base_path TEXT,
    creator_id BIGINT,
    library_report_sequence INTEGER,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);

-- Sentinel libraries table
CREATE TABLE IF NOT EXISTS sentinel_libraries (
    id BIGINT PRIMARY KEY,
    sentinel_id BIGINT NOT NULL,
    reported_id INTEGER,
    download_base_path TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Sentinel app binaries table
CREATE TABLE IF NOT EXISTS sentinel_app_binaries (
    id BIGINT PRIMARY KEY,
    union_id VARCHAR(255),
    sentinel_library_id BIGINT NOT NULL,
    generated_id VARCHAR(255),
    size_bytes BIGINT,
    need_token BOOLEAN,
    name VARCHAR(500),
    version VARCHAR(100),
    developer VARCHAR(255),
    publisher VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Sentinel app binary files table
CREATE TABLE IF NOT EXISTS sentinel_app_binary_files (
    id BIGINT PRIMARY KEY,
    sentinel_app_binary_id BIGINT NOT NULL,
    name VARCHAR(500),
    size_bytes BIGINT,
    sha256 BYTEA,
    server_file_path TEXT,
    chunks_info TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Sentinel sessions table
CREATE TABLE IF NOT EXISTS sentinel_sessions (
    id BIGINT PRIMARY KEY,
    sentinel_id BIGINT NOT NULL,
    refresh_token VARCHAR(500) NOT NULL,
    status VARCHAR(50),
    creator_id BIGINT,
    expire_at TIMESTAMP,
    last_used_at TIMESTAMP,
    last_refreshed_at TIMESTAMP,
    refresh_count INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Store apps table
CREATE TABLE IF NOT EXISTS store_apps (
    id BIGINT PRIMARY KEY,
    source VARCHAR(100) NOT NULL,
    source_app_id VARCHAR(255) NOT NULL,
    name VARCHAR(500),
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Store app binaries table
CREATE TABLE IF NOT EXISTS store_app_binaries (
    id BIGINT PRIMARY KEY,
    app_id BIGINT NOT NULL,
    union_id VARCHAR(255),
    size_bytes BIGINT,
    need_token BOOLEAN,
    name VARCHAR(500),
    version VARCHAR(100),
    developer VARCHAR(255),
    publisher VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Feeds table
CREATE TABLE IF NOT EXISTS feeds (
    id BIGINT PRIMARY KEY,
    title TEXT,
    description TEXT,
    link TEXT,
    authors TEXT,
    language VARCHAR(50),
    image TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);

-- Feed items table
CREATE TABLE IF NOT EXISTS feed_items (
    id BIGINT PRIMARY KEY,
    feed_id BIGINT NOT NULL,
    title TEXT,
    description TEXT,
    content TEXT,
    link TEXT,
    updated TEXT,
    updated_parsed TIMESTAMP,
    published TEXT,
    published_parsed TIMESTAMP,
    authors TEXT,
    guid VARCHAR(500),
    image TEXT,
    enclosures TEXT,
    publish_platform VARCHAR(100),
    read_count INTEGER,
    digest_description TEXT,
    digest_images TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
);
