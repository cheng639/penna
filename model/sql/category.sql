CREATE TABLE `category` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
    `status` tinyint(4) NOT NULL DEFAULT '1',
    `sort` int(11) NOT NULL DEFAULT '0',
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `category` VALUES (1, 'foo', 0, 1, 1, '2026-05-09 16:47:57', '2026-05-09 16:47:57', NULL);
INSERT INTO `category` VALUES (2, 'bar', 0, 1, 1, '2026-05-15 17:39:41', '2026-05-15 17:45:40', NULL);
INSERT INTO `category` VALUES (3, 'level2', 2, 1, 1, '2026-05-17 14:02:30', '2026-05-17 14:11:19', NULL);