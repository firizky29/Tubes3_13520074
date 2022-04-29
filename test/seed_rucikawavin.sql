-- MariaDB dump 10.19  Distrib 10.4.22-MariaDB, for Win64 (AMD64)
--
-- Host: rucikawavin.mysql.database.azure.com    Database: rucikawavin
-- ------------------------------------------------------
-- Server version	5.6.47.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `disease_predictions`
--

DROP TABLE IF EXISTS `disease_predictions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `disease_predictions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` date NOT NULL,
  `user_name` varchar(100) NOT NULL,
  `dna` text NOT NULL,
  `disease_prediction` varchar(255) NOT NULL,
  `similarity_level` decimal(5,2) NOT NULL,
  `status` varchar(10) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_disease_predictions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `disease_predictions`
--

LOCK TABLES `disease_predictions` WRITE;
/*!40000 ALTER TABLE `disease_predictions` DISABLE KEYS */;
INSERT INTO `disease_predictions` VALUES (1,'2022-04-28','Firizky','AAAAAAAAAAAAAAAAAAAATCG','Virus',100.00,'true','2022-04-28 12:29:34.516',NULL),(2,'2022-04-28','Firizky','AAAAAAAAAAAAAAAAAAAATCG','Virus',100.00,'true','2022-04-28 12:29:40.253',NULL),(3,'2022-04-28','Hambin','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:30:18.561',NULL),(4,'2022-04-28','Hambin','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:30:24.767',NULL),(5,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:31:00.415',NULL),(6,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:31:07.085',NULL),(7,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:31:12.087',NULL),(8,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:31:46.245',NULL),(9,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAATCA','Virus',75.00,'false','2022-04-28 12:31:59.392',NULL),(10,'2022-04-28','Eiffel','AAAAAAAAAAAAAAAAAAAAAA','Virus',25.00,'false','2022-04-28 12:32:42.493',NULL),(11,'2022-04-28','Hambin','AAAAAAAAAAAAAAAAAAAAAA','Virus',25.00,'false','2022-04-28 12:36:15.811',NULL),(12,'2022-04-28','Hambin','AAAAAAAAAAAAAAAAAAAAAA','Virus',25.00,'false','2022-04-28 12:36:43.661',NULL),(13,'2022-04-28','Anak 1','AAAAAAAAAAAAAAAATAAAAA','Virus',50.00,'false','2022-04-28 12:41:14.683',NULL),(14,'2022-04-28','Anak 2','AAAAAAAAAAAAAAAATAAAAA','Virus',50.00,'false','2022-04-28 12:41:57.848',NULL),(15,'2022-04-28','Anak 3','TCGAAAAAAAAAAAAAAAAAAAAA','Virus',75.00,'false','2022-04-28 12:46:53.495',NULL),(16,'2022-04-28','Anak 4','TCGAAAAAAAAAAAAAAAAAAAAA','Virus',75.00,'false','2022-04-28 12:47:16.694',NULL),(17,'2022-04-28','Firizky','AATAA','Covid',100.00,'true','2022-04-28 13:12:44.367',NULL),(18,'2022-04-28','Cobain','AATAA','Covid',100.00,'true','2022-04-28 13:22:25.069',NULL),(19,'2022-04-29','Ilham','AAA','Covid',100.00,'true','2022-04-29 22:59:25.600',NULL),(20,'2022-04-29','Eiffel','TTTTTTTTTTTTTTTTTTTTTTTT','Covid',0.00,'false','2022-04-29 23:00:53.246',NULL),(21,'2022-04-29','Eiffel','TTTTTTTTTTTTTTTTTTTTTTTT','Covid',0.00,'false','2022-04-29 23:01:15.465',NULL),(22,'2022-04-29','Eiffel 2','TTTTTTTTTTTTTTTTTTTTTTTT','Covid',0.00,'false','2022-04-29 23:01:32.648',NULL),(23,'2022-04-29','Ilham Bintang','ATCGACGAGACTACGTACACACTTTACGAGCG','Covid',100.00,'true','2022-04-29 23:19:02.201',NULL),(24,'2022-04-29','Ilham Bintang','ATCGACGTTTTAAGGATACGTACACACTTTACGAGCG','Covid',100.00,'true','2022-04-29 23:22:50.347',NULL),(25,'2022-04-29','Ilham Bintang','ATCGACGTTTTAAGGATACGTACACACTTTACGAGCG','Covid',100.00,'true','2022-04-29 23:24:03.690',NULL),(26,'2022-04-29','Ilham Bintang','TATATAGGGTAGCCACAG','Covid',100.00,'true','2022-04-29 23:25:15.013',NULL),(27,'2022-04-29','Ilham Bintang','TATATAGGGTAGCCACAG','Covid',100.00,'true','2022-04-29 23:26:08.544',NULL),(28,'2022-04-29','Eiffel','TATATAGGGTAGCCACAG','Covid',100.00,'true','2022-04-29 23:27:07.946',NULL),(29,'2022-04-29','Ilham Bintang','TATATAGGGTAGCCACAG','Covid',100.00,'true','2022-04-29 23:27:35.445',NULL),(30,'2022-04-29','Hambin','TATATAGGGTAGCCACAG','Covid 2',72.00,'false','2022-04-29 23:28:08.598',NULL),(31,'2022-04-29','Hambin','TATATAGGGTAGCCACAG','Covid',100.00,'true','2022-04-29 23:29:57.993',NULL),(32,'2022-04-29','Hambin','TATATAGGGTAGCCACAG','Covid 2',72.00,'false','2022-04-29 23:33:35.033',NULL);
/*!40000 ALTER TABLE `disease_predictions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `diseases`
--

DROP TABLE IF EXISTS `diseases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `diseases` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `dna` text NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_diseases_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `diseases`
--

LOCK TABLES `diseases` WRITE;
/*!40000 ALTER TABLE `diseases` DISABLE KEYS */;
INSERT INTO `diseases` VALUES (2,'Virus','ATCG',NULL,NULL,NULL),(3,'Covid','AAA','2022-04-28 13:12:25.818','2022-04-28 13:12:25.818',NULL),(4,'Virus 1','ACACAG','2022-04-28 15:02:05.245','2022-04-28 15:02:05.245',NULL),(5,'Virus 2','GCACAG','2022-04-28 15:02:33.245','2022-04-28 15:02:33.245',NULL),(6,'Virus 3','AATAA','2022-04-29 22:58:34.117','2022-04-29 22:58:34.117',NULL),(7,'Virus Covid','AGACTACGTAC','2022-04-29 23:18:00.698','2022-04-29 23:18:00.698',NULL),(8,'Covid 2','AGACTACGTAC','2022-04-29 23:27:51.264','2022-04-29 23:27:51.264',NULL);
/*!40000 ALTER TABLE `diseases` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-29 23:47:10
