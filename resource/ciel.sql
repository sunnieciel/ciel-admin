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
) ENGINE=InnoDB AUTO_INCREMENT=441 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `f_node`
--

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
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel','$2a$10$VrWCJx9plc8gCeV19/owB.Prh1A7Kgcx1Fqlnhmd3MC2J5s3zwNNW',1,'2022-03-08 08:59:33','2022-07-17 11:07:09'),(42,2,'admin','$2a$10$4ELJmBB5FhqvIl0oCgstHeYrpW79g4C3.Xf6541lxtBXbBZFbtVk6',1,'2022-07-02 11:28:52','2022-07-15 15:56:37');
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `ip`, `area`, `status`, `created_at`, `updated_at`) VALUES (3,1,'222','44423',1,'2022-07-18 16:24:19','2022-07-19 13:00:42'),(7,1,'2','',1,'2022-07-19 13:12:16','2022-07-19 13:12:16'),(8,1,'::1',NULL,1,'2022-07-19 14:09:48','2022-07-19 14:09:48'),(9,1,'::1',NULL,1,'2022-07-19 15:35:48','2022-07-19 15:35:48'),(10,1,'::1',NULL,1,'2022-07-19 15:42:30','2022-07-19 15:42:30'),(11,1,'::1',NULL,1,'2022-07-19 16:40:37','2022-07-19 16:40:37');
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
) ENGINE=InnoDB AUTO_INCREMENT=218 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (56,'/file/path','1','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(57,'/file','1','文件','查询文件集合',1,'2022-06-12 14:10:23','2022-06-14 13:44:54'),(58,'/file/:id','1','文件','查询文件详情',1,'2022-06-12 14:10:23','2022-06-14 13:44:49'),(59,'/file/:id','4','文件','删除文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:35'),(61,'/file','3','文件','修改文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:26'),(62,'/roleApi/path','1','角色','禁用角色API页面',1,'2022-06-12 17:02:13','2022-07-10 20:43:49'),(63,'/roleApi','1','角色','查询角色禁用API集合',1,'2022-06-12 17:02:13','2022-07-10 20:44:39'),(65,'/roleApi/:id','4','角色','删除角色禁用API',1,'2022-06-12 17:02:13','2022-07-10 20:45:03'),(66,'/roleApi','2','角色','添加角色禁用API',1,'2022-06-12 17:02:13','2022-07-10 20:44:53'),(68,'/roleMenu/path','1','菜单','角色菜单页面',1,'2022-06-13 15:19:45','2022-06-14 13:43:53'),(69,'/roleMenu','1','菜单','查询角色菜单集合',1,'2022-06-13 15:19:45','2022-06-14 13:43:44'),(71,'/roleMenu/:id','4','菜单','删除角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:43:39'),(72,'/roleMenu','2','菜单','添加角色菜单',1,'2022-06-13 15:19:45','2022-07-11 16:24:31'),(144,'/api/path','1','API','api页面',1,'2022-06-26 15:14:58','2022-07-10 20:43:29'),(145,'/api/path/add','1','API','api添加页面',1,'2022-06-26 15:14:58','2022-07-15 23:50:16'),(146,'/api/path/edit/:id','1','API','api修改页面',1,'2022-06-26 15:14:58','2022-07-10 20:42:50'),(147,'/api/path/del/:id','4','API','api删除操作',1,'2022-06-26 15:14:58','2022-07-10 20:43:04'),(148,'/api','2','API','添加api',1,'2022-06-26 15:14:58','2022-07-10 20:42:03'),(149,'/api','3','API','修改api',1,'2022-06-26 15:14:58','2022-07-10 20:32:13'),(151,'/dict/path','1','字典','字典页面',1,'2022-06-26 15:27:04','2022-07-10 13:47:37'),(152,'/dict/path/add','1','字典','添加字典页面',1,'2022-06-26 15:27:04','2022-07-10 13:47:31'),(153,'/dict/path/edit/:id','1','字典','修改字典页面',1,'2022-06-26 15:27:04','2022-07-10 13:47:26'),(154,'/dict/path/del/:id','4','字典','删除字典',1,'2022-06-26 15:27:04','2022-07-10 13:47:19'),(155,'/dict','2','字典','添加字典',1,'2022-06-26 15:27:04','2022-07-10 13:47:14'),(156,'/dict','3','字典','修改字典',1,'2022-06-26 15:27:04','2022-07-10 13:47:08'),(157,'/operationLog/path','1','操作日志','操作日志页面',1,'2022-06-26 20:30:22','2022-07-12 10:44:09'),(160,'/operationLog/path/del/:id','4','操作日志','删除操作日志',1,'2022-06-26 20:30:22','2022-07-12 10:44:01'),(163,'/admin/path','1','管理员','管理员页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:37'),(164,'/admin/path/add','1','管理员','管理员添加页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:32'),(165,'/admin/path/edit/:id','1','管理员','管理员修改页面',1,'2022-06-27 16:21:07','2022-07-10 13:46:27'),(166,'/admin/path/del/:id','4','管理员','删除管理员',1,'2022-06-27 16:21:07','2022-07-10 13:46:22'),(167,'/admin','2','管理员','添加管理员',1,'2022-06-27 16:21:07','2022-07-10 13:46:17'),(168,'/admin','3','管理员','修改管理员',1,'2022-06-27 16:21:07','2022-07-10 13:45:46'),(169,'/role/path','1','角色','角色页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:38'),(170,'/role/path/add','1','角色','角色添加页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:34'),(171,'/role/path/edit/:id','1','角色','角色修改页面',1,'2022-06-27 18:27:46','2022-07-10 13:45:30'),(172,'/role/path/del/:id','4','角色','角色删除操作',1,'2022-06-27 18:27:46','2022-07-10 13:45:24'),(173,'/role','2','角色','添加角色',1,'2022-06-27 18:27:46','2022-07-10 13:45:19'),(174,'/role','3','角色','修改角色',1,'2022-06-27 18:27:46','2022-07-10 13:45:14'),(175,'/roleMenu/path/add','1','角色','角色菜单添加页面',1,'2022-06-27 18:38:32','2022-07-10 13:45:09'),(177,'/roleMenu/path/del/:id','4','角色','角色菜单删除操作',1,'2022-06-27 18:38:32','2022-07-10 13:45:04'),(178,'/roleMenu','3','角色','修改角色菜单',1,'2022-06-27 18:38:32','2022-07-10 13:44:57'),(179,'/roleApi/path/add','1','角色','角色api添加页面',1,'2022-06-27 19:10:03','2022-07-10 13:44:50'),(181,'/roleApi/path/del/:id','4','角色','角色api删除操作',1,'2022-06-27 19:10:03','2022-07-10 13:44:42'),(182,'/roleApi','3','角色','修改角色api',1,'2022-06-27 19:10:03','2022-07-10 13:44:34'),(183,'/file/path/add','1','文件','文件添加页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:38'),(184,'/file/path/edit/:id','1','文件','文件修改页面',1,'2022-06-27 19:54:22','2022-07-10 13:42:26'),(185,'/file/path/del/:id','4','文件','文件删除操作',1,'2022-06-27 19:54:22','2022-07-10 13:42:20'),(186,'/file','2','文件','添加文件',1,'2022-06-27 19:54:22','2022-07-10 13:42:06'),(193,'/node/path','1','备忘录','记事本页面',1,'2022-07-05 19:01:03','2022-07-10 13:41:55'),(194,'/node/path/add','1','备忘录','记事本添加页面',1,'2022-07-05 19:01:03','2022-07-10 13:41:46'),(195,'/node/path/edit/:id','1','备忘录','记事本修改页面',1,'2022-07-05 19:01:03','2022-07-10 13:41:39'),(196,'/node/path/del/:id','4','备忘录','删除记事本',1,'2022-07-05 19:01:03','2022-07-10 13:41:30'),(197,'/node','2','备忘录','添加记事本',1,'2022-07-05 19:01:03','2022-07-10 13:41:21'),(198,'/node','3','备忘录','修改记事本',1,'2022-07-05 19:01:03','2022-07-10 13:41:12'),(199,'/operationLog/clear','1','操作日志','清空操作日志',1,'2022-07-10 12:39:19','2022-07-12 10:43:45'),(202,'/adminLoginLog/path','1','登陆日志','登陆日志页面',1,'2022-07-11 19:06:26','2022-07-12 10:44:34'),(205,'/adminLoginLog/path/del/:id','4','登陆日志','登陆日志删除操作',1,'2022-07-11 19:06:26','2022-07-12 10:44:26'),(208,'/adminLoginLog/claer','1','登陆日志','清空登陆日志',1,'2022-07-14 09:49:26','2022-07-14 09:49:26'),(209,'/adminLoginLog/path/add','1','sys','登录日志添加页面',1,'2022-07-18 16:07:40','2022-07-18 16:07:40'),(210,'/adminLoginLog/path/edit/:id','1','sys','登录日志修改页面',1,'2022-07-18 16:07:40','2022-07-18 16:07:40'),(211,'/adminLoginLog','2','sys','添加登录日志',1,'2022-07-18 16:07:40','2022-07-18 16:07:40'),(212,'/adminLoginLog','3','sys','修改登录日志',1,'2022-07-18 16:07:40','2022-07-18 16:07:40'),(213,'/user/path','1','user','用户列表页面',1,'2022-07-19 15:21:19','2022-07-19 15:21:19'),(214,'/user/path/add','1','user','用户列表添加页面',1,'2022-07-19 15:21:19','2022-07-19 15:21:19'),(215,'/user/path/edit/:id','1','user','用户列表修改页面',1,'2022-07-19 15:21:19','2022-07-19 15:21:19'),(216,'/user/path/del/:id','1','user','用户列表删除操作',1,'2022-07-19 15:21:19','2022-07-19 15:21:19'),(217,'/user','2','user','添加用户列表',1,'2022-07-19 15:21:19','2022-07-19 15:21:19');
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
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'api_group','菜单\r\nAPI\r\n角色\r\n管理员\r\n字典\r\n文件\r\n操作日志\r\n登陆日志\r\n备忘录\r\n','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-07-12 10:43:17'),(22,'music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-07-14 15:47:17'),(33,'node-category','1.记事\r\n3.mysql\r\n5.english\r\n6.freekey\r\n8.go\r\n9.idea\r\n10.js\r\n12.linux\r\n15.nginx\r\n16.error\r\n17.quotations\r\n18.异星工厂','备忘录分类','1',1,1,'2022-07-07 20:18:58','2022-07-18 17:09:46');
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
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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
) ENGINE=InnoDB AUTO_INCREMENT=151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-07-09 08:09:47'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/menu/path',1.10,1,'这里是菜单页面','\"\"',1,'2022-02-16 11:14:13','2022-07-08 15:24:50'),(3,1,'1/2022/03/IdGUqj.png','','角色','/role/path',1.30,1,'','\"\"',1,'2022-03-04 08:57:14','2022-07-09 08:08:31'),(4,1,'1/2022/03/BYFY4d.gif','','API','/api/path',1.20,1,'','',1,'2022-07-03 06:25:52','2022-07-09 10:41:50'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/path',1.40,1,'','',1,'2022-03-08 07:45:04','2022-07-09 08:10:07'),(28,1,'1/2022/03/mAMoWX.png',NULL,'字典','/dict/path',1.50,1,'字典页面',NULL,1,'2022-03-08 07:45:04','2022-07-02 08:06:55'),(30,1,'1/2022/03/2S41in.png','','文件','/file/path',1.60,1,'',NULL,1,'2022-03-08 08:05:30','2022-07-03 05:12:16'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/operationLog/path',1.80,1,'','',1,'2022-06-13 11:59:57','2022-07-09 08:19:58'),(132,-1,'','','工具','',2.00,2,'','',1,'2022-07-03 06:25:52','2022-07-03 06:25:52'),(136,132,'2/2022/07/SHf1y4.webp','2/2022/07/SHf1y4.webp','站点导航','/to/urls',2.20,1,'','/sys/tool/urls.html',1,'2022-07-03 06:25:52','2022-07-15 15:44:55'),(137,132,'2/2022/07/SHf1y4.webp','2/2022/07/SHf1y4.webp','语录','/to/quotations',2.30,1,'','/sys/tool/quotations.html',1,'2022-07-03 06:29:23','2022-07-16 12:37:47'),(138,132,'2/2022/07/CQVqgn.webp','','备忘录','/node/path',2.40,1,'我的备忘录','',1,'2022-07-05 11:01:03','2022-07-18 03:58:01'),(139,1,'','','登录日志','/adminLoginLog/path',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2022-07-16 15:09:08'),(140,132,'1/2022/07/IJoBIZ.png','','文档','/to/document',2.50,1,'','/sys/tool//document.html',1,'2022-07-12 11:46:48','2022-07-16 12:38:28'),(145,132,'https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y','','前端常用资料','/to/knowledge_web',2.70,1,'js，css，html相关的常用查询资料可以放在此页面，方便进行查阅。','/sys/tool/knowledgeWeb.html',1,'2022-07-16 06:31:16','2022-07-16 07:45:53');
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
) ENGINE=InnoDB AUTO_INCREMENT=676 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (653,1,'http://localhost:1211/operationLog/clear','','GET','/operationLog/clear','::1',3,'2022-07-16 13:49:58'),(654,1,'http://localhost:1211/adminLoginLog/clear','','GET','/adminLoginLog/clear','::1',10,'2022-07-16 13:50:06'),(655,1,'http://localhost:1211/menu/path/del/144?','','GET','/menu/path/del/:id','::1',10,'2022-07-16 14:31:44'),(656,1,'map[mid:[145] rid:1]','','POST','/roleMenu/post','::1',11,'2022-07-16 14:32:04'),(657,1,'map[bg_img:https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y desc:js，css，html相关的常用查询资料可以放在此页面，方便进行查阅。 file_path:/sys/tool/knowledgeWeb.html icon:https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y id:145 name:web端常用资料 path:/to/knowledge_web pid:132 sort:2.7 status:1 type:1]','','POST','/menu/put','::1',10,'2022-07-16 14:33:01'),(658,1,'map[bg_img:https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y desc:js，css，html相关的常用查询资料可以放在此页面，方便进行查阅。 file_path:/sys/tool/knowledgeWeb.html icon:https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y id:145 name:前端常用资料 path:/to/knowledge_web pid:132 sort:2.7 status:1 type:1]','','POST','/menu/put','::1',9,'2022-07-16 14:34:38'),(659,1,'map[bg_img: desc:js，css，html相关的常用查询资料可以放在此页面，方便进行查阅。 file_path:/sys/tool/knowledgeWeb.html icon:https://www.gravatar.com/avatar/1826024425?d=monsterid&f=y id:145 name:前端常用资料 path:/to/knowledge_web pid:132 sort:2.7 status:1 type:1]','','POST','/menu/put','::1',48,'2022-07-16 15:45:53'),(660,1,'map[bg_img:2/2022/07/SHf1y4.webp desc: file_path:/sys/tool/quotations.html icon:2/2022/07/SHf1y4.webp id:137 name:语录 path:/to/quotations pid:132 sort:2.3 status:1 type:1]','','POST','/menu/put','::1',9,'2022-07-16 20:37:47'),(661,1,'map[bg_img:2/2022/07/SHf1y4.webp desc: file_path:/sys/tool//document.html icon:1/2022/07/IJoBIZ.png id:140 name:文档 path:/to/document pid:132 sort:2.5 status:1 type:1]','','POST','/menu/put','::1',10,'2022-07-16 20:38:12'),(662,1,'map[bg_img: desc: file_path:/sys/tool//document.html icon:1/2022/07/IJoBIZ.png id:140 name:文档 path:/to/document pid:132 sort:2.5 status:1 type:1]','','POST','/menu/put','::1',11,'2022-07-16 20:38:28'),(663,1,'map[aid:[56] rid:2]','','POST','/roleApi/post','::1',2,'2022-07-16 22:48:51'),(664,1,'http://localhost:1211/roleApi/path/del/237?rid=2&','','GET','/roleApi/path/del/:id','::1',2,'2022-07-16 22:48:56'),(665,1,'map[bg_img: desc:这里是登陆日志页面可以对数据进行相应的操作。 file_path: icon: id:139 name:登录日志 path:/adminLoginLog/path pid:1 sort:1.9 status:1 type:1]','','POST','/menu/put','::1',9,'2022-07-16 23:09:08'),(666,1,'map[newPwd:1 oldPwd:1]','','PUT','/admin/updatePwd','::1',168,'2022-07-17 19:07:09'),(667,1,'http://localhost:1211/menu/path/del/146?','','GET','/menu/path/del/:id','::1',2,'2022-07-17 21:00:22'),(668,1,'map[bg_img: desc:我的备忘录 file_path: icon:2/2022/07/CQVqgn.webp id:138 name:备忘录 path:/node/path pid:132 sort:2.4 status:1 type:1]','','POST','/menu/put','::1',10,'2022-07-18 11:58:01'),(669,1,'http://localhost:1211/menu/path/del/147?','','GET','/menu/path/del/:id','::1',9,'2022-07-18 12:29:29'),(670,1,'http://localhost:1211/menu/path/del/148?','','GET','/menu/path/del/:id','::1',10,'2022-07-18 15:36:32'),(671,1,'map[desc:备忘录分类 group:1 id:33 k:node-category status:1 type:1 v:1.记事\r\n3.mysql\r\n5.english\r\n6.freekey\r\n8.go\r\n9.idea\r\n10.js\r\n12.linux\r\n15.nginx\r\n16.error\r\n17.quotations\r\n18.异星工厂]','','POST','/dict/put','::1',9,'2022-07-18 17:09:46'),(672,1,'http://localhost:1211/menu/path/del/59?','','GET','/menu/path/del/:id','::1',10,'2022-07-19 14:09:39'),(673,1,'map[mid:[149 150] rid:1]','','POST','/roleMenu/post','::1',10,'2022-07-19 15:41:15'),(674,1,'http://localhost:1211/menu/path/del/150?','','GET','/menu/path/del/:id','::1',10,'2022-07-19 16:48:39'),(675,1,'http://localhost:1211/menu/path/del/149?','','GET','/menu/path/del/:id','::1',9,'2022-07-19 16:48:42');
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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'Super Admin','2022-02-16 11:12:41','2022-02-21 04:46:24'),(2,'Admin','2022-02-16 11:13:11','2022-06-27 10:36:21');
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
) ENGINE=InnoDB AUTO_INCREMENT=238 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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
) ENGINE=InnoDB AUTO_INCREMENT=167 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(100,1,78),(136,1,132),(142,2,28),(145,1,136),(146,1,137),(147,1,138),(148,1,139),(149,1,140),(150,2,1),(151,2,2),(152,2,3),(153,2,4),(154,2,5),(155,2,30),(157,2,78),(158,2,132),(159,2,136),(160,2,137),(161,2,138),(162,2,139),(163,2,140),(164,1,145);
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
  `uname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '{"label":"用户名","searchType":2,"required":1,"disabled":1}',
  `pass` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '{"hide":1,"editHide":1}',
  `nickname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '{"label":"昵称","required":1,"comment":"取一个昵称吧"}',
  `description` text COLLATE utf8mb4_general_ci COMMENT '{"fieldType":"markdown"}',
  `status` int DEFAULT '1' COMMENT '{"searchType":2,"fieldType":"select","options":"1:正常:tag-info,2:禁用:tag-danger"}',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` (`id`, `uname`, `pass`, `nickname`, `description`, `status`, `created_at`, `updated_at`) VALUES (1,'ciel','','ciel','33',2,'2022-07-19 15:45:32','2022-07-19 15:58:33'),(2,'morri','123','morri','morri',1,'2022-07-19 15:46:14','2022-07-19 15:46:14');
/*!40000 ALTER TABLE `u_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_user_login_log`
--

DROP TABLE IF EXISTS `u_user_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_user_login_log` (
  `id` bigint unsigned DEFAULT NULL,
  `uid` bigint unsigned DEFAULT NULL,
  `ip` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  KEY `uid` (`uid`),
  CONSTRAINT `u_user_login_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `u_user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user_login_log`
--

LOCK TABLES `u_user_login_log` WRITE;
/*!40000 ALTER TABLE `u_user_login_log` DISABLE KEYS */;
INSERT INTO `u_user_login_log` (`id`, `uid`, `ip`, `created_at`) VALUES (1,1,'127.0.0.1','2022-07-19 15:45:45'),(2,2,'localhost','2022-07-19 15:46:33');
/*!40000 ALTER TABLE `u_user_login_log` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-19 17:54:18
