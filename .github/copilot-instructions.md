# GitHub Copilot Instructions

本文件为 GitHub Copilot 提供项目特定的编码指导。

## 项目基础

这是 Librarian - TuiHub 的标准服务端实现。使用 Go 1.25+，基于 Kratos 框架。

## 代码风格

### Go 代码规范

- 使用 `golangci-lint` 进行代码检查
- 最大行长度: 120 字符
- 函数最大长度: 100 行, 50 语句
- 认知复杂度限制: 20

### Import 语句格式

按以下顺序组织 import:

```go
import (
    // 标准库
    "context"
    "fmt"

    // tuihub 内部包
    "github.com/tuihub/librarian/internal/model"

    // 第三方包
    "github.com/google/wire"
)
```

### 错误处理

- 总是检查错误返回值
- 使用 `errors.New()` 或 `fmt.Errorf()` 创建错误
- 避免使用 panic 处理业务错误

### 结构体初始化

项目启用了 `exhaustruct` lint 规则，需要显式初始化所有结构体字段:

```go
// 正确
user := &model.User{
    ID:       0,
    Username: "",
    Password: "",
    Type:     model.UserTypeNormal,
    Status:   model.UserStatusActive,
}

// 错误 - 会被 lint 拒绝
user := &model.User{
    Username: "test",
}
```

## 架构约定

### 分层架构

遵循 `cmd → service → biz → data` 的调用顺序。

- **service 层**: 仅处理 protobuf 类型转换，不包含业务逻辑
- **biz 层**: 实现业务逻辑，定义领域模型
- **data 层**: 数据库操作，实现 Repository 接口

### 依赖注入 (Wire)

每个包需要定义 ProviderSet:

```go
var ProviderSet = wire.NewSet(
    NewSomeService,
    NewSomeRepository,
)
```

修改依赖后运行: `make generate`

### 类型转换

使用 Goverter 生成类型转换代码:

- 转换器定义在 `internal/service/sephirah/converter/`
- PB → Biz: `pb_to_biz.go`
- Biz → PB: `biz_to_pb.go`

## 常用模式

### 创建新的业务方法

```go
func (t *Tiphereth) DoSomething(ctx context.Context, req *model.SomeRequest) error {
    // 1. 参数验证
    if req.ID == 0 {
        return errors.New("invalid ID")
    }

    // 2. 业务逻辑
    result, err := t.repo.GetSomething(ctx, req.ID)
    if err != nil {
        return err
    }

    // 3. 返回结果
    return nil
}
```

### 添加新的消息队列 Topic

```go
func NewSomeTopic(
    k *KetherBase,
) *libmq.Topic[model.SomeMessage] {
    return libmq.NewTopic[model.SomeMessage](
        "some-topic",
        func(ctx context.Context, msg model.SomeMessage) error {
            // 处理消息
            return nil
        },
    )
}
```

### 添加新的缓存

```go
func NewSomeCache(
    repo *data.SomeRepo,
    store libcache.Store,
) *libcache.Key[model.SomeData] {
    return libcache.NewKey[model.SomeData](
        store,
        "SomeCache",
        func(ctx context.Context) (*model.SomeData, error) {
            return repo.GetData(ctx)
        },
        libcache.WithExpiration(libtime.OneDay),
    )
}
```

## 禁止事项

### 禁止导入的包

```go
// ❌ 禁止
import "github.com/golang/protobuf/..."

// ✅ 使用
import "google.golang.org/protobuf/..."
```

```go
// ❌ 禁止 (非测试文件)
import "math/rand"

// ✅ 使用
import "math/rand/v2"
```

```go
// ❌ 禁止 (非 main.go)
import "log"

// ✅ 使用
import "log/slog"
// 或项目内部 logger
import "github.com/tuihub/librarian/internal/lib/logger"
```

### Porter 插件限制

`pkg/tuihub-*` 目录下的代码禁止导入:
- `github.com/tuihub/librarian/internal`
- `github.com/tuihub/tuihub-go`

## 数据库操作 (Ent)

### 查询模式

```go
// 单条查询
user, err := client.User.
    Query().
    Where(user.ID(id)).
    Only(ctx)

// 列表查询
users, err := client.User.
    Query().
    Where(user.StatusEQ(model.UserStatusActive)).
    Limit(10).
    Offset(0).
    All(ctx)

// 创建
user, err := client.User.
    Create().
    SetUsername(username).
    SetPassword(password).
    Save(ctx)

// 更新
user, err := client.User.
    UpdateOneID(id).
    SetUsername(newUsername).
    Save(ctx)
```

### 事务

```go
tx, err := client.Tx(ctx)
if err != nil {
    return err
}
defer func() {
    if v := recover(); v != nil {
        tx.Rollback()
        panic(v)
    }
}()

// 执行操作
if err := tx.Commit(); err != nil {
    return err
}
```

## 测试

### 单元测试

```go
func TestSomething(t *testing.T) {
    // Arrange
    input := &model.SomeInput{}

    // Act
    result, err := DoSomething(input)

    // Assert
    require.NoError(t, err)
    assert.Equal(t, expected, result)
}
```

使用 `testify` 包进行断言。

## 命令行

### 常用命令

```bash
# 构建
make build

# 运行
make run

# Lint
make lint

# 单元测试
make test-unit

# 代码生成
make generate
```

## 模块名称参考

| 模块 | 职责 |
|------|------|
| Tiphereth | 用户/账户管理 |
| Gebura | 应用/游戏管理 |
| Binah | 文件存储 |
| Chesed | 图片管理 |
| Netzach | 通知系统 |
| Yesod | RSS/Feed |
| Kether | 后台任务 |
| Angela | 服务器管理 |
