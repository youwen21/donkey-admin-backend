# 快速开始指南

本文档将帮助您快速下载、安装和启动 donkey-admin-backend 管理后台系统。

## 目录

- [环境要求](#环境要求)
- [下载安装](#下载安装)
- [数据库初始化](#数据库初始化)
- [配置文件](#配置文件)
- [启动服务](#启动服务)
- [验证安装](#验证安装)

## 环境要求

在开始之前，请确保您的系统已安装以下软件：

### 必需软件

- **Go**: 1.19 或更高版本
  - 下载地址: https://golang.org/dl/
  - 验证安装: `go version`

- **MySQL**: 5.7 或更高版本（推荐 8.0+）
  - 下载地址: https://dev.mysql.com/downloads/mysql/
  - 验证安装: `mysql --version`

## 下载安装

### 方式一：从源码编译（推荐）

1. **克隆代码仓库**

```bash
// git clone <repository-url>
git clone git@github.com:youwen21/donkey-admin-backend.git
cd donkey-admin-backend
```

2. **安装依赖**

```bash
go mod download
```

3. **编译项目**

```bash
# 直接使用 go build
go build -o bin/donkey-admin
```

编译成功后，可执行文件位于 `bin/` 目录下。


## 数据库初始化

### 步骤 1: 创建数据库

登录 MySQL，创建数据库：

```bash
mysql -u root -p
```

在 MySQL 命令行中执行：

```sql
CREATE DATABASE IF NOT EXISTS `donkey_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 步骤 2: 导入数据库结构

有两种方式初始化数据库：

#### 方式一：使用 SQL 脚本（推荐）

1. **执行初始化脚本**

```bash
mysql -u root -p donkey_admin < docs/sql/init.sql
```

或者登录 MySQL 后执行：

```bash
mysql -u root -p
```

```sql
USE donkey_admin;
SOURCE docs/sql/init.sql;
```

### 步骤 3: 验证数据库

检查表是否创建成功：

```bash
mysql -u root -p donkey_admin -e "SHOW TABLES;"
```

应该看到以下表：
- `t_subsystem` - 子系统表
- `t_organization` - 组织表
- `t_user` - 用户表
- `t_role` - 角色表
- `t_user_role` - 用户角色关联表
- `t_menu` - 菜单表
- `t_operation` - 操作表
- `t_role_menu` - 角色菜单关联表
- `t_user_permission` - 用户权限表

### 默认数据说明

初始化脚本会自动创建以下默认数据：

- **默认管理员用户**
  - 用户名: `admin`
  - 密码: `admin123`
  - 类型: Root 用户（拥有所有权限）

- **默认子系统**
  - 名称: 管理后台
  - 系统标识: `admin`

- **默认组织**
  - 名称: 总公司

- **默认菜单和操作**
  - 系统管理菜单及其子菜单
  - 基本的增删改查操作

> **安全提示**: 首次登录后，请立即修改默认管理员密码！

## 配置文件

### 步骤 1: 复制配置文件

```bash
cp conf/config.toml.example conf/config.toml
```

### 步骤 2: 编辑配置文件

编辑 `conf/config.toml`，配置数据库连接信息：

```toml
[mysql_default]
host = "127.0.0.1"                # MySQL 地址
username = "root"
password = "your_mysql_password"  # 修改为您的MySQL密码
port = "3306"
database = "donkey_admin"          # 数据库名称
charset = "utf8mb4"

[redis_default]
host = "127.0.0.1"
port = "6379"                      # Redis端口，如果未安装Redis可留空
password = ""                       # Redis密码，如果没有密码可留空
db = 1

[smtp]
host = ""                          # SMTP服务器（可选）
port = 465
user = ""
password = ""
```

### 配置说明

- **mysql_default**: MySQL 数据库连接配置（必需）
- **redis_default**: Redis 缓存配置（可选，如果未安装 Redis 可以不配置）
- **smtp**: 邮件服务配置（可选，用于发送邮件通知）

## 启动服务

### 方式一：直接运行（开发环境）

```bash
# 使用默认配置
./bin/donkey-admin

# 或指定端口
./bin/donkey-admin -port=8000

# 或指定配置文件路径
./bin/donkey-admin -runmode=dev
```

### 启动参数

- `-host`: 监听地址（默认: 0.0.0.0）
- `-port`: 监听端口（默认: 8000）
- `-runmode`: 运行模式（dev/prod）
- `-logfile`: 日志文件路径
- `-loglevel`: 日志级别

示例：

```bash
./bin/donkey-admin -host=0.0.0.0 -port=8000 -runmode=dev
```

## 验证安装

### 1. 检查服务是否启动

访问健康检查接口：

```bash
curl http://localhost:8000/ping
```

应该返回：`pong`

### 2. 检查 API 登录接口

测试登录接口（需要先获取 Token）：

```bash
curl -X POST http://localhost:8000/admin-api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```


## 下一步

安装成功后，您可以：

1. **修改默认密码**: 登录后立即修改管理员密码
2. **配置组织架构**: 根据实际情况创建组织架构
3. **创建用户和角色**: 添加其他用户并分配角色
4. **配置菜单权限**: 设置菜单和操作权限
5. **阅读架构文档**: 了解系统架构设计，参考 [架构文档](./architecture.md)

---

**祝您使用愉快！**

