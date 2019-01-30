/*
Navicat MySQL Data Transfer

Source Server         : local_mysql8
Source Server Version : 80013
Source Host           : localhost:3308
Source Database       : godatabase

Target Server Type    : MYSQL
Target Server Version : 80013
File Encoding         : 65001

Date: 2018-11-28 17:42:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for application
-- ----------------------------
DROP TABLE IF EXISTS `application`;
CREATE TABLE `application` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `SysName` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '系统应用名称',
  `SysUrl` varchar(255) NOT NULL COMMENT '系统访问路径',
  `CreationTime` datetime DEFAULT NULL COMMENT '创建时间',
  `CreatorUserId` bigint(20) DEFAULT NULL COMMENT '创建人',
  `IsDeleted` int(11) DEFAULT '0' COMMENT '是否删除 0 未删除 1 删除',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of application
-- ----------------------------
INSERT INTO `application` VALUES ('1', 'TMS', '', null, null, null);
INSERT INTO `application` VALUES ('2', '应用2', '', null, null, null);

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreationTime` datetime NOT NULL,
  `CreatorUserId` bigint(20) DEFAULT NULL,
  `Discriminator` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'RolePermissionSetting',
  `IsGranted` tinyint(4) NOT NULL DEFAULT '1',
  `Name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限简称 用于控制操作权限',
  `TenantId` int(11) DEFAULT NULL,
  `RoleId` int(11) DEFAULT NULL,
  `UserId` bigint(20) DEFAULT NULL,
  `DisplayName` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台展示名称',
  `SysId` int(11) DEFAULT NULL COMMENT '应用id',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES ('2', '2018-09-11 07:51:17', '0', '', '0', 'platform.tenant', '0', '0', '1', '平台租户管理', '1');
INSERT INTO `permission` VALUES ('3', '2018-09-11 07:55:20', '0', '', '0', 'platform.basepermission', '0', '0', '1', '平台基础权限管理', '1');
INSERT INTO `permission` VALUES ('4', '2018-09-11 07:56:01', '0', '', '0', 'platform.role', '0', '0', '1', '平台角色管理', '1');
INSERT INTO `permission` VALUES ('9', '2018-09-11 08:21:51', '0', '', '0', 'platform.permission', '0', '0', '0', '平台权限管理', null);
INSERT INTO `permission` VALUES ('32', '2018-09-13 07:13:30', '0', '', '0', 'platform.basepermission', '0', '16', '0', '平台基础权限管理', null);
INSERT INTO `permission` VALUES ('33', '2018-09-13 07:13:34', '0', '', '0', 'platform.tenant', '0', '16', '0', '平台租户管理', null);
INSERT INTO `permission` VALUES ('37', '2018-09-14 05:50:33', '0', '', '0', 'platform.role', '12', '0', '0', '平台角色管理', null);
INSERT INTO `permission` VALUES ('38', '2018-09-14 05:50:33', '0', '', '0', 'platform.permission', '12', '0', '0', '平台权限管理', null);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `ConcurrencyStamp` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  `CreatorUserId` bigint(20) DEFAULT NULL,
  `DeleterUserId` bigint(20) DEFAULT NULL,
  `DeletionTime` datetime DEFAULT NULL,
  `DisplayName` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `IsDefault` tinyint(4) DEFAULT NULL,
  `IsDeleted` tinyint(4) NOT NULL DEFAULT '0',
  `IsStatic` tinyint(4) DEFAULT NULL,
  `LastModificationTime` datetime DEFAULT NULL,
  `LastModifierUserId` bigint(20) DEFAULT NULL,
  `Name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `NormalizedName` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `TenantId` int(11) DEFAULT NULL,
  `Description` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '', '2018-09-12 02:16:49', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '0', '');
INSERT INTO `role` VALUES ('8', '', '2018-09-12 02:45:25', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '12', '');
INSERT INTO `role` VALUES ('16', '', '2018-09-12 23:13:30', '0', '0', null, '操作员', '0', '1', '0', null, '0', 'oper', 'OPER', '0', '');
INSERT INTO `role` VALUES ('17', '', '2018-09-28 08:18:52', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '13', '');
INSERT INTO `role` VALUES ('18', '', '2018-09-28 08:34:15', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '14', '');
INSERT INTO `role` VALUES ('19', '', '2018-09-28 08:55:27', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '15', '');
INSERT INTO `role` VALUES ('20', '', '2018-09-28 08:57:25', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '16', '');
INSERT INTO `role` VALUES ('21', '', '2018-09-28 08:58:41', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '17', '');
INSERT INTO `role` VALUES ('22', '', '2018-09-28 09:07:53', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '18', '');
INSERT INTO `role` VALUES ('23', '', '2018-09-28 09:10:32', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '19', '');
INSERT INTO `role` VALUES ('24', '', '2018-09-28 09:14:43', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '20', '');
INSERT INTO `role` VALUES ('25', '', '2018-09-28 09:17:06', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '21', '');
INSERT INTO `role` VALUES ('26', '', '2018-09-28 17:29:04', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '22', '');
INSERT INTO `role` VALUES ('27', '', '2018-09-28 17:33:41', '0', '0', null, '管理员', '0', '0', '0', null, '0', 'admin', 'ADMIN', '23', '');

-- ----------------------------
-- Table structure for ssouser
-- ----------------------------
DROP TABLE IF EXISTS `ssouser`;
CREATE TABLE `ssouser` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '手机号码 当用户使用手机号登陆时，直接走ssouser，邮箱等其他方式走 user--->ssouser',
  `Passwd` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登陆密码',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='sso系统验证用户';

-- ----------------------------
-- Records of ssouser
-- ----------------------------
INSERT INTO `ssouser` VALUES ('1', '13299998888', '123456');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `AccessFailedCount` int(11) NOT NULL DEFAULT '0',
  `AuthenticationSource` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `ConcurrencyStamp` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  `CreatorUserId` bigint(20) DEFAULT NULL,
  `DeleterUserId` bigint(20) DEFAULT NULL,
  `DeletionTime` datetime DEFAULT NULL,
  `EmailAddress` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `EmailConfirmationCode` varchar(328) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `IsActive` tinyint(4) NOT NULL DEFAULT '1',
  `IsDeleted` tinyint(4) NOT NULL DEFAULT '0',
  `IsEmailConfirmed` tinyint(4) NOT NULL DEFAULT '1',
  `IsLockoutEnabled` tinyint(4) NOT NULL DEFAULT '1',
  `IsPhoneNumberConfirmed` tinyint(4) NOT NULL DEFAULT '1',
  `IsTwoFactorEnabled` tinyint(4) NOT NULL DEFAULT '1',
  `LastLoginTime` datetime DEFAULT NULL,
  `LastModificationTime` datetime DEFAULT NULL,
  `LastModifierUserId` bigint(20) DEFAULT NULL,
  `LockoutEndDateUtc` datetime DEFAULT NULL,
  `Name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `NormalizedEmailAddress` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `NormalizedUserName` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `PasswordResetCode` varchar(328) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `PhoneNumber` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `SecurityStamp` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `Surname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `TenantId` int(11) DEFAULT NULL,
  `UserName` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `SysId` int(11) NOT NULL,
  `SsoId` int(11) DEFAULT NULL COMMENT 'sso验证id',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '0', null, null, '2018-11-08 13:29:42', null, null, '2018-11-06 13:29:48', 'xsungroup@163.com', null, '1', '0', '1', '1', '1', '1', null, null, null, null, '测试', '1123131', '王虎', null, '13299998888', null, '12312', '1', 'jack', '1', '1');

-- ----------------------------
-- Table structure for userrole
-- ----------------------------
DROP TABLE IF EXISTS `userrole`;
CREATE TABLE `userrole` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreationTime` datetime NOT NULL,
  `CreatorUserId` bigint(20) DEFAULT NULL,
  `RoleId` int(11) NOT NULL,
  `TenantId` int(11) DEFAULT NULL,
  `UserId` bigint(20) NOT NULL,
  `SysId` int(11) DEFAULT NULL COMMENT '系统id',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of userrole
-- ----------------------------
INSERT INTO `userrole` VALUES ('7', '2018-09-18 02:09:15', '0', '8', '12', '12', null);
INSERT INTO `userrole` VALUES ('8', '2018-09-28 08:18:52', '0', '17', '13', '13', null);
INSERT INTO `userrole` VALUES ('9', '2018-09-28 08:34:15', '0', '18', '14', '14', null);
INSERT INTO `userrole` VALUES ('10', '2018-09-28 08:55:27', '0', '19', '15', '15', null);
INSERT INTO `userrole` VALUES ('11', '2018-09-28 08:57:25', '0', '20', '16', '16', null);
INSERT INTO `userrole` VALUES ('12', '2018-09-28 08:58:41', '0', '21', '17', '17', null);
INSERT INTO `userrole` VALUES ('13', '2018-09-28 09:07:53', '0', '22', '18', '18', null);
INSERT INTO `userrole` VALUES ('14', '2018-09-28 09:10:32', '0', '23', '19', '19', null);
INSERT INTO `userrole` VALUES ('15', '2018-09-28 09:14:43', '0', '24', '20', '20', null);
INSERT INTO `userrole` VALUES ('16', '2018-09-28 09:17:06', '0', '25', '21', '21', null);
INSERT INTO `userrole` VALUES ('17', '2018-09-28 17:29:04', '0', '26', '22', '22', null);
INSERT INTO `userrole` VALUES ('18', '2018-09-28 17:33:41', '0', '27', '23', '23', null);
