# AGENTS.md

本文档为 AI 代理提供项目上下文和编码指南。

## 项目概述

Librarian 是 TuiHub 的标准服务端实现，使用 Go 语言编写。这是一个经过长期开发的成熟项目，遵循清洁架构原则。

- **语言**: Go 1.25+
- **框架**: [Kratos v2](https://github.com/go-kratos/kratos) (微服务框架)
- **依赖注入**: [Wire](https://github.com/google/wire)
- **ORM**: [Ent](https://entgo.io/)
- **RPC**: [Connect-RPC](https://connectrpc.com/) (gRPC-Web 兼容)
- **消息队列**: [Watermill](https://github.com/ThreeDotsLabs/watermill)
- **配置格式**: TOML

## 目录结构

```
.
├── cmd/                    # 应用入口和命令行定义
├── configs/                # 配置文件模板
├── internal/               # 私有应用代码
│   ├── biz/               # 业务逻辑层
│   │   ├── bizangela/     # Angela 模块 (服务器实例管理)
│   │   ├── bizbinah/      # Binah 模块 (文件存储)
│   │   ├── bizchesed/     # Chesed 模块 (图片管理)
│   │   ├── bizgebura/     # Gebura 模块 (应用/游戏管理)
│   │   ├── bizkether/     # Kether 模块 (后台任务)
│   │   ├── biznetzach/    # Netzach 模块 (通知系统)
│   │   ├── bizsupervisor/ # Supervisor 模块 (功能特性管理)
│   │   ├── biztiphereth/  # Tiphereth 模块 (用户/账户管理)
│   │   └── bizyesod/      # Yesod 模块 (RSS/Feed 管理)
│   ├── client/            # 外部服务客户端
│   ├── conf/              # 配置结构定义
│   ├── data/              # 数据访问层
│   │   └── internal/ent/  # Ent 生成的代码
│   ├── lib/               # 内部工具库
│   │   ├── libapp/        # 应用设置
│   │   ├── libauth/       # 认证/授权
│   │   ├── libcache/      # 缓存抽象
│   │   ├── libcron/       # 定时任务
│   │   ├── libidgenerator/# ID 生成器
│   │   ├── libmq/         # 消息队列抽象
│   │   ├── libobserve/    # 可观测性 (OpenTelemetry)
│   │   ├── libs3/         # S3 存储抽象
│   │   ├── libsearch/     # 搜索抽象 (Bleve/Meilisearch)
│   │   └── libzap/        # 日志配置
│   ├── model/             # 领域模型定义
│   ├── server/            # 服务器配置 (gRPC, Connect)
│   └── service/           # 服务层 (gRPC 处理器)
│       ├── angelaweb/     # Web 管理界面
│       ├── porter/        # Porter 插件服务
│       ├── sentinel/      # Sentinel 客户端服务
│       ├── sephirah/      # 主 API 服务
│       └── supervisor/    # Supervisor 服务
├── pkg/                   # 公共库/插件
│   ├── tuihub-go/         # Go 客户端/插件助手库
│   ├── tuihub-rss/        # RSS 插件
│   ├── tuihub-steam/      # Steam 插件
│   └── tuihub-telegram/   # Telegram 插件
├── tests/                 # 集成测试
├── main.go                # 主入口
├── Makefile               # 构建脚本
└── go.mod                 # Go 模块定义
```

## 架构分层

项目遵循清洁架构原则，分层如下:

```
cmd → service → biz → data
```

- **cmd**: 应用入口，处理命令行参数和启动流程
- **service**: gRPC/Connect 服务处理器，负责请求/响应转换
- **biz**: 业务逻辑实现，不依赖具体框架
- **data**: 数据访问层，实现仓储接口

### Wire 依赖注入

项目使用 Wire 进行依赖注入。每个包定义 `ProviderSet`:

```go
var ProviderSet = wire.NewSet(
    NewSomeService,
    NewAnotherService,
)
```

Wire 文件使用 `wireinject` build tag:

```go
//go:build wireinject
```

生成命令: `go generate ./...`

## 模块命名约定

项目使用希伯来卡巴拉生命之树 (Sephirot) 术语命名核心模块:

| 模块 | 职责 |
|------|------|
| Sephirah | 主 API 服务入口 |
| Kether | 后台任务处理 |
| Tiphereth | 用户和账户管理 |
| Gebura | 应用/游戏管理 |
| Binah | 文件存储 |
| Chesed | 图片管理 |
| Netzach | 通知系统 |
| Yesod | RSS/Feed 管理 |
| Angela | 服务器实例管理 |

## 开发工作流

### 依赖安装

```bash
# 初始化开发环境
make init

# 仅安装 lint 工具
make init-lint

# 安装测试工具
make init-test
```

### 代码生成

```bash
# 生成 Wire 和 Goverter 代码
make generate
```

### 构建

```bash
# 调试模式构建
make build

# 运行服务
make run
```

### Lint

```bash
# 运行 golangci-lint 并自动修复
make lint
```

### 测试

```bash
# 运行单元测试
make test-unit

# 运行集成测试 (需要 Docker)
make test-goc

# 运行全部测试
make test-all
```

## 代码规范

### Lint 配置

项目使用 golangci-lint v2 (`.golangci.yml`)，启用以下关键规则:

- **exhaustruct**: 检查结构体字段完整性
- **gocognit**: 认知复杂度检查 (最大 20)
- **funlen**: 函数长度限制 (100 行, 50 语句)
- **cyclop**: 圈复杂度限制 (最大 30)
- **golines**: 行长度限制 (120 字符)
- **gci**: import 分组排序

### Import 排序

import 语句按以下顺序分组:

1. 标准库
2. `github.com/tuihub` 前缀包
3. 第三方包
4. 空行分隔
5. 点导入

### 禁止使用的包

- `github.com/golang/protobuf` → 使用 `google.golang.org/protobuf`
- `github.com/satori/go.uuid` → 使用 `github.com/google/uuid`
- `math/rand` (非测试文件) → 使用 `math/rand/v2`
- `log` (非 main 文件) → 使用 `log/slog`

### Porter 插件规则

`pkg/tuihub-*` 目录下的插件代码:
- 禁止导入 `github.com/tuihub/librarian/internal`
- 禁止导入 `github.com/tuihub/tuihub-go`

## 配置

### 配置文件格式

使用 TOML 格式，支持多种驱动:

```toml
[database]
driver = "postgres"  # memory, sqlite3, postgres

[storage]
driver = "s3"  # memory, file, s3

[mq]
driver = "redis"  # memory, sql, redis

[cache]
driver = "redis"  # memory, redis

[search]
driver = "meili"  # bleve, meili
```

### 环境变量

| 变量 | 说明 |
|------|------|
| `LOG_LEVEL` | 日志级别 (debug, info, warn, error) |
| `DEMO_MODE` | 演示模式 (true/false) |
| `CREATE_ADMIN_USER` | 启动时创建管理员用户名 |
| `CREATE_ADMIN_PASS` | 启动时创建管理员密码 |

## 数据库 (Ent)

### Schema 位置

Ent schema 定义在 `internal/data/internal/ent/schema/`。

### 生成代码

修改 schema 后运行:

```bash
go generate ./internal/data/internal/ent/...
```

## 消息队列

使用 Watermill 抽象消息队列，支持:
- 内存 (开发)
- SQL (SQLite/PostgreSQL)
- Redis Stream

Topic 定义在 `internal/lib/libmq/`。

## API 协议

使用 [TuiHub Protos](https://github.com/tuihub/protos) 定义的 protobuf 协议。

服务实现 Connect-RPC 接口，同时支持:
- gRPC (原生)
- gRPC-Web
- Connect 协议

## 测试

### 单元测试

使用 `testify` 断言库，测试文件使用 `_test.go` 后缀。

### 集成测试

位于 `tests/` 目录，使用 Docker Compose 部署依赖:

```bash
cd tests
docker-compose up -d
make all
```

## 发布

使用 GoReleaser 进行多平台构建和发布:

```bash
# 干运行
make release-dry-run

# 正式发布 (CI)
make release
```

支持平台: Linux/macOS/Windows (amd64/arm64)

## 常见任务

### 添加新的业务模块

1. 在 `internal/biz/` 创建 `biznewmodule/` 目录
2. 定义 ProviderSet 和业务逻辑
3. 在 `internal/biz/biz.go` 注册 ProviderSet
4. 在 `internal/data/` 创建对应的 Repository

### 添加新的 Ent Schema

1. 在 `internal/data/internal/ent/schema/` 创建 schema 文件
2. 运行 `go generate ./internal/data/internal/ent/...`
3. 在 Repository 中使用生成的客户端

### 添加新的 API 端点

1. 确保 proto 已在 tuihub/protos 更新
2. 更新 `go.mod` 中的 protos 版本
3. 在 `internal/service/sephirah/` 实现 handler
4. 在 converter 中添加类型转换

### 添加新的 Porter 插件

1. 在 `pkg/` 创建 `tuihub-newplugin/` 目录
2. 实现 `porter.LibrarianPorterServiceHandler`
3. 参考 `pkg/tuihub-go/README.md` 示例
