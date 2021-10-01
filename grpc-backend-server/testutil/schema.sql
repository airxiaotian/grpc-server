-- MySQL dump 10.13  Distrib 8.0.21, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: harp
-- ------------------------------------------------------
-- Server version	5.7.31
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!50503 SET NAMES utf8mb4 */
;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;
--
-- Table structure for table `acceptance_details`
--
DROP TABLE IF EXISTS `acceptance_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `acceptance_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `order_details_id` int(11) NOT NULL,
  `scheduled_acceptance_date` datetime DEFAULT NULL,
  `actual_acceptance_date` date DEFAULT NULL,
  `acceptance_quantity` decimal(13,2) DEFAULT NULL,
  `return_quantity` decimal(13,2) DEFAULT NULL,
  `acceptance_amount` int(11) DEFAULT NULL,
  `return_price` int(11) DEFAULT NULL,
  `approval_date` datetime DEFAULT NULL,
  `approval_by` int(11) DEFAULT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `suppliers_id` int(11) DEFAULT NULL,
  `scheduled_acceptance_yymm` varchar(6) DEFAULT NULL,
  `actual_acceptance_yymm` varchar(6) DEFAULT NULL,
  `module_unregister` int(11) DEFAULT NULL,
  `projects_id` varchar(128) DEFAULT NULL,
  `project_cost_id` varchar(128) DEFAULT NULL,
  `scheduled_acceptance_quantity` DECIMAL(13,2) DEFAULT NULL,
  `scheduled_acceptance_amount` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `contract_manager_details`
--
DROP TABLE IF EXISTS `contract_manager_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `contract_manager_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `manager_purchase_id` int(11) NOT NULL,
  `contract_manager_type` varchar(1) NOT NULL,
  `deputy_type` varchar(1) NOT NULL,
  `g_inout_type` varchar(1) NOT NULL,
  `staffs_id` int(11) DEFAULT NULL,
  `g_exuser_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `harp_sys_sequences`
--
DROP TABLE IF EXISTS `harp_sys_sequences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `harp_sys_sequences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sequence_name` varchar(255) NOT NULL,
  `key` varchar(255) DEFAULT NULL,
  `value` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `item_units`
--
DROP TABLE IF EXISTS `item_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `item_units` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_value` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `order_details`
--
DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `order_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `order_details_no` int(11) DEFAULT NULL,
  `product_name` varchar(20) NOT NULL,
  `specifications` varchar(300) DEFAULT NULL,
  `order_quantity` decimal(13,2) NOT NULL,
  `cancel_quantity` decimal(13,2) DEFAULT NULL,
  `order_unit_price` int(11) NOT NULL,
  `acceptance_scheduled_date` date DEFAULT NULL,
  `configuration_management_target_flag` varchar(1) NOT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `quotations_id` int(11) DEFAULT NULL,
  `quotation_details_id` int(11) DEFAULT NULL,
  `order_unit_classification` int(11) DEFAULT NULL,
  `projects_id` VARCHAR(7) DEFAULT NULL,
  `cost_types_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `order_items`
--
DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `order_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `order_details_id` int(11) NOT NULL,
  `product_name` varchar(30) NOT NULL,
  `order_quantity` int(11) NOT NULL,
  `order_price` int(11) NOT NULL,
  `order_details_no` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `order_states`
--
DROP TABLE IF EXISTS `order_states`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `order_states` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_value` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `order_types`
--
DROP TABLE IF EXISTS `order_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `order_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_value` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `orders`
--
DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(8) DEFAULT NULL,
  `suppliers_id` int(11) NOT NULL,
  `company_group_type` varchar(3) DEFAULT NULL,
  `subject` varchar(50) NOT NULL,
  `request_organization_id` VARCHAR(128) NOT NULL,
  `request_date` date NOT NULL,
  `request_by` int(11) NOT NULL,
  `approval_file` varchar(200) DEFAULT NULL,
  `derivation_source_order_id` int(11) DEFAULT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `superior_approval_date` datetime DEFAULT NULL,
  `purchasing_dept_approval_date` datetime DEFAULT NULL,
  `order_issue_date` datetime DEFAULT NULL,
  `final_acceptance_date` datetime DEFAULT NULL,
  `acceptance_completed_date` datetime DEFAULT NULL,
  `cancel_date` datetime DEFAULT NULL,
  `order_case_cd` int(11) NOT NULL,
  `order_status` int(11) NOT NULL,
  `jira_no` varchar(20) DEFAULT NULL,
  `quotations_id` int(11) DEFAULT NULL,
  `order_approval_staffs_id` varchar(128) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `projects_id` varchar(128) DEFAULT NULL,
  `project_cost_id` varchar(128) DEFAULT NULL,
  `cost_types_id` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `project_cost_details`
--
DROP TABLE IF EXISTS `project_cost_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `project_cost_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `order_details_id` int(11) NOT NULL,
  `acceptance_details_id` int(11) NOT NULL,
  `stocking_cost` int(11) DEFAULT NULL,
  `cost_entry_yymm` varchar(6) NOT NULL,
  `projects_id` VARCHAR(7) NOT NULL,
  `cost_types_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `quotation_details`
--
DROP TABLE IF EXISTS `quotation_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `quotation_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quotations_id` int(11) NOT NULL,
  `version_number` int(11) DEFAULT NULL,
  `order_details_id` int(11) DEFAULT NULL,
  `product_name` varchar(20) NOT NULL,
  `specifications` varchar(300) DEFAULT NULL,
  `order_quantity` decimal(13,2) NOT NULL,
  `order_price` int(11) NOT NULL,
  `unit_classification` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `quotation_histories`
--
DROP TABLE IF EXISTS `quotation_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `quotation_histories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quotations_id` int(11) NOT NULL,
  `quotation_data` json NOT NULL,
  `create_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `quotation_items`
--
DROP TABLE IF EXISTS `quotation_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `quotation_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `version_number` int(11) NOT NULL,
  `quotation_details_id` int(11) NOT NULL,
  `product_name` varchar(30) NOT NULL,
  `order_quantity` int(11) NOT NULL,
  `order_price` int(11) NOT NULL,
  `order_details_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `quotations`
--
DROP TABLE IF EXISTS `quotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `quotations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quotation_no` varchar(6) DEFAULT NULL,
  `version_number` int(11) DEFAULT NULL,
  `suppliers_id` int(11) NOT NULL,
  `company_group_classification` varchar(3) NOT NULL,
  `subject` varchar(50) NOT NULL,
  `supplier_quotation_no` varchar(12) DEFAULT NULL,
  `request_organization_id` int(11) NOT NULL,
  `request_date` date NOT NULL,
  `request_by` int(11) NOT NULL,
  `remarks` varchar(300) DEFAULT NULL,
  `quotation_effective_date` date DEFAULT NULL,
  `quotation_invalid_date` date DEFAULT NULL,
  `jira_no` varchar(20) DEFAULT NULL,
  `order_classification` int(11) DEFAULT NULL,
  `orders_id` int(11) DEFAULT NULL,
  `quotation_status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Table structure for table `order_histories`
--
DROP TABLE IF EXISTS `order_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `order_histories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `orders_id` int(11) NOT NULL,
  `order_data` json NOT NULL,
  `create_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */
;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;
-- Dump completed on 2020-09-24 11:07:50