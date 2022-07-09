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
) ENGINE=InnoDB AUTO_INCREMENT=427 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `status`, `created_at`, `updated_at`) VALUES (1,1,'ciel','$2a$10$vdvJdM3HNEqZrTY7eLBJ5elORDMN4exh..aFqZ66z3Xer3UULA53q',1,'2022-03-08 08:59:33','2022-07-02 11:53:04'),(42,1,'122','$2a$10$4ELJmBB5FhqvIl0oCgstHeYrpW79g4C3.Xf6541lxtBXbBZFbtVk6',1,'2022-07-02 11:28:52','2022-07-02 11:29:39');
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
) ENGINE=InnoDB AUTO_INCREMENT=199 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (56,'/file/path','1','文件','文件页面',1,'2022-06-12 14:10:23','2022-06-14 13:45:01'),(57,'/file','1','文件','查询文件集合',1,'2022-06-12 14:10:23','2022-06-14 13:44:54'),(58,'/file/:id','1','文件','查询文件详情',1,'2022-06-12 14:10:23','2022-06-14 13:44:49'),(59,'/file/:id','4','文件','删除文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:35'),(61,'/file','3','文件','修改文件',1,'2022-06-12 14:10:23','2022-06-14 13:44:26'),(62,'/roleApi/path','1','API','禁用角色API页面',1,'2022-06-12 17:02:13','2022-06-14 13:44:19'),(63,'/roleApi','1','API','查询角色禁用API集合',1,'2022-06-12 17:02:13','2022-06-14 13:44:14'),(65,'/roleApi/:id','4','API','删除角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:08'),(66,'/roleApi','2','API','添加角色禁用API',1,'2022-06-12 17:02:13','2022-06-14 13:44:01'),(68,'/roleMenu/path','1','菜单','角色菜单页面',1,'2022-06-13 15:19:45','2022-06-14 13:43:53'),(69,'/roleMenu','1','菜单','查询角色菜单集合',1,'2022-06-13 15:19:45','2022-06-14 13:43:44'),(71,'/roleMenu/:id','4','菜单','删除角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:43:39'),(72,'/roleMenu','2','菜单','添加角色菜单',1,'2022-06-13 15:19:45','2022-06-14 13:37:48'),(144,'/api/path','1','sys','api页面',1,'2022-06-26 15:14:58','2022-06-26 15:14:58'),(145,'/api/path/add','1','sys','api添加页面',1,'2022-06-26 15:14:58','2022-06-26 15:14:58'),(146,'/api/path/edit/:id','1','sys','api修改页面',1,'2022-06-26 15:14:58','2022-06-26 15:14:58'),(147,'/api/path/del/:id','4','sys','api删除操作',1,'2022-06-26 15:14:58','2022-06-26 15:18:56'),(148,'/api','2','sys','添加api',1,'2022-06-26 15:14:58','2022-06-26 15:14:58'),(149,'/api','3','sys','修改api',1,'2022-06-26 15:14:58','2022-06-26 15:14:58'),(151,'/dict/path','1','sys','dict页面',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(152,'/dict/path/add','1','sys','dict添加页面',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(153,'/dict/path/edit/:id','1','sys','dict修改页面',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(154,'/dict/path/del/:id','4','sys','dict删除操作',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(155,'/dict','2','sys','添加dict',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(156,'/dict','3','sys','修改dict',1,'2022-06-26 15:27:04','2022-06-26 15:27:04'),(157,'/operationLog/path','1','sys','operation_log页面',1,'2022-06-26 20:30:22','2022-06-26 20:30:22'),(158,'/operationLog/path/add','1','sys','operation_log添加页面',1,'2022-06-26 20:30:22','2022-06-26 20:30:22'),(159,'/operationLog/path/edit/:id','1','sys','operation_log修改页面',1,'2022-06-26 20:30:22','2022-06-26 20:30:22'),(160,'/operationLog/path/del/:id','4','sys','operation_log删除操作',1,'2022-06-26 20:30:22','2022-06-27 15:38:02'),(161,'/operationLog','2','sys','添加operation_log',1,'2022-06-26 20:30:22','2022-06-26 20:30:22'),(162,'/operationLog','3','sys','修改operation_log',1,'2022-06-26 20:30:22','2022-06-26 20:30:22'),(163,'/admin/path','1','sys','admin页面',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(164,'/admin/path/add','1','sys','admin添加页面',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(165,'/admin/path/edit/:id','1','sys','admin修改页面',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(166,'/admin/path/del/:id','4','sys','admin删除操作',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(167,'/admin','2','sys','添加admin',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(168,'/admin','3','sys','修改admin',1,'2022-06-27 16:21:07','2022-06-27 16:21:07'),(169,'/role/path','1','sys','角色页面',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(170,'/role/path/add','1','sys','角色添加页面',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(171,'/role/path/edit/:id','1','sys','角色修改页面',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(172,'/role/path/del/:id','4','sys','角色删除操作',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(173,'/role','2','sys','添加角色',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(174,'/role','3','sys','修改角色',1,'2022-06-27 18:27:46','2022-06-27 18:27:46'),(175,'/roleMenu/path/add','1','sys','role_menu添加页面',1,'2022-06-27 18:38:32','2022-06-27 18:38:32'),(176,'/roleMenu/path/edit/:id','1','sys','role_menu修改页面',1,'2022-06-27 18:38:32','2022-06-27 18:38:32'),(177,'/roleMenu/path/del/:id','4','sys','role_menu删除操作',1,'2022-06-27 18:38:32','2022-06-27 18:38:32'),(178,'/roleMenu','3','sys','修改role_menu',1,'2022-06-27 18:38:32','2022-06-27 18:38:32'),(179,'/roleApi/path/add','1','sys','role_api添加页面',1,'2022-06-27 19:10:03','2022-06-27 19:10:03'),(180,'/roleApi/path/edit/:id','1','sys','role_api修改页面',1,'2022-06-27 19:10:03','2022-06-27 19:10:03'),(181,'/roleApi/path/del/:id','4','sys','role_api删除操作',1,'2022-06-27 19:10:03','2022-06-27 19:10:03'),(182,'/roleApi','3','sys','修改role_api',1,'2022-06-27 19:10:03','2022-06-27 19:10:03'),(183,'/file/path/add','1','sys','file添加页面',1,'2022-06-27 19:54:22','2022-06-27 19:54:22'),(184,'/file/path/edit/:id','1','sys','file修改页面',1,'2022-06-27 19:54:22','2022-06-27 19:54:22'),(185,'/file/path/del/:id','4','sys','file删除操作',1,'2022-06-27 19:54:22','2022-06-27 19:54:22'),(186,'/file','2','sys','添加file',1,'2022-06-27 19:54:22','2022-06-27 19:54:22'),(187,'/thingRecord/path','1','f','thing_record页面',1,'2022-06-29 16:58:46','2022-06-29 16:58:46'),(188,'/thingRecord/path/add','1','f','thing_record添加页面',1,'2022-06-29 16:58:46','2022-06-29 16:58:46'),(189,'/thingRecord/path/edit/:id','1','f','thing_record修改页面',1,'2022-06-29 16:58:46','2022-06-29 16:58:46'),(190,'/thingRecord/path/del/:id','4','f','thing_record删除操作',1,'2022-06-29 16:58:46','2022-06-29 16:58:46'),(191,'/thingRecord','2','f','添加thing_record',1,'2022-06-29 16:58:46','2022-07-03 13:13:03'),(192,'/thingRecord','3','f','修改thing_record',1,'2022-06-29 16:58:46','2022-06-29 16:58:46'),(193,'/node/path','1','f','node页面',1,'2022-07-05 19:01:03','2022-07-05 19:01:03'),(194,'/node/path/add','1','f','node添加页面',1,'2022-07-05 19:01:03','2022-07-05 19:01:03'),(195,'/node/path/edit/:id','1','f','node修改页面',1,'2022-07-05 19:01:03','2022-07-05 19:01:03'),(196,'/node/path/del/:id','4','f','node删除操作',1,'2022-07-05 19:01:03','2022-07-05 19:01:03'),(197,'/node','2','f','添加node',1,'2022-07-05 19:01:03','2022-07-05 19:01:03'),(198,'/node','3','f','修改node',1,'2022-07-05 19:01:03','2022-07-05 19:01:03');
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
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (11,'api_group','系统,日志,菜单,API,字典,管理员,文件','API分组选项','1',1,1,'2022-02-27 20:40:57','2022-06-28 13:38:34'),(22,'music-url','https://www.youtube.com/embed/videoseries?list=PLrzviuM_VBi2P4RQyQWGC5zZPvnEz4R62','登陆音乐列表','1',1,1,'2022-03-08 16:36:11','2022-06-29 13:55:36'),(33,'node-category','1.记事\r\n3.mysql\r\n5.english\r\n6.freekey\r\n8.go\r\n9.idea\r\n10.js\r\n12.linux\r\n15.nginx\r\n16.error\r\n17.quotations','备忘录分类','1',1,1,'2022-07-07 20:18:58','2022-07-09 01:23:38');
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
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (26,'1/2022/03/BYFY4d.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:50:13'),(27,'1/2022/03/FdI4Yw.gif',3,1,'2022-03-27 20:19:07','2022-07-09 16:51:51'),(29,'1/2022/03/mAMoWX.png',2,1,NULL,'2022-07-09 16:49:57'),(30,'1/2022/03/2S41in.png',2,1,'2022-03-28 15:32:00','2022-07-09 16:49:47'),(31,'1/2022/03/IdGUqj.png',2,1,'2022-03-28 15:36:45','2022-07-09 16:49:40'),(32,'1/2022/03/5Eoxb1.png',2,1,'2022-03-28 15:40:17','2022-07-09 16:49:33'),(77,'2/2022/07/CQVqgn.webp',2,1,'2022-07-03 12:44:29','2022-07-03 12:44:29'),(78,'2/2022/07/qMBDps.png',2,1,'2022-07-03 12:49:10','2022-07-03 12:49:10'),(79,'2/2022/07/lSCC0m.webp',2,1,'2022-07-03 13:00:15','2022-07-03 13:00:15'),(80,'2/2022/07/SHf1y4.webp',2,1,'2022-07-03 18:38:21','2022-07-03 18:38:21');
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
) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','',1.00,2,'','',1,'2022-06-24 06:18:55','2022-07-09 08:09:47'),(2,1,'1/2022/03/FdI4Yw.gif','','菜单','/menu/path',1.10,1,'这里是菜单页面','\"\"',1,'2022-02-16 11:14:13','2022-07-08 15:24:50'),(3,1,'1/2022/03/IdGUqj.png','','角色','/role/path',1.30,1,'','\"\"',1,'2022-03-04 08:57:14','2022-07-09 08:08:31'),(4,1,'1/2022/03/BYFY4d.gif','','API','/api/path',1.20,1,'','',1,NULL,'2022-07-09 10:41:50'),(5,1,'1/2022/03/5Eoxb1.png','','管理员','/admin/path',1.40,1,'','',1,'2022-03-08 07:45:04','2022-07-09 08:10:07'),(28,1,'1/2022/03/mAMoWX.png',NULL,'字典','/dict/path',1.50,1,'字典页面',NULL,1,'2022-03-08 07:45:04','2022-07-02 08:06:55'),(30,1,'1/2022/03/2S41in.png','','文件','/file/path',1.60,1,'',NULL,1,'2022-03-08 08:05:30','2022-07-03 05:12:16'),(59,1,'1/2022/03/5Eoxb1.png','','代码生成','/gen/path',1.70,1,'',NULL,1,'2022-04-01 14:41:45','2022-07-03 05:12:25'),(78,1,'2/2022/07/lSCC0m.webp','','操作日志','/operationLog/path',1.80,1,'','',1,'2022-06-13 11:59:57','2022-07-09 08:19:58'),(132,-1,'','','工具','',2.00,2,'','',1,'2022-07-09 08:45:07','2022-07-09 08:45:40'),(136,132,'2/2022/07/SHf1y4.webp','','站点导航','/to/urls',2.20,1,'','/sys/tool/urls.html',1,'2022-07-03 06:25:52','2022-07-08 17:24:33'),(137,132,'2/2022/07/SHf1y4.webp','','语录','/to/quotations',2.30,1,'','/sys/tool/quotations.html',1,'2022-07-03 06:29:23','2022-07-07 06:16:27'),(138,132,'2/2022/07/CQVqgn.webp','','备忘录','/node/path',2.40,1,'我的备忘录','',1,'2022-07-05 11:01:03','2022-07-09 08:21:04');
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
) ENGINE=InnoDB AUTO_INCREMENT=449 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (448,1,'http://localhost:1211/operationLog/clear','','GET','/operationLog/clear','::1',9,'2022-07-09 20:19:39');
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
) ENGINE=InnoDB AUTO_INCREMENT=181 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` (`id`, `rid`, `aid`) VALUES (130,2,56),(131,2,57),(132,2,58),(133,2,59),(134,2,61),(135,2,62),(136,2,63),(137,2,65),(138,2,66),(139,2,68),(140,2,69),(141,2,71),(142,2,72),(143,2,144),(144,2,145),(145,2,146),(146,2,147),(147,2,148),(148,2,149),(149,2,151),(150,2,152),(151,2,153),(152,2,154),(153,2,155),(154,2,156),(155,2,157),(156,2,158),(157,2,159),(158,2,160),(159,2,161),(160,2,162),(161,2,163),(162,2,164),(163,2,165),(164,2,166),(165,2,167),(166,2,168),(167,2,169),(168,2,170),(169,2,171),(170,2,172),(171,2,173),(172,2,174),(173,2,175),(174,2,176),(175,2,177),(176,2,178),(177,2,179),(178,2,180),(179,2,181),(180,2,182);
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
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(67,1,28),(68,1,30),(84,1,59),(100,1,78),(136,1,132),(142,2,28),(145,1,136),(146,1,137),(147,1,138);
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

-- Dump completed on 2022-07-09 20:21:02
