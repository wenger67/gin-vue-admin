-- MySQL dump 10.13  Distrib 5.7.31, for Linux (x86_64)
--
-- Host: localhost    Database: qmPlus
-- ------------------------------------------------------
-- Server version	5.7.31-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `address_tags`
--

DROP TABLE IF EXISTS `address_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `address_tags` (
  `address_id` int(10) unsigned NOT NULL,
  `category_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`address_id`,`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address_tags`
--

LOCK TABLES `address_tags` WRITE;
/*!40000 ALTER TABLE `address_tags` DISABLE KEYS */;
INSERT INTO `address_tags` VALUES (111,130),(112,124),(112,125),(112,126),(112,127),(112,128),(112,129),(112,134),(113,125),(113,126),(113,127),(113,128),(113,135);
/*!40000 ALTER TABLE `address_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `addresses`
--

DROP TABLE IF EXISTS `addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `region_id` int(10) NOT NULL COMMENT 'region id',
  `address_name` varchar(255) DEFAULT NULL COMMENT '地址，精确到类似小区级别',
  `user_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'user amount',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_addresses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
INSERT INTO `addresses` VALUES (111,'2020-08-04 17:16:12','2020-08-05 11:11:19',NULL,102,'阿大四大四大的',111111),(112,'2020-08-05 10:23:24','2020-08-05 10:23:24','2020-08-05 11:41:19',0,'阿大四大四大的',111111),(113,'2020-08-05 10:24:22','2020-08-05 10:24:22','2020-08-05 11:41:25',0,'阿大四大四大的',111111);
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `authority_menu`
--

DROP TABLE IF EXISTS `authority_menu`;
/*!50001 DROP VIEW IF EXISTS `authority_menu`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `authority_menu` AS SELECT
 1 AS `id`,
 1 AS `created_at`,
 1 AS `updated_at`,
 1 AS `deleted_at`,
 1 AS `menu_level`,
 1 AS `parent_id`,
 1 AS `path`,
 1 AS `name`,
 1 AS `hidden`,
 1 AS `component`,
 1 AS `title`,
 1 AS `icon`,
 1 AS `nick_name`,
 1 AS `sort`,
 1 AS `authority_id`,
 1 AS `menu_id`,
 1 AS `keep_alive`,
 1 AS `default_menu`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('p','8881','/base/login','POST','','',''),('p','8881','/base/register','POST','','',''),('p','8881','/api/createApi','POST','','',''),('p','8881','/api/getApiList','POST','','',''),('p','8881','/api/getApiById','POST','','',''),('p','8881','/api/deleteApi','POST','','',''),('p','8881','/api/updateApi','POST','','',''),('p','8881','/api/getAllApis','POST','','',''),('p','8881','/authority/createAuthority','POST','','',''),('p','8881','/authority/deleteAuthority','POST','','',''),('p','8881','/authority/getAuthorityList','POST','','',''),('p','8881','/authority/setDataAuthority','POST','','',''),('p','8881','/menu/getMenu','POST','','',''),('p','8881','/menu/getMenuList','POST','','',''),('p','8881','/menu/addBaseMenu','POST','','',''),('p','8881','/menu/getBaseMenuTree','POST','','',''),('p','8881','/menu/addMenuAuthority','POST','','',''),('p','8881','/menu/getMenuAuthority','POST','','',''),('p','8881','/menu/deleteBaseMenu','POST','','',''),('p','8881','/menu/updateBaseMenu','POST','','',''),('p','8881','/menu/getBaseMenuById','POST','','',''),('p','8881','/user/changePassword','POST','','',''),('p','8881','/user/uploadHeaderImg','POST','','',''),('p','8881','/user/getInfoList','POST','','',''),('p','8881','/user/getUserList','POST','','',''),('p','8881','/user/setUserAuthority','POST','','',''),('p','8881','/fileUploadAndDownload/upload','POST','','',''),('p','8881','/fileUploadAndDownload/getFileList','POST','','',''),('p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),('p','8881','/casbin/updateCasbin','POST','','',''),('p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','8881','/jwt/jsonInBlacklist','POST','','',''),('p','8881','/system/getSystemConfig','POST','','',''),('p','8881','/system/setSystemConfig','POST','','',''),('p','8881','/customer/customer','POST','','',''),('p','8881','/customer/customer','PUT','','',''),('p','8881','/customer/customer','DELETE','','',''),('p','8881','/customer/customer','GET','','',''),('p','8881','/customer/customerList','GET','','',''),('p','9528','/base/login','POST','','',''),('p','9528','/base/register','POST','','',''),('p','9528','/api/createApi','POST','','',''),('p','9528','/api/getApiList','POST','','',''),('p','9528','/api/getApiById','POST','','',''),('p','9528','/api/deleteApi','POST','','',''),('p','9528','/api/updateApi','POST','','',''),('p','9528','/api/getAllApis','POST','','',''),('p','9528','/authority/createAuthority','POST','','',''),('p','9528','/authority/deleteAuthority','POST','','',''),('p','9528','/authority/getAuthorityList','POST','','',''),('p','9528','/authority/setDataAuthority','POST','','',''),('p','9528','/menu/getMenu','POST','','',''),('p','9528','/menu/getMenuList','POST','','',''),('p','9528','/menu/addBaseMenu','POST','','',''),('p','9528','/menu/getBaseMenuTree','POST','','',''),('p','9528','/menu/addMenuAuthority','POST','','',''),('p','9528','/menu/getMenuAuthority','POST','','',''),('p','9528','/menu/deleteBaseMenu','POST','','',''),('p','9528','/menu/updateBaseMenu','POST','','',''),('p','9528','/menu/getBaseMenuById','POST','','',''),('p','9528','/user/changePassword','POST','','',''),('p','9528','/user/uploadHeaderImg','POST','','',''),('p','9528','/user/getInfoList','POST','','',''),('p','9528','/user/getUserList','POST','','',''),('p','9528','/user/setUserAuthority','POST','','',''),('p','9528','/fileUploadAndDownload/upload','POST','','',''),('p','9528','/fileUploadAndDownload/getFileList','POST','','',''),('p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),('p','9528','/casbin/updateCasbin','POST','','',''),('p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','9528','/jwt/jsonInBlacklist','POST','','',''),('p','9528','/system/getSystemConfig','POST','','',''),('p','9528','/system/setSystemConfig','POST','','',''),('p','9528','/customer/customer','POST','','',''),('p','9528','/customer/customer','PUT','','',''),('p','9528','/customer/customer','DELETE','','',''),('p','9528','/customer/customer','GET','','',''),('p','9528','/customer/customerList','GET','','',''),('p','9528','/autoCode/createTemp','POST','','',''),('p','888','/base/login','POST','','',''),('p','888','/base/register','POST','','',''),('p','888','/api/createApi','POST','','',''),('p','888','/api/getApiList','POST','','',''),('p','888','/api/getApiById','POST','','',''),('p','888','/api/deleteApi','POST','','',''),('p','888','/api/updateApi','POST','','',''),('p','888','/api/getAllApis','POST','','',''),('p','888','/authority/createAuthority','POST','','',''),('p','888','/authority/deleteAuthority','POST','','',''),('p','888','/authority/getAuthorityList','POST','','',''),('p','888','/authority/setDataAuthority','POST','','',''),('p','888','/authority/updateAuthority','PUT','','',''),('p','888','/authority/copyAuthority','POST','','',''),('p','888','/menu/getMenu','POST','','',''),('p','888','/menu/getMenuList','POST','','',''),('p','888','/menu/addBaseMenu','POST','','',''),('p','888','/menu/getBaseMenuTree','POST','','',''),('p','888','/menu/addMenuAuthority','POST','','',''),('p','888','/menu/getMenuAuthority','POST','','',''),('p','888','/menu/deleteBaseMenu','POST','','',''),('p','888','/menu/updateBaseMenu','POST','','',''),('p','888','/menu/getBaseMenuById','POST','','',''),('p','888','/user/changePassword','POST','','',''),('p','888','/user/uploadHeaderImg','POST','','',''),('p','888','/user/getInfoList','POST','','',''),('p','888','/user/getUserList','POST','','',''),('p','888','/user/setUserAuthority','POST','','',''),('p','888','/user/deleteUser','DELETE','','',''),('p','888','/fileUploadAndDownload/upload','POST','','',''),('p','888','/fileUploadAndDownload/getFileList','POST','','',''),('p','888','/fileUploadAndDownload/deleteFile','POST','','',''),('p','888','/casbin/updateCasbin','POST','','',''),('p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),('p','888','/casbin/casbinTest/:pathParam','GET','','',''),('p','888','/jwt/jsonInBlacklist','POST','','',''),('p','888','/system/getSystemConfig','POST','','',''),('p','888','/system/setSystemConfig','POST','','',''),('p','888','/customer/customer','POST','','',''),('p','888','/customer/customer','PUT','','',''),('p','888','/customer/customer','DELETE','','',''),('p','888','/customer/customer','GET','','',''),('p','888','/customer/customerList','GET','','',''),('p','888','/autoCode/createTemp','POST','','',''),('p','888','/autoCode/getTables','GET','','',''),('p','888','/autoCode/getDB','GET','','',''),('p','888','/autoCode/getColume','GET','','',''),('p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),('p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),('p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),('p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),('p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),('p','888','/sysDictionary/createSysDictionary','POST','','',''),('p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),('p','888','/sysDictionary/updateSysDictionary','PUT','','',''),('p','888','/sysDictionary/findSysDictionary','GET','','',''),('p','888','/sysDictionary/getSysDictionaryList','GET','','',''),('p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),('p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),('p','888','/sysOperationRecord/updateSysOperationRecord','PUT','','',''),('p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),('p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),('p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),('p','888','/subject/getSubjectById','POST','','',''),('p','888','/subject/getSubjectList','POST','','',''),('p','888','/subject/createSubject','POST','','',''),('p','888','/subject/updateSubject','POST','','',''),('p','888','/subject/deleteSubject','POST','','',''),('p','888','/subject/getAllSubjects','POST','','',''),('p','888','/categories/createCategories','POST','','',''),('p','888','/categories/deleteCategories','DELETE','','',''),('p','888','/categories/deleteCategoriesByIds','DELETE','','',''),('p','888','/categories/updateCategories','PUT','','',''),('p','888','/categories/findCategories','GET','','',''),('p','888','/categories/getCategoriesList','GET','','',''),('p','888','/company/createCompany','POST','','',''),('p','888','/company/deleteCompany','DELETE','','',''),('p','888','/company/deleteCompanyByIds','DELETE','','',''),('p','888','/company/updateCompany','PUT','','',''),('p','888','/company/findCompany','GET','','',''),('p','888','/company/getCompanyList','GET','','',''),('p','888','/region/createRegion','POST','','',''),('p','888','/region/deleteRegion','DELETE','','',''),('p','888','/region/deleteRegionByIds','DELETE','','',''),('p','888','/region/updateRegion','PUT','','',''),('p','888','/region/findRegion','GET','','',''),('p','888','/region/getRegionList','GET','','',''),('p','888','/address/createAddress','POST','','',''),('p','888','/address/deleteAddress','DELETE','','',''),('p','888','/address/deleteAddressByIds','DELETE','','',''),('p','888','/address/updateAddress','PUT','','',''),('p','888','/address/findAddress','GET','','',''),('p','888','/address/getAddressList','GET','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `category_subject_id` int(10) NOT NULL COMMENT '类别主体id',
  `category_name` varchar(50) DEFAULT NULL COMMENT '公司类别名',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯类别表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (101,'2020-07-30 14:39:09','2020-07-30 14:39:09',NULL,100,'住宅客梯'),(102,'2020-07-30 17:05:44','2020-07-30 17:05:44',NULL,100,'住宅货梯'),(103,'2020-07-30 17:06:15','2020-07-30 17:06:15',NULL,100,'住宅扶梯'),(104,'2020-07-30 17:06:57','2020-07-30 17:06:57',NULL,100,'商场客梯'),(105,'2020-07-30 17:07:06','2020-07-30 17:07:06',NULL,100,'商场货梯'),(106,'2020-07-30 17:07:19','2020-07-30 17:07:19',NULL,100,'商场扶梯'),(107,'2020-07-30 17:08:45','2020-07-30 17:08:45',NULL,100,'医院客梯'),(108,'2020-07-30 17:09:00','2020-07-30 17:09:00',NULL,100,'医院货梯'),(109,'2020-07-30 17:09:11','2020-07-30 17:09:11',NULL,100,'医院扶梯'),(110,'2020-07-30 17:10:24','2020-07-30 17:10:24',NULL,100,'写字楼客梯'),(111,'2020-07-30 17:10:33','2020-07-30 17:10:33',NULL,100,'写字楼货梯'),(112,'2020-07-30 17:11:54','2020-07-30 17:11:54',NULL,100,'政府单位客梯'),(113,'2020-07-30 17:12:03','2020-07-30 17:12:03',NULL,100,'政府单位货梯'),(114,'2020-07-30 17:14:50','2020-07-30 17:14:50',NULL,101,'运营平台'),(115,'2020-07-30 17:16:21','2020-07-30 17:16:21',NULL,101,'物业'),(116,'2020-07-30 17:16:40','2020-07-30 17:16:40',NULL,101,'维保公司'),(117,'2020-07-30 17:17:15','2020-07-30 17:17:15',NULL,101,'电梯生产商'),(118,'2020-07-30 17:18:20','2020-07-30 17:18:20',NULL,101,'电梯年检单位'),(119,'2020-07-30 17:18:37','2020-07-30 17:18:37',NULL,101,'电梯安装公司'),(120,'2020-07-30 17:19:01','2020-07-30 17:19:01',NULL,101,'监督单位'),(121,'2020-07-30 17:19:21','2020-07-30 17:19:21',NULL,101,'电梯使用单位'),(122,'2020-08-03 19:57:12','2020-08-03 19:57:12',NULL,0,''),(123,'2020-08-04 13:33:54','2020-08-04 13:33:54',NULL,107,'公务员'),(124,'2020-08-04 13:58:15','2020-08-05 10:23:24',NULL,107,'教师'),(125,'2020-08-04 13:58:25','2020-08-05 10:24:22',NULL,107,'白领'),(126,'2020-08-04 13:58:34','2020-08-05 10:24:22',NULL,107,'交警'),(127,'2020-08-04 13:58:43','2020-08-05 10:24:22',NULL,107,'老年人'),(128,'2020-08-04 13:58:59','2020-08-05 10:24:22',NULL,107,'婴幼儿'),(129,'2020-08-04 13:59:09','2020-08-05 10:23:24',NULL,107,'中学生'),(130,'2020-08-04 13:59:16','2020-08-05 11:11:19',NULL,107,'大学生'),(131,'2020-08-04 15:56:36','2020-08-04 15:56:36',NULL,0,''),(132,'2020-08-04 16:22:53','2020-08-04 16:22:53',NULL,0,''),(133,'2020-08-04 16:26:40','2020-08-04 16:26:40',NULL,0,''),(134,'2020-08-05 10:23:24','2020-08-05 10:23:24',NULL,0,''),(135,'2020-08-05 10:24:22','2020-08-05 10:24:22',NULL,0,'');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_subjects`
--

DROP TABLE IF EXISTS `category_subjects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category_subjects` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `subject_name` varchar(50) DEFAULT NULL COMMENT '类别主体名',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_category_subjects_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='类别主体表，电梯/公司/记录/......';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_subjects`
--

LOCK TABLES `category_subjects` WRITE;
/*!40000 ALTER TABLE `category_subjects` DISABLE KEYS */;
INSERT INTO `category_subjects` VALUES (100,'2020-07-30 10:22:39','2020-07-30 17:12:55',NULL,'电梯使用场景'),(101,'2020-07-30 10:23:20','2020-07-30 10:23:20',NULL,'公司类别'),(102,'2020-07-30 10:24:48','2020-07-30 16:51:36',NULL,'用户电梯关系'),(103,'2020-07-30 10:25:12','2020-07-30 17:13:10',NULL,'电梯记录'),(104,'2020-07-30 10:25:35','2020-07-30 16:51:41',NULL,'电梯故障来源'),(105,'2020-07-30 10:26:03','2020-07-30 16:51:46',NULL,'电梯故障解除方式'),(106,'2020-07-30 10:26:22','2020-07-30 16:51:49',NULL,'电梯故障原因'),(107,'2020-08-04 13:33:32','2020-08-05 10:24:22',NULL,'地址标签');
/*!40000 ALTER TABLE `category_subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `companies`
--

DROP TABLE IF EXISTS `companies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `companies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `full_name` varchar(50) DEFAULT NULL COMMENT '公司全称',
  `alias` varchar(50) DEFAULT NULL COMMENT '公司简称',
  `legal_person` varchar(50) DEFAULT NULL COMMENT '法人信息',
  `phone_number` varchar(20) DEFAULT NULL COMMENT '联系方式',
  `status` varchar(20) DEFAULT NULL COMMENT '公司状态',
  `reg_code` varchar(20) DEFAULT NULL COMMENT '工商注册号',
  `org_code` varchar(20) DEFAULT NULL COMMENT '组织机构号',
  `credit_code` varchar(20) DEFAULT NULL COMMENT '统一信用代码',
  `tax_code` varchar(20) DEFAULT NULL COMMENT '纳税人识别号',
  `address` varchar(255) DEFAULT NULL COMMENT '注册地址',
  `category_id` int(10) NOT NULL COMMENT '电梯类别',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_companies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='公司表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `companies`
--

LOCK TABLES `companies` WRITE;
/*!40000 ALTER TABLE `companies` DISABLE KEYS */;
/*!40000 ALTER TABLE `companies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_customers`
--

DROP TABLE IF EXISTS `exa_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exa_customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `customer_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '客户电话',
  `sys_user_id` int(10) unsigned DEFAULT NULL COMMENT '负责员工id',
  `sys_user_authority_id` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '负责员工角色',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_customers_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_customers`
--

LOCK TABLES `exa_customers` WRITE;
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
INSERT INTO `exa_customers` VALUES (1,'2020-02-25 10:01:48','2020-04-10 04:29:29',NULL,'测试客户','1761111111',10,'888');
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_chunks`
--

DROP TABLE IF EXISTS `exa_file_chunks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exa_file_chunks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `exa_file_id` int(10) unsigned DEFAULT NULL COMMENT '文件id',
  `file_chunk_path` varchar(255) DEFAULT NULL COMMENT '切片路径',
  `file_chunk_number` int(11) DEFAULT NULL COMMENT '切片标号',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_chunks`
--

LOCK TABLES `exa_file_chunks` WRITE;
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '文件名',
  `url` varchar(255) DEFAULT NULL COMMENT '文件URL',
  `tag` varchar(255) DEFAULT NULL COMMENT '文件类型',
  `key` varchar(255) DEFAULT NULL COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_file_upload_and_downloads_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (17,'2020-04-26 03:51:39','2020-04-26 03:51:39',NULL,'10.png','http://qmplusimg.henrongyi.top/158787308910.png','png','158787308910.png'),(19,'2020-04-27 07:48:38','2020-04-27 07:48:38',NULL,'logo.png','http://qmplusimg.henrongyi.top/1587973709logo.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_files`
--

DROP TABLE IF EXISTS `exa_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `exa_files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL COMMENT '文件名',
  `file_md5` varchar(255) DEFAULT NULL COMMENT '文件md5',
  `file_path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `chunk_total` int(11) DEFAULT NULL COMMENT '切片总数',
  `is_finish` tinyint(1) DEFAULT NULL COMMENT '是否完整',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_files_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_files`
--

LOCK TABLES `exa_files` WRITE;
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jwt_blacklists` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lift_changes`
--

DROP TABLE IF EXISTS `lift_changes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lift_changes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `lift_id` int(10) unsigned NOT NULL COMMENT '电梯id',
  `content` varchar(255) DEFAULT NULL COMMENT '变更内容',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_lift_changes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯变更表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lift_changes`
--

LOCK TABLES `lift_changes` WRITE;
/*!40000 ALTER TABLE `lift_changes` DISABLE KEYS */;
/*!40000 ALTER TABLE `lift_changes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lift_models`
--

DROP TABLE IF EXISTS `lift_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lift_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `factory_id` int(10) unsigned DEFAULT NULL COMMENT '电梯生产商id',
  `brand` varchar(50) DEFAULT NULL COMMENT '电梯品牌',
  `modal` varchar(50) DEFAULT NULL COMMENT '电梯型号',
  `load` int(11) DEFAULT NULL COMMENT '电梯载重',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_lift_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯型号表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lift_models`
--

LOCK TABLES `lift_models` WRITE;
/*!40000 ALTER TABLE `lift_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `lift_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lift_records`
--

DROP TABLE IF EXISTS `lift_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lift_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `lift_id` int(10) unsigned NOT NULL COMMENT '电梯id',
  `category_id` int(10) unsigned NOT NULL COMMENT '记录类别',
  `images` varchar(255) DEFAULT NULL COMMENT '图片记录',
  `content` varchar(255) DEFAULT NULL COMMENT '文字记录',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '结束时间',
  `worker_id` int(10) unsigned NOT NULL COMMENT '操作人员',
  `recorder_id` int(10) unsigned NOT NULL COMMENT '记录人员',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_lift_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯操作记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lift_records`
--

LOCK TABLES `lift_records` WRITE;
/*!40000 ALTER TABLE `lift_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `lift_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lift_troubles`
--

DROP TABLE IF EXISTS `lift_troubles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lift_troubles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `lift_id` int(10) unsigned NOT NULL COMMENT '电梯id',
  `from_category_id` int(10) unsigned NOT NULL COMMENT '故障来源类别',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '故障开始时间',
  `start_user_id` int(10) DEFAULT NULL COMMENT '发起故障人员',
  `response_time` timestamp NULL DEFAULT NULL COMMENT '故障响应时间',
  `response_user_id` int(10) DEFAULT NULL COMMENT '故障响应人员',
  `scene_time` timestamp NULL DEFAULT NULL COMMENT '达到现场时间',
  `scene_user_id` int(10) DEFAULT NULL COMMENT '达到现场人员',
  `fix_time` timestamp NULL DEFAULT NULL COMMENT '解除故障时间',
  `fix_user_id` int(10) DEFAULT NULL COMMENT '解除故障人员',
  `fix_category_id` int(10) unsigned DEFAULT NULL COMMENT '解除故障方式类别',
  `reason_category_id` int(10) unsigned DEFAULT NULL COMMENT '故障原因类别',
  `content` varchar(255) DEFAULT NULL COMMENT '故障详情',
  `progress` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '故障进度',
  `recorder_id` int(10) unsigned NOT NULL COMMENT '记录人员',
  `feedback_content` varchar(255) DEFAULT NULL COMMENT '反馈内容',
  `feedback_rate` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '反馈评分',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_lift_troubles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯故障记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lift_troubles`
--

LOCK TABLES `lift_troubles` WRITE;
/*!40000 ALTER TABLE `lift_troubles` DISABLE KEYS */;
/*!40000 ALTER TABLE `lift_troubles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lifts`
--

DROP TABLE IF EXISTS `lifts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lifts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL COMMENT '电梯别名',
  `code` varchar(255) DEFAULT NULL COMMENT '电梯编号',
  `installer_id` int(10) unsigned DEFAULT NULL COMMENT '电梯安装公司id',
  `maintainer_id` int(10) unsigned DEFAULT NULL COMMENT '电梯维保公司id',
  `checker_id` int(10) unsigned DEFAULT NULL COMMENT '电梯年检公司id',
  `owner_id` int(10) unsigned DEFAULT NULL COMMENT '电梯使用公司id',
  `factory_time` timestamp NULL DEFAULT NULL COMMENT '电梯出厂时间',
  `install_time` timestamp NULL DEFAULT NULL COMMENT '电梯安装时间',
  `check_time` timestamp NULL DEFAULT NULL COMMENT '电梯年检时间',
  `lift_model_id` int(10) unsigned DEFAULT NULL COMMENT '电梯型号',
  `category_id` int(10) NOT NULL COMMENT '电梯类别',
  `floor_count` int(11) NOT NULL COMMENT '总楼层',
  `latitude` decimal(10,7) DEFAULT NULL COMMENT '电梯位置地理经度',
  `longitude` decimal(10,7) DEFAULT NULL COMMENT '电梯位置地理纬度',
  `address_id` int(10) DEFAULT NULL COMMENT '地址id',
  `region_id` int(11) DEFAULT NULL COMMENT '区域id',
  `building` varchar(50) DEFAULT NULL COMMENT '楼栋',
  `cell` int(11) DEFAULT NULL COMMENT '单元',
  `ad_device_id` int(10) unsigned DEFAULT NULL COMMENT '广告机设备id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_lifts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='电梯表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lifts`
--

LOCK TABLES `lifts` WRITE;
/*!40000 ALTER TABLE `lifts` DISABLE KEYS */;
/*!40000 ALTER TABLE `lifts` ENABLE KEYS */;
UNLOCK TABLES;


