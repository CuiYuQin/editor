-- 创建数据库
DROP DATABASE `editor`;
CREATE DATABASE `editor`;

-- 选择数据库
use `editor`;


CREATE TABLE IF NOT EXISTS `user`(
   `user_id` INT NOT NULL AUTO_INCREMENT,
   `user_name` VARCHAR(100) NOT NULL,
   `tree` MEDIUMTEXT,
   `delete_tree` MEDIUMTEXT,
   `create_time` DATETIME NOT NULL,
   PRIMARY KEY ( `user_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into user (user_name,tree,delete_tree,create_time) values ('starSea','[{\"children\":[{\"children\":[],\"id\":1619886619095,\"label\":\"Untitled\",\"type\":false,\"url\":\"1\"}],\"id\":1,\"label\":\"我的文档\",\"type\":true,\"url\":\"0\"}]','[{\"children\":[{\"children\":[],\"id\":1619886619095,\"label\":\"Untitled\",\"type\":false,\"url\":\"1\"}],\"id\":1,\"label\":\"我的文档\",\"type\":true,\"url\":\"0\"}]',now());
-- insert into user (user_name,tree,delete_tree,create_time) values ('我愿编写一颗恒星','tree','delete_tree',now());
-- insert into user (user_name,tree,delete_tree,create_time) values ('遥遥微风与我同行','tree','delete_tree',now());

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

insert into ebook (title,author,createTime,content,contentHtml) values ('Untitled','starSea',now(),'## 请仔细阅读帮助文档','');

select * from ebook;
