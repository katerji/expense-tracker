CREATE TABLE `account`
(
    `id`              int unsigned NOT NULL AUTO_INCREMENT,
    `name`            varchar(255) NOT NULL,
    `primary_user_id` int unsigned NOT NULL,
    `created_on`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on`      datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;