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
-- Table structure for table `f_node`
--

DROP TABLE IF EXISTS `f_node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `f_node` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category` tinyint unsigned DEFAULT '1',
  `year` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `month` tinyint DEFAULT '0',
  `day` int DEFAULT '0',
  `uid` int NOT NULL,
  `level` int DEFAULT '1',
  `tag` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `summary` text COLLATE utf8mb4_general_ci,
  `main_things` longtext COLLATE utf8mb4_general_ci,
  `other_info` longtext COLLATE utf8mb4_general_ci,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `f_node_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=459 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `unread_msg_count`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel','$2a$10$Vo3PfpAzeBI2Dj4R0mKK7.B0o89ob154qF1RfcduamimLqbMU39fe',0,1,'2022-03-08 08:59:33','2022-07-24 15:05:25'),(42,1,'admin','$2a$10$VCrmz3RzJmO2aSX2CQSqsunK59fkkc4otXVWzPLCxsTRqSn/ybNzC',0,1,'2022-07-02 11:28:52','2022-07-24 15:04:56');
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
  `ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT ' {"label":"登录IP","notShow":0,"fieldType":"text","editHide":0,"editDisabled":0,"required":1}',
  `area` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '{"searchType":2,"hide":1}',
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_admin_login_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
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
  `from_uid` int NOT NULL,
  `to_uid` int NOT NULL,
  `group` int NOT NULL COMMENT '{"Label":"分组","FieldType":"select","SearchType":1,"Required":1,"Options":"1:系统:tag-info,2:管理员:tag-warning,3:其他:tag-danger"}',
  `type` int NOT NULL COMMENT '{"Label":"类型","FieldType":"select","SearchType":1,"Required":1,"Options":"1:文字:tag-info,2:图片:tag-primary,3:语音:tag-warning,4:视频:tag-success,5:链接:tag-danger"}',
  `content` text COLLATE utf8mb4_general_ci COMMENT '{"Label":"内容","SearchType":2,"Required":1}',
  `link` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '{"Label":"链接","SearchType":2}',
  `status` int DEFAULT '1' COMMENT '{"Label":"状态","FieldType":"select","Options":"1:未读:tag-warning,2已读:tag-info"}',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `from_uid` (`from_uid`),
  KEY `to_uid` (`to_uid`),
  CONSTRAINT `s_admin_message_ibfk_1` FOREIGN KEY (`from_uid`) REFERENCES `s_admin` (`id`),
  CONSTRAINT `s_admin_message_ibfk_2` FOREIGN KEY (`to_uid`) REFERENCES `s_admin` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_message`
--

LOCK TABLES `s_admin_message` WRITE;
/*!40000 ALTER TABLE `s_admin_message` DISABLE KEYS */;
INSERT INTO `s_admin_message` (`id`, `from_uid`, `to_uid`, `group`, `type`, `content`, `link`, `status`, `created_at`, `updated_at`) VALUES (139,42,42,1,1,'欢迎你的到来','',1,'2022-07-23 20:55:03','2022-07-23 20:55:03');
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
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=235 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (56,'/admin/file/path','1','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(61,'/admin/file/put','2','文件','文件修改',1,'2022-06-12 14:10:23','2022-07-23 20:08:04'),(62,'/admin/roleApi/path','1','角色','角色禁用API页面',1,'2022-06-12 17:02:13','2022-07-23 20:15:56'),(65,'/admin/roleApi/:id','1','角色','删除角色禁用API',1,'2022-06-12 17:02:13','2022-07-20 20:18:08'),(66,'/admin/roleApi/post','2','角色','角色禁用API添加',1,'2022-06-12 17:02:13','2022-07-23 20:15:37'),(68,'/admin/roleMenu/path','1','角色','角色菜单页面',1,'2022-06-13 15:19:45','2022-07-23 19:33:45'),(71,'/admin/roleMenu/:id','1','角色','角色菜单删除',1,'2022-06-13 15:19:45','2022-07-23 20:14:58'),(72,'/admin/roleMenu/post','2','角色','角色菜单添加',1,'2022-06-13 15:19:45','2022-07-23 20:14:48'),(144,'/admin/api/path','1','API','API页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:54'),(145,'/admin/api/path/add','1','API','API添加页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:47'),(146,'/admin/api/path/edit/:id','1','API','API修改页面',1,'2022-06-26 15:14:58','2022-07-23 20:03:40'),(147,'/admin/api/path/del/:id','1','API','API删除操作',1,'2022-06-26 15:14:58','2022-07-23 20:03:32'),(148,'/admin/api/post','2','API','API添加',1,'2022-06-26 15:14:58','2022-07-23 20:03:24'),(149,'/admin/api/put','2','API','API修改',1,'2022-06-26 15:14:58','2022-07-23 20:03:17'),(151,'/admin/dict/path','1','字典','字典页面',1,'2022-06-26 15:27:04','2022-07-10 13:47:37'),(152,'/admin/dict/path/add','1','字典','字典修改页面',1,'2022-06-26 15:27:04','2022-07-23 20:07:26'),(153,'/admin/dict/path/edit/:id','1','字典','字典修改页面',1,'2022-06-26 15:27:04','2022-07-23 20:07:19'),(154,'/admin/dict/path/del/:id','1','字典','字典删除',1,'2022-06-26 15:27:04','2022-07-23 20:07:05'),(155,'/admin/dict/post','2','字典','字典添加',1,'2022-06-26 15:27:04','2022-07-23 20:06:54'),(156,'/admin/dict/put','2','字典','字典修改',1,'2022-06-26 15:27:04','2022-07-23 20:06:46'),(157,'/admin/operationLog/path','1','操作日志','操作日志页面',1,'2022-06-26 20:30:22','2022-07-12 10:44:09'),(160,'/admin/operationLog/path/del/:id','1','操作日志','操作日志删除',1,'2022-06-26 20:30:22','2022-07-23 20:07:45'),(163,'/admin/admin/path','1','管理员','管理员页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:37'),(164,'/admin/admin/path/add','1','管理员','管理员添加页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:32'),(165,'/admin/admin/path/edit/:id','1','管理员','管理员修改页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:27'),(166,'/admin/admin/path/del/:id','1','管理员','管理员删除',1,'2022-06-27 16:21:07','2022-07-23 20:11:41'),(167,'/admin/admin/post','2','管理员','管理员添加',1,'2022-06-27 16:21:07','2022-07-23 20:11:34'),(168,'/admin/admin/put','2','管理员','管理员修改',1,'2022-06-27 16:21:07','2022-07-23 20:11:28'),(169,'/admin/role/path','1','角色','角色页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:38'),(170,'/admin/role/path/add','1','角色','角色添加页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:34'),(171,'/admin/role/path/edit/:id','1','角色','角色修改页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:30'),(172,'/admin/role/path/del/:id','1','角色','角色删除操作',1,'2022-06-27 18:27:46','2022-07-20 20:18:03'),(173,'/admin/role/post','2','角色','角色添加',1,'2022-06-27 18:27:46','2022-07-23 20:14:31'),(174,'/admin/role/put','2','角色','角色修改',1,'2022-06-27 18:27:46','2022-07-23 20:14:26'),(175,'/admin/roleMenu/path/add','1','角色','角色菜单添加页面',1,'2022-06-27 18:38:32','2022-07-10 13:45:09'),(177,'/admin/roleMenu/path/del/:id','1','角色','角色菜单删除操作',1,'2022-06-27 18:38:32','2022-07-20 20:17:57'),(179,'/admin/roleApi/path/add','1','角色','角色API添加页面',1,'2022-06-27 19:10:03','2022-07-23 20:14:11'),(181,'/admin/roleApi/path/del/:id','1','角色','角色API删除操作',1,'2022-06-27 19:10:03','2022-07-23 20:14:05'),(183,'/admin/file/path/add','1','文件','文件添加页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:38'),(184,'/admin/file/path/edit/:id','1','文件','文件修改页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:26'),(185,'/admin/file/path/del/:id','1','文件','文件删除操作',1,'2022-06-27 19:54:22','2022-07-20 20:17:19'),(186,'/admin/file/post','2','文件','文件添加',1,'2022-06-27 19:54:22','2022-07-23 20:07:54'),(193,'/admin/node/path','1','备忘录','备忘录页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:37'),(194,'/admin/node/path/add','1','备忘录','备忘录添加页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:26'),(195,'/admin/node/path/edit/:id','1','备忘录','备忘录修改页面',1,'2022-07-05 19:01:03','2022-07-23 20:06:17'),(196,'/admin/node/path/del/:id','1','备忘录','备忘录删除',1,'2022-07-05 19:01:03','2022-07-23 20:06:06'),(197,'/admin/node/post','2','备忘录','备忘录添加',1,'2022-07-05 19:01:03','2022-07-23 20:05:59'),(198,'/admin/node/put','2','备忘录','备忘录修改',1,'2022-07-05 19:01:03','2022-07-23 20:05:48'),(199,'/admin/operationLog/clear','1','操作日志','操作日志清空',1,'2022-07-10 12:39:19','2022-07-23 20:07:36'),(202,'/admin/adminLoginLog/path','1','登陆日志','登陆日志页面',1,'2022-07-11 19:06:26','2022-07-12 10:44:34'),(205,'/admin/adminLoginLog/path/del/:id','1','登陆日志','登陆日志删除操作',1,'2022-07-11 19:06:26','2022-07-20 20:17:34'),(208,'/admin/adminLoginLog/claer','1','登陆日志','登陆日志清空',1,'2022-07-14 09:49:26','2022-07-23 20:11:16'),(218,'/admin/adminMessage/path','1','管理员消息','消息页面',1,'2022-07-22 20:14:53','2022-07-23 20:12:37'),(220,'/admin/adminMessage/path/edit/:id','1','管理员消息','信息修改页面',1,'2022-07-22 20:14:53','2022-07-23 20:04:18'),(221,'/admin/adminMessage/path/del/:id','1','管理员消息','信息删除操作',1,'2022-07-22 20:14:53','2022-07-23 20:04:08'),(222,'/admin/adminMessage/post','2','管理员消息','消息发送',1,'2022-07-22 20:14:53','2022-07-23 20:12:25'),(224,'/admin/menu/path','1','菜单','菜单页面',1,'2022-07-23 19:34:26','2022-07-23 19:34:26'),(225,'/admin/menu/put','2','菜单','菜单修改',1,'2022-07-23 19:35:30','2022-07-23 19:37:19'),(226,'/admin/menu/post','2','菜单','菜单添加',1,'2022-07-23 19:35:46','2022-07-23 19:37:07'),(227,'/admin/menu/path/del/:id','1','菜单','菜单删除',1,'2022-07-23 19:36:08','2022-07-23 19:37:30'),(228,'/admin/menu/path/edit/:id','1','菜单','菜单修改页面',1,'2022-07-23 19:36:36','2022-07-23 19:36:36'),(230,'/admin/menu/path/add','1','菜单','菜单添加页面',1,'2022-07-23 19:37:57','2022-07-23 19:37:57'),(231,'/admin/admin/updateUname','3','管理员','管理员用户名修改',1,'2022-07-23 19:40:42','2022-07-23 19:42:27'),(232,'/admin/admin/updatePwdWithoutOldPwd','3','管理员','管理员密码修改',1,'2022-07-23 19:57:23','2022-07-23 19:57:23'),(233,'/admin/roleApi/clear','1','角色','角色API清空',1,'2022-07-23 20:02:23','2022-07-23 20:02:34');
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
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'api_group','菜单\nAPI\n角色\n管理员\n字典\n文件\n操作日志\n登陆日志\n备忘录\n管理员消息\nhaha','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-07-24 12:37:40'),(22,'music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-07-14 15:47:17'),(33,'node-category','1.记事\r\n3.Mysql\r\n5.English\r\n6.Freekey\r\n8.Golang\r\n9.Idea\r\n10.js\r\n11.jquery\r\n12.Linux\r\n15.Nginx\r\n16.Errors\r\n17.Quotations','备忘录分类','1',1,1,'2022-07-07 20:18:58','2022-07-21 21:08:55'),(42,'white_ips','','系统白名单','1',1,1,'2022-07-23 19:04:44','2022-07-23 19:04:44');
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
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (26,'1/2022/03/BYFY4d.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:50:13'),(27,'1/2022/03/FdI4Yw.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:51:51'),(29,'1/2022/03/mAMoWX.png',2,1,'2022-03-27 20:19:07','2022-07-09 16:49:57'),(30,'1/2022/03/2S41in.png',2,1,'2022-03-28 15:32:00','2022-07-11 18:42:26'),(31,'1/2022/03/IdGUqj.png',2,1,'2022-03-28 15:36:45','2022-07-09 16:49:40'),(32,'1/2022/03/5Eoxb1.png',2,1,'2022-03-28 15:40:17','2022-07-09 16:49:33'),(77,'2/2022/07/CQVqgn.webp',2,1,'2022-07-03 12:44:29','2022-07-03 12:44:29'),(78,'2/2022/07/qMBDps.png',2,1,'2022-07-03 12:49:10','2022-07-03 12:49:10'),(79,'2/2022/07/lSCC0m.webp',2,1,'2022-07-03 13:00:15','2022-07-03 13:00:15'),(80,'2/2022/07/SHf1y4.webp',2,1,'2022-07-03 18:38:21','2022-07-15 23:55:13'),(81,'1/2022/07/IJoBIZ.png',1,1,'2022-07-13 18:25:37','2022-07-13 18:25:37');
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
  `icon` text COLLATE utf8mb4_general_ci,
  `bg_img` text COLLATE utf8mb4_general_ci,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sort` decimal(7,2) NOT NULL DEFAULT '0.00',
  `type` int NOT NULL DEFAULT '1' COMMENT '1normal 2group',
  `desc` text COLLATE utf8mb4_general_ci,
  `file_path` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-07-09 08:09:47'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/admin/menu/path',1.10,1,'这里是菜单页面','\"\"',1,'2022-02-16 11:14:13','2022-07-08 15:24:50'),(3,1,'1/2022/03/IdGUqj.png','','角色','/admin/role/path',1.30,1,'','\"\"',1,'2022-03-04 08:57:14','2022-07-09 08:08:31'),(4,1,'1/2022/03/BYFY4d.gif','','API','/admin/api/path',1.20,1,'','',1,'2022-07-03 06:25:52','2022-07-09 10:41:50'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/admin/path',1.40,1,'','',1,'2022-03-08 07:45:04','2022-07-09 08:10:07'),(28,1,'1/2022/03/mAMoWX.png',NULL,'字典','/admin/dict/path',1.50,1,'字典页面',NULL,1,'2022-03-08 07:45:04','2022-07-02 08:06:55'),(30,1,'1/2022/03/2S41in.png','','文件','/admin/file/path',1.60,1,'',NULL,1,'2022-03-08 08:05:30','2022-07-03 05:12:16'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/admin/operationLog/path',1.80,1,'','',1,'2022-06-13 11:59:57','2022-07-09 08:19:58'),(132,-1,'','','工具','',2.00,2,'','',1,'2022-07-03 06:25:52','2022-07-03 06:25:52'),(136,132,'2/2022/07/SHf1y4.webp','','站点导航','/admin/to/urls',2.20,1,'','/sys/tool/urls.html',1,'2022-07-03 06:25:52','2022-07-20 12:15:54'),(138,132,'2/2022/07/CQVqgn.webp','','备忘录','/admin/node/path',2.40,1,'我的备忘录','',1,'2022-07-05 11:01:03','2022-07-18 03:58:01'),(139,1,'','','登录日志','/admin/adminLoginLog/path',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2022-07-16 15:09:08'),(151,1,'https://www.gravatar.com/avatar/1017420430?d=retro&f=y','','信息','/admin/adminMessage/path',1.92,1,'管理员长链接通知列表','',1,'2022-07-22 12:14:53','2022-07-22 12:16:36'),(152,132,'','','文档','/admin/to/document',2.50,1,'','/sys/tool/document.html',1,'2022-07-23 11:27:26','2022-07-23 11:29:50');
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
  `content` text COLLATE utf8mb4_general_ci NOT NULL,
  `response` text COLLATE utf8mb4_general_ci NOT NULL,
  `method` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `uri` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `ip` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `use_time` int NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_operation_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1019 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (1017,42,'http://localhost:1211/admin/operationLog/clear','','GET','/admin/operationLog/clear','::1',9,'2022-07-24 23:05:43'),(1018,42,'http://localhost:1211/admin/adminLoginLog/clear','','GET','/admin/adminLoginLog/clear','::1',9,'2022-07-24 23:05:47');
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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'Super Admin','2022-02-16 11:12:41','2022-02-21 04:46:24'),(22,'过客','2022-07-23 08:45:05','2022-07-24 12:51:22');
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
) ENGINE=InnoDB AUTO_INCREMENT=316 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` (`id`, `rid`, `aid`) VALUES (282,22,147),(283,22,148),(284,22,149),(285,22,198),(286,22,197),(287,22,196),(288,22,156),(289,22,155),(291,22,154),(292,22,199),(293,22,160),(294,22,61),(295,22,186),(296,22,185),(297,22,208),(298,22,205),(299,22,166),(300,22,168),(301,22,167),(302,22,231),(303,22,232),(304,22,226),(305,22,227),(306,22,173),(307,22,233),(308,22,181),(309,22,177),(310,22,174),(311,22,172),(312,22,72),(313,22,71),(314,22,66),(315,22,65);
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
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(100,1,78),(136,1,132),(145,1,136),(147,1,138),(148,1,139),(167,1,151),(172,22,1),(173,22,2),(174,22,3),(175,22,4),(176,22,5),(177,22,28),(178,22,30),(179,22,78),(180,22,132),(181,22,136),(182,22,138),(183,22,139),(185,22,151),(186,1,152);
/*!40000 ALTER TABLE `s_role_menu` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-24 23:06:10
