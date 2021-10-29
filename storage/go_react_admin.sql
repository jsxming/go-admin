/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : localhost:3306
 Source Schema         : go_react_admin

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 27/10/2021 20:36:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of authority
-- ----------------------------
BEGIN;
INSERT INTO `authority` VALUES (1, '/home', '首页', 0);
INSERT INTO `authority` VALUES (2, '/root', 'root', 0);
INSERT INTO `authority` VALUES (3, '/common', '公共页面', 0);
INSERT INTO `authority` VALUES (4, '/home/child1', 'home子页面1', 1);
INSERT INTO `authority` VALUES (5, '/home/child2', 'home子页面2', 1);
INSERT INTO `authority` VALUES (6, '/home/child1/a', 'home-a', 4);
INSERT INTO `authority` VALUES (7, '/home/child1/b', 'home-b', 4);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL,
  `role_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

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
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `auth_id` int(10) unsigned NOT NULL,
  `role_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role_auth
-- ----------------------------
BEGIN;
INSERT INTO `role_auth` VALUES (6, 1, 4);
INSERT INTO `role_auth` VALUES (11, 1, 1);
INSERT INTO `role_auth` VALUES (12, 2, 1);
INSERT INTO `role_auth` VALUES (13, 3, 1);
INSERT INTO `role_auth` VALUES (14, 4, 1);
INSERT INTO `role_auth` VALUES (15, 5, 1);
INSERT INTO `role_auth` VALUES (16, 6, 1);
INSERT INTO `role_auth` VALUES (17, 7, 1);
INSERT INTO `role_auth` VALUES (29, 2, 2);
INSERT INTO `role_auth` VALUES (30, 3, 2);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tel` char(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

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
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `role_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `urindex` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

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