drop table if exists `ad_devices`;
create table `ad_devices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `model_id` int(10) not null comment '电梯型号',
  `factory_id` int(10) not null comment '生产商',
  `factory_time` datetime default null comment '出厂时间',
  `install_time` datetime default null comment '安装时间',
  `status_id` int(10) default null comment '设备状态类别',
  `online` boolean default true comment '设备是否在线',
  `last_offline_time` datetime default null comment '上次离线时间',
  `last_online_time` datetime default null comment '上次上线时间',
  `owner_id` int(10) default null comment '责任人',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_ad_devices_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='广告机设备表';

drop table if exists `ad_device_events`;
create table `ad_device_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `device_id` int(10) not null comment '设备id',
  `type` int(10) not null comment '事件类型category',
  `content` char(255) default null comment '事件内容',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_ad_device_events_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='广告机设备事件表';

drop table if exists `ad_device_configs`;
create table `ad_device_configs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `device_id` int(10) not null comment '设备id',
  `global` boolean default true comment '全局配置',
  `key` char(50) not null comment '配置键值',
  `value` char(255) not null  comment '配置内容',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_ad_device_configs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='广告机设备配置表';

drop table if exists `ad_device_datas`;
create table `ad_device_datas` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `device_id` int(10) not null comment '设备id',
  `accx` float default 0.0 comment 'x轴加速度',
  `accy` float default 0.0 comment 'y轴加速度',
  `accz` float default 0.0 comment 'z轴加速度',
  `degx` float default 0.0 comment 'x轴倾斜角',
  `degy` float default 0.0 comment 'y轴倾斜角',
  `degz` float default 0.0 comment 'z轴倾斜角',
  `speedz` float default 0.0 comment '电梯速度',
  `floor` float default 1 comment '电梯当前楼层',
  `door_state` int(10) default null comment '电梯门状态类型',
  `people_inside` boolean default true comment '电梯内是否有人',
  `trouble_type` int(10) default null comment '电梯状态/故障类型',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_ad_device_datas_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='广告机设备数据表';

--
-- Table structure for table `regions`
--

DROP TABLE IF EXISTS `regions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `regions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `province` varchar(50) DEFAULT NULL COMMENT '省份',
  `city` varchar(50) DEFAULT NULL COMMENT '城市',
  `district` varchar(50) DEFAULT NULL COMMENT '行政区',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_regions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='区域表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `regions`
--

LOCK TABLES `regions` WRITE;
/*!40000 ALTER TABLE `regions` DISABLE KEYS */;
INSERT INTO `regions` VALUES (100,'2020-07-31 15:46:33','2020-07-31 17:46:02',NULL,'湖北省','武汉市','东西湖区'),(101,'2020-07-31 16:46:50','2020-07-31 16:46:50',NULL,'湖北省','武汉市','江岸区'),(102,'2020-08-04 14:13:02','2020-08-05 11:11:18',NULL,'湖北省','武汉市','东湖高新区');
/*!40000 ALTER TABLE `regions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `authority_id` int(10) unsigned DEFAULT NULL COMMENT '角色id',
  `path` varchar(255) DEFAULT NULL COMMENT '路由path',
  `description` varchar(255) DEFAULT NULL COMMENT '路由描述',
  `api_group` varchar(255) DEFAULT NULL COMMENT '路由分组',
  `method` varchar(255) DEFAULT 'POST' COMMENT '请求方法',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_apis_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2019-09-28 03:23:49','2019-09-28 09:06:16',NULL,NULL,'/base/login','用户登录','base','POST'),(2,'2019-09-28 03:32:46','2019-09-28 09:06:11',NULL,NULL,'/base/register','用户注册','base','POST'),(3,'2019-09-28 03:33:41','2020-05-10 08:33:48',NULL,NULL,'/api/createApi','创建api','api','POST'),(4,'2019-09-28 06:09:04','2019-09-28 09:05:59',NULL,NULL,'/api/getApiList','获取api列表','api','POST'),(5,'2019-09-28 06:15:50','2019-09-28 09:05:53',NULL,NULL,'/api/getApiById','获取api详细信息','api','POST'),(7,'2019-09-28 06:19:26','2019-09-28 09:05:44',NULL,NULL,'/api/deleteApi','删除Api','api','POST'),(8,'2019-09-28 06:19:48','2019-09-28 09:05:39',NULL,NULL,'/api/updateApi','更新Api','api','POST'),(10,'2019-09-30 07:05:38','2019-09-30 07:05:38',NULL,NULL,'/api/getAllApis','获取所有api','api','POST'),(11,'2019-09-30 07:23:09','2019-09-30 07:23:09',NULL,NULL,'/authority/createAuthority','创建角色','authority','POST'),(12,'2019-09-30 07:23:33','2019-09-30 07:23:33',NULL,NULL,'/authority/deleteAuthority','删除角色','authority','POST'),(13,'2019-09-30 07:23:57','2019-09-30 07:23:57',NULL,NULL,'/authority/getAuthorityList','获取角色列表','authority','POST'),(14,'2019-09-30 07:24:20','2019-09-30 07:24:20',NULL,NULL,'/menu/getMenu','获取菜单树','menu','POST'),(15,'2019-09-30 07:24:50','2019-09-30 07:24:50',NULL,NULL,'/menu/getMenuList','分页获取基础menu列表','menu','POST'),(16,'2019-09-30 07:25:07','2019-09-30 07:25:07',NULL,NULL,'/menu/addBaseMenu','新增菜单','menu','POST'),(17,'2019-09-30 07:25:25','2019-09-30 07:25:25',NULL,NULL,'/menu/getBaseMenuTree','获取用户动态路由','menu','POST'),(18,'2019-09-30 07:25:53','2019-09-30 07:25:53',NULL,NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','menu','POST'),(19,'2019-09-30 07:26:20','2019-09-30 07:26:20',NULL,NULL,'/menu/getMenuAuthority','获取指定角色menu','menu','POST'),(20,'2019-09-30 07:26:43','2019-09-30 07:26:43',NULL,NULL,'/menu/deleteBaseMenu','删除菜单','menu','POST'),(21,'2019-09-30 07:28:05','2019-09-30 07:28:05',NULL,NULL,'/menu/updateBaseMenu','更新菜单','menu','POST'),(22,'2019-09-30 07:28:21','2019-09-30 07:28:21',NULL,NULL,'/menu/getBaseMenuById','根据id获取菜单','menu','POST'),(23,'2019-09-30 07:29:19','2019-09-30 07:29:19',NULL,NULL,'/user/changePassword','修改密码','user','POST'),(24,'2019-09-30 07:29:33','2019-09-30 07:29:33',NULL,NULL,'/user/uploadHeaderImg','上传头像','user','POST'),(25,'2019-09-30 07:30:00','2019-09-30 07:30:00',NULL,NULL,'/user/getInfoList','分页获取用户列表','user','POST'),(28,'2019-10-09 07:15:17','2019-10-09 07:17:07',NULL,NULL,'/user/getUserList','获取用户列表','user','POST'),(29,'2019-10-09 15:01:40','2019-10-09 15:01:40',NULL,NULL,'/user/setUserAuthority','修改用户角色','user','POST'),(30,'2019-10-26 12:14:38','2019-10-26 12:14:38',NULL,NULL,'/fileUploadAndDownload/upload','文件上传示例','fileUploadAndDownload','POST'),(31,'2019-10-26 12:14:59','2019-10-26 12:14:59',NULL,NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','fileUploadAndDownload','POST'),(32,'2019-12-12 05:28:47','2019-12-12 05:28:47',NULL,NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(33,'2019-12-12 05:28:59','2019-12-12 05:28:59',NULL,NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(34,'2019-12-12 09:02:15','2019-12-12 09:02:15',NULL,NULL,'/fileUploadAndDownload/deleteFile','删除文件','fileUploadAndDownload','POST'),(35,'2019-12-28 10:18:07','2019-12-28 10:18:07',NULL,NULL,'/jwt/jsonInBlacklist','jwt加入黑名单','jwt','POST'),(36,'2020-01-06 09:56:36','2020-01-06 09:56:36',NULL,NULL,'/authority/setDataAuthority','设置角色资源权限','authority','POST'),(37,'2020-01-13 06:04:05','2020-01-13 06:04:05',NULL,NULL,'/system/getSystemConfig','获取配置文件内容','system','POST'),(38,'2020-01-13 07:02:06','2020-01-13 07:02:06',NULL,NULL,'/system/setSystemConfig','设置配置文件内容','system','POST'),(39,'2020-02-25 07:32:39','2020-02-25 07:32:39',NULL,NULL,'/customer/customer','创建客户','customer','POST'),(40,'2020-02-25 07:32:51','2020-02-25 07:34:56',NULL,NULL,'/customer/customer','更新客户','customer','PUT'),(41,'2020-02-25 07:33:57','2020-02-25 07:33:57',NULL,NULL,'/customer/customer','删除客户','customer','DELETE'),(42,'2020-02-25 07:36:48','2020-02-25 07:37:16',NULL,NULL,'/customer/customer','获取单一客户','customer','GET'),(43,'2020-02-25 07:37:06','2020-02-25 07:37:06',NULL,NULL,'/customer/customerList','获取客户列表','customer','GET'),(44,'2020-03-12 06:36:54','2020-03-12 06:56:50',NULL,NULL,'/casbin/casbinTest/:pathParam','RESTFUL模式测试','casbin','GET'),(45,'2020-03-29 15:01:28','2020-03-29 15:01:28',NULL,NULL,'/autoCode/createTemp','自动化代码','autoCode','POST'),(46,'2020-04-15 04:46:58','2020-04-15 04:46:58',NULL,NULL,'/authority/updateAuthority','更新角色信息','authority','PUT'),(47,'2020-04-20 07:14:25','2020-04-20 07:14:25',NULL,NULL,'/authority/copyAuthority','拷贝角色','authority','POST'),(64,'2020-05-10 08:44:25','2020-05-10 08:44:25',NULL,NULL,'/user/deleteUser','删除用户','user','DELETE'),(81,'2020-06-23 10:40:50','2020-06-23 10:40:50',NULL,NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','sysDictionaryDetail','POST'),(82,'2020-06-23 10:40:50','2020-06-23 10:40:50',NULL,NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','sysDictionaryDetail','DELETE'),(83,'2020-06-23 10:40:50','2020-06-23 10:40:50',NULL,NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','sysDictionaryDetail','PUT'),(84,'2020-06-23 10:40:50','2020-06-23 10:40:50',NULL,NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','sysDictionaryDetail','GET'),(85,'2020-06-23 10:40:50','2020-06-23 10:40:50',NULL,NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','sysDictionaryDetail','GET'),(86,'2020-06-23 10:48:13','2020-06-23 10:48:13',NULL,NULL,'/sysDictionary/createSysDictionary','新增字典','sysDictionary','POST'),(87,'2020-06-23 10:48:13','2020-06-23 10:48:13',NULL,NULL,'/sysDictionary/deleteSysDictionary','删除字典','sysDictionary','DELETE'),(88,'2020-06-23 10:48:13','2020-06-23 10:48:13',NULL,NULL,'/sysDictionary/updateSysDictionary','更新字典','sysDictionary','PUT'),(89,'2020-06-23 10:48:13','2020-06-23 10:48:13',NULL,NULL,'/sysDictionary/findSysDictionary','根据ID获取字典','sysDictionary','GET'),(90,'2020-06-23 10:48:13','2020-06-23 10:48:13',NULL,NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','sysDictionary','GET'),(91,'2020-06-29 05:21:35','2020-06-29 05:21:35',NULL,NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','sysOperationRecord','POST'),(92,'2020-06-29 05:21:35','2020-06-29 05:21:35',NULL,NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','sysOperationRecord','DELETE'),(93,'2020-06-29 05:21:35','2020-06-29 05:21:35',NULL,NULL,'/sysOperationRecord/updateSysOperationRecord','更新操作记录','sysOperationRecord','PUT'),(94,'2020-06-29 05:21:35','2020-06-29 05:21:35',NULL,NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','sysOperationRecord','GET'),(95,'2020-06-29 05:21:35','2020-06-29 05:21:35',NULL,NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','sysOperationRecord','GET'),(96,'2020-07-05 06:34:20','2020-07-05 06:34:20',NULL,NULL,'/autoCode/getTables','获取数据库表','autoCode','GET'),(97,'2020-07-05 07:02:07','2020-07-05 07:02:07',NULL,NULL,'/autoCode/getDB','获取所有数据库','autoCode','GET'),(98,'2020-07-05 08:32:08','2020-07-05 08:32:08',NULL,NULL,'/autoCode/getColume','获取所选table的所有字段','autoCode','GET'),(99,'2020-07-07 07:59:53','2020-07-07 07:59:53',NULL,NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','sysOperationRecord','DELETE'),(106,'2020-07-30 02:13:21','2020-07-30 02:13:21',NULL,NULL,'/subject/getSubjectById','getSubjectById','subject','POST'),(107,'2020-07-30 02:13:38','2020-07-30 02:13:38',NULL,NULL,'/subject/getSubjectList','getSubjectList','subject','POST'),(108,'2020-07-30 02:13:54','2020-07-30 02:13:54',NULL,NULL,'/subject/createSubject','createSubject','subject','POST'),(109,'2020-07-30 02:14:12','2020-07-30 02:14:12',NULL,NULL,'/subject/updateSubject','updateSubject','subject','POST'),(110,'2020-07-30 02:14:26','2020-07-30 02:14:26',NULL,NULL,'/subject/deleteSubject','deleteSubject','subject','POST'),(111,'2020-07-30 02:14:41','2020-07-30 02:14:41',NULL,NULL,'/subject/getAllSubjects','getAllSubjects','subject','POST'),(112,'2020-07-30 02:31:55','2020-07-30 02:31:55',NULL,NULL,'/categories/createCategories','新增categories表','categories','POST'),(113,'2020-07-30 02:31:55','2020-07-30 02:31:55',NULL,NULL,'/categories/deleteCategories','删除categories表','categories','DELETE'),(114,'2020-07-30 02:31:55','2020-07-30 02:31:55',NULL,NULL,'/categories/deleteCategoriesByIds','批量删除categories表','categories','DELETE'),(115,'2020-07-30 02:31:55','2020-07-30 02:31:55',NULL,NULL,'/categories/updateCategories','更新categories表','categories','PUT'),(116,'2020-07-30 02:31:56','2020-07-30 02:31:56',NULL,NULL,'/categories/findCategories','根据ID获取categories表','categories','GET'),(117,'2020-07-30 02:31:56','2020-07-30 02:31:56',NULL,NULL,'/categories/getCategoriesList','获取categories表列表','categories','GET'),(118,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/createCompany','新增company表','company','POST'),(119,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/deleteCompany','删除company表','company','DELETE'),(120,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/deleteCompanyByIds','批量删除company表','company','DELETE'),(121,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/updateCompany','更新company表','company','PUT'),(122,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/findCompany','根据ID获取company表','company','GET'),(123,'2020-07-30 09:53:38','2020-07-30 09:53:38',NULL,NULL,'/company/getCompanyList','获取company表列表','company','GET'),(124,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/createRegion','新增region表','region','POST'),(125,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/deleteRegion','删除region表','region','DELETE'),(126,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/deleteRegionByIds','批量删除region表','region','DELETE'),(127,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/updateRegion','更新region表','region','PUT'),(128,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/findRegion','根据ID获取region表','region','GET'),(129,'2020-07-31 05:47:42','2020-07-31 05:47:42',NULL,NULL,'/region/getRegionList','获取region表列表','region','GET'),(130,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/createAddress','新增address表','address','POST'),(131,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/deleteAddress','删除address表','address','DELETE'),(132,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/deleteAddressByIds','批量删除address表','address','DELETE'),(133,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/updateAddress','更新address表','address','PUT'),(134,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/findAddress','根据ID获取address表','address','GET'),(135,'2020-08-03 10:25:35','2020-08-03 10:25:35',NULL,NULL,'/address/getAddressList','获取address表列表','address','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_authorities` (
  `authority_id` varchar(255) NOT NULL COMMENT '角色id',
  `authority_name` varchar(255) DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(255) DEFAULT NULL COMMENT '父角色',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `authority_id` (`authority_id`) USING BTREE,
  KEY `idx_sys_authorities_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('888','普通用户','0','2020-04-04 11:44:56','2020-06-13 16:07:37',NULL),('8881','普通用户子角色','888','2020-04-04 11:44:56','2020-04-24 10:16:42',NULL),('9528','测试角色','0','2020-04-04 11:44:56','2020-04-24 10:16:42',NULL);
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_authority_menus` (
  `sys_authority_authority_id` varchar(255) NOT NULL COMMENT '角色id',
  `sys_base_menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  PRIMARY KEY (`sys_authority_authority_id`,`sys_base_menu_id`) USING BTREE,
  KEY `sys_authority_authority_id` (`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES ('888',1),('888',2),('888',3),('888',4),('888',5),('888',6),('888',17),('888',18),('888',19),('888',20),('888',21),('888',22),('888',23),('888',26),('888',33),('888',34),('888',38),('888',40),('888',41),('888',42),('888',50),('888',51),('888',52),('888',54),('888',55),('888',56),('888',57),('888',58),('888',59),('888',60),('888',61),('888',62),('8881',1),('8881',2),('8881',18),('8881',38),('8881',40),('8881',41),('8881',42),('9528',1),('9528',2),('9528',3),('9528',4),('9528',5),('9528',6),('9528',17),('9528',18),('9528',19),('9528',20),('9528',21),('9528',22),('9528',23),('9528',26),('9528',33),('9528',34),('9528',38),('9528',40),('9528',41),('9528',42);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_base_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `menu_level` int(10) unsigned DEFAULT NULL COMMENT '菜单等级（暂未使用）',
  `parent_id` int(10) unsigned DEFAULT NULL COMMENT '父菜单id',
  `path` varchar(255) DEFAULT NULL COMMENT '菜单path（路由path）',
  `name` varchar(255) DEFAULT NULL COMMENT '菜单name（路由name）',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(255) DEFAULT NULL COMMENT '组件位置',
  `title` varchar(255) DEFAULT NULL COMMENT '显示名字',
  `icon` varchar(255) DEFAULT NULL COMMENT '显示图标',
  `nick_name` varchar(255) DEFAULT NULL COMMENT '菜单别名',
  `sort` int(255) DEFAULT NULL COMMENT '排序',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存菜单内容',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '默认菜单（暂未使用）',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_base_menus_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2019-09-19 14:05:18','2020-05-30 07:43:06',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue','仪表盘','setting','仪表盘',1,0,0),(2,'2019-09-19 14:06:17','2020-07-31 01:49:33',NULL,0,0,'about','about',0,'view/about/index.vue','关于我们','info','测试菜单',80,0,0),(3,'2019-09-19 14:06:38','2020-07-31 01:49:07',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue','超级管理员','user-solid','超级管理员',30,0,0),(4,'2019-09-19 14:11:53','2020-05-30 07:43:25',NULL,0,3,'authority','authority',0,'view/superAdmin/authority/authority.vue','角色管理','s-custom','角色管理',1,0,0),(5,'2019-09-19 14:13:18','2020-04-30 09:45:27',NULL,0,3,'menu','menu',0,'view/superAdmin/menu/menu.vue','菜单管理','s-order','菜单管理',2,1,0),(6,'2019-09-19 14:13:36','2020-04-24 02:16:43',NULL,0,3,'api','api',0,'view/superAdmin/api/api.vue','api管理','s-platform','api管理',3,1,0),(17,'2019-10-09 07:12:29','2020-04-24 02:16:43',NULL,0,3,'user','user',0,'view/superAdmin/user/user.vue','用户管理','coordinate','用户管理',4,0,0),(18,'2019-10-15 14:27:22','2020-07-31 01:49:19',NULL,0,0,'person','person',1,'view/person/person.vue','个人信息','message-solid','个人信息',50,0,0),(19,'2019-10-20 03:14:42','2020-07-31 01:49:29',NULL,0,0,'example','example',0,'view/example/index.vue','示例文件','s-management','示例文件',70,0,0),(20,'2019-10-20 03:18:11','2020-04-24 02:16:42',NULL,0,19,'table','table',0,'view/example/table/table.vue','表格示例','s-order','表格示例',1,0,0),(21,'2019-10-20 03:19:52','2020-04-24 02:16:43',NULL,0,19,'form','form',0,'view/example/form/form.vue','表单示例','document','表单示例',2,0,0),(22,'2019-10-20 03:22:19','2020-04-24 02:16:43',NULL,0,19,'rte','rte',0,'view/example/rte/rte.vue','富文本编辑器','reading','富文本编辑器',3,0,0),(23,'2019-10-20 03:23:39','2020-04-24 02:16:43',NULL,0,19,'excel','excel',0,'view/example/excel/excel.vue','excel导入导出','s-marketing','excel导入导出',4,0,0),(26,'2019-10-20 03:27:02','2020-04-24 02:16:43',NULL,0,19,'upload','upload',0,'view/example/upload/upload.vue','上传下载','upload','上传下载',5,0,0),(33,'2020-02-17 08:20:47','2020-04-24 02:16:43',NULL,0,19,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue','断点续传','upload','断点续传',6,0,0),(34,'2020-02-24 11:48:37','2020-04-24 02:16:43',NULL,0,19,'customer','customer',0,'view/example/customer/customer.vue','客户列表（资源示例）','s-custom','客户列表（资源示例）',7,0,0),(38,'2020-03-29 13:31:03','2020-07-31 01:49:24',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue','系统工具','s-cooperation','系统工具',60,0,0),(40,'2020-03-29 13:35:10','2020-05-03 13:38:49',NULL,0,38,'autoCode','autoCode',0,'view/systemTools/autoCode/index.vue','代码生成器','cpu','代码生成器',1,1,0),(41,'2020-03-29 13:36:26','2020-05-03 13:38:43',NULL,0,38,'formCreate','formCreate',0,'view/systemTools/formCreate/index.vue','表单生成器','magic-stick','表单生成器',2,1,0),(42,'2020-04-02 06:19:36','2020-04-24 02:16:43',NULL,0,38,'system','system',0,'view/systemTools/system/system.vue','系统配置','s-operation','系统配置',3,0,0),(45,'2020-04-29 09:19:34','2020-07-31 01:48:56',NULL,0,0,'iconList','iconList',0,'view/iconList/index.vue','图标集合','star-on',NULL,20,0,0),(50,'2020-06-24 11:49:54','2020-06-28 12:34:47',NULL,0,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue','字典管理','notebook-2',NULL,5,0,0),(51,'2020-06-24 11:51:33','2020-06-28 12:35:04',NULL,0,3,'dictionaryDetail/:id','dictionaryDetail',1,'view/superAdmin/dictionary/sysDictionaryDetail.vue','字典详情','s-order',NULL,1,0,0),(52,'2020-06-29 05:31:17','2020-07-07 08:05:34',NULL,0,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue','操作历史','time',NULL,6,0,0),(54,'2020-07-30 02:05:03','2020-07-31 01:49:13',NULL,0,0,'category','category',0,'view/category/index.vue','类别管理','s-order',NULL,40,0,0),(55,'2020-07-30 02:06:20','2020-07-30 02:18:19',NULL,0,54,'subject','categorySubject',0,'view/category/subject/subject.vue','类别主体','s-shop',NULL,1,0,0),(56,'2020-07-30 02:08:20','2020-07-30 08:56:59',NULL,0,54,'categories','categoryItem',0,'view/category/categories/categories.vue','类别列表','money',NULL,2,0,0),(57,'2020-07-31 01:50:56','2020-07-31 01:51:10',NULL,0,0,'company','company',0,'view/company/index.vue','公司管理','school',NULL,35,0,0),(58,'2020-07-31 01:53:14','2020-07-31 01:55:10',NULL,0,57,'list','list',0,'view/company/list/list.vue','公司列表','s-fold',NULL,1,0,0),(59,'2020-07-31 01:58:12','2020-07-31 01:58:12',NULL,0,57,'employee','employee',0,'view/company/employee/employee.vue','员工管理','s-custom',NULL,2,0,0),(60,'2020-07-31 05:53:49','2020-07-31 05:53:49',NULL,0,0,'address','address',0,'view/address/index.vue','地址管理','add-location',NULL,45,0,0),(61,'2020-07-31 05:54:58','2020-07-31 05:54:58',NULL,0,60,'region','region',0,'view/address/region/region.vue','区域划分','notebook-1',NULL,1,0,0),(62,'2020-07-31 06:00:03','2020-07-31 06:00:03',NULL,0,60,'addressList','addressList',0,'view/address/list/address.vue','常用地址','medal-1',NULL,2,0,0);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(255) NOT NULL COMMENT '角色id',
  `data_authority_id` varchar(255) NOT NULL COMMENT '拥有的资源角色id',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id`) USING BTREE,
  KEY `sys_authority_authority_id` (`sys_authority_authority_id`) USING BTREE,
  KEY `data_authority_id` (`data_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES ('888','888'),('888','8881'),('888','9528'),('888222','888'),('888222','8881'),('888222','9528'),('8883','888'),('8883','8881'),('8883','9528'),('9528','8881'),('9528','9528');
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionaries`
--

DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dictionaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(255) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionaries`
--

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` VALUES (2,'2020-06-24 20:44:00','2020-06-24 20:44:00',NULL,'性别','sex',1,'性别字典'),(3,'2020-07-05 15:27:31','2020-07-05 15:27:31',NULL,'数据库int类型','int',1,'int类型对应的数据库类型'),(4,'2020-07-05 15:33:07','2020-07-05 16:07:18',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型'),(5,'2020-07-05 15:34:23','2020-07-05 15:52:45',NULL,'数据库浮点型','float64',1,'数据库浮点型'),(6,'2020-07-05 15:35:05','2020-07-05 15:35:05',NULL,'数据库字符串','string',1,'数据库字符串'),(7,'2020-07-05 15:36:48','2020-07-05 15:36:48',NULL,'数据库bool类型','bool',1,'数据库bool类型');
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dictionary_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `label` varchar(255) DEFAULT NULL COMMENT '展示值',
  `value` int(11) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` int(11) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` int(11) DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (12,'2020-07-05 15:31:41','2020-07-05 15:31:41',NULL,'smallint',1,1,1,3),(13,'2020-07-05 15:31:52','2020-07-05 15:31:52',NULL,'mediumint',2,1,2,3),(14,'2020-07-05 15:32:04','2020-07-05 15:32:04',NULL,'int',3,1,3,3),(15,'2020-07-05 15:32:11','2020-07-05 15:32:11',NULL,'bigint',4,1,4,3),(19,'2020-07-05 15:33:16','2020-07-05 15:33:16',NULL,'data',0,1,0,4),(20,'2020-07-05 15:33:21','2020-07-05 15:33:21',NULL,'time',1,1,1,4),(21,'2020-07-05 15:33:25','2020-07-05 15:33:25',NULL,'year',2,1,2,4),(22,'2020-07-05 15:33:35','2020-07-05 15:33:35',NULL,'datetime',3,1,3,4),(23,'2020-07-05 15:33:42','2020-07-05 15:33:42',NULL,'timestamp',5,1,5,4),(24,'2020-07-05 15:34:30','2020-07-05 15:34:30',NULL,'float',0,1,0,5),(25,'2020-07-05 15:34:35','2020-07-05 15:34:35',NULL,'double',1,1,1,5),(26,'2020-07-05 15:34:41','2020-07-05 15:34:41',NULL,'decimal',2,1,2,5),(27,'2020-07-05 15:37:45','2020-07-05 15:37:45',NULL,'tinyint',0,1,0,7),(28,'2020-07-05 15:53:25','2020-07-05 15:53:25',NULL,'char',0,1,0,6),(29,'2020-07-05 15:53:29','2020-07-05 15:53:29',NULL,'varchar',1,1,1,6),(30,'2020-07-05 15:53:35','2020-07-05 15:53:35',NULL,'tinyblob',2,1,2,6),(31,'2020-07-05 15:53:40','2020-07-05 15:53:40',NULL,'tinytext',3,1,3,6),(32,'2020-07-05 15:53:48','2020-07-05 15:53:48',NULL,'text',4,1,4,6),(33,'2020-07-05 15:53:55','2020-07-05 15:53:55',NULL,'blob',5,1,5,6),(34,'2020-07-05 15:54:02','2020-07-05 15:54:02',NULL,'mediumblob',6,1,6,6),(35,'2020-07-05 15:54:09','2020-07-05 15:54:09',NULL,'mediumtext',7,1,7,6),(36,'2020-07-05 15:54:16','2020-07-05 15:54:16',NULL,'longblob',8,1,8,6),(37,'2020-07-05 15:54:24','2020-07-05 15:54:24',NULL,'longtext',9,1,9,6);
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operation_records`
--

DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_operation_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求ip',
  `method` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求路由',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `latency` bigint(20) DEFAULT NULL,
  `agent` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `error_message` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `body` text COLLATE utf8mb4_bin COMMENT '请求Body',
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `resp` text COLLATE utf8mb4_bin COMMENT '响应Body',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1478 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varbinary(255) DEFAULT NULL COMMENT 'uuid',
  `nick_name` varchar(255) DEFAULT 'QMPlusUser' COMMENT '用户昵称',
  `header_img` varchar(255) DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg' COMMENT '用户头像',
  `authority_id` double DEFAULT '888' COMMENT '用户角色',
  `username` varchar(255) DEFAULT NULL COMMENT '登录用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '登录密码',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_users_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (10,'2019-09-13 09:23:46','2020-06-26 13:17:50',NULL,_binary 'ce0d6685-c15f-4126-a5b4-890bc9d2356d','超级管理员','http://qmplusimg.henrongyi.top/1571627762timg.jpg',888,'admin','e10adc3949ba59abbe56e057f20f883e'),(11,'2019-09-13 09:27:29','2019-09-13 09:27:29',NULL,_binary 'fd6ef79b-944c-4888-8377-abe2d2608858','QMPlusUser','http://qmplusimg.henrongyi.top/1572075907logo.png',9528,'a303176530','3ec063004a6f31642261936a379fde3d');
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_workflow_step_infos`
--

DROP TABLE IF EXISTS `sys_workflow_step_infos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_workflow_step_infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `workflow_id` int(10) unsigned DEFAULT NULL COMMENT '节点id',
  `is_strat` tinyint(1) DEFAULT NULL COMMENT '是否是开始节点',
  `step_name` varchar(255) DEFAULT NULL COMMENT '步骤name',
  `step_no` double DEFAULT NULL COMMENT '第几步',
  `step_authority_id` varchar(255) DEFAULT NULL COMMENT '可操作者角色',
  `is_end` tinyint(1) DEFAULT NULL COMMENT '是否是结尾',
  `sys_workflow_id` int(10) unsigned DEFAULT NULL COMMENT '关联工作流id',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_workflow_step_infos_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_workflow_step_infos_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_workflow_step_infos`
--

LOCK TABLES `sys_workflow_step_infos` WRITE;
/*!40000 ALTER TABLE `sys_workflow_step_infos` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_workflow_step_infos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_workflows`
--

DROP TABLE IF EXISTS `sys_workflows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_workflows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `workflow_nick_name` varchar(255) DEFAULT NULL COMMENT '工作流中文名称',
  `workflow_name` varchar(255) DEFAULT NULL COMMENT '工作流英文名称',
  `workflow_description` varchar(255) DEFAULT NULL COMMENT '工作流描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_workflows_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_workflows`
--

LOCK TABLES `sys_workflows` WRITE;
/*!40000 ALTER TABLE `sys_workflows` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_workflows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_admins`
--

DROP TABLE IF EXISTS `user_admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varbinary(255) DEFAULT NULL COMMENT 'uuid',
  `phone_number` varchar(20) NOT NULL COMMENT '登陆手机号',
  `password` varchar(255) DEFAULT NULL COMMENT '登陆密码',
  `real_name` varchar(255) DEFAULT NULL COMMENT '用户真名',
  `nick_name` varchar(255) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `role_id` int(10) unsigned NOT NULL COMMENT '角色id',
  `role_name` varchar(255) DEFAULT NULL COMMENT '角色名称',
  `company_id` int(10) unsigned NOT NULL COMMENT '用户所属公司id',
  `address` varchar(255) DEFAULT NULL COMMENT '用户住址',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_admins_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_admins`
--

LOCK TABLES `user_admins` WRITE;
/*!40000 ALTER TABLE `user_admins` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_lifts`
--

DROP TABLE IF EXISTS `user_lifts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_lifts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户',
  `lift_id` int(10) unsigned NOT NULL COMMENT '电梯',
  `category_id` int(10) NOT NULL COMMENT '用户电梯关系类型',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_lifts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户电梯关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_lifts`
--

LOCK TABLES `user_lifts` WRITE;
/*!40000 ALTER TABLE `user_lifts` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_lifts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `authority_menu`
--

/*!50001 DROP VIEW IF EXISTS `authority_menu`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`nick_name` AS `nick_name`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-05 12:07:10