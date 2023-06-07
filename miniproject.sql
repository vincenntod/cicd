-- CREATE DATABASE  IF NOT EXISTS `miniproject` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
-- USE `miniproject`;
-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: miniproject
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `role_id` int DEFAULT '2',
  `verified` enum('true','false') DEFAULT NULL,
  `active` enum('true','false') DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `actors_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'vincennte','$2a$10$tmgUTu32nG.QmlH2JEik2uQzefCyuvlG61g./s1AUHHQHdMPSLL3q',1,'true','true','2023-06-01 17:00:00','2023-06-01 17:00:00'),(5,'zefanya','$2a$10$lgqF8iPc9/L8nwUGMdSuaOLNCAlE.EhiIsXggtYRfLLOa5VOx6PRC',2,'true','true','2023-06-03 02:44:35','2023-06-03 02:44:35'),(6,'maudy','$2a$10$YTgEZCnsbiuyCHk2X5eHwORrO5VhJSwp6Ue/xqWR.FBPtF7tTR7Xy',2,'false','false','2023-06-03 03:46:09','2023-06-03 03:46:09'),(7,'ayunda','$2a$10$EsYBM2kEtHGt5G/X1MCLvOt/20Ci.xXgYkIpR4UBueeiogTuMcLAi',2,NULL,NULL,'2023-06-03 04:02:53','2023-06-03 04:02:53'),(8,'vincenntez','$2a$10$V.YeKuincepNTZ3fdGVDMu2NpMx6VwynyGO/taDMyP2BaV3rqee..',2,NULL,NULL,'2023-06-03 04:53:15','2023-06-03 04:53:15');
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `avatar` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (1,'Vincen','Maxwell','@gmail.com','cek12',NULL,NULL),(3,'Maudy','Ayunda','@gmail.com',' 12','2023-02-05 17:00:00','2023-02-05 17:00:00'),(7,'Michael','Lawson','michael.lawson@reqres.in','https://reqres.in/img/faces/7-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46'),(8,'Lindsay','Ferguson','lindsay.ferguson@reqres.in','https://reqres.in/img/faces/8-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46'),(9,'Tobias','Funke','tobias.funke@reqres.in','https://reqres.in/img/faces/9-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46'),(10,'Byron','Fields','byron.fields@reqres.in','https://reqres.in/img/faces/10-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46'),(11,'George','Edwards','george.edwards@reqres.in','https://reqres.in/img/faces/11-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46'),(12,'Rachel','Howell','rachel.howell@reqres.in','https://reqres.in/img/faces/12-image.jpg','2023-06-03 01:07:46','2023-06-03 01:07:46');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `register_approval`
--

DROP TABLE IF EXISTS `register_approval`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `register_approval` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `admin_id` bigint DEFAULT '0',
  `super_admin_id` bigint DEFAULT '0',
  `status` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `register_approval`
--

LOCK TABLES `register_approval` WRITE;
/*!40000 ALTER TABLE `register_approval` DISABLE KEYS */;
/*!40000 ALTER TABLE `register_approval` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int NOT NULL,
  `role_name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'superadmin'),(2,'admin');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-03 15:32:06
