CREATE TABLE `user_account`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int unsigned NOT NULL,
    `account_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_id` (`user_id`,`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;