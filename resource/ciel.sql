-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: 127.0.0.1    Database: ciel
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
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `nickname`, `email`, `phone`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel','$2a$10$OAp3RJVKv6WhAX3o.fY/A.R0jUOyzvtlfxpS3DgHtEVkLx/lY6b4.','I\'m ciel','','',1,'2022-03-08 08:59:33','2022-09-03 12:27:59'),(42,1,'admin','$2a$10$lxEEnsF7201zWPQilY6rZ.eLRkS89wVqNmSXPIw6t3emlyfOctgcy','morri','','',1,'2022-07-02 11:28:52','2022-09-03 12:36:58');
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
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `ip`, `area`, `status`, `created_at`, `updated_at`) VALUES (100,42,'::1',NULL,1,'2022-09-03 20:36:53','2022-09-03 20:36:53'),(101,42,'::1',NULL,1,'2022-09-03 20:37:05','2022-09-03 20:37:05'),(102,42,'::1',NULL,1,'2022-09-03 22:40:45','2022-09-03 22:40:45'),(103,42,'::1',NULL,1,'2022-09-04 14:28:25','2022-09-04 14:28:25');
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
  `type` tinyint unsigned DEFAULT '4' COMMENT '类型 1添加 2删除 3修改 4查看 5 页面 ',
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=253 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `type`, `desc`, `created_at`, `updated_at`) VALUES (56,'/admin/file/path','1','文件',5,'文件页面','2022-06-12 14:10:23','2022-06-14 13:45:01'),(61,'/admin/file/put','2','文件',3,'文件修改','2022-06-12 14:10:23','2022-07-23 20:08:04'),(62,'/admin/roleApi/path','1','角色',5,'角色禁用API','2022-06-12 17:02:13','2022-09-03 15:40:54'),(65,'/admin/roleApi/:id','1','角色',2,'删除角色禁用API','2022-06-12 17:02:13','2022-07-20 20:18:08'),(66,'/admin/roleApi/post','2','角色',1,'角色禁用API添加','2022-06-12 17:02:13','2022-07-23 20:15:37'),(68,'/admin/roleMenu/path','1','角色',5,'角色菜单页面','2022-06-13 15:19:45','2022-07-23 19:33:45'),(71,'/admin/roleMenu/:id','1','角色',2,'角色菜单删除','2022-06-13 15:19:45','2022-07-23 20:14:58'),(72,'/admin/roleMenu/post','2','角色',1,'角色菜单添加','2022-06-13 15:19:45','2022-07-23 20:14:48'),(144,'/admin/api/path','1','API',5,'API页面','2022-06-26 15:14:58','2022-07-23 20:03:54'),(145,'/admin/api/path/add','1','API',5,'API添加页面','2022-06-26 15:14:58','2022-07-23 20:03:47'),(146,'/admin/api/path/edit/:id','1','API',5,'API修改页面','2022-06-26 15:14:58','2022-07-23 20:03:40'),(147,'/admin/api/path/del/:id','1','API',2,'API删除操作','2022-06-26 15:14:58','2022-07-23 20:03:32'),(148,'/admin/api/post','2','API',1,'API添加','2022-06-26 15:14:58','2022-07-23 20:03:24'),(149,'/admin/api/put','2','API',3,'API修改','2022-06-26 15:14:58','2022-08-30 22:12:36'),(151,'/admin/dict/path','1','字典',5,'字典页面','2022-06-26 15:27:04','2022-07-10 13:47:37'),(152,'/admin/dict/path/add','1','字典',5,'字典修改页面','2022-06-26 15:27:04','2022-07-23 20:07:26'),(153,'/admin/dict/path/edit/:id','1','字典',5,'字典修改页面','2022-06-26 15:27:04','2022-07-23 20:07:19'),(154,'/admin/dict/path/del/:id','1','字典',2,'字典删除','2022-06-26 15:27:04','2022-07-23 20:07:05'),(155,'/admin/dict/post','2','字典',1,'字典添加','2022-06-26 15:27:04','2022-07-23 20:06:54'),(156,'/admin/dict/put','2','字典',3,'字典修改','2022-06-26 15:27:04','2022-07-23 20:06:46'),(157,'/admin/operationLog/path','1','操作日志',5,'操作日志页面','2022-06-26 20:30:22','2022-07-12 10:44:09'),(160,'/admin/operationLog/path/del/:id','1','操作日志',2,'操作日志删除','2022-06-26 20:30:22','2022-07-23 20:07:45'),(163,'/admin/admin/path','1','管理员',5,'管理员页面','2022-06-27 16:21:07','2022-07-10 13:46:37'),(164,'/admin/admin/path/add','1','管理员',5,'管理员添加页面','2022-06-27 16:21:07','2022-07-10 13:46:32'),(165,'/admin/admin/path/edit/:id','1','管理员',5,'管理员修改页面','2022-06-27 16:21:07','2022-07-10 13:46:27'),(166,'/admin/admin/path/del/:id','1','管理员',2,'管理员删除','2022-06-27 16:21:07','2022-07-23 20:11:41'),(167,'/admin/admin/post','2','管理员',1,'管理员添加','2022-06-27 16:21:07','2022-07-23 20:11:34'),(168,'/admin/admin/put','2','管理员',3,'管理员修改','2022-06-27 16:21:07','2022-07-23 20:11:28'),(169,'/admin/role/path','1','角色',5,'角色页面','2022-06-27 18:27:46','2022-07-10 13:45:38'),(170,'/admin/role/path/add','1','角色',5,'角色添加页面','2022-06-27 18:27:46','2022-07-10 13:45:34'),(171,'/admin/role/path/edit/:id','1','角色',5,'角色修改页面','2022-06-27 18:27:46','2022-07-10 13:45:30'),(172,'/admin/role/path/del/:id','1','角色',2,'角色删除操作','2022-06-27 18:27:46','2022-07-20 20:18:03'),(173,'/admin/role/post','2','角色',1,'角色添加','2022-06-27 18:27:46','2022-07-23 20:14:31'),(174,'/admin/role/put','2','角色',3,'角色修改','2022-06-27 18:27:46','2022-07-23 20:14:26'),(175,'/admin/roleMenu/path/add','1','角色',5,'角色菜单添加页面','2022-06-27 18:38:32','2022-07-10 13:45:09'),(177,'/admin/roleMenu/path/del/:id','1','角色',2,'角色菜单删除操作','2022-06-27 18:38:32','2022-07-20 20:17:57'),(179,'/admin/roleApi/path/add','1','角色',5,'角色API添加页面','2022-06-27 19:10:03','2022-07-23 20:14:11'),(181,'/admin/roleApi/path/del/:id','1','角色',2,'角色API删除操作','2022-06-27 19:10:03','2022-07-23 20:14:05'),(183,'/admin/file/path/add','1','文件',5,'文件添加页面','2022-06-27 19:54:22','2022-07-10 13:42:38'),(184,'/admin/file/path/edit/:id','1','文件',5,'文件修改页面','2022-06-27 19:54:22','2022-07-10 13:42:26'),(185,'/admin/file/path/del/:id','1','文件',2,'文件删除操作','2022-06-27 19:54:22','2022-07-20 20:17:19'),(186,'/admin/file/post','2','文件',1,'文件添加','2022-06-27 19:54:22','2022-07-23 20:07:54'),(199,'/admin/operationLog/clear','1','操作日志',2,'操作日志清空','2022-07-10 12:39:19','2022-07-23 20:07:36'),(202,'/admin/adminLoginLog/path','1','登陆日志',5,'登陆日志页面','2022-07-11 19:06:26','2022-07-12 10:44:34'),(205,'/admin/adminLoginLog/path/del/:id','1','登陆日志',2,'登陆日志删除操作','2022-07-11 19:06:26','2022-07-20 20:17:34'),(208,'/admin/adminLoginLog/claer','1','登陆日志',2,'登陆日志清空','2022-07-14 09:49:26','2022-07-23 20:11:16'),(224,'/admin/menu/path','1','菜单',5,'菜单页面','2022-07-23 19:34:26','2022-07-23 19:34:26'),(225,'/admin/menu/put','2','菜单',3,'菜单修改','2022-07-23 19:35:30','2022-07-23 19:37:19'),(226,'/admin/menu/post','2','菜单',1,'菜单添加','2022-07-23 19:35:46','2022-07-23 19:37:07'),(227,'/admin/menu/path/del/:id','1','菜单',2,'菜单删除','2022-07-23 19:36:08','2022-07-23 19:37:30'),(228,'/admin/menu/path/edit/:id','1','菜单',5,'菜单修改页面','2022-07-23 19:36:36','2022-07-23 19:36:36'),(230,'/admin/menu/path/add','1','菜单',5,'菜单添加页面','2022-07-23 19:37:57','2022-07-23 19:37:57'),(231,'/admin/admin/updateUname','3','管理员',3,'管理员用户名修改','2022-07-23 19:40:42','2022-07-23 19:42:27'),(232,'/admin/admin/updatePwdWithoutOldPwd','3','管理员',3,'管理员密码修改','2022-07-23 19:57:23','2022-07-23 19:57:23'),(233,'/admin/roleApi/clear','1','角色',2,'角色API清空','2022-07-23 20:02:23','2022-07-23 20:02:34'),(242,'/user','1','user',5,'用户列表页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(243,'/user/add','1','user',5,'用户列表添加页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(244,'/user/edit/:id','1','user',5,'用户列表修改页面','2022-09-01 14:36:51','2022-09-01 14:36:51'),(245,'/user/del/:id','1','user',2,'用户列表删除操作','2022-09-01 14:36:51','2022-09-01 14:36:51'),(247,'/userLoginLog','1','用户登录日志',5,'用户登录日志页面','2022-09-03 14:47:17','2022-09-03 14:47:17'),(248,'/userLoginLog/add','1','用户登录日志',5,'用户登录日志添加页面','2022-09-03 14:47:17','2022-09-03 14:47:17'),(249,'/userLoginLog/edit/:id','1','用户登录日志',5,'用户登录日志修改页面','2022-09-03 14:47:17','2022-09-03 14:47:17'),(250,'/userLoginLog/del/:id','1','用户登录日志',2,'用户登录日志删除操作','2022-09-03 14:47:17','2022-09-03 14:47:17'),(252,'/admin/user/updateUname','3','用户',3,'修改用户名','2022-09-03 23:19:07','2022-09-03 23:19:07');
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
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `title`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'API分组选项','api_group','菜单\r\nAPI\r\n角色\r\n管理员\r\n字典\r\n文件\r\n操作日志\r\n登陆日志\r\n备忘录\r\n管理员消息\r\n用户\r\n用户登录日志','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-09-03 14:47:17'),(22,'登陆音乐列表','music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-07-14 15:47:17'),(42,'系统白名单','white_ips','','多个ip用小写逗号隔开','1',1,1,'2022-07-23 19:04:44','2022-09-03 20:42:02');
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
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (26,'1/2022/03/BYFY4d.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:50:13'),(27,'1/2022/03/FdI4Yw.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:51:51'),(29,'1/2022/03/mAMoWX.png',2,1,'2022-03-27 20:19:07','2022-07-09 16:49:57'),(30,'1/2022/03/2S41in.png',2,1,'2022-03-28 15:32:00','2022-07-11 18:42:26'),(31,'1/2022/03/IdGUqj.png',2,1,'2022-03-28 15:36:45','2022-07-09 16:49:40'),(32,'1/2022/03/5Eoxb1.png',2,1,'2022-03-28 15:40:17','2022-07-09 16:49:33'),(77,'2/2022/07/CQVqgn.webp',2,1,'2022-07-03 12:44:29','2022-07-03 12:44:29'),(78,'2/2022/07/qMBDps.png',2,1,'2022-07-03 12:49:10','2022-07-03 12:49:10'),(79,'2/2022/07/lSCC0m.webp',2,1,'2022-07-03 13:00:15','2022-07-03 13:00:15'),(80,'2/2022/07/SHf1y4.webp',2,1,'2022-07-03 18:38:21','2022-07-15 23:55:13'),(81,'1/2022/07/IJoBIZ.png',1,1,'2022-07-13 18:25:37','2022-08-31 14:14:58');
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
) ENGINE=InnoDB AUTO_INCREMENT=174 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-07-09 08:09:47'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/admin/menu',1.10,1,'这里是菜单页面','\"\"',1,'2022-02-16 11:14:13','2022-07-08 15:24:50'),(3,1,'1/2022/03/IdGUqj.png','','角色','/admin/role',1.30,1,'角色权限管理，在这里可以创建新的角色','',1,'2022-03-04 08:57:14','2022-09-03 12:12:03'),(4,1,'1/2022/03/BYFY4d.gif','','API','/admin/api',1.20,1,'系统所有的操作api在此','',1,'2022-07-03 06:25:52','2022-09-03 12:13:16'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/admin',1.40,1,'','',1,'2022-03-08 07:45:04','2022-07-09 08:10:07'),(28,1,'1/2022/03/mAMoWX.png',NULL,'字典','/admin/dict',1.50,1,'字典页面',NULL,1,'2022-03-08 07:45:04','2022-07-02 08:06:55'),(30,1,'1/2022/03/2S41in.png','','文件','/admin/file',1.60,1,'',NULL,1,'2022-03-08 08:05:30','2022-07-03 05:12:16'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/admin/operationLog',1.80,1,'','',1,'2022-06-13 11:59:57','2022-07-09 08:19:58'),(139,1,'','','登录日志','/admin/adminLoginLog',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2022-07-16 15:09:08'),(156,1,'','','代码生成','/admin/gen',1.91,1,'','',1,'2022-09-01 04:44:11','2022-09-01 13:36:44'),(171,-1,'','','用户','',3.00,2,'','',1,'2022-09-01 14:24:18','2022-09-01 14:24:18'),(172,171,'','','用户列表','/admin/user',3.10,1,'','',1,'2022-09-01 14:24:18','2022-09-01 14:24:18'),(173,171,'','','登录日志','/admin/userLoginLog',3.20,1,'','',1,'2022-09-03 06:47:17','2022-09-03 06:51:21');
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
) ENGINE=InnoDB AUTO_INCREMENT=1344 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
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
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'超级管理员','2022-02-16 11:12:41','2022-09-02 12:22:24'),(22,'系统管理员','2022-07-23 08:45:05','2022-09-02 12:22:31'),(24,'临时管理员','2022-09-03 14:50:33','2022-09-03 14:50:33');
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
) ENGINE=InnoDB AUTO_INCREMENT=905 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` (`id`, `rid`, `aid`) VALUES (530,22,147),(531,22,148),(534,22,245),(536,22,155),(537,22,154),(542,22,160),(543,22,199),(546,22,185),(547,22,186),(552,22,250),(554,22,205),(555,22,208),(556,22,167),(557,22,166),(562,22,227),(563,22,226),(566,22,233),(567,22,172),(568,22,181),(570,22,177),(572,22,173),(576,22,72),(577,22,71),(579,22,66),(580,22,65),(640,22,149),(641,22,156),(642,22,61),(643,22,168),(644,22,232),(645,22,231),(646,22,225),(647,22,174);
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
) ENGINE=InnoDB AUTO_INCREMENT=213 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (197,1,1),(198,1,2),(199,1,3),(201,1,5),(202,1,28),(203,1,30),(204,1,78),(207,1,139),(208,1,156),(209,1,171),(210,1,172),(211,1,4),(212,1,173);
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
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `uname_2` (`uname`),
  KEY `join_ip` (`join_ip`),
  KEY `status` (`status`),
  KEY `phone` (`phone`),
  KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` (`id`, `uname`, `pass`, `nickname`, `icon`, `summary`, `desc`, `join_ip`, `device`, `phone`, `email`, `status`, `created_at`, `updated_at`) VALUES (1,'ciel','$2a$10$22gvVm5gNCE2zun7e3fJ/ujXdAWzVv.M5.Rp7UwW3hZkFBkTOztvu','ciel','','测试完成啦','这是一个测试用户','127.0.0.1','','','',1,'2022-09-01 22:38:06','2022-09-03 23:45:33'),(3,'morri33','$2a$10$XU7ft6XGnu5l2Wo5JSKCYutEA1G1rlPyVF7bcs8mgbHOzoqs2AaGG','morri','','','','','','','',1,'2022-09-03 23:12:47','2022-09-03 23:45:13');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user_login_log`
--

LOCK TABLES `u_user_login_log` WRITE;
/*!40000 ALTER TABLE `u_user_login_log` DISABLE KEYS */;
INSERT INTO `u_user_login_log` (`id`, `uid`, `ip`, `created_at`, `updated_at`) VALUES (1,1,'1','2022-09-03 14:49:17','2022-09-03 14:49:17');
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

-- Dump completed on 2022-09-04 14:31:02
