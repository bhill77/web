CREATE TABLE if not exists `articles` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `author` varchar(255) DEFAULT NULL,
    `title` text DEFAULT NULL,
    `body` text DEFAULT NULL,
    `created` timestamp NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`),
    FULLTEXT KEY `author` (`author`),
    FULLTEXT KEY `title` (`title`,`body`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
