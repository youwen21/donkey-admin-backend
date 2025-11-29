-- ============================================
-- donkey-admin-backend 数据库初始化脚本
-- ============================================

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `donkey_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `donkey_admin`;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_menu
-- ----------------------------
DROP TABLE IF EXISTS `t_menu`;
CREATE TABLE `t_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `system_id` int NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '上级菜单ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单名',
  `level` int NOT NULL DEFAULT '0' COMMENT '菜单级别',
  `node_path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '节点路径',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '对应的uri',
  `status` int NOT NULL DEFAULT '1' COMMENT '1:有效 2:禁用',
  `order_no` int NOT NULL DEFAULT '0' COMMENT '菜单序号',
  `create_uid` int NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sys_id_pid_name` (`system_id`,`parent_id`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=277 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_menu
-- ----------------------------
BEGIN;
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (1, 1, 0, '基础管理', 0, '/', '', 1, 11, 1, '2018-03-21 10:14:31', 1, '2018-03-21 11:46:59');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (2, 1, 0, '菜单和操作管理', 0, '/', '', 1, 12, 1, '2018-03-21 10:14:48', 1, '2025-11-28 08:10:39');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (3, 1, 0, '系统管理', 0, '/', '', 1, 13, 1, '2018-03-21 10:15:05', 1, '2018-03-21 11:47:22');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (4, 1, 1, '组织', 1, '/1/', '/admin/org/list', 1, 111, 1, '2018-03-21 10:16:23', 1, '2025-11-22 21:14:40');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (5, 1, 1, '用户', 1, '/1/', '/admin/user/list', 1, 112, 1, '2018-03-21 10:17:13', 1, '2018-03-21 11:47:42');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (6, 1, 1, '角色', 1, '/1/', '/admin/role/list', 2, 113, 1, '2018-03-21 10:17:51', 1, '2025-11-28 21:24:23');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (7, 1, 2, '菜单', 1, '/2/', '/admin/menu/list', 1, 121, 1, '2018-03-21 10:19:07', 1, '2025-11-28 08:10:23');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (8, 1, 2, '操作', 1, '/2/', '/admin/operation/list', 1, 122, 1, '2018-03-21 10:30:15', 1, '2025-11-28 08:10:30');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (9, 1, 3, '子系统', 1, '/3/', '/admin/subsystem/list', 1, 131, 1, '2018-03-21 10:31:10', 1, '2025-11-28 08:10:48');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (93, 7, 0, '产品管理', 0, '/', '', 1, 71, 1, '2019-10-24 14:09:48', 1, '2019-10-24 14:19:09');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (94, 7, 93, '订单列表', 1, '/93/', '/order/list', 1, 711, 1, '2019-10-24 14:20:12', 1, '2019-10-24 14:20:12');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (95, 7, 93, '产品分类列表', 1, '/93/', '/classify/list', 1, 712, 1, '2019-10-24 14:20:42', 1, '2019-10-24 14:20:42');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (96, 7, 93, '标签列表', 1, '/93/', '/tag/list', 1, 713, 1, '2019-10-24 14:21:07', 1, '2019-10-24 14:21:07');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (200, 7, 93, '支出账单', 1, '/93/', '/bill/list', 1, 717, 22, '2020-11-20 14:47:00', 22, '2020-11-20 14:47:00');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (230, 7, 0, '统计数据', 0, '/', '', 1, 91, 1, '2021-04-27 17:03:55', 1, '2021-04-27 17:03:55');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (234, 7, 230, '产品销售率', 1, '/230/', '/statistics/selling/list', 1, 911, 1, '2021-04-27 17:07:09', 1, '2021-05-11 18:46:36');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (275, 7, 234, 'sss', 2, '/230/234/', '/a/a/a/a/a', 1, 0, 0, '2025-11-23 21:42:53', 0, '2025-11-23 21:42:53');
INSERT INTO `t_menu` (`id`, `system_id`, `parent_id`, `name`, `level`, `node_path`, `url`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (276, 7, 275, 'bbbbb', 3, '/230/234/275/', '/b/b/b/b/b/b', 1, 0, 0, '2025-11-23 21:47:45', 0, '2025-11-23 21:47:45');
COMMIT;

-- ----------------------------
-- Table structure for t_operation
-- ----------------------------
DROP TABLE IF EXISTS `t_operation`;
CREATE TABLE `t_operation` (
  `id` int NOT NULL AUTO_INCREMENT,
  `system_id` int NOT NULL DEFAULT '0' COMMENT '子系统ID',
  `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '操作名称',
  `code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '操作编号,此编号对应页面class, 用来控制按钮是否显示，也对应后台接口操作编号，校验用户是否有此操作权限',
  `status` int NOT NULL DEFAULT '1' COMMENT '1:有效 2:禁用',
  `order_no` int NOT NULL DEFAULT '1' COMMENT '排序',
  `create_uid` int NOT NULL DEFAULT '0' COMMENT '创建人id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0' COMMENT '更新用户id',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_operation
-- ----------------------------
BEGIN;
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (1, 1, 4, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:33:37', 0, '2025-11-28 23:33:37');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (2, 1, 4, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:34:07', 0, '2025-11-28 23:34:07');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (3, 1, 4, '删除', 'btn-del', 1, 0, 1, '2025-11-28 23:34:33', 0, '2025-11-28 23:34:33');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (4, 1, 4, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:35:19', 0, '2025-11-28 23:35:19');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (5, 1, 5, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:35:55', 0, '2025-11-28 23:35:55');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (6, 1, 5, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:36:15', 0, '2025-11-28 23:36:15');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (7, 1, 5, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:36:50', 0, '2025-11-28 23:36:50');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (8, 1, 5, '授权', 'config-perm', 1, 0, 1, '2025-11-28 23:37:40', 0, '2025-11-28 23:37:40');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (9, 1, 6, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:33:37', 0, '2025-11-28 23:40:26');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (10, 1, 6, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:34:07', 0, '2025-11-28 23:40:27');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (11, 1, 6, '删除', 'btn-del', 1, 0, 1, '2025-11-28 23:34:33', 0, '2025-11-28 23:40:30');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (12, 1, 6, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:35:19', 0, '2025-11-28 23:40:32');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (13, 1, 7, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:33:37', 0, '2025-11-28 23:40:44');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (14, 1, 7, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:34:07', 0, '2025-11-28 23:40:46');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (15, 1, 7, '删除', 'btn-del', 1, 0, 1, '2025-11-28 23:34:33', 0, '2025-11-28 23:40:48');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (16, 1, 7, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:35:19', 0, '2025-11-28 23:40:53');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (17, 1, 8, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:33:37', 0, '2025-11-28 23:41:01');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (18, 1, 8, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:34:07', 0, '2025-11-28 23:41:03');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (19, 1, 8, '删除', 'btn-del', 1, 0, 1, '2025-11-28 23:34:33', 0, '2025-11-28 23:41:06');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (20, 1, 8, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:35:19', 0, '2025-11-28 23:41:10');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (21, 1, 9, '新增', 'btn-add', 1, 0, 1, '2025-11-28 23:33:37', 0, '2025-11-28 23:42:39');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (22, 1, 9, '编辑', 'btn-edit', 1, 0, 1, '2025-11-28 23:34:07', 0, '2025-11-28 23:42:41');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (23, 1, 9, '删除', 'btn-del', 1, 0, 1, '2025-11-28 23:34:33', 0, '2025-11-28 23:42:43');
INSERT INTO `t_operation` (`id`, `system_id`, `menu_id`, `name`, `code`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (24, 1, 9, '变更状态', 'btn-status', 1, 0, 1, '2025-11-28 23:35:19', 0, '2025-11-28 23:42:45');
COMMIT;

-- ----------------------------
-- Table structure for t_organization
-- ----------------------------
DROP TABLE IF EXISTS `t_organization`;
CREATE TABLE `t_organization` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '上级组织ID',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组织名称',
  `level` int NOT NULL DEFAULT '0' COMMENT '组织级别',
  `node_path` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组织节点路径',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态 1启用 2禁用 0删除',
  `order_no` int NOT NULL DEFAULT '0' COMMENT '序号',
  `create_uid` int NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `parent_name` (`parent_id`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_organization
-- ----------------------------
BEGIN;
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (1, 0, 'x公司', 0, '/', 1, 1999, 1, '2018-03-21 10:25:43', 1, '2018-03-21 10:25:43');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (2, 1, '技术部', 1, '/1/', 1, 12, 1, '2018-03-21 10:26:03', 1, '2018-03-21 10:26:03');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (3, 1, '运营部', 1, '/1/', 1, 11, 1, '2018-03-21 10:26:25', 1, '2018-03-21 10:26:25');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (4, 1, '客户部', 1, '/', 1, 13, 1, '2018-03-22 14:56:07', 1, '2025-11-24 18:11:07');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (5, 1, '市场部门', 0, '/', 1, 14, 1, '2021-06-10 10:06:32', 1, '2021-06-10 10:06:32');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (13, 1, '内容部', 1, '/1/', 1, 0, 1, '2021-07-19 17:26:58', 1, '2021-07-19 17:26:58');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (14, 13, '编辑', 2, '/1/13/', 1, 0, 1, '2021-07-19 17:27:26', 1, '2021-07-19 17:27:26');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (20, 3, '编辑', 2, '/1/3/', 1, 100, 0, '2022-04-10 17:22:29', 0, '2022-04-10 17:22:29');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (23, 20, 'yyyyy', 3, '/1/3/20/', 1, 0, 0, '2025-11-20 21:26:03', 0, '2025-11-20 21:26:03');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (24, 26, 'sdfdsfsfs', 2, '/', 2, 0, 0, '2025-11-20 21:40:44', 0, '2025-11-24 18:11:11');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (26, 5, 'hhh', 1, '/5/', 1, 0, 0, '2025-11-20 21:52:34', 0, '2025-11-20 21:52:34');
INSERT INTO `t_organization` (`id`, `parent_id`, `name`, `level`, `node_path`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (27, 25, 'jjjjj', 2, '/5/25/', 1, 0, 0, '2025-11-20 21:52:48', 0, '2025-11-20 21:52:48');
COMMIT;

-- ----------------------------
-- Table structure for t_subsystem
-- ----------------------------
DROP TABLE IF EXISTS `t_subsystem`;
CREATE TABLE `t_subsystem` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `domain` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '域名',
  `syskey` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'key',
  `secret` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'secret',
  `status` int NOT NULL DEFAULT '1' COMMENT '1:有效 2:禁用',
  `order_no` int NOT NULL DEFAULT '1' COMMENT '排序',
  `create_uid` int NOT NULL DEFAULT '0' COMMENT '创建人id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0' COMMENT '更新用户id',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `domain` (`domain`) USING BTREE,
  UNIQUE KEY `syskey` (`syskey`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_subsystem
-- ----------------------------
BEGIN;
INSERT INTO `t_subsystem` (`id`, `name`, `domain`, `syskey`, `secret`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (1, '登录系统', 'http://127.0.0.1:8200/', 'a8f391aa78d852f31e73189e49f12968', '6ad1411abfa8f00c5b0fedf03f9d8229', 1, 1, 1, '2018-03-21 10:11:51', 1, '2022-04-14 12:04:27');
INSERT INTO `t_subsystem` (`id`, `name`, `domain`, `syskey`, `secret`, `status`, `order_no`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (7, 'x后台cc', 'http://www.qq.com', '9dd285060a7e0b427d5cb921d33a7496', '1e51d45ed174f929af478468e5988373', 1, 7, 1, '2019-10-24 12:32:13', 1, '2025-11-22 20:13:57');
COMMIT;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登陆名',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `real_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '真实名字',
  `is_root` tinyint NOT NULL DEFAULT '2' COMMENT '是否root用户',
  `is_staff` tinyint NOT NULL DEFAULT '1' COMMENT '是否内部员工',
  `staff_no` int NOT NULL DEFAULT '0' COMMENT '员工号',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '员工邮箱',
  `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '员工手机号',
  `status` int NOT NULL DEFAULT '1' COMMENT '是否在职，1:在职，0:离职',
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `org_id` int NOT NULL DEFAULT '0' COMMENT '所属组织',
  `create_uid` int NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_name` (`name`) USING BTREE,
  UNIQUE KEY `idx_phone` (`phone`) USING BTREE,
  UNIQUE KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_user
-- ----------------------------
BEGIN;
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (1, 'admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', '管理员', 1, 1, 1, 'root@qq.com', '15900000000', 1, '', 0, 1, 1, '2018-03-15 12:17:29', 1, '2025-11-20 09:32:49');
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (111, 'tryt', '935edf6db5192fbe3b953c75c98f26a2534e3c57', 'realWe', 0, 1, 0, 'you@yeah.net', '18500888888', 1, '', 0, 20, 0, '2022-04-14 18:11:01', 0, '2025-11-26 15:22:42');
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (112, 'test2', '7c4a8d09ca3762af61e59520943dc26494f8941b', 'realname-test2', 1, 1, 0, 'qqq@s.com', '2131231', 1, '', 0, 16, 0, '2022-04-14 18:20:54', 0, '2025-11-23 19:43:01');
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (113, 'demo555555', '7c4a8d09ca3762af61e59520943dc26494f8941b', 'realname-demo1', 0, 0, 0, 'demo1@ss.com', '1234561', 1, '', 0, 20, 0, '2022-04-15 15:18:57', 0, '2025-11-23 19:43:05');
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (114, 'demo2', '7c4a8d09ca3762af61e59520943dc26494f8941b', 'realname-demo2', 0, 1, 0, 'demo2@xx.com', '18590099999', 1, '', 0, 2, 0, '2022-04-15 15:55:00', 0, '2025-11-23 19:43:08');
INSERT INTO `t_user` (`id`, `name`, `password`, `real_name`, `is_root`, `is_staff`, `staff_no`, `email`, `phone`, `status`, `avatar`, `role_id`, `org_id`, `create_uid`, `create_time`, `update_uid`, `update_time`) VALUES (115, 'trytry', '95c145a2dbcf4130af906e7d29aae0d54abe073b', 'realname', 0, 0, 0, '', '', 0, '', 0, 0, 0, '2025-11-22 08:23:13', 0, '2025-11-23 19:43:10');
COMMIT;

-- ----------------------------
-- Table structure for t_user_permission
-- ----------------------------
DROP TABLE IF EXISTS `t_user_permission`;
CREATE TABLE `t_user_permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT '0' COMMENT '角色ID',
  `system_id` int NOT NULL DEFAULT '0' COMMENT '系统id',
  `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单id',
  `opera_ids` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单下的可用操作',
  `create_uid` int NOT NULL DEFAULT '0' COMMENT '创建人id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_uid` int NOT NULL DEFAULT '0' COMMENT '更新用户id',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of t_user_permission
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

-- ============================================
-- 初始化完成
-- ============================================

