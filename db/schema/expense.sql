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