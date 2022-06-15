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
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `status`, `created_at`, `updated_at`) VALUES (1,1,'admin','$2a$10$2wtzucESYl2r5eARjdcefeg2/4bbaWxFqyeG1C5clSPOYjrowXG.G',1,'2022-03-08 08:59:33','2022-06-12 05:54:00');
/*!40000 ALTER TABLE `s_admin` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (32,'/user/path','GET','用户','用户列表页面',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(33,'/user','GET','用户','查询用户列表集合',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(34,'/user/:id','GET','用户','查询用户列表详情',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(35,'/user/:id','DELETE','用户','删除用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(36,'/user','POST','用户','添加用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(37,'/user','PUT','用户','修改用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(38,'/loginLog/path','GET','用户','login_log页面',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(39,'/loginLog','GET','用户','查询login_log集合',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(40,'/loginLog/:id','GET','用户','查询login_log详情',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(41,'/loginLog/:id','DELETE','用户','删除login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(42,'/loginLog','POST','用户','添加login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(43,'/loginLog','PUT','用户','修改login_log',1,'2022-06-11 18:48:30','2022-06-14 13:00:11'),(44,'/dict/path','GET','系统','字典页面',1,'2022-06-11 20:13:38','2022-06-14 13:29:13'),(45,'/dict','GET','字典','查询字典集合',1,'2022-06-11 20:13:38','2022-06-14 13:48:41'),(46,'/dict/:id','GET','字典','查询字典详情',1,'2022-06-11 20:13:38','2022-06-14 13:48:34'),(47,'/dict/:id','DELETE','字典','删除字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:29'),(48,'/dict','POST','字典','添加字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:19'),(49,'/dict','PUT','字典','修改字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:11'),(50,'/admin/path','GET','管理员','管理员页面',1,'2022-06-11 20:26:21','2022-06-14 13:48:03'),(51,'/admin','GET','管理员','查询管理员集合',1,'2022-06-11 20:26:21','2022-06-14 13:47:58'),(52,'/admin/:id','GET','管理员','查询管理员详情',1,'2022-06-11 20:26:21','2022-06-14 13:47:49'),(53,'/admin/:id','DELETE','管理员','删除管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:44'),(54,'/admin','POST','管理员','添加管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:35'),(55,'/admin','PUT','管理员','修改管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:29'),(56,'/file/path','GET','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(57,'/file','GET','文件','查询文件集合',1,'2022-06-12 14:10:23','2022-06-14 13:44:54'),(58,'/file/:id','GET','文件','查询文件详情',1,'2022-06-12 14:10:23','2022-06-14 13:44:49'),(59,'/file/:id','DELETE','文件','删除文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:35'),(61,'/file','PUT','文件','修改文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:26'),(62,'/roleApi/path','GET','API','禁用角色API页面',1,'2022-06-12 17:02:13','2022-06-14 13:44:19'),(63,'/roleApi','GET','API','查询角色禁用API集合',1,'2022-06-12 17:02:13','2022-06-14 13:44:14'),(65,'/roleApi/:id','DELETE','API','删除角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:08'),(66,'/roleApi','POST','API','添加角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:01'),(68,'/roleMenu/path','GET','菜单','角色菜单页面',1,'2022-06-13 15:19:45','2022-06-14 13:43:53'),(69,'/roleMenu','GET','菜单','查询角色菜单集合',1,'2022-06-13 15:19:45','2022-06-14 13:43:44'),(71,'/roleMenu/:id','DELETE','菜单','删除角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:43:39'),(72,'/roleMenu','POST','菜单','添加角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:37:48'),(74,'/operationLog/path','GET','日志','操作日志页面',1,'2022-06-13 19:59:57','2022-06-14 13:37:43'),(75,'/operationLog','GET','日志','查询操作日志集合',1,'2022-06-13 19:59:57','2022-06-14 13:37:32'),(76,'/operationLog/:id','GET','日志','查询操作日志详情',1,'2022-06-13 19:59:57','2022-06-14 13:36:29'),(77,'/operationLog/:id','DELETE','日志','删除操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:50'),(78,'/operationLog','POST','日志','添加操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:22'),(79,'/operationLog','PUT','日志','修改操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:11');
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
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'api_group','系统,日志,菜单,API,字典,管理员,文件,用户','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-06-14 13:47:08'),(22,'221','test','2','2',1,2,'2022-03-08 16:36:11','2022-06-14 13:43:18');
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
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (26,'1/2022/03/BYFY4d.gif',1,1,'2022-03-27 20:19:07','2022-03-27 20:19:07'),(27,'1/2022/03/FdI4Yw.gif',1,1,'2022-03-27 20:19:07','2022-03-27 20:19:07'),(29,'1/2022/03/mAMoWX.png',1,1,'2022-03-28 15:28:47','2022-03-28 15:28:47'),(30,'1/2022/03/2S41in.png',1,1,'2022-03-28 15:32:00','2022-03-28 15:32:00'),(31,'1/2022/03/IdGUqj.png',1,1,'2022-03-28 15:36:45','2022-03-28 15:36:45'),(32,'1/2022/03/5Eoxb1.png',1,1,'2022-03-28 15:40:17','2022-06-12 15:08:09');
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
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` int NOT NULL DEFAULT '1' COMMENT '1normal 2group',
  `sort` decimal(7,2) NOT NULL DEFAULT '0.00',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `name`, `path`, `type`, `sort`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','系统','',2,1.00,1,'2022-02-27 06:15:01','2022-04-01 07:51:06'),(2,1,'1/2022/03/FdI4Yw.gif','菜单','/menu/path',1,1.10,1,'2022-02-16 11:14:13','2022-03-27 12:19:17'),(3,1,'1/2022/03/IdGUqj.png','角色','/role/path',1,1.30,1,'2022-03-04 08:57:14','2022-03-28 07:36:59'),(4,1,'1/2022/03/BYFY4d.gif','API','/api/path',1,1.20,1,'2022-02-27 05:04:56','2022-03-27 12:19:27'),(5,1,'1/2022/03/5Eoxb1.png','管理员','/admin/path',1,1.40,1,'2022-03-08 07:45:04','2022-06-03 12:06:38'),(28,1,'1/2022/03/mAMoWX.png','字典','/dict/path',1,1.50,1,'2022-02-27 12:51:48','2022-04-01 11:32:15'),(30,1,'1/2022/03/2S41in.png','文件','/file/path',1,1.60,1,'2022-03-08 08:05:30','2022-04-01 11:23:43'),(59,1,'1/2022/03/5Eoxb1.png','代码生成','/gen/path',1,1.70,1,'2022-04-01 14:41:45','2022-04-01 15:16:27'),(73,-1,'','用户','',2,2.00,1,'2022-06-11 12:04:09','2022-06-11 12:04:09'),(74,73,'https://i.scdn.co/image/ab67706f000000035b2f3fe92b9ba1f4d536e231','用户列表','/user/path',1,2.10,1,'2022-06-11 12:04:09','2022-06-11 12:04:09'),(75,73,'https://www.gravatar.com/avatar/2442588627?d=monsterid&f=y','登陆日志','/loginLog/path',1,2.20,1,'2022-06-11 12:05:08','2022-06-13 12:17:54'),(78,1,'https://www.gravatar.com/avatar/302255901?d=robohash&f=y','操作日志','/operationLog/path',1,1.92,1,'2022-06-13 11:59:57','2022-06-13 12:00:26');
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
) ENGINE=InnoDB AUTO_INCREMENT=138 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (137,1,'http://localhost:1211/operationLog//batch?ids[]=136&ids[]=135&ids[]=134&ids[]=133&ids[]=132&ids[]=131&ids[]=130&ids[]=129','{\"code\":0,\"msg\":\"ok\"}','DELETE','/operationLog/batch','::1',9,'2022-06-15 20:36:44');
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
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'Super Admin','2022-02-16 11:12:41','2022-02-21 04:46:24'),(2,'Admin','2022-02-16 11:13:11','2022-06-13 09:06:46');
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
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(84,1,59),(91,1,73),(92,1,74),(93,1,75),(100,1,78);
/*!40000 ALTER TABLE `s_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_login_log`
--

DROP TABLE IF EXISTS `u_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_login_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `ip` char(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_login_log`
--

LOCK TABLES `u_login_log` WRITE;
/*!40000 ALTER TABLE `u_login_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `u_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `u_user`
--

DROP TABLE IF EXISTS `u_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `u_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uname` char(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pwd` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` (`id`, `uname`, `nickname`, `pwd`, `email`, `status`, `created_at`, `updated_at`) VALUES (33,'test','22','','33',1,'2022-06-15 12:30:47','2022-06-15 12:30:54');
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

-- Dump completed on 2022-06-15 20:37:01
