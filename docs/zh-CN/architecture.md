# 系统架构文档

## 1. 概述

donkey-admin-backend 是一个基于 Go 语言开发的企业级管理后台系统，采用分层架构设计，提供用户管理、组织管理、角色权限管理、菜单管理等核心功能。

## 2. 技术栈

### 2.1 核心技术
- **编程语言**: Go 1.19+
- **Web 框架**: Gin
- **ORM 框架**: GORM
- **数据库**: MySQL
- **缓存**: Redis (可选)
- **认证**: JWT (golang-jwt/jwt/v4)

### 2.2 主要依赖
- `github.com/gin-gonic/gin` - HTTP Web 框架
- `gorm.io/gorm` - ORM 框架
- `gorm.io/driver/mysql` - MySQL 驱动
- `github.com/go-redis/redis/v8` - Redis 客户端
- `github.com/golang-jwt/jwt/v4` - JWT 认证
- `github.com/spf13/viper` - 配置管理
- `gopkg.in/gomail.v2` - 邮件发送

## 3. 项目结构

```
donkey-admin-backend/
├── app/                    # 应用核心代码
│   ├── handler/           # 处理器层（Controller）
│   │   ├── admin/         # 管理后台接口处理器
│   │   └── api/           # 外部API接口处理器
│   ├── model/             # 数据模型层
│   └── service/           # 业务逻辑层
│       ├── auth/          # 认证服务
│       ├── acl/           # ACL权限服务
│       ├── rbac/          # RBAC权限服务
│       ├── iuser/         # 用户服务
│       ├── irole/         # 角色服务
│       ├── imenu/         # 菜单服务
│       ├── ioperation/    # 操作服务
│       ├── iorganization/ # 组织服务
│       ├── isubsystem/    # 子系统服务
│       └── tree/          # 树形结构服务
├── conf/                  # 配置管理
│   ├── config.go          # 配置结构定义
│   ├── config.toml        # 配置文件
│   └── flag_vars/         # 命令行参数
├── middleware/            # 中间件
│   ├── jwt-token.go       # JWT认证中间件
│   ├── base-auth.go       # 基础认证中间件
│   ├── cors.go            # CORS跨域中间件
│   └── middle_auth/       # 认证相关定义
├── router/                # 路由定义
│   ├── init.go            # 路由初始化
│   ├── admin.go           # 管理后台路由
│   └── api.go             # API路由
├── serverx/               # 服务器封装
│   ├── gin_server.go      # Gin服务器
│   └── static_server.go   # 静态文件服务器
├── lib/                   # 公共库
│   ├── libutils/          # 工具函数
│   └── libdto/            # 数据传输对象
├── req-resp/              # 请求响应处理
│   ├── appresp/           # 响应封装
│   └── gin_req/           # 请求处理
├── apperror/              # 错误处理
├── asset/                 # 静态资源
├── main.go                # 程序入口
└── go.mod                 # Go模块定义
```

## 4. 架构设计

### 4.1 分层架构

系统采用经典的分层架构，从上到下分为：

```
┌─────────────────────────────────┐
│      Handler Layer              │  处理HTTP请求，参数验证
│   (app/handler/admin|api)       │
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│      Service Layer              │  业务逻辑处理
│   (app/service/*)               │
│   - admin: 管理后台业务逻辑       │
│   - api: 外部API业务逻辑          │
│   - def: 业务定义和表单           │
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│      DML Layer                  │  数据访问层
│   (app/service/*/internal/dml)  │
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│      Model Layer                │  数据模型
│   (app/model)                   │
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│      Database                   │  MySQL数据库
└─────────────────────────────────┘
```

### 4.2 服务层设计模式

每个业务模块（如 user、role、menu）都遵循统一的目录结构：

```
service/iuser/
├── user.go                    # 服务主入口，提供基础CRUD
├── user_def/                  # 业务定义
│   └── user.go                # 查询表单、响应结构等
├── user_admin/                # 管理后台业务逻辑
│   └── user.go                # 管理后台特有的业务处理
├── user_api/                  # 外部API业务逻辑
│   └── user.go                # API特有的业务处理
└── internal/                  # 内部实现
    └── dml/                   # 数据访问层
        └── user.go            # 数据库操作
```

**设计优势**：
- 清晰的职责分离：基础服务、管理后台服务、API服务分离
- 易于扩展：新增业务逻辑不影响基础服务
- 统一接口：所有模块提供一致的CRUD接口

### 4.3 权限系统架构

系统采用 **ACL（访问控制列表）** 和 **RBAC（基于角色的访问控制）** 混合模式：

