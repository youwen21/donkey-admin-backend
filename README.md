# donkey-admin-backend

> 基于 Go + Gin + GORM 开发的企业级管理后台系统，提供完整的用户权限管理、组织架构管理、菜单权限控制等功能。

**构建的二进制可执行文件包含config配置文件，勿要对外传播**

## 简介

donkey-admin-backend 是一套功能完善的后台服务接口系统，采用分层架构设计，支持灵活的权限配置和多系统统一管理。系统提供 RESTful API 接口，可配合 donkey-admin-front 前端项目使用，也支持自定义前端开发。

**主要特性**：
- 🔐 完整的用户认证与权限管理系统（ACL + RBAC）
- 🏢 组织架构管理，支持树形结构
- 📋 菜单与操作权限的细粒度控制
- 🔄 多子系统统一权限管理
- 🚀 基于 Gin 框架的高性能 API 服务
- 📦 清晰的分层架构，易于扩展和维护

## 技术栈

- **语言**: Go 1.19+
- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **数据库**: MySQL 5.7+
- **认证**: JWT

## 快速开始

### 环境要求

- Go 1.19 或更高版本
- MySQL 5.7 或更高版本

### 安装步骤

1. **克隆项目**
```bash
git clone git@github.com:youwen21/donkey-admin-backend.git
cd donkey-admin-backend
```

2. **安装依赖**
```bash
go mod download
```

3. **初始化数据库**
```bash
mysql -u root -p < docs/sql/init.sql
```

4. **配置数据库**
```bash
cp conf/config.toml.example conf/config.toml
# 编辑 conf/config.toml 配置数据库连接信息
```

5. **编译运行**
```bash
go build -o bin/donkey-admin
./bin/donkey-admin
```

6. **服务**
- API 地址: http://localhost:8000
- 默认管理员: `admin` / `admin123`

> 📖 详细的安装和配置说明请参考 [快速开始指南](./docs/zh-CN/guide.md)

## 功能模块

### 用户认证
- 用户登录/登出

### 组织管理
- 组织架构分页查询
- 组织新增/编辑/删除

### 用户管理
- 用户列表分页查询
- 用户新增/编辑/删除
- 用户状态管理（启用/禁用）
- 用户授权管理（ACL），支持菜单和按钮级权限控制

### 菜单管理
- 菜单列表分页查询
- 菜单新增/编辑/删除
- 动态菜单加载，根据用户权限显示可见菜单

### 操作管理
- 操作按钮列表分页查询
- 操作新增/编辑/删除
- 与菜单关联，控制页面操作权限

### 子系统管理
集中统一管理企业多系统的菜单和按钮权限，其他系统只关系业务，基础的员工、菜单、组织等数据，API请求本系统
- 子系统分页查询
- 子系统新增/编辑/删除

### ACL权限
- **页面权限**：基于菜单路由的访问控制
- **按钮权限**：基于操作的细粒度权限控制
- **动态权限**：从后台获取权限数据，实时更新

### 用户通知
- TODO


## 项目结构

```
donkey-admin-backend/
├── app/                    # 应用核心代码
│   ├── handler/           # 处理器层（Controller）
│   ├── model/             # 数据模型层
│   └── service/           # 业务逻辑层
├── conf/                  # 配置管理
├── middleware/            # 中间件
├── router/                # 路由定义
├── lib/                   # 公共库
├── docs/                  # 文档
│   ├── sql/              # 数据库初始化脚本
│   └── zh-CN/            # 中文文档
└── main.go               # 程序入口
```

## 文档

详细的开发文档请参考 [docs](./docs/) 目录：

- 📘 [快速开始指南](./docs/zh-CN/guide.md) - 安装、配置、启动指南
- 🏗️ [架构设计](./docs/zh-CN/architecture.md) - 系统架构说明
- 🔑 [用户权限文档](./docs/zh-CN/user-permission/permission.md) - 权限管理详细说明

## 许可证

[LICENSE](./LICENSE)

## 贡献

欢迎提交 Issue 和 Pull Request！

---

**注意**: 本项目可配合 [donkey-admin-front](https://github.com/youwen21/donkey-admin-front) 前端项目使用，也支持自定义前端开发。
