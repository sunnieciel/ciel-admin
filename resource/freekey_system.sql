-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: 127.0.0.1    Database: freekey_system
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `c_banner`
--

DROP TABLE IF EXISTS `c_banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `c_banner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL,
  `image` varchar(64) NOT NULL,
  `link` varchar(128) DEFAULT NULL,
  `desc` varchar(64) DEFAULT NULL,
  `sort` int DEFAULT '0',
  `status` tinyint unsigned DEFAULT '1' COMMENT '1open 2close',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `c_banner`
--

LOCK TABLES `c_banner` WRITE;
/*!40000 ALTER TABLE `c_banner` DISABLE KEYS */;
INSERT INTO `c_banner` (`id`, `title`, `image`, `link`, `desc`, `sort`, `status`, `created_at`) VALUES (2,'test','icon/icon43.png','https://google.com','',0,1,'2022-12-29 00:58:38');
/*!40000 ALTER TABLE `c_banner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_admin`
--

DROP TABLE IF EXISTS `s_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rid` int NOT NULL,
  `uname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `unread_msg_num` int unsigned DEFAULT '0',
  `pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  KEY `unread_msg_num` (`unread_msg_num`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `unread_msg_num`, `pwd`, `nickname`, `email`, `phone`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel',0,'$2a$10$OAp3RJVKv6WhAX3o.fY/A.R0jUOyzvtlfxpS3DgHtEVkLx/lY6b4.','I\'m ciel','','',1,'2022-03-08 08:59:33','2022-12-27 12:25:10'),(42,1,'admin',0,'$2a$10$rLhxw1XB0lbj8YcHmkMqVOkhdduvoqhf3zdo9TKBgAA8jYoasL3AG','admin','','',1,'2022-07-02 11:28:52','2022-12-28 07:50:52');
/*!40000 ALTER TABLE `s_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_admin_login_log`
--

DROP TABLE IF EXISTS `s_admin_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_admin_login_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL COMMENT '{"label":"用户id","searchType":1,"hide":1,"disabled":1,"required":1}',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT ' {"label":"登录IP","notShow":0,"fieldType":"text","editHide":0,"editDisabled":0,"required":1}',
  `area` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '{"searchType":2,"hide":1}',
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_admin_login_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=215 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `ip`, `area`, `status`, `created_at`, `updated_at`) VALUES (207,42,'::1',NULL,1,'2022-12-27 20:41:11','2022-12-27 20:41:11'),(208,42,'::1',NULL,1,'2022-12-27 23:26:22','2022-12-27 23:26:22'),(209,42,'::1',NULL,1,'2022-12-27 23:32:04','2022-12-27 23:32:04'),(210,42,'::1',NULL,1,'2022-12-28 20:00:24','2022-12-28 20:00:24'),(211,42,'::1',NULL,1,'2022-12-29 00:50:14','2022-12-29 00:50:14'),(212,42,'::1',NULL,1,'2022-12-30 14:38:56','2022-12-30 14:38:56'),(213,42,'::1',NULL,1,'2022-12-30 20:30:08','2022-12-30 20:30:08'),(214,42,'::1',NULL,1,'2022-12-30 20:47:19','2022-12-30 20:47:19');
/*!40000 ALTER TABLE `s_admin_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_admin_message`
--

DROP TABLE IF EXISTS `s_admin_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_admin_message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `aid` bigint unsigned NOT NULL,
  `type` tinyint unsigned DEFAULT '1' COMMENT '1 系统消息',
  `title` varchar(128) NOT NULL,
  `content` text,
  `url` varchar(64) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `aid` (`aid`),
  KEY `type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_message`
--

LOCK TABLES `s_admin_message` WRITE;
/*!40000 ALTER TABLE `s_admin_message` DISABLE KEYS */;
INSERT INTO `s_admin_message` (`id`, `aid`, `type`, `title`, `content`, `url`, `created_at`) VALUES (13,1,1,'用户user33创建了10.00充值订单','','/admin/walletTopUpApplication?trans_id=MR-tK9lQIqj9SfSL','2022-12-27 20:24:39');
/*!40000 ALTER TABLE `s_admin_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_api`
--

DROP TABLE IF EXISTS `s_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `method` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` tinyint unsigned DEFAULT '4' COMMENT '类型 1添加 2删除 3修改 4查看 5 页面 ',
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=371 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `type`, `desc`, `created_at`, `updated_at`) VALUES (56,'/admin/file','1','文件',5,'文件页面','2022-06-12 14:10:23','2022-06-14 13:45:01'),(61,'/admin/file/put','2','文件',3,'文件修改','2022-06-12 14:10:23','2022-07-23 20:08:04'),(62,'/admin/roleApi/path','1','角色',5,'角色禁用API','2022-06-12 17:02:13','2022-09-03 15:40:54'),(65,'/admin/roleApi/:id','1','角色',2,'删除角色禁用API','2022-06-12 17:02:13','2022-07-20 20:18:08'),(66,'/admin/roleApi/post','2','角色',1,'角色禁用API添加','2022-06-12 17:02:13','2022-07-23 20:15:37'),(68,'/admin/roleMenu/path','1','角色',5,'角色菜单页面','2022-06-13 15:19:45','2022-07-23 19:33:45'),(71,'/admin/roleMenu/:id','1','角色',2,'角色菜单删除','2022-06-13 15:19:45','2022-07-23 20:14:58'),(72,'/admin/roleMenu/post','2','角色',1,'角色菜单添加','2022-06-13 15:19:45','2022-07-23 20:14:48'),(144,'/admin/api','1','API',5,'API页面','2022-06-26 15:14:58','2022-07-23 20:03:54'),(145,'/admin/api/path/add','1','API',5,'API添加页面','2022-06-26 15:14:58','2022-07-23 20:03:47'),(146,'/admin/api/path/edit/:id','1','API',5,'API修改页面','2022-06-26 15:14:58','2022-07-23 20:03:40'),(147,'/admin/api/path/del/:id','1','API',2,'API删除操作','2022-06-26 15:14:58','2022-07-23 20:03:32'),(148,'/admin/api/post','2','API',1,'API添加','2022-06-26 15:14:58','2022-07-23 20:03:24'),(149,'/admin/api/put','2','API',3,'API修改','2022-06-26 15:14:58','2022-08-30 22:12:36'),(151,'/admin/dict/path','1','字典',5,'字典页面','2022-06-26 15:27:04','2022-07-10 13:47:37'),(152,'/admin/dict/path/add','1','字典',5,'字典修改页面','2022-06-26 15:27:04','2022-07-23 20:07:26'),(153,'/admin/dict/path/edit/:id','1','字典',5,'字典修改页面','2022-06-26 15:27:04','2022-07-23 20:07:19'),(154,'/admin/dict/path/del/:id','1','字典',2,'字典删除','2022-06-26 15:27:04','2022-07-23 20:07:05'),(155,'/admin/dict/post','2','字典',1,'字典添加','2022-06-26 15:27:04','2022-07-23 20:06:54'),(156,'/admin/dict/put','2','字典',3,'字典修改','2022-06-26 15:27:04','2022-07-23 20:06:46'),(157,'/admin/operationLog/path','1','操作日志',5,'操作日志页面','2022-06-26 20:30:22','2022-07-12 10:44:09'),(160,'/admin/operationLog/path/del/:id','1','操作日志',2,'操作日志删除','2022-06-26 20:30:22','2022-07-23 20:07:45'),(163,'/admin/admin/path','1','管理员',5,'管理员页面','2022-06-27 16:21:07','2022-07-10 13:46:37'),(164,'/admin/admin/path/add','1','管理员',5,'管理员添加页面','2022-06-27 16:21:07','2022-07-10 13:46:32'),(165,'/admin/admin/path/edit/:id','1','管理员',5,'管理员修改页面','2022-06-27 16:21:07','2022-07-10 13:46:27'),(166,'/admin/admin/path/del/:id','1','管理员',2,'管理员删除','2022-06-27 16:21:07','2022-07-23 20:11:41'),(167,'/admin/admin/post','2','管理员',1,'管理员添加','2022-06-27 16:21:07','2022-07-23 20:11:34'),(168,'/admin/admin/put','2','管理员',3,'管理员修改','2022-06-27 16:21:07','2022-07-23 20:11:28'),(169,'/admin/role/path','1','角色',5,'角色页面','2022-06-27 18:27:46','2022-07-10 13:45:38'),(170,'/admin/role/path/add','1','角色',5,'角色添加页面','2022-06-27 18:27:46','2022-07-10 13:45:34'),(171,'/admin/role/path/edit/:id','1','角色',5,'角色修改页面','2022-06-27 18:27:46','2022-07-10 13:45:30'),(172,'/admin/role/path/del/:id','1','角色',2,'角色删除操作','2022-06-27 18:27:46','2022-07-20 20:18:03'),(173,'/admin/role/post','2','角色',1,'角色添加','2022-06-27 18:27:46','2022-07-23 20:14:31'),(174,'/admin/role/put','2','角色',3,'角色修改','2022-06-27 18:27:46','2022-07-23 20:14:26'),(175,'/admin/roleMenu/path/add','1','角色',5,'角色菜单添加页面','2022-06-27 18:38:32','2022-07-10 13:45:09'),(177,'/admin/roleMenu/path/del/:id','1','角色',2,'角色菜单删除操作','2022-06-27 18:38:32','2022-07-20 20:17:57'),(179,'/admin/roleApi/path/add','1','角色',5,'角色API添加页面','2022-06-27 19:10:03','2022-07-23 20:14:11'),(181,'/admin/roleApi/path/del/:id','1','角色',2,'角色API删除操作','2022-06-27 19:10:03','2022-07-23 20:14:05'),(183,'/admin/file/path/add','1','文件',5,'文件添加页面','2022-06-27 19:54:22','2022-07-10 13:42:38'),(184,'/admin/file/path/edit/:id','1','文件',5,'文件修改页面','2022-06-27 19:54:22','2022-07-10 13:42:26'),(185,'/admin/file/path/del/:id','1','文件',2,'文件删除操作','2022-06-27 19:54:22','2022-07-20 20:17:19'),(186,'/admin/file/post','2','文件',1,'文件添加','2022-06-27 19:54:22','2022-07-23 20:07:54'),(199,'/admin/operationLog/clear','1','操作日志',2,'操作日志清空','2022-07-10 12:39:19','2022-07-23 20:07:36'),(202,'/admin/adminLoginLog/path','1','登陆日志',5,'登陆日志页面','2022-07-11 19:06:26','2022-07-12 10:44:34'),(205,'/admin/adminLoginLog/path/del/:id','1','登陆日志',2,'登陆日志删除操作','2022-07-11 19:06:26','2022-07-20 20:17:34'),(224,'/admin/menu','1','菜单',5,'菜单页面','2022-07-23 19:34:26','2022-12-30 19:59:30'),(225,'/admin/menu','2','菜单',3,'菜单修改','2022-07-23 19:35:30','2022-12-30 19:59:40'),(226,'/admin/menu/post','2','菜单',1,'菜单添加','2022-07-23 19:35:46','2022-07-23 19:37:07'),(227,'/admin/menu/path/del/:id','1','菜单',2,'菜单删除','2022-07-23 19:36:08','2022-07-23 19:37:30'),(228,'/admin/menu/path/edit/:id','1','菜单',5,'菜单修改页面','2022-07-23 19:36:36','2022-07-23 19:36:36'),(230,'/admin/menu/path/add','1','菜单',5,'菜单添加页面','2022-07-23 19:37:57','2022-07-23 19:37:57'),(231,'/admin/admin/updateUname','3','管理员',3,'管理员用户名修改','2022-07-23 19:40:42','2022-07-23 19:42:27'),(232,'/admin/admin/updatePwdWithoutOldPwd','3','管理员',3,'管理员密码修改','2022-07-23 19:57:23','2022-07-23 19:57:23'),(233,'/admin/roleApi/clear','1','角色',2,'角色API清空','2022-07-23 20:02:23','2022-07-23 20:02:34'),(242,'/user','1','user',5,'用户列表页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(243,'/user/add','1','user',5,'用户列表添加页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(244,'/user/edit/:id','1','user',5,'用户列表修改页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(245,'/user/del/:id','1','user',2,'用户列表删除操作','2022-09-01 14:36:51','2022-09-01 14:36:51'),(247,'/userLoginLog','1','用户登录日志',5,'用户登录日志页面','2022-09-03 14:47:17','2022-09-03 14:47:17'),(250,'/userLoginLog/del/:id','1','用户登录日志',2,'用户登录日志删除操作','2022-09-03 14:47:17','2022-09-03 14:47:17'),(252,'/admin/user/updateUname','3','用户',3,'修改用户名','2022-09-03 23:19:07','2022-09-03 23:19:07'),(253,'/gold','1','金币',5,'用户金币页面','2022-09-04 20:32:44','2022-09-04 20:32:44'),(255,'/gold/edit/:id','1','金币',5,'用户金币修改页面','2022-09-04 20:32:44','2022-09-04 20:32:44'),(256,'/gold/del/:id','1','金币',2,'用户金币删除操作','2022-09-04 20:32:44','2022-09-04 20:32:44'),(257,'/gold','2','金币',1,'添加用户金币','2022-09-04 20:32:44','2022-09-04 20:32:44'),(258,'/topUpCategory','1','充值类型',5,'充值类型页面','2022-09-05 07:02:32','2022-09-05 07:02:32'),(259,'/topUpCategory/add','1','充值类型',5,'充值类型添加页面','2022-09-05 07:02:32','2022-09-05 07:02:32'),(260,'/topUpCategory/edit/:id','1','充值类型',5,'充值类型修改页面','2022-09-05 07:02:32','2022-09-05 07:02:32'),(261,'/topUpCategory/del/:id','1','充值类型',2,'充值类型删除操作','2022-09-05 07:02:32','2022-09-05 07:02:32'),(262,'/topUpCategory','2','充值类型',1,'添加充值类型','2022-09-05 07:02:32','2022-09-05 07:02:32'),(263,'/goldChangeLog','1','账变纪录',5,'账变记录页面','2022-09-05 11:09:03','2022-09-05 11:09:03'),(265,'/goldChangeLog/edit/:id','1','账变纪录',5,'账变记录修改页面','2022-09-05 11:09:03','2022-09-05 11:09:03'),(266,'/goldChangeLog/del/:id','1','账变纪录',2,'账变记录删除操作','2022-09-05 11:09:03','2022-09-05 11:09:03'),(267,'/goldChangeLog','2','账变纪录',1,'添加账变记录','2022-09-05 11:09:03','2022-09-05 11:09:03'),(268,'/goldStatisticsLog','1','账变统计',5,'账变统计页面','2022-09-05 11:15:18','2022-09-05 11:15:18'),(271,'/goldStatisticsLog/del/:id','1','账变统计',2,'账变统计删除操作','2022-09-05 11:15:18','2022-09-05 11:15:18'),(273,'/admin/gold/topUpByAdmin','2','金币',1,'人工充值','2022-09-05 13:27:12','2022-09-05 13:27:12'),(274,'/admin/gold/deductByAdmin','2','金币',1,'扣除用户金币','2022-09-05 14:10:52','2022-09-05 14:10:52'),(277,'/admin/goldStatisticsLog/clear','1','账变统计',2,'清除账变统计','2022-09-05 20:22:35','2022-09-05 20:22:35'),(278,'/admin/goldChangeLog/clear','1','账变纪录',2,'清空账变记录','2022-09-05 22:43:12','2022-09-05 22:43:12'),(279,'admin/goldChangeLog/clear','1','账变纪录',2,'清空','2022-09-05 22:49:22','2022-09-05 22:50:24'),(280,'/admin/goldChangeLog/clear','1','登陆日志',2,'清空','2022-09-05 22:50:49','2022-09-05 22:51:04'),(356,'/walletTopUpApplication','1','wallet_top_up_application',5,'充值订单页面','2022-12-26 19:35:13','2022-12-26 19:35:13'),(357,'/walletTopUpApplication/add','1','wallet_top_up_application',5,'充值订单添加页面','2022-12-26 19:35:13','2022-12-26 19:35:13'),(358,'/walletTopUpApplication/edit/:id','1','wallet_top_up_application',5,'充值订单修改页面','2022-12-26 19:35:13','2022-12-26 19:35:13'),(359,'/walletTopUpApplication/del/:id','1','wallet_top_up_application',2,'充值订单删除操作','2022-12-26 19:35:13','2022-12-26 19:35:13'),(360,'/walletTopUpApplication','2','wallet_top_up_application',1,'添加充值订单','2022-12-26 19:35:13','2022-12-26 19:35:13'),(361,'/adminMessage','1','admin_message',5,'管理员信息页面','2022-12-27 13:59:40','2022-12-27 13:59:40'),(362,'/adminMessage/add','1','admin_message',5,'管理员信息添加页面','2022-12-27 13:59:40','2022-12-27 13:59:40'),(363,'/adminMessage/edit/:id','1','admin_message',5,'管理员信息修改页面','2022-12-27 13:59:40','2022-12-27 13:59:40'),(364,'/adminMessage/del/:id','1','admin_message',2,'管理员信息删除操作','2022-12-27 13:59:40','2022-12-27 13:59:40'),(365,'/adminMessage','2','admin_message',1,'添加管理员信息','2022-12-27 13:59:40','2022-12-27 13:59:40'),(366,'/banner','1','banner',5,'banner图页面','2022-12-29 00:46:37','2022-12-29 00:46:37'),(367,'/banner/add','1','banner',5,'banner图添加页面','2022-12-29 00:46:37','2022-12-29 00:46:37'),(368,'/banner/edit/:id','1','banner',5,'banner图修改页面','2022-12-29 00:46:37','2022-12-29 00:46:37'),(369,'/banner/del/:id','1','banner',2,'banner图删除操作','2022-12-29 00:46:37','2022-12-29 00:46:37'),(370,'/banner','2','banner',1,'添加banner图','2022-12-29 00:46:37','2022-12-29 00:46:37');
/*!40000 ALTER TABLE `s_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_dict`
--

DROP TABLE IF EXISTS `s_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `k` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `v` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'sys',
  `status` int DEFAULT NULL,
  `type` int NOT NULL DEFAULT '1' COMMENT '1 文本，2 img',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `title`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'API分组选项','api_group','菜单\r\nAPI\r\n角色\r\n管理员\r\n字典\r\n文件\r\n操作日志\r\n登陆日志\r\n备忘录\r\n管理员消息\r\n用户\r\n用户登录日志\r\n金币\r\n充值类型\r\n账变纪录\r\n账变统计\n书籍分类\n书籍作者\n书籍\n章节\n书籍列表\n百科\n语录\n外链\n阅读\n阅读分类\n英语文章\n英语关卡\n英语文章段落\n英语文章单词\n问与答\n问与答分类\nquestion\nwallet_top_up_application\nadmin_message\nbanner','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-12-29 00:46:37'),(22,'登陆音乐列表','music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-07-14 15:47:17'),(42,'系统白名单','white_ips','','多个ip用小写逗号隔开','1',1,1,'2022-07-23 19:04:44','2022-09-03 20:42:02'),(44,'欢迎语','great','hello',NULL,'2',1,1,'2022-12-28 20:25:41','2022-12-28 20:25:41');
/*!40000 ALTER TABLE `s_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_file`
--

DROP TABLE IF EXISTS `s_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (26,'1/2022/03/BYFY4d.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:50:13'),(27,'1/2022/03/FdI4Yw.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:51:51'),(29,'1/2022/03/mAMoWX.png',2,1,'2022-03-27 20:19:07','2022-07-09 16:49:57'),(30,'1/2022/03/2S41in.png',2,1,'2022-03-28 15:32:00','2022-07-11 18:42:26'),(31,'1/2022/03/IdGUqj.png',2,1,'2022-03-28 15:36:45','2022-07-09 16:49:40'),(32,'1/2022/03/5Eoxb1.png',2,1,'2022-03-28 15:40:17','2022-07-09 16:49:33'),(77,'2/2022/07/CQVqgn.webp',2,1,'2022-07-03 12:44:29','2022-07-03 12:44:29'),(78,'2/2022/07/qMBDps.png',2,1,'2022-07-03 12:49:10','2022-07-03 12:49:10'),(79,'2/2022/07/lSCC0m.webp',2,1,'2022-07-03 13:00:15','2022-07-03 13:00:15'),(80,'2/2022/07/SHf1y4.webp',2,1,'2022-07-03 18:38:21','2022-07-15 23:55:13'),(81,'1/2022/07/IJoBIZ.png',1,1,'2022-07-13 18:25:37','2022-08-31 14:14:58'),(87,'icon/icon01.png',1,1,'2022-12-25 23:13:53','2022-12-25 23:13:53'),(88,'icon/icon02.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(103,'icon/icon03.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(104,'icon/icon04.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(105,'icon/icon05.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(106,'icon/icon06.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(107,'icon/icon07.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(108,'icon/icon08.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(109,'icon/icon09.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(110,'icon/icon10.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(111,'icon/icon11.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(112,'icon/icon12.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(113,'icon/icon13.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(114,'icon/icon14.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(115,'icon/icon15.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(116,'icon/icon16.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(117,'icon/icon17.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(118,'icon/icon18.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(119,'icon/icon19.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(120,'icon/icon20.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(121,'icon/icon21.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(122,'icon/icon22.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(123,'icon/icon23.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(124,'icon/icon24.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(125,'icon/icon25.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(126,'icon/icon26.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(127,'icon/icon27.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(128,'icon/icon28.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(129,'icon/icon29.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(130,'icon/icon30.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(131,'icon/icon31.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(132,'icon/icon32.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(133,'icon/icon33.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(134,'icon/icon34.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(135,'icon/icon35.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(136,'icon/icon36.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(137,'icon/icon37.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(138,'icon/icon38.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(139,'icon/icon39.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(140,'icon/icon40.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(141,'icon/icon41.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(142,'icon/icon42.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(143,'icon/icon43.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(144,'icon/icon44.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(145,'icon/icon45.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(146,'icon/icon46.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(147,'icon/icon47.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(148,'icon/icon48.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(149,'icon/icon49.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(150,'icon/icon50.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50');
/*!40000 ALTER TABLE `s_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_menu`
--

DROP TABLE IF EXISTS `s_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL,
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `bg_img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sort` decimal(7,2) NOT NULL DEFAULT '0.00',
  `type` int NOT NULL DEFAULT '1' COMMENT '1normal 2group',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `file_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-12-30 12:02:12'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/admin/menu',1.10,1,'这里是菜单页面','',1,'2022-02-16 11:14:13','2022-12-30 12:02:18'),(3,1,'1/2022/03/IdGUqj.png','','角色','/admin/role',1.30,1,'角色权限管理，在这里可以创建新的角色','',1,'2022-03-04 08:57:14','2022-09-06 06:12:36'),(4,1,'1/2022/03/BYFY4d.gif','','API','/admin/api',1.20,1,'系统所有的操作api在此','',1,'2022-07-03 06:25:52','2022-09-06 06:12:36'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/admin',1.40,1,'','',1,'2022-03-08 07:45:04','2022-09-06 06:12:36'),(28,1,'1/2022/03/mAMoWX.png','','字典','/admin/dict',1.50,1,'字典页面','',1,'2022-03-08 07:45:04','2022-09-06 06:12:36'),(30,1,'1/2022/03/2S41in.png','','文件','/admin/file',1.60,1,'','',1,'2022-03-08 08:05:30','2022-09-06 06:12:36'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/admin/operationLog',1.80,1,'','',1,'2022-06-13 11:59:57','2022-09-06 06:12:36'),(139,1,'','','登录日志','/admin/adminLoginLog',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2022-09-06 06:12:36'),(156,1,'','','代码生成','/admin/gen',1.92,1,'','',1,'2022-09-01 04:44:11','2022-12-27 06:01:51'),(171,-1,'','','用户','',5.00,2,'','',1,'2022-09-01 14:24:18','2022-12-28 16:49:42'),(172,171,'','','用户列表','/admin/user',5.10,1,'','',1,'2022-09-01 14:24:18','2022-12-28 16:49:42'),(173,171,'','','登录日志','/admin/userLoginLog',5.20,1,'','',1,'2022-09-03 06:47:17','2022-12-28 16:49:42'),(174,171,'','','钱包','/admin/wallet',5.30,1,'','',1,'2022-09-04 12:32:44','2022-12-28 16:49:42'),(175,-1,'','','配置','',2.00,2,'','',1,'2022-09-04 23:02:32','2022-09-06 06:12:29'),(176,175,'','','账变类型','/admin/walletChangeType',2.10,1,'','',1,'2022-09-04 23:02:32','2022-12-24 11:47:26'),(177,171,'','','账变记录','/admin/walletChangeLog',5.40,1,'','',1,'2022-09-05 03:09:03','2022-12-28 16:49:42'),(178,-1,'','','统计','',4.00,2,'','',1,'2022-09-05 03:15:18','2022-12-28 16:49:59'),(179,178,'','','账变统计','/admin/walletStatisticsLog',4.10,1,'','',1,'2022-09-05 03:15:18','2022-12-28 16:49:59'),(180,178,'','','账变报表','/admin/walletReport',4.20,1,'','',1,'2022-09-06 06:14:24','2022-12-28 16:49:59'),(202,171,'','','充值订单','/admin/walletTopUpApplication',5.50,1,'','',1,'2022-12-26 11:35:13','2022-12-28 16:49:42'),(203,1,'','','管理员信息','/admin/adminMessage',1.91,1,'','',1,'2022-12-27 05:59:40','2022-12-27 06:01:43'),(204,-1,'','','通用','',3.00,2,'','',1,'2022-12-28 16:46:37','2022-12-28 16:48:10'),(205,204,'','','Banner图','/admin/banner',3.10,1,'','',1,'2022-12-28 16:46:37','2022-12-28 16:50:52');
/*!40000 ALTER TABLE `s_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_operation_log`
--

DROP TABLE IF EXISTS `s_operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_operation_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `uri` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `use_time` int NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_operation_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2494 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (2467,42,'http://localhost:2033/admin/operationLog/clear','','GET','/admin/operationLog/clear','::1',3,'2022-12-28 16:11:11'),(2468,42,'http://localhost:2033/admin/walletStatisticsLog/del/2?','','GET','/admin/walletStatisticsLog/del/:id','::1',2,'2022-12-28 16:15:20'),(2469,42,'map[name:超级管理员 rid:1]','','POST','/admin/roleMenu/post','::1',1,'2022-12-29 00:47:13'),(2470,42,'map[mid:[204 205] name:超级管理员 rid:1]','','POST','/admin/roleMenu/post','::1',5,'2022-12-29 00:47:18'),(2471,42,'map[id:204 sort:3]','','PUT','/admin/menu/setGroupSort','::1',8,'2022-12-29 00:48:10'),(2472,42,'map[id:199 sort:11]','','PUT','/admin/menu/setGroupSort','::1',9,'2022-12-29 00:49:01'),(2473,42,'map[id:194 sort:10]','','PUT','/admin/menu/setGroupSort','::1',12,'2022-12-29 00:49:11'),(2474,42,'map[id:191 sort:9]','','PUT','/admin/menu/setGroupSort','::1',10,'2022-12-29 00:49:18'),(2475,42,'map[id:188 sort:8]','','PUT','/admin/menu/setGroupSort','::1',8,'2022-12-29 00:49:25'),(2476,42,'map[id:186 sort:7]','','PUT','/admin/menu/setGroupSort','::1',9,'2022-12-29 00:49:30'),(2477,42,'map[id:181 sort:6]','','PUT','/admin/menu/setGroupSort','::1',17,'2022-12-29 00:49:35'),(2478,42,'map[id:171 sort:5]','','PUT','/admin/menu/setGroupSort','::1',18,'2022-12-29 00:49:42'),(2479,42,'map[id:178 sort:4]','','PUT','/admin/menu/setGroupSort','::1',12,'2022-12-29 00:49:59'),(2480,42,'map[bg_img: desc: icon: id:205 name:Banner图 path:/admin/banner pid:204 sort:3.1 status:1 type:1]','','POST','/admin/menu/put','::1',2,'2022-12-29 00:50:52'),(2481,42,'map[desc: image:icon/icon43.png link:https://google.com sort:0 status:1 title:test]','','POST','/admin/banner/post','::1',16,'2022-12-29 00:55:26'),(2482,42,'map[desc: image: link: sort:0 status:1 title:]','','POST','/admin/banner/post','::1',6,'2022-12-29 00:56:24'),(2483,42,'map[desc: image: link: sort:0 status:1 title:]','','POST','/admin/banner/post','::1',5,'2022-12-29 00:58:03'),(2484,42,'map[created_at:2022-12-29 00:58:03 desc: id:1 image:icon/icon43.png link: sort:0 status:1 title:]','','POST','/admin/banner/put','::1',2,'2022-12-29 00:58:17'),(2485,42,'http://localhost:2033/admin/banner/del/1?','','GET','/admin/banner/del/:id','::1',5,'2022-12-29 00:58:27'),(2486,42,'map[desc: image:icon/icon43.png link:https://google.com sort:0 status:1 title:test]','','POST','/admin/banner/post','::1',3,'2022-12-29 00:58:38'),(2487,42,'http://localhost:2033/admin/file/del/151?','','GET','/admin/file/del/:id','::1',4,'2022-12-29 01:35:51'),(2488,42,'map[desc:菜单页面 group:菜单 id:224 method:1 type:5 url:/admin/menu]','','POST','/admin/api/put','::1',4,'2022-12-30 19:59:30'),(2489,42,'map[desc:菜单修改 group:菜单 id:225 method:2 type:3 url:/admin/menu]','','POST','/admin/api/put','::1',4,'2022-12-30 19:59:40'),(2490,42,'map[bg_img: desc: icon: id:1 name:系统 path: pid:-1 sort:1 status:1 type:2]','','POST','/admin/menu/put','::1',3,'2022-12-30 20:02:12'),(2491,42,'map[bg_img: desc:这里是菜单页面 icon:1/2022/03/FdI4Yw.gif id:2 name:菜单 path:/admin/menu pid:1 sort:1.1 status:1 type:1]','','POST','/admin/menu/put','::1',4,'2022-12-30 20:02:18'),(2492,42,'map[desc:3 group:1 k:3 status:1 title:3 type:1 v:]','','POST','/admin/dict/add','::1',6,'2022-12-30 20:06:03'),(2493,42,'http://localhost:2033/admin/dict/del/45?','','GET','/admin/dict/del/:id','::1',1,'2022-12-30 20:06:07');
/*!40000 ALTER TABLE `s_operation_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role`
--

DROP TABLE IF EXISTS `s_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'超级管理员','2022-02-16 11:12:41','2022-09-02 12:22:24'),(22,'系统管理员','2022-07-23 08:45:05','2022-09-02 12:22:31');
/*!40000 ALTER TABLE `s_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role_api`
--

DROP TABLE IF EXISTS `s_role_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role_api` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rid` int DEFAULT NULL,
  `aid` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_role_api_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
/*!40000 ALTER TABLE `s_role_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role_menu`
--

DROP TABLE IF EXISTS `s_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rid` int DEFAULT NULL,
  `mid` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `rid` (`rid`),
  KEY `mid` (`mid`),
  CONSTRAINT `s_role_menu_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `s_role_menu_ibfk_2` FOREIGN KEY (`mid`) REFERENCES `s_menu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (197,1,1),(198,1,2),(199,1,3),(201,1,5),(202,1,28),(203,1,30),(204,1,78),(207,1,139),(208,1,156),(209,1,171),(210,1,172),(211,1,4),(212,1,173),(213,1,174),(214,1,175),(215,1,176),(216,1,177),(217,1,178),(218,1,179),(219,1,180),(241,1,202),(242,1,203),(243,1,204),(244,1,205);
/*!40000 ALTER TABLE `s_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_user`
--

DROP TABLE IF EXISTS `u_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pass` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `join_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '注册IP',
  `device` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备名称',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `pass_error_count` tinyint unsigned DEFAULT '0' COMMENT '密码错误次数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `uname_2` (`uname`),
  KEY `join_ip` (`join_ip`),
  KEY `status` (`status`),
  KEY `phone` (`phone`),
  KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` (`id`, `uname`, `pass`, `nickname`, `icon`, `summary`, `desc`, `join_ip`, `device`, `phone`, `email`, `status`, `pass_error_count`, `created_at`, `updated_at`) VALUES (4,'freekey','$2a$10$DyEqv63stEpvh.1QJS31N.hPzk5I62cJGYEfcRaNbpK8QAHcDblc2','freekey admin','icon/icon02.png','','','127.0.0.1','','','',1,0,'2022-09-04 16:13:08','2022-12-26 13:50:47'),(10,'ciel2','$2a$10$q.zZ/t5/6tWNZ4fIk7Zki.8IVau7h3mwZ3zjAopiAfVyEyOAOtXTu','ciel2','icon/icon01.png','','','78.138.12.130','','','',1,0,'2022-11-02 21:30:28','2022-12-26 13:50:39'),(16,'user33','$2a$10$m4EXhA9LI6l2XksNl/yeGe2/gDMD.runELtlmZsXzCXicG98t/.0G','user33','icon/icon43.png','','','127.0.0.1','','','',1,0,'2022-12-26 13:14:51','2022-12-27 20:08:11');
/*!40000 ALTER TABLE `u_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_user_login_log`
--

DROP TABLE IF EXISTS `u_user_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned NOT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user_login_log`
--

LOCK TABLES `u_user_login_log` WRITE;
/*!40000 ALTER TABLE `u_user_login_log` DISABLE KEYS */;
INSERT INTO `u_user_login_log` (`id`, `uid`, `ip`, `created_at`, `updated_at`) VALUES (45,16,'127.0.0.1','2022-12-28 19:26:13','2022-12-28 19:26:13');
/*!40000 ALTER TABLE `u_user_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_wallet`
--

DROP TABLE IF EXISTS `u_wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_wallet` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned NOT NULL,
  `balance` decimal(12,2) unsigned NOT NULL DEFAULT '0.00',
  `pass` varchar(64) DEFAULT NULL,
  `pass_err_count` tinyint unsigned DEFAULT '0' COMMENT '密码输错次数',
  `desc` text,
  `status` tinyint unsigned DEFAULT '1' COMMENT '金库状态 0 设置密码 1正常,2 锁定',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`),
  KEY `uid_2` (`uid`),
  KEY `balance` (`balance`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户金库';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet`
--

LOCK TABLES `u_wallet` WRITE;
/*!40000 ALTER TABLE `u_wallet` DISABLE KEYS */;
INSERT INTO `u_wallet` (`id`, `uid`, `balance`, `pass`, `pass_err_count`, `desc`, `status`, `created_at`, `updated_at`) VALUES (1,4,20308.00,'',0,'',0,'2022-09-04 16:13:08','2022-12-24 14:32:21'),(7,10,44.00,'',0,'',0,'2022-11-02 21:30:28','2022-12-25 16:45:47'),(13,16,10.00,'$2a$10$M.QBVqrZgLK7W.lE2OKWoeDWkua3ucSlYgGmaPBsfffNdukdmeRBu',0,'',1,'2022-12-26 13:14:51','2022-12-28 15:46:34');
/*!40000 ALTER TABLE `u_wallet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_wallet_change_log`
--

DROP TABLE IF EXISTS `u_wallet_change_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_wallet_change_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trans_id` varchar(64) NOT NULL,
  `uid` bigint unsigned NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `balance` decimal(12,2) NOT NULL,
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1人工充值,2支付宝充值,3微信充值,4paypal充值,5人工扣除',
  `desc` varchar(64) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `amount` (`amount`),
  KEY `balance` (`balance`),
  KEY `type` (`type`),
  KEY `desc` (`desc`),
  KEY `trans_id` (`trans_id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='账变记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_change_log`
--

LOCK TABLES `u_wallet_change_log` WRITE;
/*!40000 ALTER TABLE `u_wallet_change_log` DISABLE KEYS */;
INSERT INTO `u_wallet_change_log` (`id`, `trans_id`, `uid`, `amount`, `balance`, `type`, `desc`, `created_at`) VALUES (63,'1d0vq7rxoh0cp9uff5uvo7c200xyw9kl',10,33.00,33.00,1,'人工充值','2022-12-24 14:44:48'),(64,'1uj3ztpy9r0cp9vbwm117tk200h20u7o',10,-1.00,32.00,5,'人工扣除','2022-12-24 15:27:14'),(65,'1uj3ztpy9r0cp9vc7cvrt2840019805x',10,-20.00,12.00,5,'人工扣除','2022-12-24 15:27:37'),(66,'1uj3ztpyda0cp9vedk0pzug200j71jka',10,20.00,32.00,1,'人工充值','2022-12-24 15:30:27'),(67,'1uj3ztpyda0cp9vein12oh4400mm42jd',10,-20.00,12.00,5,'人工扣除3','2022-12-24 15:31:28'),(68,'kyped7015f7cparj927wkts200yg173l',10,33.00,45.00,2,'333','2022-12-25 16:41:25'),(69,'kyped7015k0cparmle8h8bk2006s4ots',10,-1.00,44.00,5,'人工扣除','2022-12-25 16:45:47'),(71,'MR-tK9lQIqj9SfSL',16,10.00,10.00,2,'','2022-12-28 15:46:34');
/*!40000 ALTER TABLE `u_wallet_change_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_wallet_change_type`
--

DROP TABLE IF EXISTS `u_wallet_change_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_wallet_change_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL,
  `sub_title` varchar(64) DEFAULT NULL,
  `type` tinyint unsigned DEFAULT '1' COMMENT '1 add; 2 reduce',
  `class` varchar(16) DEFAULT NULL,
  `desc` text,
  `status` tinyint unsigned DEFAULT '1',
  `count_status` tinyint DEFAULT '1' COMMENT 'Whether this field needs statistics, 1 true 2 false',
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='充值类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_change_type`
--

LOCK TABLES `u_wallet_change_type` WRITE;
/*!40000 ALTER TABLE `u_wallet_change_type` DISABLE KEYS */;
INSERT INTO `u_wallet_change_type` (`id`, `title`, `sub_title`, `type`, `class`, `desc`, `status`, `count_status`) VALUES (1,'支付宝充值',NULL,1,'tag-primary','',1,1),(2,'微信充值','',1,'tag-success','',1,1),(3,'银行卡充值',NULL,1,'tag-warning',NULL,1,1),(4,'PayPal充值','',1,'tag-info','',1,1),(5,'人工扣除','',2,'tag-danger','',1,1);
/*!40000 ALTER TABLE `u_wallet_change_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_wallet_statistics_log`
--

DROP TABLE IF EXISTS `u_wallet_statistics_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_wallet_statistics_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned NOT NULL,
  `t1` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t2` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t3` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t4` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t5` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t6` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t7` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t8` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t9` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t10` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t11` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t12` decimal(12,2) NOT NULL DEFAULT '0.00',
  `t13` decimal(12,2) NOT NULL DEFAULT '0.00',
  `created_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_statistics_log`
--

LOCK TABLES `u_wallet_statistics_log` WRITE;
/*!40000 ALTER TABLE `u_wallet_statistics_log` DISABLE KEYS */;
INSERT INTO `u_wallet_statistics_log` (`id`, `uid`, `t1`, `t2`, `t3`, `t4`, `t5`, `t6`, `t7`, `t8`, `t9`, `t10`, `t11`, `t12`, `t13`, `created_date`) VALUES (3,4,10063.00,33.00,33.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,'2022-10-31'),(4,10,53.00,0.00,0.00,0.00,74.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,'2022-12-24'),(5,4,0.00,0.00,0.00,0.00,33.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,'2022-12-24'),(6,10,0.00,33.00,0.00,0.00,1.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,'2022-12-25'),(8,16,0.00,10.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,0.00,'2022-12-28');
/*!40000 ALTER TABLE `u_wallet_statistics_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_wallet_top_up_application`
--

DROP TABLE IF EXISTS `u_wallet_top_up_application`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_wallet_top_up_application` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trans_id` varchar(64) DEFAULT NULL,
  `uid` bigint unsigned NOT NULL,
  `change_type` int unsigned NOT NULL,
  `money` decimal(8,2) unsigned DEFAULT NULL,
  `ip` varchar(64) NOT NULL,
  `description` varchar(64) DEFAULT NULL,
  `aid` bigint unsigned DEFAULT '0',
  `status` tinyint unsigned DEFAULT '1' COMMENT '1 wait 2 success 3 fail',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `trans_id` (`trans_id`),
  KEY `uid` (`uid`),
  KEY `status` (`status`),
  KEY `description` (`description`),
  KEY `ip` (`ip`),
  KEY `aid` (`aid`),
  KEY `change_type` (`change_type`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_top_up_application`
--

LOCK TABLES `u_wallet_top_up_application` WRITE;
/*!40000 ALTER TABLE `u_wallet_top_up_application` DISABLE KEYS */;
INSERT INTO `u_wallet_top_up_application` (`id`, `trans_id`, `uid`, `change_type`, `money`, `ip`, `description`, `aid`, `status`, `created_at`, `updated_at`) VALUES (9,'MR-tK9lQIqj9SfSL',16,2,10.00,'127.0.0.1','3',42,2,'2022-12-27 20:24:38','2022-12-28 15:51:53');
/*!40000 ALTER TABLE `u_wallet_top_up_application` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-30 20:52:01