#### 4.3.1 权限模型

*注意，本系统采用ACL权限模型*

```
用户 (User)
  ├── 角色 (Role) ──┐
  │                 │
  └── 直接权限 (UserPermission) ──┐
                                   │
                            ┌──────▼──────┐
                            │   权限控制    │
                            └──────┬──────┘
                                   │
                    ┌──────────────┼──────────────┐
                    │              │              │
            ┌───────▼──────┐ ┌─────▼─────┐ ┌─────▼─────┐
            │   菜单权限    │ │  按钮权限   │ │  数据权限   │
            │   (Menu)     │ │ (Operation)│ │  (Data)   │
            └──────────────┘ └───────────┘ └───────────┘
```

#### 4.3.2 权限类型

1. **页面权限（Menu Permission）**
   - 基于菜单路由的访问控制
   - 控制用户可访问的页面

2. **按钮权限（Operation Permission）**
   - 基于操作的细粒度权限控制
   - 控制页面内的操作按钮（增删改查等）

3. **用户直接权限（User Permission）**
   - 用户可直接分配菜单和操作权限


#### 4.3.3 Root 用户

系统支持 Root 用户概念：
- Root 用户拥有所有权限
- 不受权限系统限制
- 用于系统初始化和超级管理

### 4.4 认证系统

#### 4.4.1 JWT 认证流程

```
1. 用户登录
   └─> 验证用户名密码
       └─> 生成JWT Token
           └─> 返回Token（Cookie + JSON）

2. 请求认证
   └─> 从Cookie/Header获取Token
       └─> 验证Token有效性
           └─> 提取用户ID
               └─> 设置到Context
```

#### 4.4.2 中间件链

```
请求 → CORS中间件 → JWT认证中间件 → 业务处理器
```

- **CORS中间件**: 处理跨域请求
- **JWT认证中间件**: 验证Token，提取用户信息
- **权限中间件**: （可选）验证用户权限

### 4.5 路由设计

#### 4.5.1 路由分组

- **管理后台路由** (`/admin-api/v1/*`)
  - 需要JWT认证
  - 用于前端管理界面

- **API路由** (`/api/*`)
  - 用于外部系统调用
  - 可配置不同的认证方式

#### 4.5.2 RESTful 设计

每个资源提供标准的RESTful接口：

```
GET    /resource/query      # 分页查询列表
GET    /resource/get        # 获取单个详情
POST   /resource/add        # 新增
POST   /resource/update     # 更新
POST   /resource/del        # 删除
POST   /resource/setInfo     # 批量设置信息
```

## 5. 核心模块

### 5.1 用户管理 (User)

- **功能**: 用户CRUD、状态管理、权限分配
- **模型**: `app/model/user.go`
- **服务**: `app/service/iuser/`
- **处理器**: `app/handler/admin/user.go`

### 5.2 菜单管理 (Menu)

- **功能**: 菜单CRUD、树形结构管理
- **模型**: `app/model/menu.go`
- **服务**: `app/service/imenu/`
- **特性**: 支持树形结构，动态菜单加载

### 5.3 操作管理 (Operation)

- **功能**: 操作按钮CRUD
- **模型**: `app/model/operation.go`
- **服务**: `app/service/ioperation/`
- **关联**: 与菜单关联，控制页面操作权限

### 5.4 组织管理 (Organization)

- **功能**: 组织架构CRUD、树形结构
- **模型**: `app/model/organization.go`
- **服务**: `app/service/iorganization/`
- **特性**: 支持组织树形结构

### 5.5 子系统管理 (Subsystem)

- **功能**: 多系统统一权限管理
- **模型**: `app/model/subsystem.go`
- **服务**: `app/service/isubsystem/`
- **用途**: 集中管理企业多系统的菜单、按钮和权限

### 5.6 权限管理 (Permission)

- **功能**: 用户权限配置、权限查询
- **模型**: `app/model/user_permission.go`
- **服务**: `app/service/iuser_permission/`
- **特性**: 支持菜单+操作的组合权限

## 6. 数据模型关系

### 6.1 核心实体关系

```
User (用户)
  └── UserPermission (用户权限)
      ├── Menu (菜单)
      └── Operation (操作)

Organization (组织)
  └── User (用户) [通过org_id关联]

Subsystem (子系统)
  ├── Menu (菜单) [通过system_id关联]
  └── Operation (操作) [通过system_id关联]
```

## 7. 配置管理

### 7.1 配置文件

