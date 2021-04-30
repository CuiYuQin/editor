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

insert into user (user_name,tree,delete_tree,create_time) values ('starSea','tree','delete_tree',now());
insert into user (user_name,tree,delete_tree,create_time) values ('我愿编写一颗恒星','tree','delete_tree',now());
insert into user (user_name,tree,delete_tree,create_time) values ('遥遥微风与我同行','tree','delete_tree',now());

select * from user;