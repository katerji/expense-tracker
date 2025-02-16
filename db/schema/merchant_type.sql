CREATE TABLE `merchant_type`
(
    `id`   int unsigned NOT NULL AUTO_INCREMENT,
    `type` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;