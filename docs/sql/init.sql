-- ============================================
-- donkey-admin-backend 数据库初始化脚本
-- ============================================

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `donkey_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `donkey_admin`;

-- ============================================
-- 1. 子系统表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_subsystem` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '子系统名称',
  `domain` varchar(255) NOT NULL DEFAULT '' COMMENT '域名',
  `syskey` varchar(100) NOT NULL DEFAULT '' COMMENT '系统标识key',
  `secret` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-有效，2-禁用',
  `order_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序号',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_syskey` (`syskey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='子系统表';

-- 插入默认子系统
INSERT INTO `t_subsystem` (`id`, `name`, `domain`, `syskey`, `secret`, `status`, `order_no`, `create_uid`) VALUES
(1, '管理后台', 'http://localhost:8000', 'admin', 'admin_secret_key_2024', 1, 1, 1);

-- ============================================
-- 2. 组织表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_organization` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级组织ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '组织名称',
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '组织级别',
  `node_path` varchar(500) NOT NULL DEFAULT '' COMMENT '组织节点路径',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-启用，2-禁用，0-删除',
  `order_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序号',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_node_path` (`node_path`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='组织表';

-- 插入默认组织
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`) VALUES
(1, 0, '总公司', 1, '1', 1, 1, 1);

-- ============================================
-- 3. 用户表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '登录名',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `real_name` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `org_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属组织ID',
  `is_root` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否root用户：1-是，0-否',
  `is_staff` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否内部员工：1-是，0-否',
  `staff_no` int(11) NOT NULL DEFAULT '0' COMMENT '员工号',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态：1-在职，0-离职',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码（SHA1加密）',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_org_id` (`org_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 插入默认管理员用户（密码：admin123，SHA1加密后：f865b53623b121fd34ee5426c792e5c33af8c227）
-- 注意：实际密码是 admin123，SHA1 加密后的值是 f865b53623b121fd34ee5426c792e5c33af8c227
INSERT INTO `t_user` (`id`, `name`, `real_name`, `email`, `phone`, `role_id`, `org_id`, `is_root`, `is_staff`, `status`, `password`, `create_uid`) VALUES
(1, 'admin', '系统管理员', 'admin@example.com', '13800138000', 0, 1, 1, 1, 1, 'f865b53623b121fd34ee5426c792e5c33af8c227', 1);

-- ============================================
-- 4. 角色表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-启用，2-禁用',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 插入默认角色
INSERT INTO `t_role` (`id`, `name`, `status`, `create_uid`) VALUES
(1, '超级管理员', 1, 1),
(2, '普通管理员', 1, 1);

-- ============================================
-- 5. 用户角色关联表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_role` (`user_id`, `role_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- ============================================
-- 6. 菜单表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `system_id` int(11) NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '菜单级别',
  `node_path` varchar(500) NOT NULL DEFAULT '' COMMENT '菜单节点路径',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单URL',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-有效，2-禁用',
  `order_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序号',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_system_id` (`system_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_node_path` (`node_path`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';

-- 插入默认菜单
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`) VALUES
(1, 1, 0, '系统管理', 1, '1', '/system', 1, 1, 1),
(2, 1, 1, '用户管理', 2, '1,2', '/system/user', 1, 1, 1),
(3, 1, 1, '角色管理', 2, '1,3', '/system/role', 1, 2, 1),
(4, 1, 1, '菜单管理', 2, '1,4', '/system/menu', 1, 3, 1),
(5, 1, 1, '组织管理', 2, '1,5', '/system/organization', 1, 4, 1);

-- ============================================
-- 7. 操作表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `system_id` int(11) NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '操作名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '操作编号（用于前端控制按钮显示）',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-有效，2-禁用',
  `order_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序号',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_system_id` (`system_id`),
  KEY `idx_menu_id` (`menu_id`),
  KEY `idx_code` (`code`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作表';

-- 插入默认操作（增删改查）
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`) VALUES
(1, 1, 2, '新增', 'add', 1, 1, 1),
(2, 1, 2, '编辑', 'edit', 1, 2, 1),
(3, 1, 2, '删除', 'delete', 1, 3, 1),
(4, 1, 2, '查看', 'view', 1, 4, 1),
(5, 1, 3, '新增', 'add', 1, 1, 1),
(6, 1, 3, '编辑', 'edit', 1, 2, 1),
(7, 1, 3, '删除', 'delete', 1, 3, 1),
(8, 1, 3, '查看', 'view', 1, 4, 1);

-- ============================================
-- 8. 角色菜单关联表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `system_id` int(11) NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `opera_ids` varchar(500) NOT NULL DEFAULT '' COMMENT '菜单下的可用操作ID（逗号分隔）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu` (`role_id`, `system_id`, `menu_id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_system_id` (`system_id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';

-- ============================================
-- 9. 用户权限表
-- ============================================
CREATE TABLE IF NOT EXISTS `t_user_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `system_id` int(11) NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `opera_ids` varchar(500) NOT NULL DEFAULT '' COMMENT '菜单下的可用操作ID（逗号分隔）',
  `create_uid` int(11) NOT NULL DEFAULT '0' COMMENT '创建人ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_uid` int(11) NOT NULL DEFAULT '0' COMMENT '更新人ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_system_menu` (`user_id`, `system_id`, `menu_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_system_id` (`system_id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户权限表';

-- ============================================
-- 初始化完成
-- ============================================

