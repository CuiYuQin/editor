-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: localhost    Database: ebook
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `bbs_comment`
--

DROP DATABASE `ebook`;

CREATE DATABASE IF NOT EXISTS ebook default charset utf8 COLLATE utf8_general_ci;
 
use ebook;
 


DROP TABLE IF EXISTS `bbs_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs_comment` (
  `commentId` int NOT NULL AUTO_INCREMENT,
  `bbsId` int NOT NULL,
  `author` varchar(45) NOT NULL,
  `avatar` varchar(100) NOT NULL,
  `content` varchar(200) NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`commentId`),
  UNIQUE KEY `commentId_UNIQUE` (`commentId`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `bbs_topic`
--

DROP TABLE IF EXISTS `bbs_topic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bbs_topic` (
  `idbbs` int NOT NULL AUTO_INCREMENT,
  `author` varchar(45) NOT NULL,
  `topicName` varchar(100) NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`idbbs`),
  UNIQUE KEY `idbbs_UNIQUE` (`idbbs`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `class_notification`
--

DROP TABLE IF EXISTS `class_notification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `class_notification` (
  `notificationId` int NOT NULL AUTO_INCREMENT,
  `className` varchar(45) NOT NULL,
  `title` varchar(45) NOT NULL,
  `content` varchar(200) NOT NULL,
  `noticeTime` datetime NOT NULL,
  PRIMARY KEY (`notificationId`),
  UNIQUE KEY `notificationId_UNIQUE` (`notificationId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='发布通知';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `class_student`
--

DROP TABLE IF EXISTS `class_student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `class_student` (
  `studentName` varchar(20) NOT NULL,
  `className` varchar(20) NOT NULL,
  `status` int NOT NULL COMMENT '0 (未加入)\\n1 (已加入）',
  PRIMARY KEY (`studentName`,`className`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `class_teacher`
--

DROP TABLE IF EXISTS `class_teacher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `class_teacher` (
  `className` varchar(45) NOT NULL,
  `teacherName` varchar(45) NOT NULL,
  `description` varchar(200) NOT NULL,
  PRIMARY KEY (`className`),
  UNIQUE KEY `className_UNIQUE` (`className`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='教师班级表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ebook`
--

-- DROP TABLE IF EXISTS `ebook`;
-- /*!40101 SET @saved_cs_client     = @@character_set_client */;
-- /*!50503 SET character_set_client = utf8mb4 */;
-- CREATE TABLE `ebook` (
--   `ebookId` int NOT NULL AUTO_INCREMENT,
--   `title` varchar(40) NOT NULL,
--   `author` varchar(20) NOT NULL,
--   `createTime` datetime NOT NULL,
--   `contentHtml` varchar(20000) NOT NULL,
--   PRIMARY KEY (`ebookId`),
--   UNIQUE KEY `ebookId_UNIQUE` (`ebookId`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
-- /*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ebook_note`
--

DROP TABLE IF EXISTS `ebook_note`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ebook_note` (
  `noteId` int unsigned NOT NULL AUTO_INCREMENT,
  `ebookId` int NOT NULL,
  `username` varchar(45) NOT NULL,
  `contentHtml` varchar(20000) NOT NULL,
  `createTime` datetime NOT NULL,
  `updateTime` datetime NOT NULL,
  PRIMARY KEY (`noteId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `student_paper`
--

DROP TABLE IF EXISTS `student_paper`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `student_paper` (
  `stuPaperId` int NOT NULL AUTO_INCREMENT,
  `studentName` varchar(45) NOT NULL,
  `paperId` int NOT NULL,
  `askContent` varchar(400) DEFAULT NULL,
  `selectContent` varchar(400) DEFAULT NULL,
  `createTime` datetime NOT NULL,
  `score` float DEFAULT NULL,
  `remark` varchar(45) DEFAULT NULL,
  `status` varchar(45) NOT NULL COMMENT '0 ： 创建试卷\n1： 完成试卷\n2：批改试卷',
  PRIMARY KEY (`studentName`,`paperId`),
  UNIQUE KEY `stuPaperId_UNIQUE` (`stuPaperId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teacher_paper`
--

DROP TABLE IF EXISTS `teacher_paper`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teacher_paper` (
  `paperId` int NOT NULL AUTO_INCREMENT,
  `paperName` varchar(45) NOT NULL,
  `paperDescription` varchar(200) NOT NULL,
  `selectContent` varchar(400) DEFAULT NULL,
  `askContent` varchar(400) DEFAULT NULL,
  `createTime` datetime NOT NULL,
  `creator` varchar(45) NOT NULL,
  PRIMARY KEY (`paperId`),
  UNIQUE KEY `paperId_UNIQUE` (`paperId`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='试卷';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
-- CREATE TABLE `user` (
--   `id` int NOT NULL AUTO_INCREMENT,
--   `username` varchar(45) NOT NULL,
--   `password` varchar(64) NOT NULL,
--   `types` int NOT NULL,
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `username_UNIQUE` (`username`),
--   UNIQUE KEY `id_UNIQUE` (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 COMMENT='user';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_token`
--

DROP TABLE IF EXISTS `user_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_token` (
  `userId` int NOT NULL,
  `token` varchar(64) NOT NULL,
  PRIMARY KEY (`userId`),
  UNIQUE KEY `userId_UNIQUE` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-05-02 23:02:38


CREATE TABLE IF NOT EXISTS `user`(
   `id` INT NOT NULL AUTO_INCREMENT,
   `username` VARCHAR(100) NOT NULL,
   `password` varchar(64) NOT NULL,
   `types` int NOT NULL,
   `tree` MEDIUMTEXT,
   `delete_tree` MEDIUMTEXT,
   `create_time` DATETIME ,
  --  PRIMARY KEY ( `id` )
   PRIMARY KEY (`id`),
   UNIQUE KEY `username_UNIQUE` (`username`),
   UNIQUE KEY `id_UNIQUE` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into user (username,password,types,tree,delete_tree,create_time) values ('starSea','e10adc3949ba59abbe56e057f20f883e',1,'[{\"children\":[{\"children\":[],\"id\":1619886619095,\"label\":\"Untitled\",\"type\":false,\"url\":\"1\"}],\"id\":1,\"label\":\"我的文档\",\"type\":true,\"url\":\"0\"}]','[]',now());
insert into user (username,password,types,tree,delete_tree,create_time) values ('我愿编写一颗恒星','e10adc3949ba59abbe56e057f20f883e',2,'[{\"children\":[{\"children\":[],\"id\":1619886619095,\"label\":\"Untitled\",\"type\":false,\"url\":\"2\"}],\"id\":1,\"label\":\"我的文档\",\"type\":true,\"url\":\"0\"}]','[]',now());

-- insert into user (username,tree,delete_tree,create_time) values ('我愿编写一颗恒星','tree','delete_tree',now());
-- insert into user (username,tree,delete_tree,create_time) values ('遥遥微风与我同行','tree','delete_tree',now());

select * from user;


CREATE TABLE IF NOT EXISTS `ebook`(
   `ebookId` INT NOT NULL AUTO_INCREMENT,
   `title` VARCHAR(100) NOT NULL,
   `author` VARCHAR(100) NOT NULL,
   `createTime` DATETIME NOT NULL,
   `content` MEDIUMTEXT,
   `contentHtml` MEDIUMTEXT,
   PRIMARY KEY ( `ebookId` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into ebook (title,author,createTime,content,contentHtml) values ('test1','starSea',now(),'## 请仔细阅读帮助文档','');
insert into ebook (title,author,createTime,content,contentHtml) values ('test2','我愿编写一颗恒星',now(),'## 请仔细阅读帮助文档','');

select * from ebook;