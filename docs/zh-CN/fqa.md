
## 常见问题

### 1. 数据库连接失败

**错误信息**: `failed to connect database`

**解决方案**:
- 检查 MySQL 服务是否启动: `systemctl status mysql` 或 `service mysql status`
- 检查配置文件中的数据库连接信息是否正确
- 确认数据库用户有足够的权限
- 检查防火墙是否允许连接

### 2. 端口被占用

**错误信息**: `bind: address already in use`

**解决方案**:
- 检查端口占用: `lsof -i :8000` 或 `netstat -tulpn | grep 8000`
- 修改配置文件中的端口号
- 或使用启动参数指定其他端口: `-port=8001`

### 3. 找不到配置文件

**错误信息**: `Fatal error config file`

**解决方案**:
- 确认 `conf/config.toml` 文件存在
- 检查文件路径是否正确
- 可以从 `conf/config.toml.example` 复制并修改

### 4. 数据库表不存在

**错误信息**: `Table 'donkey_admin.xxx' doesn't exist`

**解决方案**:
- 重新执行数据库初始化脚本: `mysql -u root -p donkey_admin < docs/sql/init.sql`
- 检查数据库名称是否正确