- **主配置**: `conf/config.toml`
- **示例配置**: `conf/config.toml.example`
- **配置结构**: `conf/config.go`

### 7.2 配置项

- **MySQL配置**: 数据库连接信息
- **Redis配置**: 缓存配置（可选）
- **SMTP配置**: 邮件服务配置
- **JWT配置**: Token密钥和过期时间
- **服务器配置**: 监听地址和端口

## 8. 中间件

### 8.1 认证中间件

- **AdminToken()**: 管理后台JWT认证
- **InnerToken()**: 内部API认证
- **BaseAuth**: HTTP基础认证

### 8.2 功能中间件

- **CORS**: 跨域资源共享
- **BrowserCache**: 浏览器缓存控制
- **DumpRequest**: 请求日志记录

## 9. 工具库

### 9.1 libutils

- **JWT工具**: Token生成和验证
- **加密工具**: 密码加密、AES、RSA
- **字符串工具**: 字符串处理、驼峰转换
- **切片工具**: 切片操作工具
- **邮件工具**: 邮件发送

### 9.2 libdto

- **分页DTO**: 统一分页请求和响应
- **通用DTO**: 通用数据结构

## 10. 错误处理

### 10.1 错误定义

- **位置**: `apperror/errors.go`
- **统一错误码**: 定义系统错误码
- **错误封装**: `apperror/apperror.go`

### 10.2 响应格式

统一响应格式：
```json
{
  "code": 0,        // 0表示成功，非0表示失败
  "msg": "success", // 提示信息
  "data": {}        // 数据内容
}
```

## 11. 部署架构

### 11.1 服务启动

1. **初始化配置**: 加载配置文件
2. **初始化数据库**: 连接MySQL
3. **初始化Redis**: （可选）连接Redis
4. **初始化路由**: 注册路由和中间件
5. **启动HTTP服务**: 监听指定端口

### 11.2 静态资源

- **前端资源**: `asset/dist/` - Vue前端打包文件
- **静态文件**: `asset/static/` - 静态资源
- **支持方式**: 本地文件或嵌入到二进制文件

## 12. 扩展性设计

### 12.1 多系统支持

通过 `Subsystem` 模块实现：
- 统一管理多个子系统的权限
- 每个子系统有独立的菜单和操作
- 用户可配置不同子系统的权限

### 12.2 服务扩展

- **新增业务模块**: 按照现有模式创建service目录
- **新增API接口**: 在handler层添加处理器
- **新增中间件**: 在middleware目录添加

### 12.3 数据访问扩展

- **DML层统一接口**: 数据LUR缓存
- **DAL层统一接口**: 封装GORM, 所有模块提供一致的CRUD方法
- **支持原生SQL**: 提供Exec、RawGet、RawFind方法
- **批量操作**: 支持批量插入和更新

## 13. 安全设计

### 13.1 认证安全

- **JWT Token**: 无状态认证，支持分布式
- **Token过期**: 可配置Token过期时间
- **密码加密**: 使用加密算法存储密码

### 13.2 权限安全

- **细粒度控制**: 菜单级和按钮级权限
- **Root用户隔离**: Root用户权限独立管理
- **权限验证**: 中间件层统一验证

### 13.3 数据安全

- **参数验证**: Handler层参数校验
- **SQL注入防护**: 使用GORM参数化查询
- **CORS配置**: 控制跨域访问

## 14. 性能优化

### 14.1 数据库优化

- **连接池**: GORM自动管理连接池
- **索引优化**: 关键字段建立索引
- **查询优化**: 使用分页减少数据量

### 14.2 缓存策略

- **LRU支持**: 在DML层内存缓存
- **Redis支持**: 可选的Redis缓存
- **浏览器缓存**: 静态资源缓存中间件


## 15. 开发规范

### 15.1 代码组织

- **包命名**: 使用有意义的包名
- **文件命名**: 与功能对应
- **目录结构**: 遵循统一的分层结构

### 15.2 接口设计

- **RESTful**: 遵循RESTful设计原则
- **统一响应**: 使用统一的响应格式
- **错误处理**: 统一的错误码和错误信息

### 15.3 注释规范

- **包注释**: 每个包有功能说明
- **函数注释**: 关键函数有注释说明
- **结构体注释**: 重要结构体有字段说明

## 16. 总结

donkey-admin-backend 采用清晰的分层架构，通过 Handler-Service-DML-Model 四层设计实现了良好的代码组织和职责分离。权限系统采用 ACL + RBAC 混合模式，支持灵活的权限配置。系统设计考虑了扩展性、安全性和性能，适合作为企业级管理后台的基础框架。

