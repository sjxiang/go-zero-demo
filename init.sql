CREATE DATABASE IF NOT EXISTS `zero-demo` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `zero-demo`;

DROP TABLE IF EXISTS `user`;


-- 用户表结构设计
CREATE TABLE `user` (
    `id`          bigint(20)  NOT NULL AUTO_INCREMENT,
    `nickname`    varchar(64) NOT NULL,
    `mobile`      char(11)    NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


describe user;


DROP TABLE IF EXISTS `user_data`;


-- 用户数据表结构设计
CREATE TABLE `user_data` (
    `id`          bigint(20)  NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20)  NOT NULL,
    `data`        varchar(64) NOT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

