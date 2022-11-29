CREATE DATABASE IF NOT EXISTS `zero-demo` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `zero-demo`;

DROP TABLE IF EXISTS `user`;


-- 用户表结构设计
CREATE TABLE `user` (
    `id`          bigint(20)  NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20)  NOT NULL,
    `username`    varchar(64) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


describe user;