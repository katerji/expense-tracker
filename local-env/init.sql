DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;
CREATE TABLE `user`
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `first_name` varchar(255) NOT NULL,
    `last_name`  varchar(255) NOT NULL,
    `email`      varchar(255) NOT NULL,
    `password`   varchar(255) NOT NULL,
    `created_on` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `account`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT,
    `name`            varchar(255) NOT NULL,
    `primary_user_id` int unsigned NOT NULL,
    `created_on`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `merchant_type`
(
    `id`   int unsigned NOT NULL AUTO_INCREMENT,
    `type` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `merchant`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL,
    `type_id`    int unsigned NOT NULL,
    `created_on` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `expense`
(
    `id`               int unsigned NOT NULL AUTO_INCREMENT,
    `amount`           float        NOT NULL,
    `currency`         varchar(255) NOT NULL,
    `time_of_purchase` timestamp NULL DEFAULT NULL,
    `description`      text,
    `merchant_id`      int unsigned NOT NULL,
    `account_id`       int unsigned NOT NULL,
    `created_on`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY                `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;