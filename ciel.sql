-- MySQL dump 10.13  Distrib 8.0.24, for Linux (x86_64)
--
-- Host: localhost    Database: ciel
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` VALUES (1,1,'admin','$2a$10$2wtzucESYl2r5eARjdcefeg2/4bbaWxFqyeG1C5clSPOYjrowXG.G',1,'2022-03-08 08:59:33','2022-06-12 05:54:00'),(18,1,'2222','$2a$10$pJ5G2m8CLEHEo0etGszirOWtaMapwLCd1bHtLA/WLgHM8e7R4fdDO',1,'2022-03-08 08:57:23','2022-06-12 06:01:38');
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
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` VALUES (32,'/user/path','GET','u','用户列表页面',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(33,'/user','GET','u','查询用户列表集合',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(34,'/user/:id','GET','u','查询用户列表详情',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(35,'/user/:id','DELETE','u','删除用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(36,'/user','POST','u','添加用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(37,'/user','PUT','u','修改用户列表',1,'2022-06-10 20:23:55','2022-06-10 20:23:55'),(38,'/loginLog/path','GET','u','login_log页面',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(39,'/loginLog','GET','u','查询login_log集合',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(40,'/loginLog/:id','GET','u','查询login_log详情',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(41,'/loginLog/:id','DELETE','u','删除login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(42,'/loginLog','POST','u','添加login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(43,'/loginLog','PUT','u','修改login_log',1,'2022-06-11 18:48:30','2022-06-11 18:48:30'),(44,'/dict/path','GET','s','dict页面',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(45,'/dict','GET','s','查询dict集合',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(46,'/dict/:id','GET','s','查询dict详情',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(47,'/dict/:id','DELETE','s','删除dict',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(48,'/dict','POST','s','添加dict',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(49,'/dict','PUT','s','修改dict',1,'2022-06-11 20:13:38','2022-06-11 20:13:38'),(50,'/admin/path','GET','sys','admin页面',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(51,'/admin','GET','sys','查询admin集合',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(52,'/admin/:id','GET','sys','查询admin详情',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(53,'/admin/:id','DELETE','sys','删除admin',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(54,'/admin','POST','sys','添加admin',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(55,'/admin','PUT','sys','修改admin',1,'2022-06-11 20:26:21','2022-06-11 20:26:21'),(56,'/file/path','GET','s','file页面',1,'2022-06-12 14:10:23','2022-06-12 14:10:23'),(57,'/file','GET','s','查询file集合',1,'2022-06-12 14:10:23','2022-06-12 14:10:23'),(58,'/file/:id','GET','s','查询file详情',1,'2022-06-12 14:10:23','2022-06-12 14:10:23'),(59,'/file/:id','DELETE','s','删除file',1,'2022-06-12 14:10:23','2022-06-12 14:10:23'),(60,'/file','POST','s','添加file',1,'2022-06-12 14:10:23','2022-06-12 15:59:29'),(61,'/file','PUT','s','修改file',1,'2022-06-12 14:10:23','2022-06-12 16:01:34'),(62,'/roleApi/path','GET','sys','role_api页面',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(63,'/roleApi','GET','sys','查询role_api集合',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(64,'/roleApi/:id','GET','sys','查询role_api详情',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(65,'/roleApi/:id','DELETE','sys','删除role_api',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(66,'/roleApi','POST','sys','添加role_api',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(67,'/roleApi','PUT','sys','修改role_api',1,'2022-06-12 17:02:13','2022-06-12 17:02:13'),(68,'/roleMenu/path','GET','s','role_menu页面',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(69,'/roleMenu','GET','s','查询role_menu集合',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(70,'/roleMenu/:id','GET','s','查询role_menu详情',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(71,'/roleMenu/:id','DELETE','s','删除role_menu',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(72,'/roleMenu','POST','s','添加role_menu',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(73,'/roleMenu','PUT','s','修改role_menu',1,'2022-06-13 15:19:45','2022-06-13 15:19:45'),(74,'/operationLog/path','GET','sys','操作日志页面',1,'2022-06-13 19:59:57','2022-06-13 19:59:57'),(75,'/operationLog','GET','sys','查询操作日志集合',1,'2022-06-13 19:59:57','2022-06-13 19:59:57'),(76,'/operationLog/:id','GET','sys','查询操作日志详情',1,'2022-06-13 19:59:57','2022-06-13 19:59:57'),(77,'/operationLog/:id','DELETE','sys','删除操作日志',1,'2022-06-13 19:59:57','2022-06-13 19:59:57'),(78,'/operationLog','POST','sys','添加操作日志',1,'2022-06-13 19:59:57','2022-06-13 19:59:57'),(79,'/operationLog','PUT','sys','修改操作日志',1,'2022-06-13 19:59:57','2022-06-13 19:59:57');
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` VALUES (11,'2','22','22','1',1,4,'2022-02-27 20:40:57','2022-06-11 20:17:47'),(22,'221','test','2','2',1,2,'2022-03-08 16:36:11','2022-04-01 21:00:36');
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
INSERT INTO `s_file` VALUES (26,'1/2022/03/BYFY4d.gif',1,1,'2022-03-27 20:19:07','2022-03-27 20:19:07'),(27,'1/2022/03/FdI4Yw.gif',1,1,'2022-03-27 20:19:07','2022-03-27 20:19:07'),(29,'1/2022/03/mAMoWX.png',1,1,'2022-03-28 15:28:47','2022-03-28 15:28:47'),(30,'1/2022/03/2S41in.png',1,1,'2022-03-28 15:32:00','2022-03-28 15:32:00'),(31,'1/2022/03/IdGUqj.png',1,1,'2022-03-28 15:36:45','2022-03-28 15:36:45'),(32,'1/2022/03/5Eoxb1.png',1,1,'2022-03-28 15:40:17','2022-06-12 15:08:09');
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
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` VALUES (1,-1,'','系统','',2,1.00,1,'2022-02-27 06:15:01','2022-04-01 07:51:06'),(2,1,'1/2022/03/FdI4Yw.gif','菜单','/menu/path',1,1.10,1,'2022-02-16 11:14:13','2022-03-27 12:19:17'),(3,1,'1/2022/03/IdGUqj.png','角色','/role/path',1,1.30,1,'2022-03-04 08:57:14','2022-03-28 07:36:59'),(4,1,'1/2022/03/BYFY4d.gif','API','/api/path',1,1.20,1,'2022-02-27 05:04:56','2022-03-27 12:19:27'),(5,1,'1/2022/03/5Eoxb1.png','管理员','/admin/path',1,1.40,1,'2022-03-08 07:45:04','2022-06-03 12:06:38'),(28,1,'1/2022/03/mAMoWX.png','字典','/dict/path',1,1.50,1,'2022-02-27 12:51:48','2022-04-01 11:32:15'),(30,1,'1/2022/03/2S41in.png','文件','/file/path',1,1.60,1,'2022-03-08 08:05:30','2022-04-01 11:23:43'),(59,1,'1/2022/03/5Eoxb1.png','代码生成','/gen/path',1,1.70,1,'2022-04-01 14:41:45','2022-04-01 15:16:27'),(73,-1,'','用户','',2,2.00,1,'2022-06-11 12:04:09','2022-06-11 12:04:09'),(74,73,'https://i.scdn.co/image/ab67706f000000035b2f3fe92b9ba1f4d536e231','用户列表','/user/path',1,2.10,1,'2022-06-11 12:04:09','2022-06-11 12:04:09'),(75,73,'https://www.gravatar.com/avatar/2442588627?d=monsterid&f=y','登陆日志','/loginLog/path',1,2.20,1,'2022-06-11 12:05:08','2022-06-13 12:17:54'),(78,1,'https://www.gravatar.com/avatar/302255901?d=robohash&f=y','操作日志','/operationLog/path',1,1.92,1,'2022-06-13 11:59:57','2022-06-13 12:00:26');
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` VALUES (1,'Super Admin','2022-02-16 11:12:41','2022-02-21 04:46:24'),(2,'Admin','2022-02-16 11:13:11','2022-06-13 09:06:46');
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
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` VALUES (74,2,33);
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
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(62,2,1),(63,2,2),(64,2,3),(67,1,28),(68,1,30),(72,2,4),(74,2,5),(75,2,28),(76,2,30),(84,1,59),(91,1,73),(92,1,74),(93,1,75),(100,1,78);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_login_log`
--

LOCK TABLES `u_login_log` WRITE;
/*!40000 ALTER TABLE `u_login_log` DISABLE KEYS */;
INSERT INTO `u_login_log` VALUES (1,6,'33322','2022-04-03 14:26:21','2022-06-11 19:46:42');
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
INSERT INTO `u_user` VALUES (6,'23','22','124','123',1,'2022-06-10 12:13:06','2022-06-10 12:27:43');
/*!40000 ALTER TABLE `u_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'ciel'
--

--
-- Dumping routines for database 'ciel'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-13 20:39:58
