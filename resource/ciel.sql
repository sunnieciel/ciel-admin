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
  `uid` int NOT NULL,
  `level` int DEFAULT '1',
  `tag` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `main_things` longtext COLLATE utf8mb4_general_ci,
  `other_info` longtext COLLATE utf8mb4_general_ci,
  `happen_date` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `f_node_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=391 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `f_node`
--
--
-- Table structure for table `f_thing`
--

DROP TABLE IF EXISTS `f_thing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `f_thing` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `type` int DEFAULT NULL,
  `begin_time` datetime DEFAULT NULL,
  `status` int DEFAULT NULL COMMENT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `f_thing`
--

-- Table structure for table `f_thing_record`
--

DROP TABLE IF EXISTS `f_thing_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `f_thing_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tid` bigint NOT NULL,
  `record_date` datetime DEFAULT NULL,
  `score` decimal(10,0) DEFAULT NULL,
  `day` int DEFAULT NULL,
  `desc` text COLLATE utf8mb4_general_ci,
  `created_at` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`tid`),
  CONSTRAINT `f_thing_record_ibfk_1` FOREIGN KEY (`tid`) REFERENCES `f_thing` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `f_thing_record`
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
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel','$2a$10$2wtzucESYl2r5eARjdcefeg2/4bbaWxFqyeG1C5clSPOYjrowXG.G',1,'2022-03-08 08:59:33','2022-06-17 07:29:49');
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
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (32,'/user/path','GET','用户','用户列表页面',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(33,'/user','GET','用户','查询用户列表集合',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(34,'/user/:id','GET','用户','查询用户列表详情',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(35,'/user/:id','DELETE','用户','删除用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(36,'/user','POST','用户','添加用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(37,'/user','PUT','用户','修改用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(38,'/loginLog/path','GET','用户','login_log页面',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(39,'/loginLog','GET','用户','查询login_log集合',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(40,'/loginLog/:id','GET','用户','查询login_log详情',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(41,'/loginLog/:id','DELETE','用户','删除login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(42,'/loginLog','POST','用户','添加login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(43,'/loginLog','PUT','用户','修改login_log',1,'2022-06-11 18:48:30','2022-06-14 13:00:11'),(44,'/dict/path','GET','系统','字典页面',1,'2022-06-11 20:13:38','2022-06-14 13:29:13'),(45,'/dict','GET','字典','查询字典集合',1,'2022-06-11 20:13:38','2022-06-14 13:48:41'),(46,'/dict/:id','GET','字典','查询字典详情',1,'2022-06-11 20:13:38','2022-06-14 13:48:34'),(47,'/dict/:id','DELETE','字典','删除字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:29'),(48,'/dict','POST','字典','添加字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:19'),(49,'/dict','PUT','字典','修改字典',1,'2022-06-11 20:13:38','2022-06-14 13:48:11'),(50,'/admin/path','GET','管理员','管理员页面',1,'2022-06-11 20:26:21','2022-06-14 13:48:03'),(51,'/admin','GET','管理员','查询管理员集合',1,'2022-06-11 20:26:21','2022-06-14 13:47:58'),(52,'/admin/:id','GET','管理员','查询管理员详情',1,'2022-06-11 20:26:21','2022-06-14 13:47:49'),(53,'/admin/:id','DELETE','管理员','删除管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:44'),(54,'/admin','POST','管理员','添加管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:35'),(55,'/admin','PUT','管理员','修改管理员',1,'2022-06-11 20:26:21','2022-06-14 13:47:29'),(56,'/file/path','GET','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(57,'/file','GET','文件','查询文件集合',1,'2022-06-12 14:10:23','2022-06-14 13:44:54'),(58,'/file/:id','GET','文件','查询文件详情',1,'2022-06-12 14:10:23','2022-06-14 13:44:49'),(59,'/file/:id','DELETE','文件','删除文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:35'),(61,'/file','PUT','文件','修改文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:26'),(62,'/roleApi/path','GET','API','禁用角色API页面',1,'2022-06-12 17:02:13','2022-06-14 13:44:19'),(63,'/roleApi','GET','API','查询角色禁用API集合',1,'2022-06-12 17:02:13','2022-06-14 13:44:14'),(65,'/roleApi/:id','DELETE','API','删除角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:08'),(66,'/roleApi','POST','API','添加角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:01'),(68,'/roleMenu/path','GET','菜单','角色菜单页面',1,'2022-06-13 15:19:45','2022-06-14 13:43:53'),(69,'/roleMenu','GET','菜单','查询角色菜单集合',1,'2022-06-13 15:19:45','2022-06-14 13:43:44'),(71,'/roleMenu/:id','DELETE','菜单','删除角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:43:39'),(72,'/roleMenu','POST','菜单','添加角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:37:48'),(74,'/operationLog/path','GET','日志','操作日志页面',1,'2022-06-13 19:59:57','2022-06-14 13:37:43'),(75,'/operationLog','GET','日志','查询操作日志集合',1,'2022-06-13 19:59:57','2022-06-14 13:37:32'),(76,'/operationLog/:id','GET','日志','查询操作日志详情',1,'2022-06-13 19:59:57','2022-06-14 13:36:29'),(77,'/operationLog/:id','DELETE','日志','删除操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:50'),(78,'/operationLog','POST','日志','添加操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:22'),(79,'/operationLog','PUT','日志','修改操作日志',1,'2022-06-13 19:59:57','2022-06-14 13:35:11'),(82,'/node/path','GET','记事本','node页面',1,'2022-06-16 23:34:50','2022-06-19 15:42:22'),(83,'/node','GET','记事本','查询node集合',1,'2022-06-16 23:34:50','2022-06-19 15:42:19'),(84,'/node/:id','GET','记事本','查询node详情',1,'2022-06-16 23:34:50','2022-06-19 15:42:15'),(85,'/node/:id','DELETE','记事本','删除node',1,'2022-06-16 23:34:50','2022-06-19 15:42:12'),(86,'/node','POST','记事本','添加node',1,'2022-06-16 23:34:50','2022-06-19 15:42:09'),(87,'/node','PUT','记事本','修改node',1,'2022-06-16 23:34:50','2022-06-19 15:42:06'),(88,'/thing/path','GET','事件','thing页面',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(89,'/thing','GET','事件','查询thing集合',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(90,'/thing/:id','GET','事件','查询thing详情',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(91,'/thing/:id','DELETE','事件','删除thing',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(92,'/thing','POST','事件','添加thing',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(93,'/thing','PUT','事件','修改thing',1,'2022-06-21 19:14:43','2022-06-21 19:14:43'),(94,'/thingRecord/path','GET','事件','thing_record页面',1,'2022-06-21 19:27:17','2022-06-21 19:27:17'),(95,'/thingRecord','GET','事件','查询thing_record集合',1,'2022-06-21 19:27:17','2022-06-22 15:33:47'),(96,'/thingRecord/:id','GET','事件','查询thing_record详情',1,'2022-06-21 19:27:17','2022-06-22 15:33:34'),(97,'/thingRecord/:id','DELETE','事件','删除thing_record',1,'2022-06-21 19:27:17','2022-06-22 15:33:28'),(98,'/thingRecord','POST','事件','添加thing_record',1,'2022-06-21 19:27:17','2022-06-22 15:33:23'),(99,'/thingRecord','PUT','事件','修改thing_record',1,'2022-06-21 19:27:17','2022-06-22 15:33:13');
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
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'api_group','系统,日志,菜单,API,字典,管理员,文件,用户,记事本,事件','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-06-22 15:33:05'),(22,'music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-06-21 16:30:47');
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
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `name`, `path`, `type`, `sort`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','系统','',2,1.00,1,'2022-02-27 06:15:01','2022-04-01 07:51:06'),(2,1,'1/2022/03/FdI4Yw.gif','菜单','/menu/path',1,1.10,1,'2022-02-16 11:14:13','2022-03-27 12:19:17'),(3,1,'1/2022/03/IdGUqj.png','角色','/role/path',1,1.30,1,'2022-03-04 08:57:14','2022-03-28 07:36:59'),(4,1,'1/2022/03/BYFY4d.gif','API','/api/path',1,1.20,1,'2022-02-27 05:04:56','2022-03-27 12:19:27'),(5,1,'1/2022/03/5Eoxb1.png','管理员','/admin/path',1,1.40,1,'2022-03-08 07:45:04','2022-06-03 12:06:38'),(28,1,'1/2022/03/mAMoWX.png','字典','/dict/path',1,1.50,1,'2022-02-27 12:51:48','2022-04-01 11:32:15'),(30,1,'1/2022/03/2S41in.png','文件','/file/path',1,1.60,1,'2022-03-08 08:05:30','2022-04-01 11:23:43'),(59,1,'1/2022/03/5Eoxb1.png','代码生成','/gen/path',1,1.70,1,'2022-04-01 14:41:45','2022-04-01 15:16:27'),(73,-1,'','用户','',2,2.00,1,'2022-06-11 12:04:09','2022-06-11 12:04:09'),(74,73,'','用户列表','/user/path',1,2.10,1,'2022-06-11 12:04:09','2022-06-16 15:16:31'),(75,73,'','登陆日志','/loginLog/path',1,2.20,1,'2022-06-11 12:05:08','2022-06-16 15:16:27'),(78,1,'https://i.scdn.co/image/ab67706f000000035b2f3fe92b9ba1f4d536e231','操作日志','/operationLog/path',1,1.92,1,'2022-06-13 11:59:57','2022-06-16 06:02:22'),(84,-1,'','好玩','',2,3.00,1,'2022-06-16 15:34:50','2022-06-16 15:34:50'),(85,84,'','记事本','/node/path',1,3.10,1,'2022-06-16 15:34:50','2022-06-17 07:07:38'),(86,84,'https://www.gravatar.com/avatar/3233983405?d=identicon&f=y','计时器','/thing/path',1,3.20,1,'2022-06-21 11:14:43','2022-06-21 11:14:43'),(87,84,'https://www.gravatar.com/avatar/3014341940?d=wavatar&f=y','计时列表','/thingRecord/path',1,3.30,1,'2022-06-21 11:27:17','2022-06-21 11:27:17');
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
) ENGINE=InnoDB AUTO_INCREMENT=172 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (137,1,'http://localhost:1211/operationLog//batch?ids[]=136&ids[]=135&ids[]=134&ids[]=133&ids[]=132&ids[]=131&ids[]=130&ids[]=129','{\"code\":0,\"msg\":\"ok\"}','DELETE','/operationLog/batch','::1',9,'2022-06-15 20:36:44'),(138,1,'created_at=2022-06-11+20%3A04%3A09&icon=&id=74&name=%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8&path=%2Fuser%2Fpath&pid=73&sort=2.1&status=1&type=1&updated_at=2022-06-11+20%3A04%3A09','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',10,'2022-06-16 14:02:17'),(139,1,'created_at=2022-06-13+19%3A59%3A57&icon=https%3A%2F%2Fi.scdn.co%2Fimage%2Fab67706f000000035b2f3fe92b9ba1f4d536e231&id=78&name=%E6%93%8D%E4%BD%9C%E6%97%A5%E5%BF%97&path=%2FoperationLog%2Fpath&pid=1&sort=1.92&status=1&type=1&updated_at=2022-06-13+20%3A00%3A26','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',9,'2022-06-16 14:02:22'),(140,1,'created_at=2022-06-11+20%3A04%3A09&icon=https%3A%2F%2Fwww.gravatar.com%2Favatar%2F302255901%3Fd%3Drobohash%26f%3Dy&id=74&name=%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8&path=%2Fuser%2Fpath&pid=73&sort=2.1&status=1&type=1&updated_at=2022-06-16+14%3A02%3A17','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',10,'2022-06-16 14:02:30'),(141,1,'created_at=2022-06-11+20%3A05%3A08&icon=&id=75&name=%E7%99%BB%E9%99%86%E6%97%A5%E5%BF%97&path=%2FloginLog%2Fpath&pid=73&sort=2.2&status=1&type=1&updated_at=2022-06-13+20%3A17%3A54','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',5,'2022-06-16 23:16:27'),(142,1,'created_at=2022-06-11+20%3A04%3A09&icon=&id=74&name=%E7%94%A8%E6%88%B7%E5%88%97%E8%A1%A8&path=%2Fuser%2Fpath&pid=73&sort=2.1&status=1&type=1&updated_at=2022-06-16+14%3A02%3A30','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',9,'2022-06-16 23:16:31'),(143,1,'fields%5B0%5D%5BComment%5D=&fields%5B0%5D%5BDefault%5D=&fields%5B0%5D%5BEditDisabled%5D=0&fields%5B0%5D%5BEditHide%5D=1&fields%5B0%5D%5BExtra%5D=auto_increment&fields%5B0%5D%5BFieldType%5D=text&fields%5B0%5D%5BIndex%5D=0&fields%5B0%5D%5BKey%5D=PRI&fields%5B0%5D%5BLabel%5D=id&fields%5B0%5D%5BName%5D=id&fields%5B0%5D%5BNotShow%5D=0&fields%5B0%5D%5BNull%5D=false&fields%5B0%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B0%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B0%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B0%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B0%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B0%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B0%5D%5BQueryName%5D=id&fields%5B0%5D%5BSearchType%5D=0&fields%5B0%5D%5BType%5D=bigint&fields%5B1%5D%5BComment%5D=&fields%5B1%5D%5BDefault%5D=&fields%5B1%5D%5BEditDisabled%5D=0&fields%5B1%5D%5BEditHide%5D=0&fields%5B1%5D%5BExtra%5D=&fields%5B1%5D%5BFieldType%5D=text&fields%5B1%5D%5BIndex%5D=1&fields%5B1%5D%5BKey%5D=MUL&fields%5B1%5D%5BLabel%5D=%E7%94%A8%E6%88%B7ID&fields%5B1%5D%5BName%5D=uid&fields%5B1%5D%5BNotShow%5D=1&fields%5B1%5D%5BNull%5D=false&fields%5B1%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B1%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B1%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B1%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B1%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B1%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B1%5D%5BQueryName%5D=uid&fields%5B1%5D%5BSearchType%5D=0&fields%5B1%5D%5BType%5D=int&fields%5B2%5D%5BComment%5D=&fields%5B2%5D%5BEditDisabled%5D=0&fields%5B2%5D%5BEditHide%5D=1&fields%5B2%5D%5BFieldType%5D=text&fields%5B2%5D%5BIndex%5D=2&fields%5B2%5D%5BLabel%5D=%E7%94%A8%E6%88%B7%E5%90%8D&fields%5B2%5D%5BName%5D=t2.uname&fields%5B2%5D%5BNotShow%5D=0&fields%5B2%5D%5BOptions%5D%5B0%5D%5BName%5D=&fields%5B2%5D%5BOptions%5D%5B0%5D%5BType%5D=primary&fields%5B2%5D%5BOptions%5D%5B0%5D%5BValue%5D=&fields%5B2%5D%5BQueryName%5D=uname&fields%5B2%5D%5BSearchType%5D=2&fields%5B3%5D%5BComment%5D=&fields%5B3%5D%5BDefault%5D=1&fields%5B3%5D%5BEditDisabled%5D=0&fields%5B3%5D%5BEditHide%5D=0&fields%5B3%5D%5BExtra%5D=&fields%5B3%5D%5BFieldType%5D=select&fields%5B3%5D%5BIndex%5D=2&fields%5B3%5D%5BKey%5D=&fields%5B3%5D%5BLabel%5D=%E7%AD%89%E7%BA%A7&fields%5B3%5D%5BName%5D=level&fields%5B3%5D%5BNotShow%5D=0&fields%5B3%5D%5BNull%5D=true&fields%5B3%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%99%AE%E9%80%9A&fields%5B3%5D%5BOptions%5D%5B0%5D%5BType%5D=info&fields%5B3%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B3%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A8%80%E6%9C%89&fields%5B3%5D%5BOptions%5D%5B1%5D%5BType%5D=success&fields%5B3%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B3%5D%5BOptions%5D%5B2%5D%5BName%5D=%E4%BC%A0%E6%89%BF&fields%5B3%5D%5BOptions%5D%5B2%5D%5BType%5D=primary&fields%5B3%5D%5BOptions%5D%5B2%5D%5BValue%5D=3&fields%5B3%5D%5BOptions%5D%5B3%5D%5BName%5D=%E5%94%AF%E4%B8%80&fields%5B3%5D%5BOptions%5D%5B3%5D%5BType%5D=warning&fields%5B3%5D%5BOptions%5D%5B3%5D%5BValue%5D=4&fields%5B3%5D%5BOptions%5D%5B4%5D%5BName%5D=%E5%8F%B2%E8%AF%97&fields%5B3%5D%5BOptions%5D%5B4%5D%5BType%5D=danger&fields%5B3%5D%5BOptions%5D%5B4%5D%5BValue%5D=5&fields%5B3%5D%5BQueryName%5D=level&fields%5B3%5D%5BSearchType%5D=1&fields%5B3%5D%5BType%5D=int&fields%5B4%5D%5BComment%5D=&fields%5B4%5D%5BDefault%5D=&fields%5B4%5D%5BEditDisabled%5D=0&fields%5B4%5D%5BEditHide%5D=0&fields%5B4%5D%5BExtra%5D=&fields%5B4%5D%5BFieldType%5D=text&fields%5B4%5D%5BIndex%5D=3&fields%5B4%5D%5BKey%5D=&fields%5B4%5D%5BLabel%5D=%E6%A0%87%E7%AD%BE&fields%5B4%5D%5BName%5D=tag&fields%5B4%5D%5BNotShow%5D=0&fields%5B4%5D%5BNull%5D=true&fields%5B4%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B4%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B4%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B4%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B4%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B4%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B4%5D%5BQueryName%5D=tag&fields%5B4%5D%5BSearchType%5D=2&fields%5B4%5D%5BType%5D=varchar%2864%29&fields%5B5%5D%5BComment%5D=&fields%5B5%5D%5BDefault%5D=&fields%5B5%5D%5BEditDisabled%5D=0&fields%5B5%5D%5BEditHide%5D=0&fields%5B5%5D%5BExtra%5D=&fields%5B5%5D%5BFieldType%5D=text&fields%5B5%5D%5BIndex%5D=4&fields%5B5%5D%5BKey%5D=&fields%5B5%5D%5BLabel%5D=%E4%B8%BB%E8%A6%81%E4%BA%8B%E4%BB%B6&fields%5B5%5D%5BName%5D=main_things&fields%5B5%5D%5BNotShow%5D=0&fields%5B5%5D%5BNull%5D=true&fields%5B5%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B5%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B5%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B5%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B5%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B5%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B5%5D%5BQueryName%5D=main_things&fields%5B5%5D%5BSearchType%5D=2&fields%5B5%5D%5BType%5D=longtext&fields%5B6%5D%5BComment%5D=&fields%5B6%5D%5BDefault%5D=&fields%5B6%5D%5BEditDisabled%5D=0&fields%5B6%5D%5BEditHide%5D=0&fields%5B6%5D%5BExtra%5D=&fields%5B6%5D%5BFieldType%5D=text&fields%5B6%5D%5BIndex%5D=5&fields%5B6%5D%5BKey%5D=&fields%5B6%5D%5BLabel%5D=%E5%85%B6%E4%BB%96%E4%BF%A1%E6%81%AF&fields%5B6%5D%5BName%5D=other_info&fields%5B6%5D%5BNotShow%5D=0&fields%5B6%5D%5BNull%5D=true&fields%5B6%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B6%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B6%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B6%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B6%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B6%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B6%5D%5BQueryName%5D=other_info&fields%5B6%5D%5BSearchType%5D=2&fields%5B6%5D%5BType%5D=longtext&fields%5B7%5D%5BComment%5D=&fields%5B7%5D%5BDefault%5D=&fields%5B7%5D%5BEditDisabled%5D=0&fields%5B7%5D%5BEditHide%5D=1&fields%5B7%5D%5BExtra%5D=&fields%5B7%5D%5BFieldType%5D=text&fields%5B7%5D%5BIndex%5D=6&fields%5B7%5D%5BKey%5D=&fields%5B7%5D%5BLabel%5D=created_at&fields%5B7%5D%5BName%5D=created_at&fields%5B7%5D%5BNotShow%5D=0&fields%5B7%5D%5BNull%5D=true&fields%5B7%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B7%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B7%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B7%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B7%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B7%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B7%5D%5BQueryName%5D=created_at&fields%5B7%5D%5BSearchType%5D=0&fields%5B7%5D%5BType%5D=datetime&fields%5B8%5D%5BComment%5D=&fields%5B8%5D%5BDefault%5D=&fields%5B8%5D%5BEditDisabled%5D=0&fields%5B8%5D%5BEditHide%5D=1&fields%5B8%5D%5BExtra%5D=&fields%5B8%5D%5BFieldType%5D=text&fields%5B8%5D%5BIndex%5D=7&fields%5B8%5D%5BKey%5D=&fields%5B8%5D%5BLabel%5D=updated_at&fields%5B8%5D%5BName%5D=updated_at&fields%5B8%5D%5BNotShow%5D=0&fields%5B8%5D%5BNull%5D=true&fields%5B8%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B8%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B8%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B8%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B8%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B8%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B8%5D%5BQueryName%5D=updated_at&fields%5B8%5D%5BSearchType%5D=0&fields%5B8%5D%5BType%5D=datetime&genConf%5BAddBtn%5D=0&genConf%5BDelBtn%5D=0&genConf%5BHtmlGroup%5D=f&genConf%5BIcon%5D=&genConf%5BMenuLevel1%5D=%E5%A5%BD%E7%8E%A9&genConf%5BMenuLevel2%5D=%E6%97%A5%E8%AE%B0&genConf%5BOrderBy%5D=t1.id+desc&genConf%5BPageDesc%5D=%E8%BF%99%E9%87%8C%E6%98%AFnode%E9%A1%B5%E9%9D%A2%E5%8F%AF%E4%BB%A5%E5%AF%B9%E6%95%B0%E6%8D%AE%E8%BF%9B%E8%A1%8C%E7%9B%B8%E5%BA%94%E7%9A%84%E6%93%8D%E4%BD%9C%E3%80%82&genConf%5BPageName%5D=node&genConf%5BQueryField%5D=t1.%2A%2Ct2.uname&genConf%5BStructName%5D=node&genConf%5BT1%5D=f_node&genConf%5BT2%5D=s_admin+t2+on+t1.uid+%3D+t2.id&genConf%5BT3%5D=&genConf%5BT4%5D=&genConf%5BT5%5D=&genConf%5BT6%5D=&genConf%5BUpdateBtn%5D=0&genConf%5BUrlPrefix%5D=','{\"code\":0,\"msg\":\"ok\"}','POST','/gen','::1',24,'2022-06-16 23:34:50'),(144,1,'mid%5B%5D=84&mid%5B%5D=85&rid=1','{\"code\":0,\"msg\":\"ok\"}','POST','/roleMenu','::1',9,'2022-06-16 23:35:05'),(145,1,'created_at=2022-06-16+23%3A34%3A50&icon=&id=85&name=%E6%97%A5%E8%AE%B0&path=%2Fnode%2Fpath&pid=84&sort=3.1&status=1&type=1&updated_at=2022-06-16+23%3A34%3A50','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',2,'2022-06-16 23:35:46'),(146,1,'created_at=2022-06-16+23%3A34%3A50&icon=&id=85&name=%E8%AE%B0%E4%BA%8B%E6%9C%AC&path=%2Fnode%2Fpath&pid=84&sort=3.1&status=1&type=1&updated_at=2022-06-16+23%3A35%3A46','{\"code\":0,\"msg\":\"ok\"}','PUT','/menu','::1',10,'2022-06-17 15:07:38'),(147,1,'id=1&uname=ciel','{\"code\":0,\"msg\":\"ok\"}','PUT','/admin/updateUname','::1',46,'2022-06-17 15:29:49'),(148,1,'created_at=2022-02-27+20%3A40%3A57&desc=API%E5%88%86%E7%BB%84%E9%80%89%E9%A1%B9&group=1&id=11&k=api_group&status=1&type=1&updated_at=2022-06-14+13%3A47%3A08&v=%E7%B3%BB%E7%BB%9F%2C%E6%97%A5%E5%BF%97%2C%E8%8F%9C%E5%8D%95%2CAPI%2C%E5%AD%97%E5%85%B8%2C%E7%AE%A1%E7%90%86%E5%91%98%2C%E6%96%87%E4%BB%B6%2C%E7%94%A8%E6%88%B7%2C%E8%AE%B0%E4%BA%8B%E6%9C%AC','{\"code\":0,\"msg\":\"ok\"}','PUT','/dict','::1',5,'2022-06-19 15:42:01'),(149,1,'created_at=2022-06-16+23%3A34%3A50&desc=%E4%BF%AE%E6%94%B9node&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=87&method=PUT&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',10,'2022-06-19 15:42:06'),(150,1,'created_at=2022-06-16+23%3A34%3A50&desc=%E6%B7%BB%E5%8A%A0node&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=86&method=POST&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',2,'2022-06-19 15:42:09'),(151,1,'created_at=2022-06-16+23%3A34%3A50&desc=%E5%88%A0%E9%99%A4node&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=85&method=DELETE&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode%2F%3Aid','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-19 15:42:12'),(152,1,'created_at=2022-06-16+23%3A34%3A50&desc=%E6%9F%A5%E8%AF%A2node%E8%AF%A6%E6%83%85&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=84&method=GET&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode%2F%3Aid','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-19 15:42:15'),(153,1,'created_at=2022-06-16+23%3A34%3A50&desc=%E6%9F%A5%E8%AF%A2node%E9%9B%86%E5%90%88&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=83&method=GET&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-19 15:42:19'),(154,1,'created_at=2022-06-16+23%3A34%3A50&desc=node%E9%A1%B5%E9%9D%A2&group=%E8%AE%B0%E4%BA%8B%E6%9C%AC&id=82&method=GET&status=1&updated_at=2022-06-16+23%3A34%3A50&url=%2Fnode%2Fpath','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-19 15:42:22'),(155,1,'created_at=2022-03-08+16%3A36%3A11&desc=%E7%99%BB%E9%99%86%E9%9F%B3%E4%B9%90%E5%88%97%E8%A1%A8&group=1&id=22&k=music-url&status=1&type=1&updated_at=2022-06-14+13%3A43%3A18&v=https%3A%2F%2Fwww.youtube.com%2Fembed%2Fvideoseries%3Flist%3DPLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','{\"code\":0,\"msg\":\"ok\"}','PUT','/dict','::1',9,'2022-06-21 16:19:31'),(156,1,'created_at=2022-03-08+16%3A36%3A11&desc=%E7%99%BB%E9%99%86%E9%9F%B3%E4%B9%90%E5%88%97%E8%A1%A8&group=1&id=22&k=music-url&status=1&type=1&updated_at=2022-06-21+16%3A19%3A31&v=https%3A%2F%2Fwww.bilibili.com%2Fvideo%2FBV1Dh41167W4%3Fshare_source%3Dcopy_web','{\"code\":0,\"msg\":\"ok\"}','PUT','/dict','::1',10,'2022-06-21 16:30:47'),(157,1,'fields%5B0%5D%5BComment%5D=&fields%5B0%5D%5BDefault%5D=&fields%5B0%5D%5BEditDisabled%5D=0&fields%5B0%5D%5BEditHide%5D=1&fields%5B0%5D%5BExtra%5D=auto_increment&fields%5B0%5D%5BFieldType%5D=text&fields%5B0%5D%5BIndex%5D=0&fields%5B0%5D%5BKey%5D=PRI&fields%5B0%5D%5BLabel%5D=id&fields%5B0%5D%5BName%5D=id&fields%5B0%5D%5BNotShow%5D=0&fields%5B0%5D%5BNull%5D=false&fields%5B0%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B0%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B0%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B0%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B0%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B0%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B0%5D%5BQueryName%5D=id&fields%5B0%5D%5BSearchType%5D=0&fields%5B0%5D%5BType%5D=bigint&fields%5B1%5D%5BComment%5D=&fields%5B1%5D%5BDefault%5D=&fields%5B1%5D%5BEditDisabled%5D=0&fields%5B1%5D%5BEditHide%5D=0&fields%5B1%5D%5BExtra%5D=&fields%5B1%5D%5BFieldType%5D=text&fields%5B1%5D%5BIndex%5D=1&fields%5B1%5D%5BKey%5D=&fields%5B1%5D%5BLabel%5D=name&fields%5B1%5D%5BName%5D=name&fields%5B1%5D%5BNotShow%5D=0&fields%5B1%5D%5BNull%5D=false&fields%5B1%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B1%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B1%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B1%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B1%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B1%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B1%5D%5BQueryName%5D=name&fields%5B1%5D%5BSearchType%5D=0&fields%5B1%5D%5BType%5D=varchar%2864%29&fields%5B2%5D%5BComment%5D=&fields%5B2%5D%5BDefault%5D=&fields%5B2%5D%5BEditDisabled%5D=0&fields%5B2%5D%5BEditHide%5D=0&fields%5B2%5D%5BExtra%5D=&fields%5B2%5D%5BFieldType%5D=text&fields%5B2%5D%5BIndex%5D=2&fields%5B2%5D%5BKey%5D=&fields%5B2%5D%5BLabel%5D=type&fields%5B2%5D%5BName%5D=type&fields%5B2%5D%5BNotShow%5D=0&fields%5B2%5D%5BNull%5D=true&fields%5B2%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B2%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B2%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B2%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B2%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B2%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B2%5D%5BQueryName%5D=type&fields%5B2%5D%5BSearchType%5D=0&fields%5B2%5D%5BType%5D=int&fields%5B3%5D%5BComment%5D=&fields%5B3%5D%5BDefault%5D=CURRENT_TIMESTAMP&fields%5B3%5D%5BEditDisabled%5D=0&fields%5B3%5D%5BEditHide%5D=1&fields%5B3%5D%5BExtra%5D=DEFAULT_GENERATED&fields%5B3%5D%5BFieldType%5D=text&fields%5B3%5D%5BIndex%5D=3&fields%5B3%5D%5BKey%5D=&fields%5B3%5D%5BLabel%5D=created_at&fields%5B3%5D%5BName%5D=created_at&fields%5B3%5D%5BNotShow%5D=0&fields%5B3%5D%5BNull%5D=true&fields%5B3%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B3%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B3%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B3%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B3%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B3%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B3%5D%5BQueryName%5D=created_at&fields%5B3%5D%5BSearchType%5D=0&fields%5B3%5D%5BType%5D=datetime&fields%5B4%5D%5BComment%5D=&fields%5B4%5D%5BDefault%5D=CURRENT_TIMESTAMP&fields%5B4%5D%5BEditDisabled%5D=0&fields%5B4%5D%5BEditHide%5D=1&fields%5B4%5D%5BExtra%5D=DEFAULT_GENERATED+on+update+CURRENT_TIMESTAMP&fields%5B4%5D%5BFieldType%5D=text&fields%5B4%5D%5BIndex%5D=4&fields%5B4%5D%5BKey%5D=&fields%5B4%5D%5BLabel%5D=updated_at&fields%5B4%5D%5BName%5D=updated_at&fields%5B4%5D%5BNotShow%5D=0&fields%5B4%5D%5BNull%5D=true&fields%5B4%5D%5BOptions%5D%5B0%5D%5BName%5D=%E6%AD%A3%E5%B8%B8&fields%5B4%5D%5BOptions%5D%5B0%5D%5BType%5D=success&fields%5B4%5D%5BOptions%5D%5B0%5D%5BValue%5D=1&fields%5B4%5D%5BOptions%5D%5B1%5D%5BName%5D=%E7%A6%81%E7%94%A8&fields%5B4%5D%5BOptions%5D%5B1%5D%5BType%5D=danger&fields%5B4%5D%5BOptions%5D%5B1%5D%5BValue%5D=2&fields%5B4%5D%5BQueryName%5D=updated_at&fields%5B4%5D%5BSearchType%5D=0&fields%5B4%5D%5BType%5D=datetime&genConf%5BAddBtn%5D=0&genConf%5BDelBtn%5D=0&genConf%5BHtmlGroup%5D=f&genConf%5BIcon%5D=&genConf%5BMenuLevel1%5D=%E5%A5%BD%E7%8E%A9&genConf%5BMenuLevel2%5D=%E8%AE%A1%E6%97%B6%E5%99%A8&genConf%5BOrderBy%5D=t1.id+desc&genConf%5BPageDesc%5D=%E8%BF%99%E9%87%8C%E6%98%AFthing%E9%A1%B5%E9%9D%A2%E5%8F%AF%E4%BB%A5%E5%AF%B9%E6%95%B0%E6%8D%AE%E8%BF%9B%E8%A1%8C%E7%9B%B8%E5%BA%94%E7%9A%84%E6%93%8D%E4%BD%9C%E3%80%82&genConf%5BPageName%5D=thing&genConf%5BQueryField%5D=t1.%2A&genConf%5BStructName%5D=thing&genConf%5BT1%5D=f_thing&genConf%5BT2%5D=&genConf%5BT3%5D=&genConf%5BT4%5D=&genConf%5BT5%5D=&genConf%5BT6%5D=&genConf%5BUpdateBtn%5D=0&genConf%5BUrlPrefix%5D=','{\"code\":0,\"msg\":\"ok\"}','POST','/gen','::1',23,'2022-06-21 19:14:43'),(158,1,'mid%5B%5D=86&rid=1','{\"code\":0,\"msg\":\"ok\"}','POST','/roleMenu','::1',9,'2022-06-21 19:14:59'),(162,1,'mid%5B%5D=87&rid=1','{\"code\":0,\"msg\":\"ok\"}','POST','/roleMenu','::1',10,'2022-06-21 19:27:38'),(163,1,'http://localhost:1211/operationLog//batch?ids[]=159','{\"code\":0,\"msg\":\"ok\"}','DELETE','/operationLog/batch','::1',3,'2022-06-21 20:46:30'),(164,1,'http://localhost:1211/operationLog//batch?ids[]=161','{\"code\":0,\"msg\":\"ok\"}','DELETE','/operationLog/batch','::1',2,'2022-06-21 20:46:33'),(165,1,'http://localhost:1211/operationLog//batch?ids[]=160','{\"code\":0,\"msg\":\"ok\"}','DELETE','/operationLog/batch','::1',3,'2022-06-21 20:46:35'),(166,1,'created_at=2022-02-27+20%3A40%3A57&desc=API%E5%88%86%E7%BB%84%E9%80%89%E9%A1%B9&group=1&id=11&k=api_group&status=1&type=1&updated_at=2022-06-19+15%3A42%3A01&v=%E7%B3%BB%E7%BB%9F%2C%E6%97%A5%E5%BF%97%2C%E8%8F%9C%E5%8D%95%2CAPI%2C%E5%AD%97%E5%85%B8%2C%E7%AE%A1%E7%90%86%E5%91%98%2C%E6%96%87%E4%BB%B6%2C%E7%94%A8%E6%88%B7%2C%E8%AE%B0%E4%BA%8B%E6%9C%AC%2C%E4%BA%8B%E4%BB%B6','{\"code\":0,\"msg\":\"ok\"}','PUT','/dict','::1',10,'2022-06-22 15:33:05'),(167,1,'created_at=2022-06-21+19%3A27%3A17&desc=%E4%BF%AE%E6%94%B9thing_record&group=%E4%BA%8B%E4%BB%B6&id=99&method=PUT&status=1&updated_at=2022-06-21+19%3A27%3A17&url=%2FthingRecord','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-22 15:33:13'),(168,1,'created_at=2022-06-21+19%3A27%3A17&desc=%E6%B7%BB%E5%8A%A0thing_record&group=%E4%BA%8B%E4%BB%B6&id=98&method=POST&status=1&updated_at=2022-06-21+19%3A27%3A17&url=%2FthingRecord','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-22 15:33:23'),(169,1,'created_at=2022-06-21+19%3A27%3A17&desc=%E5%88%A0%E9%99%A4thing_record&group=%E4%BA%8B%E4%BB%B6&id=97&method=DELETE&status=1&updated_at=2022-06-21+19%3A27%3A17&url=%2FthingRecord%2F%3Aid','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',2,'2022-06-22 15:33:28'),(170,1,'created_at=2022-06-21+19%3A27%3A17&desc=%E6%9F%A5%E8%AF%A2thing_record%E8%AF%A6%E6%83%85&group=%E4%BA%8B%E4%BB%B6&id=96&method=GET&status=1&updated_at=2022-06-21+19%3A27%3A17&url=%2FthingRecord%2F%3Aid','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-22 15:33:34'),(171,1,'created_at=2022-06-21+19%3A27%3A17&desc=%E6%9F%A5%E8%AF%A2thing_record%E9%9B%86%E5%90%88&group=%E4%BA%8B%E4%BB%B6&id=95&method=GET&status=1&updated_at=2022-06-21+19%3A27%3A17&url=%2FthingRecord','{\"code\":0,\"msg\":\"ok\"}','PUT','/api','::1',9,'2022-06-22 15:33:47');
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
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(84,1,59),(91,1,73),(92,1,74),(93,1,75),(100,1,78),(116,1,84),(117,1,85),(118,1,86),(119,1,87);
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

-- Dump completed on 2022-06-22 15:34:35
