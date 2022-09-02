-- MySQL dump 10.13  Distrib 8.0.24, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: ciel
-- ------------------------------------------------------
-- Server version	8.0.24

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
-- Table structure for table `s_admin`
--

DROP TABLE IF EXISTS `s_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rid` int NOT NULL,
  `uname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `unread_msg_count` int DEFAULT '0',
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` VALUES (1,1,'ciel','$2a$10$OAp3RJVKv6WhAX3o.fY/A.R0jUOyzvtlfxpS3DgHtEVkLx/lY6b4.',0,1,'2022-03-08 08:59:33','2022-08-30 22:48:31'),(42,1,'admin','$2a$10$VCrmz3RzJmO2aSX2CQSqsunK59fkkc4otXVWzPLCxsTRqSn/ybNzC',0,1,'2022-07-02 11:28:52','2022-08-30 22:57:00');
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
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
INSERT INTO `s_admin_login_log` VALUES (74,42,'::1',NULL,1,'2022-08-31 15:22:31','2022-08-31 15:22:31'),(75,42,'::1',NULL,1,'2022-09-01 12:25:15','2022-09-01 12:25:15'),(76,42,'::1',NULL,1,'2022-09-01 12:26:00','2022-09-01 12:26:00'),(77,42,'::1',NULL,1,'2022-09-01 12:44:36','2022-09-01 12:44:36'),(78,42,'::1',NULL,1,'2022-09-01 12:46:03','2022-09-01 12:46:03'),(79,42,'::1',NULL,1,'2022-09-01 16:35:58','2022-09-01 16:35:58'),(80,42,'::1',NULL,1,'2022-09-01 16:38:11','2022-09-01 16:38:11'),(81,42,'::1',NULL,1,'2022-09-01 16:38:44','2022-09-01 16:38:44'),(82,42,'::1',NULL,1,'2022-09-01 19:23:53','2022-09-01 19:23:53'),(83,42,'::1',NULL,1,'2022-09-01 21:35:06','2022-09-01 21:35:06'),(84,42,'::1',NULL,1,'2022-09-01 21:36:55','2022-09-01 21:36:55'),(85,42,'::1',NULL,1,'2022-09-01 22:11:22','2022-09-01 22:11:22'),(86,42,'::1',NULL,1,'2022-09-01 22:14:01','2022-09-01 22:14:01'),(87,42,'::1',NULL,1,'2022-09-01 22:36:11','2022-09-01 22:36:11'),(88,42,'::1',NULL,1,'2022-09-01 23:15:35','2022-09-01 23:15:35'),(89,42,'::1',NULL,1,'2022-09-02 09:27:20','2022-09-02 09:27:20');
/*!40000 ALTER TABLE `s_admin_login_log` ENABLE KEYS */;
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
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=247 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` VALUES (56,'/admin/file/path','1','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(61,'/admin/file/put','2','文件','文件修改',1,'2022-06-12 14:10:23','2022-07-23 20:08:04'),(62,'/admin/roleApi/path','1','角色','角色禁用API页面',1,'2022-06-12 17:02:13','2022-07-23 20:15:56'),(65,'/admin/roleApi/:id','1','角色','删除角色禁用API',1,'2022-06-12 17:02:13','2022-07-20 20:18:08'),(66,'/admin/roleApi/post','2','角色','角色禁用API添加',1,'2022-06-12 17:02:13','2022-07-23 20:15:37'),(68,'/admin/roleMenu/path','1','角色','角色菜单页面',1,'2022-06-13 15:19:45','2022-07-23 19:33:45'),(71,'/admin/roleMenu/:id','1','角色','角色菜单删除',1,'2022-06-13 15:19:45','2022-07-23 20:14:58'),(72,'/admin/roleMenu/post','2','角色','角色菜单添加',1,'2022-06-13 15:19:45','2022-07-23 20:14:48'),(144,'/admin/api/path','1','API','API页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:54'),(145,'/admin/api/path/add','1','API','API添加页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:47'),(146,'/admin/api/path/edit/:id','1','API','API修改页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:40'),(147,'/admin/api/path/del/:id','1','API','API删除操作',1,'2022-06-26 15:14:58','2022-07-23 20:03:32'),(148,'/admin/api/post','2','API','API添加',1,'2022-06-26 15:14:58','2022-07-23 20:03:24'),(149,'/admin/api/put','2','API','API修改',1,'2022-06-26 15:14:58','2022-08-30 22:12:36'),(151,'/admin/dict/path','1','字典','字典页面',1,'2022-06-26 15:27:04','2022-07-10 13:47:37'),(152,'/admin/dict/path/add','1','字典','字典修改页面',1,'2022-06-26 15:27:04','2022-07-23 20:07:26'),(153,'/admin/dict/path/edit/:id','1','字典','字典修改页面',1,'2022-06-26 15:27:04','2022-07-23 20:07:19'),(154,'/admin/dict/path/del/:id','1','字典','字典删除',1,'2022-06-26 15:27:04','2022-07-23 20:07:05'),(155,'/admin/dict/post','2','字典','字典添加',1,'2022-06-26 15:27:04','2022-07-23 20:06:54'),(156,'/admin/dict/put','2','字典','字典修改',1,'2022-06-26 15:27:04','2022-07-23 20:06:46'),(157,'/admin/operationLog/path','1','操作日志','操作日志页面',1,'2022-06-26 20:30:22','2022-07-12 10:44:09'),(160,'/admin/operationLog/path/del/:id','1','操作日志','操作日志删除',1,'2022-06-26 20:30:22','2022-07-23 20:07:45'),(163,'/admin/admin/path','1','管理员','管理员页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:37'),(164,'/admin/admin/path/add','1','管理员','管理员添加页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:32'),(165,'/admin/admin/path/edit/:id','1','管理员','管理员修改页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:27'),(166,'/admin/admin/path/del/:id','1','管理员','管理员删除',1,'2022-06-27 16:21:07','2022-07-23 20:11:41'),(167,'/admin/admin/post','2','管理员','管理员添加',1,'2022-06-27 16:21:07','2022-07-23 20:11:34'),(168,'/admin/admin/put','2','管理员','管理员修改',1,'2022-06-27 16:21:07','2022-07-23 20:11:28'),(169,'/admin/role/path','1','角色','角色页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:38'),(170,'/admin/role/path/add','1','角色','角色添加页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:34'),(171,'/admin/role/path/edit/:id','1','角色','角色修改页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:30'),(172,'/admin/role/path/del/:id','1','角色','角色删除操作',1,'2022-06-27 18:27:46','2022-07-20 20:18:03'),(173,'/admin/role/post','2','角色','角色添加',1,'2022-06-27 18:27:46','2022-07-23 20:14:31'),(174,'/admin/role/put','2','角色','角色修改',1,'2022-06-27 18:27:46','2022-07-23 20:14:26'),(175,'/admin/roleMenu/path/add','1','角色','角色菜单添加页面',1,'2022-06-27 18:38:32','2022-07-10 13:45:09'),(177,'/admin/roleMenu/path/del/:id','1','角色','角色菜单删除操作',1,'2022-06-27 18:38:32','2022-07-20 20:17:57'),(179,'/admin/roleApi/path/add','1','角色','角色API添加页面',1,'2022-06-27 19:10:03','2022-07-23 20:14:11'),(181,'/admin/roleApi/path/del/:id','1','角色','角色API删除操作',1,'2022-06-27 19:10:03','2022-07-23 20:14:05'),(183,'/admin/file/path/add','1','文件','文件添加页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:38'),(184,'/admin/file/path/edit/:id','1','文件','文件修改页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:26'),(185,'/admin/file/path/del/:id','1','文件','文件删除操作',1,'2022-06-27 19:54:22','2022-07-20 20:17:19'),(186,'/admin/file/post','2','文件','文件添加',1,'2022-06-27 19:54:22','2022-07-23 20:07:54'),(193,'/admin/node/path','1','备忘录','备忘录页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:37'),(194,'/admin/node/path/add','1','备忘录','备忘录添加页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:26'),(195,'/admin/node/path/edit/:id','1','备忘录','备忘录修改页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:17'),(196,'/admin/node/path/del/:id','1','备忘录','备忘录删除',1,'2022-07-05 19:01:03','2022-07-23 20:06:06'),(197,'/admin/node/post','2','备忘录','备忘录添加',1,'2022-07-05 19:01:03','2022-07-23 20:05:59'),(198,'/admin/node/put','2','备忘录','备忘录修改',1,'2022-07-05 19:01:03','2022-07-23 20:05:48'),(199,'/admin/operationLog/clear','1','操作日志','操作日志清空',1,'2022-07-10 12:39:19','2022-07-23 20:07:36'),(202,'/admin/adminLoginLog/path','1','登陆日志','登陆日志页面',1,'2022-07-11 19:06:26','2022-07-12 10:44:34'),(205,'/admin/adminLoginLog/path/del/:id','1','登陆日志','登陆日志删除操作',1,'2022-07-11 19:06:26','2022-07-20 20:17:34'),(208,'/admin/adminLoginLog/claer','1','登陆日志','登陆日志清空',1,'2022-07-14 09:49:26','2022-07-23 20:11:16'),(218,'/admin/adminMessage/path','1','管理员消息','消息页面',1,'2022-07-22 20:14:53','2022-07-23 20:12:37'),(220,'/admin/adminMessage/path/edit/:id','1','管理员消息','信息修改页面',1,'2022-07-22 20:14:53','2022-07-23 20:04:18'),(221,'/admin/adminMessage/path/del/:id','1','管理员消息','信息删除操作',1,'2022-07-22 20:14:53','2022-07-23 20:04:08'),(222,'/admin/adminMessage/post','2','管理员消息','消息发送',1,'2022-07-22 20:14:53','2022-07-23 20:12:25'),(224,'/admin/menu/path','1','菜单','菜单页面',1,'2022-07-23 19:34:26','2022-07-23 19:34:26'),(225,'/admin/menu/put','2','菜单','菜单修改',1,'2022-07-23 19:35:30','2022-07-23 19:37:19'),(226,'/admin/menu/post','2','菜单','菜单添加',1,'2022-07-23 19:35:46','2022-07-23 19:37:07'),(227,'/admin/menu/path/del/:id','1','菜单','菜单删除',1,'2022-07-23 19:36:08','2022-07-23 19:37:30'),(228,'/admin/menu/path/edit/:id','1','菜单','菜单修改页面',1,'2022-07-23 19:36:36','2022-07-23 19:36:36'),(230,'/admin/menu/path/add','1','菜单','菜单添加页面',1,'2022-07-23 19:37:57','2022-07-23 19:37:57'),(231,'/admin/admin/updateUname','3','管理员','管理员用户名修改',1,'2022-07-23 19:40:42','2022-07-23 19:42:27'),(232,'/admin/admin/updatePwdWithoutOldPwd','3','管理员','管理员密码修改',1,'2022-07-23 19:57:23','2022-07-23 19:57:23'),(233,'/admin/roleApi/clear','1','角色','角色API清空',1,'2022-07-23 20:02:23','2022-07-23 20:02:34'),(242,'/user','1','user','用户列表页面',1,'2022-09-01 14:36:51','2022-09-01 14:36:51'),(243,'/user/add','1','user','用户列表添加页面',1,'2022-09-01 14:36:51','2022-09-01 14:36:51'),(244,'/user/edit/:id','1','user','用户列表修改页面',1,'2022-09-01 14:36:51','2022-09-01 14:36:51'),(245,'/user/del/:id','1','user','用户列表删除操作',1,'2022-09-01 14:36:51','2022-09-01 14:36:51'),(246,'/user','2','user','添加用户列表',1,'2022-09-01 14:36:51','2022-09-01 14:36:51');
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
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` VALUES (11,'api_group','菜单\r\nAPI\r\n角色\r\n管理员\r\n字典\r\n文件\r\n操作日志\r\n登陆日志\r\n备忘录\r\n管理员消息\n用户','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-09-01 22:24:18'),(22,'music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-07-14 15:47:17'),(33,'node-category','1.记事\r\n3.Mysql\r\n5.English\r\n6.Freekey\r\n8.Golang\r\n9.Idea\r\n10.js\r\n11.jquery\r\n12.Linux\r\n15.Nginx\r\n16.Errors\r\n17.Quotations','备忘录分类','1',1,1,'2022-07-07 20:18:58','2022-08-31 11:10:48'),(42,'white_ips','','系统白名单','1',1,1,'2022-07-23 19:04:44','2022-08-31 11:13:45');
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
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` VALUES (26,'1/2022/03/BYFY4d.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:50:13'),(27,'1/2022/03/FdI4Yw.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:51:51'),(29,'1/2022/03/mAMoWX.png',2,1,'2022-03-27 20:19:07','2022-07-09 16:49:57'),(30,'1/2022/03/2S41in.png',2,1,'2022-03-28 15:32:00','2022-07-11 18:42:26'),(31,'1/2022/03/IdGUqj.png',2,1,'2022-03-28 15:36:45','2022-07-09 16:49:40'),(32,'1/2022/03/5Eoxb1.png',2,1,'2022-03-28 15:40:17','2022-07-09 16:49:33'),(77,'2/2022/07/CQVqgn.webp',2,1,'2022-07-03 12:44:29','2022-07-03 12:44:29'),(78,'2/2022/07/qMBDps.png',2,1,'2022-07-03 12:49:10','2022-07-03 12:49:10'),(79,'2/2022/07/lSCC0m.webp',2,1,'2022-07-03 13:00:15','2022-07-03 13:00:15'),(80,'2/2022/07/SHf1y4.webp',2,1,'2022-07-03 18:38:21','2022-07-15 23:55:13'),(81,'1/2022/07/IJoBIZ.png',1,1,'2022-07-13 18:25:37','2022-08-31 14:14:58');
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
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-07-09 08:09:47'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/admin/menu',1.10,1,'这里是菜单页面','\"\"',1,'2022-02-16 11:14:13','2022-07-08 15:24:50'),(3,1,'1/2022/03/IdGUqj.png','','角色','/admin/role',1.30,1,'','\"\"',1,'2022-03-04 08:57:14','2022-07-09 08:08:31'),(4,1,'1/2022/03/BYFY4d.gif','','API','/admin/api',1.20,1,'','',1,'2022-07-03 06:25:52','2022-07-09 10:41:50'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/admin',1.40,1,'','',1,'2022-03-08 07:45:04','2022-07-09 08:10:07'),(28,1,'1/2022/03/mAMoWX.png',NULL,'字典','/admin/dict',1.50,1,'字典页面',NULL,1,'2022-03-08 07:45:04','2022-07-02 08:06:55'),(30,1,'1/2022/03/2S41in.png','','文件','/admin/file',1.60,1,'',NULL,1,'2022-03-08 08:05:30','2022-07-03 05:12:16'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/admin/operationLog',1.80,1,'','',1,'2022-06-13 11:59:57','2022-07-09 08:19:58'),(132,-1,'','','工具','',2.00,2,'','',1,'2022-07-03 06:25:52','2022-07-03 06:25:52'),(136,132,'2/2022/07/SHf1y4.webp','','站点导航','/admin/to/urls',2.20,1,'','/sys/tool/urls.html',1,'2022-07-03 06:25:52','2022-07-20 12:15:54'),(139,1,'','','登录日志','/admin/adminLoginLog',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2022-07-16 15:09:08'),(156,1,'','','代码生成','/admin/gen',1.91,1,'','',1,'2022-09-01 04:44:11','2022-09-01 13:36:44'),(171,-1,'','','用户','',3.00,2,'','',1,'2022-09-01 14:24:18','2022-09-01 14:24:18'),(172,171,'','','用户列表','/admin/user',3.10,1,'','',1,'2022-09-01 14:24:18','2022-09-01 14:24:18');
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
) ENGINE=InnoDB AUTO_INCREMENT=1150 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` VALUES (1107,42,'http://localhost:1211/admin/adminLoginLog/clear','','GET','/admin/adminLoginLog/clear','::1',9,'2022-08-31 14:31:03'),(1108,42,'map[bg_img: desc: icon: name:基础代码生成 path: pid:1 sort:1.91 status:1 type:2]','','POST','/admin/menu/post','::1',2,'2022-09-01 12:44:11'),(1109,42,'map[mid:[156] rid:1]','','POST','/admin/roleMenu/post','::1',24,'2022-09-01 12:44:29'),(1110,42,'map[bg_img: desc: icon: id:156 name:基础代码生成 path: pid:1 sort:1.91 status:1 type:1]','','POST','/admin/menu/put','::1',9,'2022-09-01 12:45:23'),(1111,42,'map[bg_img: desc: icon: id:156 name:基础代码生成 path:/admin/gen pid:1 sort:1.91 status:1 type:1]','','POST','/admin/menu/put','::1',3,'2022-09-01 12:45:43'),(1112,42,'http://localhost:1211/admin/menu/del/157?page=1&','','GET','/admin/menu/del/:id','::1',9,'2022-09-01 13:47:28'),(1113,42,'http://localhost:1211/admin/menu/del/159?','','GET','/admin/menu/del/:id','::1',10,'2022-09-01 13:59:43'),(1114,42,'http://localhost:1211/admin/menu/del/158?','','GET','/admin/menu/del/:id','::1',9,'2022-09-01 13:59:46'),(1115,42,'http://localhost:1211/admin/menu/del/162?','','GET','/admin/menu/del/:id','::1',10,'2022-09-01 14:30:52'),(1116,42,'http://localhost:1211/admin/menu/del/163?page=2&','','GET','/admin/menu/del/:id','::1',46,'2022-09-01 14:31:25'),(1117,42,'http://localhost:1211/admin/menu/del/164?page=2&','','GET','/admin/menu/del/:id','::1',3,'2022-09-01 14:31:46'),(1118,42,'http://localhost:1211/admin/menu/del/160?page=1&','','GET','/admin/menu/del/:id','::1',46,'2022-09-01 14:31:54'),(1119,42,'http://localhost:1211/admin/menu/del/165?page=2&','','GET','/admin/menu/del/:id','::1',10,'2022-09-01 14:32:03'),(1120,42,'http://localhost:1211/admin/menu/del/166?page=1&','','GET','/admin/menu/del/:id','::1',9,'2022-09-01 14:35:21'),(1121,42,'http://localhost:1211/admin/api/del/241?page=1&','','GET','/admin/api/del/:id','::1',9,'2022-09-01 14:36:21'),(1122,42,'http://localhost:1211/admin/api/del/240?page=1&','','GET','/admin/api/del/:id','::1',9,'2022-09-01 14:36:24'),(1123,42,'http://localhost:1211/admin/api/del/239?page=1&','','GET','/admin/api/del/:id','::1',3,'2022-09-01 14:36:25'),(1124,42,'http://localhost:1211/admin/api/del/238?page=1&','','GET','/admin/api/del/:id','::1',2,'2022-09-01 14:36:26'),(1125,42,'http://localhost:1211/admin/api/del/237?page=1&','','GET','/admin/api/del/:id','::1',2,'2022-09-01 14:36:28'),(1126,42,'http://localhost:1211/admin/menu/del/167?','','GET','/admin/menu/del/:id','::1',9,'2022-09-01 14:54:14'),(1127,42,'map[mid:[161 168] rid:1]','','POST','/admin/roleMenu/post','::1',10,'2022-09-01 16:35:51'),(1128,42,'http://localhost:1211/admin/menu/del/168?','','GET','/admin/menu/del/:id','::1',46,'2022-09-01 16:37:25'),(1129,42,'http://localhost:1211/admin/menu/del/169?','','GET','/admin/menu/del/:id','::1',10,'2022-09-01 16:37:42'),(1130,42,'map[mid:[170] rid:1]','','POST','/admin/roleMenu/post','::1',14,'2022-09-01 16:38:30'),(1131,42,'map[status:1]','','POST','/admin/user/post','::1',12,'2022-09-01 16:44:32'),(1132,42,'map[bg_img: desc: icon: id:170 name:用户列表 path:/admin/user pid:161 sort:3.1 status:1 type:1]','','POST','/admin/menu/put','::1',3,'2022-09-01 19:25:36'),(1133,42,'map[pass:33 status:1 uname:wer]','','POST','/admin/user/post','::1',3,'2022-09-01 20:12:17'),(1134,42,'http://localhost:1211/admin/user/del/1?','','GET','/admin/user/del/:id','::1',10,'2022-09-01 20:12:21'),(1135,42,'map[created_at:2022-09-01 20:12:17 id:2 pass:33 status:1 uname:wer3 updated_at:2022-09-01 20:12:17]','','POST','/admin/user/put','::1',3,'2022-09-01 20:23:03'),(1136,42,'map[created_at:2022-09-01 20:12:17 id:2 pass:33 status:2 uname:wer3 updated_at:2022-09-01 20:23:03]','','POST','/admin/user/put','::1',2,'2022-09-01 20:23:07'),(1137,42,'map[created_at:2022-09-01 20:12:17 id:2 pass:33 status:2 uname:wer3 updated_at:2022-09-01 20:23:07]','','POST','/admin/user/put','::1',46,'2022-09-01 20:23:26'),(1138,42,'map[created_at:2022-09-01 20:12:17 id:2 pass:3333 status:2 uname:wer3 updated_at:2022-09-01 20:23:26]','','POST','/admin/user/put','::1',2,'2022-09-01 20:23:29'),(1139,42,'map[pass:33 status:1 uname:33]','','POST','/admin/user/post','::1',3,'2022-09-01 20:26:57'),(1140,42,'map[created_at:2022-09-01 20:26:57 id:3 pass:33 status:2 uname:33 updated_at:2022-09-01 20:26:57]','','POST','/admin/user/put','::1',46,'2022-09-01 20:27:04'),(1141,42,'map[created_at:2022-09-01 20:26:57 id:3 pass:33 status:1 uname:33 updated_at:2022-09-01 20:27:04]','','POST','/admin/user/put','::1',2,'2022-09-01 20:27:12'),(1142,42,'map[bg_img: desc: icon: id:156 name:代码生成 path:/admin/gen pid:1 sort:1.91 status:1 type:1]','','POST','/admin/menu/put','::1',2,'2022-09-01 21:36:44'),(1143,42,'map[created_at:2022-09-01 20:12:17 id:2 pass:3333 status:1 uname:wer3 updated_at:2022-09-01 20:23:29]','','POST','/admin/user/put','::1',9,'2022-09-01 21:38:08'),(1144,42,'map[desc:API分组选项 group:1 id:11 k:api_group status:1 type:1 v:菜单\r\nAPI\r\n角色\r\n管理员\r\n字典\r\n文件\r\n操作日志\r\n登陆日志\r\n备忘录\r\n管理员消息]','','POST','/admin/dict/put','::1',9,'2022-09-01 22:11:01'),(1145,42,'http://localhost:1211/admin/menu/del/161?','','GET','/admin/menu/del/:id','::1',10,'2022-09-01 22:11:12'),(1146,42,'http://localhost:1211/admin/menu/del/170?','','GET','/admin/menu/del/:id','::1',13,'2022-09-01 22:11:14'),(1147,42,'map[mid:[171 172] rid:1]','','POST','/admin/roleMenu/post','::1',3,'2022-09-01 22:34:38'),(1148,42,'map[desc:这是一个测试用户 device: email: icon: join_ip:127.0.0.1 nickname:ciel pass:123 phone: status:1 summary:测试完成啦 uname:ciel]','','POST','/admin/user/post','::1',3,'2022-09-01 22:38:06'),(1149,42,'map[desc:你好吗 device: email: icon: join_ip:127.0.0.1 nickname:morri pass:123 phone: status:1 summary:我是morri uname:morri]','','POST','/admin/user/post','::1',3,'2022-09-01 22:39:23');
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` VALUES (1,'Super Admin','2022-02-16 11:12:41','2022-02-21 04:46:24'),(22,'过客','2022-07-23 08:45:05','2022-08-30 15:45:52');
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
) ENGINE=InnoDB AUTO_INCREMENT=384 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` VALUES (320,22,145),(321,22,146),(322,22,147),(323,22,148),(324,22,149),(325,22,198),(326,22,197),(327,22,196),(331,22,156),(332,22,155),(333,22,154),(337,22,160),(338,22,199),(340,22,61),(341,22,186),(342,22,185),(343,22,184),(344,22,183),(345,22,56),(346,22,202),(347,22,208),(348,22,205),(349,22,168),(350,22,232),(351,22,231),(352,22,163),(353,22,164),(354,22,165),(355,22,166),(356,22,167),(357,22,218),(358,22,220),(359,22,221),(360,22,222),(361,22,224),(362,22,227),(363,22,228),(364,22,230),(365,22,225),(366,22,226),(367,22,233),(368,22,181),(369,22,179),(370,22,177),(371,22,175),(372,22,174),(373,22,173),(374,22,172),(375,22,171),(376,22,170),(377,22,169),(378,22,72),(379,22,71),(380,22,68),(381,22,66),(382,22,65),(383,22,62);
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
) ENGINE=InnoDB AUTO_INCREMENT=196 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(100,1,78),(136,1,132),(145,1,136),(148,1,139),(172,22,1),(173,22,2),(174,22,3),(175,22,4),(176,22,5),(177,22,28),(178,22,30),(179,22,78),(180,22,132),(181,22,136),(190,1,156),(194,1,171),(195,1,172);
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
  `uname` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
  `pass` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `icon` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `summary` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` text COLLATE utf8mb4_general_ci,
  `join_ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '注册IP',
  `device` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备名称',
  `phone` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `uname_2` (`uname`),
  KEY `join_ip` (`join_ip`),
  KEY `status` (`status`),
  KEY `phone` (`phone`),
  KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` VALUES (1,'ciel','123','ciel','','测试完成啦','这是一个测试用户','127.0.0.1','','','',1,'2022-09-01 22:38:06','2022-09-01 22:38:06'),(2,'morri','123','morri','','我是morri','你好吗','127.0.0.1','','','',1,'2022-09-01 22:39:23','2022-09-01 22:39:23');
/*!40000 ALTER TABLE `u_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-02 10:37:08
