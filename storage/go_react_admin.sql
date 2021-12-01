/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : go_react_admin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 06/11/2021 18:06:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '页面路径',
  `label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `is_page` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否页面级别',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of authority
-- ----------------------------
BEGIN;
INSERT INTO `authority` VALUES (1, '/home', '首页', 0, 1);
INSERT INTO `authority` VALUES (2, '/root', 'root', 0, 1);
INSERT INTO `authority` VALUES (3, '/common', '公共页面', 0, 1);
INSERT INTO `authority` VALUES (4, '/home/child1', 'home子页面1', 1, 1);
INSERT INTO `authority` VALUES (5, '/home/child2', 'home子页面2', 1, 1);
INSERT INTO `authority` VALUES (6, '/home/child1/a', 'home-a', 4, 1);
INSERT INTO `authority` VALUES (7, '/home/child1/b', 'home-b', 4, 1);
INSERT INTO `authority` VALUES (8, '/rbac', 'RBAC', 0, 1);
INSERT INTO `authority` VALUES (9, '/rbac/user', '用户管理', 8, 1);
INSERT INTO `authority` VALUES (10, '/rbac/role', '角色管理 ', 8, 1);
INSERT INTO `authority` VALUES (11, '', '绑定权限', 10, 0);
INSERT INTO `authority` VALUES (12, '', '删除角色', 10, 0);
INSERT INTO `authority` VALUES (13, '/test', '测试页面', 0, 1);
INSERT INTO `authority` VALUES (14, '', '测试添加', 13, 0);
INSERT INTO `authority` VALUES (15, '', '测试删除', 13, 0);
INSERT INTO `authority` VALUES (16, '/rbac/auth', '权限管理', 8, 1);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL,
  `role_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 'root');
INSERT INTO `role` VALUES (2, 'admin');
INSERT INTO `role` VALUES (3, 'user');
INSERT INTO `role` VALUES (4, 'common');
COMMIT;

-- ----------------------------
-- Table structure for role_auth
-- ----------------------------
DROP TABLE IF EXISTS `role_auth`;
CREATE TABLE `role_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `auth_id` int unsigned NOT NULL,
  `role_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role_auth
-- ----------------------------
BEGIN;
INSERT INTO `role_auth` VALUES (36, 3, 2);
INSERT INTO `role_auth` VALUES (37, 1, 2);
INSERT INTO `role_auth` VALUES (40, 1, 3);
INSERT INTO `role_auth` VALUES (41, 3, 3);
INSERT INTO `role_auth` VALUES (42, 2, 3);
INSERT INTO `role_auth` VALUES (178, 1, 1);
INSERT INTO `role_auth` VALUES (179, 2, 1);
INSERT INTO `role_auth` VALUES (180, 3, 1);
INSERT INTO `role_auth` VALUES (181, 4, 1);
INSERT INTO `role_auth` VALUES (182, 5, 1);
INSERT INTO `role_auth` VALUES (183, 6, 1);
INSERT INTO `role_auth` VALUES (184, 7, 1);
INSERT INTO `role_auth` VALUES (185, 8, 1);
INSERT INTO `role_auth` VALUES (186, 10, 1);
INSERT INTO `role_auth` VALUES (187, 11, 1);
INSERT INTO `role_auth` VALUES (188, 12, 1);
INSERT INTO `role_auth` VALUES (189, 13, 1);
INSERT INTO `role_auth` VALUES (190, 14, 1);
INSERT INTO `role_auth` VALUES (191, 15, 1);
INSERT INTO `role_auth` VALUES (192, 9, 1);
INSERT INTO `role_auth` VALUES (193, 16, 1);
INSERT INTO `role_auth` VALUES (194, 1, 4);
INSERT INTO `role_auth` VALUES (195, 3, 4);
INSERT INTO `role_auth` VALUES (196, 13, 4);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tel` char(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'u1', '19381609624', '123456');
INSERT INTO `user` VALUES (2, 'u2', '13344445555', '123456');
INSERT INTO `user` VALUES (3, 'u3', '', '');
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `role_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `urindex` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (17, 1, 1);
INSERT INTO `user_role` VALUES (18, 2, 3);
INSERT INTO `user_role` VALUES (19, 2, 4);
INSERT INTO `user_role` VALUES (20, 3, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
