-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: 127.0.0.1    Database: freekey_admin
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
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `link` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` int DEFAULT '0',
  `status` tinyint unsigned DEFAULT '1' COMMENT '1open 2close',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `c_banner`
--

LOCK TABLES `c_banner` WRITE;
/*!40000 ALTER TABLE `c_banner` DISABLE KEYS */;
INSERT INTO `c_banner` (`id`, `title`, `image`, `link`, `desc`, `sort`, `status`, `created_at`) VALUES (3,'banner1','/image/p003.png','','',0,1,'2023-04-02 15:24:32');
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
  `uname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uname` (`uname`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `pwd`, `nickname`, `email`, `phone`, `status`, `created_at`, `updated_at`) VALUES (42,1,'admin','$2a$10$d2PbGXoRkbOZ.VVTHRgc6umhihYTZIqCRCoP1/v.vf7f9tIhbzW8q','admin','','',1,'2022-07-02 11:28:52','2023-04-02 08:11:12'),(52,22,'guest','$2a$10$bbqk5HOJvtTdtFFaiCfWV.OpKAp.pVho1izAoL.JrCHwi3P0Hx.n2','guest','','',1,'2023-04-04 05:49:59','2023-04-04 11:37:13');
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
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT ' {"label":"登录IP","notShow":0,"fieldType":"text","editHide":0,"editDisabled":0,"required":1}',
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  CONSTRAINT `s_admin_login_log_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `s_admin` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=364 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin_login_log`
--

LOCK TABLES `s_admin_login_log` WRITE;
/*!40000 ALTER TABLE `s_admin_login_log` DISABLE KEYS */;
INSERT INTO `s_admin_login_log` (`id`, `uid`, `ip`, `status`, `created_at`, `updated_at`) VALUES (363,42,'::1',1,'2023-04-07 19:28:03','2023-04-07 19:28:03');
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
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '1 get 2 post 3 put 4 delete',
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `group`, `url`, `method`, `desc`, `created_at`, `updated_at`) VALUES (1,'管理员','/backend/getCaptcha','GET','获取验证码','2023-01-16 17:04:25','2023-01-19 15:35:00'),(2,'管理员','/backend/login','POST','管理员登录','2023-01-16 17:07:50','2023-01-16 17:07:50'),(3,'管理员','/backend/admin/getInfo','GET','获取登录信息','2023-01-16 17:08:19','2023-01-16 17:08:19'),(4,'管理员','/backend/admin/getMenu','GET','获取菜单信息','2023-01-16 09:08:55','2023-01-17 15:23:33'),(5,'菜单','/backend/menu','GET',NULL,'2023-01-16 17:08:55','2023-01-16 17:08:55'),(6,'菜单','/backend/menu/list','GET',NULL,'2023-01-16 17:08:55','2023-01-16 17:08:55'),(7,'菜单','/backend/menu','POST',NULL,'2023-01-16 17:08:55','2023-01-16 17:08:55'),(8,'菜单','/backend/menu','PUT',NULL,'2023-01-16 17:08:55','2023-01-16 17:08:55'),(9,'菜单','/backend/menu','DELETE',NULL,'2023-01-16 17:08:55','2023-01-16 17:08:55'),(10,'菜单','/backend/menu/sort','GET','菜单排序','2023-01-16 17:08:55','2023-01-16 17:08:55'),(11,'API','/backend/api','GET',NULL,'2023-01-16 09:12:32','2023-01-18 08:18:09'),(12,'API','/backend/api/list','GET',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(13,'API','/backend/api','POST',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(14,'API','/backend/api','PUT',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(15,'API','/backend/api','DELETE',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(16,'角色','/backend/role','GET',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(17,'角色','/backend/role/list','GET',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(18,'角色','/backend/role','POST',NULL,'2023-01-16 17:12:32','2023-01-16 17:12:32'),(19,'角色','/backend/role','PUT',NULL,'2023-01-16 17:13:02','2023-01-16 17:13:02'),(20,'角色','/backend/role','DELETE',NULL,'2023-01-16 17:13:02','2023-01-16 17:13:02'),(21,'角色','/backend/roleApi','GET',NULL,'2023-01-16 17:13:02','2023-01-16 17:13:02'),(22,'角色','/backend/roleApi/list','GET',NULL,'2023-01-16 17:13:02','2023-01-16 17:13:02'),(23,'角色','/backend/roleApi','POST',NULL,'2023-01-16 17:13:02','2023-01-16 17:13:02'),(24,'角色','/backend/roleApi','PUT',NULL,'2023-01-16 17:14:31','2023-01-16 17:14:31'),(25,'角色','/backend/roleApi','DELETE',NULL,'2023-01-16 17:14:31','2023-01-16 17:14:31'),(26,'角色','/backend/roleMenu','GET',NULL,'2023-01-16 19:36:56','2023-01-16 19:36:56'),(27,'角色','/backend/roleMenu','GET',NULL,'2023-01-16 19:36:56','2023-01-16 19:36:56'),(28,'角色','/backend/roleMenu','POST',NULL,'2023-01-16 19:36:56','2023-01-16 19:36:56'),(29,'角色','/backend/roleMenu','PUT',NULL,'2023-01-16 19:36:56','2023-01-16 19:36:56'),(30,'角色','/backend/roleMenu','DELETE',NULL,'2023-01-16 19:36:56','2023-01-19 15:35:05'),(31,'管理员','/backend/admin/updateUname','PUT','修改管理员用户名','2023-01-16 23:10:52','2023-04-04 12:35:22'),(32,'管理员','/backend/adminn/updatePass','PUT','修改管理员密码','2023-01-17 15:20:31','2023-01-17 15:20:31'),(43,'角色','/backend/role','GET','','2023-01-19 16:48:39','2023-01-19 16:48:39'),(44,'角色','/backend/role/list','GET','','2023-01-19 16:48:39','2023-01-19 16:48:39'),(45,'角色','/backend/role','POST','','2023-01-19 16:48:39','2023-01-19 16:48:39'),(46,'角色','/backend/role','PUT','','2023-01-19 16:48:39','2023-01-19 16:48:39'),(47,'角色','/backend/role','DELETE','','2023-01-19 16:48:39','2023-01-19 16:48:39'),(48,'字典','/backend/dict','GET','','2023-01-19 16:59:11','2023-01-19 16:59:11'),(49,'字典','/backend/dict/list','GET','','2023-01-19 16:59:11','2023-01-19 16:59:11'),(50,'字典','/backend/dict','POST','','2023-01-19 16:59:11','2023-01-19 16:59:11'),(51,'字典','/backend/dict','PUT','','2023-01-19 16:59:11','2023-01-19 16:59:11'),(52,'字典','/backend/dict','DELETE','','2023-01-19 16:59:11','2023-01-19 16:59:11'),(53,'文件','/backend/file','GET','','2023-01-19 16:59:57','2023-01-19 16:59:57'),(54,'文件','/backend/file/list','GET','','2023-01-19 16:59:57','2023-01-19 16:59:57'),(55,'文件','/backend/file','POST','','2023-01-19 16:59:57','2023-01-19 16:59:57'),(56,'文件','/backend/file','PUT','','2023-01-19 16:59:57','2023-01-19 16:59:57'),(57,'文件','/backend/file','DELETE','','2023-01-19 16:59:57','2023-01-19 16:59:57'),(58,'文件','/backend/file/upload','POST','上传文件','2023-01-19 17:01:43','2023-01-19 17:01:43'),(59,'操作日志','/backend/operationLog','GET','','2023-01-19 17:02:02','2023-01-19 17:02:02'),(60,'操作日志','/backend/operationLog/list','GET','','2023-01-19 17:02:02','2023-01-19 17:02:02'),(61,'操作日志','/backend/operationLog','POST','','2023-01-19 17:02:02','2023-01-19 17:02:02'),(62,'操作日志','/backend/operationLog','PUT','','2023-01-19 17:02:02','2023-01-19 17:02:02'),(63,'操作日志','/backend/operationLog','DELETE','','2023-01-19 17:02:02','2023-01-19 17:02:02'),(64,'操作日志','/backend/operationLog/delClear','DELETE','清空操作日志','2023-01-19 17:02:23','2023-04-04 12:35:59'),(65,'登陆日志','/backend/adminLoginLog','GET','','2023-01-19 17:02:52','2023-01-19 17:02:52'),(66,'登陆日志','/backend/adminLoginLog/list','GET','','2023-01-19 17:02:52','2023-01-19 17:02:52'),(67,'登陆日志','/backend/adminLoginLog','POST','','2023-01-19 17:02:52','2023-01-19 17:02:52'),(68,'登陆日志','/backend/adminLoginLog','PUT','','2023-01-19 17:02:52','2023-01-19 17:02:52'),(69,'登陆日志','/backend/adminLoginLog','DELETE','','2023-01-19 17:02:52','2023-01-19 17:02:52'),(70,'登陆日志','/backend/adminLoginLog/delClear','DELETE','清空管理员登录日志','2023-01-19 17:03:51','2023-04-04 12:36:13'),(71,'banner','/backend/banner','GET','','2023-01-19 17:04:02','2023-01-19 17:04:02'),(72,'banner','/backend/banner/list','GET','','2023-01-19 17:04:02','2023-01-19 17:04:02'),(73,'banner','/backend/banner','POST','','2023-01-19 17:04:02','2023-01-19 17:04:02'),(74,'banner','/backend/banner','PUT','','2023-01-19 17:04:02','2023-01-19 17:04:02'),(75,'banner','/backend/banner','DELETE','','2023-01-19 17:04:02','2023-01-19 17:04:02'),(76,'用户','/backend/user','GET','','2023-01-19 17:04:39','2023-01-19 17:04:39'),(77,'用户','/backend/user/list','GET','','2023-01-19 17:04:39','2023-01-19 17:04:39'),(78,'用户','/backend/user','POST','','2023-01-19 17:04:39','2023-01-19 17:04:39'),(79,'用户','/backend/user','PUT','','2023-01-19 17:04:39','2023-01-19 17:04:39'),(80,'用户','/backend/user','DELETE','','2023-01-19 17:04:39','2023-01-19 17:04:39'),(81,'用户','/backend/user/updateUname','PUT','修改用户名','2023-01-19 17:05:03','2023-01-19 17:05:03'),(82,'用户','/backend/user/updatePass','PUT','修改用户密码','2023-01-19 17:05:34','2023-04-04 12:38:40'),(83,'用户登录日志','/backend/userLoginLog','GET','','2023-01-19 17:06:30','2023-01-19 17:06:30'),(84,'用户登录日志','/backend/userLoginLog/list','GET','','2023-01-19 17:06:30','2023-01-19 17:06:30'),(85,'用户登录日志','/backend/userLoginLog','POST','','2023-01-19 17:06:30','2023-01-19 17:06:30'),(86,'用户登录日志','/backend/userLoginLog','PUT','','2023-01-19 17:06:30','2023-01-19 17:06:30'),(87,'用户登录日志','/backend/userLoginLog','DELETE','','2023-01-19 17:06:30','2023-01-19 17:06:30'),(88,'钱包','/backend/wallet','GET','','2023-01-19 17:07:15','2023-01-19 17:07:15'),(89,'钱包','/backend/wallet/list','GET','','2023-01-19 17:07:15','2023-01-19 17:07:15'),(90,'钱包','/backend/wallet','POST','','2023-01-19 17:07:15','2023-01-19 17:07:15'),(91,'钱包','/backend/wallet','PUT','','2023-01-19 17:07:15','2023-01-19 17:07:15'),(92,'钱包','/backend/wallet','DELETE','','2023-01-19 17:07:15','2023-01-19 17:07:15'),(93,'钱包','/backend/wallet/updatePass','PUT','修改钱包密码','2023-01-19 17:07:57','2023-04-04 12:39:32'),(94,'钱包','/backend/wallet/updateByAdmin','PUT','管理员充值提现','2023-01-19 17:08:21','2023-04-04 12:39:21'),(95,'账变类型','/backend/walletChangeType','GET','','2023-01-19 17:11:30','2023-01-19 17:11:30'),(96,'账变类型','/backend/walletChangeType/list','GET','','2023-01-19 17:11:30','2023-01-19 17:11:30'),(97,'账变类型','/backend/walletChangeType','POST','','2023-01-19 17:11:30','2023-01-19 17:11:30'),(98,'账变类型','/backend/walletChangeType','PUT','','2023-01-19 17:11:30','2023-01-19 17:11:30'),(99,'账变类型','/backend/walletChangeType','DELETE','','2023-01-19 17:11:30','2023-01-19 17:11:30'),(100,'账变记录','/backend/walletChangeLog','GET','','2023-01-19 17:11:55','2023-01-19 17:11:55'),(101,'账变记录','/backend/walletChangeLog/list','GET','','2023-01-19 17:11:55','2023-01-19 17:11:55'),(102,'账变记录','/backend/walletChangeLog','POST','','2023-01-19 17:11:55','2023-01-19 17:11:55'),(103,'账变记录','/backend/walletChangeLog','PUT','','2023-01-19 17:11:55','2023-01-19 17:11:55'),(104,'账变记录','/backend/walletChangeLog','DELETE','','2023-01-19 17:11:55','2023-01-19 17:11:55'),(105,'充值订单','/backend/walletTopUpApplication','GET','','2023-01-19 17:13:59','2023-01-19 17:13:59'),(106,'充值订单','/backend/walletTopUpApplication/list','GET','','2023-01-19 17:13:59','2023-01-19 17:13:59'),(107,'充值订单','/backend/walletTopUpApplication','POST','','2023-01-19 17:13:59','2023-01-19 17:13:59'),(108,'充值订单','/backend/walletTopUpApplication','PUT','','2023-01-19 17:13:59','2023-01-19 17:13:59'),(109,'充值订单','/backend/walletTopUpApplication','DELETE','','2023-01-19 17:13:59','2023-01-19 17:13:59'),(110,'菜单','/backend/api/addGroup','POST','','2023-01-19 18:50:11','2023-01-19 18:50:11'),(137,'用户登录日志','/userLoginLog/delClear','DELETE','清空用户登录日志','2023-04-04 12:26:22','2023-04-04 12:26:42'),(138,'API','/api/addGroup','POST','添加一组API\n','2023-04-04 12:31:45','2023-04-04 12:31:45'),(139,'角色','/roleApi/addRoleApis','POST','添加角色api权限','2023-04-04 12:33:36','2023-04-04 12:33:36'),(140,'角色','/roleApi/clear','DELETE','清除角色api权限','2023-04-04 12:34:05','2023-04-04 12:34:05'),(141,'角色','/roleMenu/addRoleMenus','POST','添加角色菜单','2023-04-04 12:34:36','2023-04-04 12:34:36'),(142,'角色','/roleMenu/clear','DELETE','清除角色菜单','2023-04-04 12:34:57','2023-04-04 12:34:57'),(148,'充值订单','/walletTopUpApplication/updateByAdmin','PUT','审核充值订单','2023-04-04 12:41:45','2023-04-04 12:41:45'),(149,'提现订单','/backend/walletWithdrawApplication','GET','','2023-04-04 12:42:01','2023-04-04 12:42:01'),(150,'提现订单','/backend/walletWithdrawApplication/list','GET','','2023-04-04 12:42:01','2023-04-04 12:42:01'),(151,'提现订单','/backend/walletWithdrawApplication','POST','','2023-04-04 12:42:01','2023-04-04 12:42:01'),(152,'提现订单','/backend/walletWithdrawApplication','PUT','','2023-04-04 12:42:01','2023-04-04 12:42:01'),(153,'提现订单','/backend/walletWithdrawApplication','DELETE','','2023-04-04 12:42:01','2023-04-04 12:42:01'),(154,'提现订单','/walletWithdrawApplication/updateByAdmin','PUT','审核提现订单','2023-04-04 12:43:04','2023-04-04 12:43:04');
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
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `k` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `v` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `desc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'sys',
  `status` int DEFAULT NULL,
  `type` int NOT NULL DEFAULT '1' COMMENT '1 文本，2 img',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `title`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (42,'系统白名单','white_ips','::2','多个ip用小写逗号隔开','1',1,1,'2022-07-23 19:04:44','2023-04-07 15:32:45'),(44,'欢迎语','great','hello2','','2',1,1,'2022-12-28 12:25:41','2023-01-17 15:46:08');
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
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `url`, `group`, `status`, `created_at`, `updated_at`) VALUES (87,'/icon/icon01.png',1,1,'2022-12-25 23:13:53','2022-12-25 23:13:53'),(88,'/icon/icon02.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(103,'/icon/icon03.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(104,'/icon/icon04.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(105,'/icon/icon05.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(106,'/icon/icon06.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(107,'/icon/icon07.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(108,'/icon/icon08.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(109,'/icon/icon09.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(110,'/icon/icon10.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(111,'/icon/icon11.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(112,'/icon/icon12.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(113,'/icon/icon13.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(114,'/icon/icon14.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(115,'/icon/icon15.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(116,'/icon/icon16.png',1,1,'2022-12-26 13:17:06','2022-12-26 13:17:06'),(117,'/icon/icon17.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(118,'/icon/icon18.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(119,'/icon/icon19.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(120,'/icon/icon20.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(121,'/icon/icon21.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(122,'/icon/icon22.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(123,'/icon/icon23.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(124,'/icon/icon24.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(125,'/icon/icon25.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(126,'/icon/icon26.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(127,'/icon/icon27.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(128,'/icon/icon28.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(129,'/icon/icon29.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(130,'/icon/icon30.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(131,'/icon/icon31.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(133,'/icon/icon33.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(134,'/icon/icon34.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(135,'/icon/icon35.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(136,'/icon/icon36.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(137,'/icon/icon37.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(138,'/icon/icon38.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(139,'/icon/icon39.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(140,'/icon/icon40.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(141,'/icon/icon41.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(142,'/icon/icon42.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(143,'/icon/icon43.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(144,'/icon/icon44.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(145,'/icon/icon45.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(146,'/icon/icon46.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(147,'/icon/icon47.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(148,'/icon/icon48.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(149,'/icon/icon49.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50'),(150,'/icon/icon50.png',1,1,'2022-12-26 13:35:50','2022-12-26 13:35:50');
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
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `bg_img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` decimal(7,2) NOT NULL DEFAULT '0.00',
  `type` int NOT NULL DEFAULT '1' COMMENT '1normal 2group',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `file_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=261 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `icon`, `bg_img`, `name`, `path`, `sort`, `type`, `desc`, `file_path`, `status`, `created_at`, `updated_at`) VALUES (1,-1,'','','系统','/',1.00,2,'','',1,'2022-06-18 06:18:55','2023-04-04 12:40:08'),(2,1,'','','菜单','/backend/sys/menu',1.10,1,'这里是菜单页面','',2,'2022-02-15 03:14:13','2023-04-07 11:21:58'),(3,1,'','','角色','/backend/sys/role',1.30,1,'角色权限管理，在这里可以创建新的角色','',1,'2022-03-04 08:57:14','2023-04-07 11:21:45'),(4,1,'','','API','/backend/sys/api',1.20,1,'','',1,'2022-07-03 06:25:52','2023-04-07 11:21:55'),(5,1,'','','管理员','/backend/sys/admin',1.40,1,'','',1,'2022-03-07 23:45:04','2023-04-07 11:21:42'),(28,1,'','','字典','/backend/sys/dict',1.50,1,'字典页面','',1,'2022-03-06 23:45:04','2023-04-07 11:21:37'),(30,1,'','','文件','/backend/sys/file',1.60,1,'','',1,'2022-03-08 08:05:30','2023-04-07 11:21:40'),(78,1,'','','操作日志','/backend/sys/operationLog',1.80,1,'','',1,'2022-06-13 11:59:57','2023-04-07 11:21:34'),(139,1,'','','登录日志','/backend/sys/adminLoginLog',1.90,1,'这里是登陆日志页面可以对数据进行相应的操作。','',1,'2022-07-11 11:06:26','2023-04-07 11:21:30'),(171,-1,'','','用户','',5.00,2,'','',1,'2022-09-01 14:24:18','2023-01-14 09:10:02'),(172,171,'','','用户列表','/backend/user',5.10,1,'这里是用户管理页面','',1,'2022-08-31 22:24:18','2023-04-07 11:21:05'),(173,171,'','','登录日志','/backend/user/userLoginLog',5.20,1,'','',1,'2022-09-02 14:47:17','2023-04-07 11:21:02'),(174,171,'','','钱包','/backend/user/wallet',5.30,1,'','',1,'2022-09-03 20:32:44','2023-04-07 11:20:59'),(175,-1,'','','配置','/',2.00,2,'','',1,'2022-09-04 23:02:32','2023-01-19 11:24:33'),(176,175,'','','账变类型','/backend/setting/walletChangeType',2.10,1,'','',1,'2022-09-04 07:02:32','2023-04-07 11:21:24'),(177,171,'','','账变记录','/backend/user/walletChangeLog',5.40,1,'','',1,'2022-09-04 19:09:03','2023-04-07 11:38:00'),(178,-1,'','','统计','',4.00,2,'','',1,'2022-09-05 03:15:18','2022-12-28 16:49:59'),(179,178,'','','账变统计','/backend/user/walletStatisticsLog',4.10,1,'','',1,'2022-09-04 19:15:18','2023-04-07 11:21:11'),(180,178,'','','账变报表','/backend/user/walletReport',4.20,1,'','',1,'2022-09-05 22:14:24','2023-04-07 11:21:08'),(204,-1,'','','通用','',3.00,2,'','',1,'2022-12-28 16:46:37','2022-12-28 16:48:10'),(205,204,'','','Banner','/backend/common/banner',3.10,1,'','',1,'2022-12-28 00:46:37','2023-04-07 11:21:19'),(230,1,'','','文档与测试','/backend/sys/test?name=CSS',1.91,1,'','',0,'2023-01-19 11:42:54','2023-01-20 12:08:58');
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
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `use_time` int NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4835 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_operation_log`
--

LOCK TABLES `s_operation_log` WRITE;
/*!40000 ALTER TABLE `s_operation_log` DISABLE KEYS */;
INSERT INTO `s_operation_log` (`id`, `uid`, `content`, `response`, `method`, `uri`, `ip`, `use_time`, `created_at`) VALUES (4828,42,'/backend/admin?id=51','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/admin','::1',13,'2023-04-07 19:31:33'),(4829,42,'/backend/dict?id=53','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/dict','::1',8,'2023-04-07 19:31:40'),(4830,42,'/backend/dict?id=52','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/dict','::1',12,'2023-04-07 19:31:41'),(4831,42,'/backend/dict?id=51','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/dict','::1',8,'2023-04-07 19:31:42'),(4832,42,'/backend/dict?id=50','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/dict','::1',9,'2023-04-07 19:31:45'),(4833,42,'/backend/dict?id=49','{\"code\":0,\"message\":\"\",\"data\":null}','DELETE','/backend/dict','::1',6,'2023-04-07 19:31:46'),(4834,42,'{\"id\":177,\"pid\":171,\"icon\":\"\",\"bgImg\":\"\",\"name\":\"账变记录\",\"path\":\"/backend/user/walletChangeLog\",\"sort\":5.4,\"type\":1,\"desc\":\"\",\"filePath\":\"\",\"status\":1,\"createdAt\":\"2022-09-05 03:09:03\",\"updatedAt\":\"2023-04-07 19:20:55\"}','{\"code\":0,\"message\":\"\",\"data\":null}','PUT','/backend/menu','::1',11,'2023-04-07 19:38:00');
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
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` (`id`, `name`, `created_at`, `updated_at`) VALUES (1,'超级管理员','2022-02-16 11:12:41','2022-09-02 12:22:24'),(22,'临时访客','2022-07-22 16:45:05','2023-04-04 05:42:41'),(28,'系统管理员','2023-01-15 22:10:16','2023-04-04 05:42:57');
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
) ENGINE=InnoDB AUTO_INCREMENT=1331 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
INSERT INTO `s_role_api` (`id`, `rid`, `aid`) VALUES (1329,1,94);
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
) ENGINE=InnoDB AUTO_INCREMENT=409 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (372,22,204),(374,22,1),(375,22,230),(376,1,1),(377,1,2),(378,1,4),(379,1,3),(380,1,5),(381,1,28),(382,1,30),(383,1,78),(384,1,139),(385,1,230),(386,1,175),(387,1,176),(389,1,204),(390,1,205),(392,1,178),(393,1,179),(394,1,180),(395,1,171),(396,1,172),(397,1,173),(398,1,174),(399,1,177);
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
  `uname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `join_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '注册IP',
  `device` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备名称',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `pass_error_count` tinyint unsigned DEFAULT '0' COMMENT '密码错误次数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone_2` (`phone`),
  KEY `uname_2` (`uname`),
  KEY `join_ip` (`join_ip`),
  KEY `status` (`status`),
  KEY `phone` (`phone`),
  KEY `email` (`email`),
  KEY `join_ip_2` (`join_ip`),
  KEY `status_2` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user`
--

LOCK TABLES `u_user` WRITE;
/*!40000 ALTER TABLE `u_user` DISABLE KEYS */;
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
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_user_login_log`
--

LOCK TABLES `u_user_login_log` WRITE;
/*!40000 ALTER TABLE `u_user_login_log` DISABLE KEYS */;
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
  `pass` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pass_err_count` tinyint unsigned DEFAULT '0' COMMENT '密码输错次数',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `status` tinyint unsigned DEFAULT '1' COMMENT '金库状态 0 设置密码 1正常,2 锁定',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`),
  KEY `uid_2` (`uid`),
  KEY `balance` (`balance`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户金库';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet`
--

LOCK TABLES `u_wallet` WRITE;
/*!40000 ALTER TABLE `u_wallet` DISABLE KEYS */;
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
  `trans_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uid` bigint unsigned NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `balance` decimal(12,2) NOT NULL,
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1人工充值,2支付宝充值,3微信充值,4paypal充值,5人工扣除',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `amount` (`amount`),
  KEY `balance` (`balance`),
  KEY `type` (`type`),
  KEY `trans_id` (`trans_id`)
) ENGINE=InnoDB AUTO_INCREMENT=607 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账变记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_change_log`
--

LOCK TABLES `u_wallet_change_log` WRITE;
/*!40000 ALTER TABLE `u_wallet_change_log` DISABLE KEYS */;
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
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `sub_title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` tinyint unsigned DEFAULT '1' COMMENT '1 add; 2 reduce',
  `class` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `status` tinyint unsigned DEFAULT '1',
  `count_status` tinyint DEFAULT '1' COMMENT 'Whether this field needs statistics, 1 true 2 false',
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='充值类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_change_type`
--

LOCK TABLES `u_wallet_change_type` WRITE;
/*!40000 ALTER TABLE `u_wallet_change_type` DISABLE KEYS */;
INSERT INTO `u_wallet_change_type` (`id`, `title`, `sub_title`, `type`, `class`, `desc`, `status`, `count_status`) VALUES (1,'支付宝充值',NULL,1,'tag-primary','',1,1),(2,'微信充值','',1,'tag-success','',1,1),(3,'银行卡充值',NULL,1,'tag-warning',NULL,1,1),(4,'PayPal充值','',1,'tag-info','',1,1);
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
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `u_wallet_statistics_log`
--

LOCK TABLES `u_wallet_statistics_log` WRITE;
/*!40000 ALTER TABLE `u_wallet_statistics_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `u_wallet_statistics_log` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-07 20:22:28
