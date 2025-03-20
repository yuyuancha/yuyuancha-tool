DROP TABLE IF EXISTS `series`;
CREATE TABLE `series` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL COMMENT '系列名稱',
    `code` varchar(100) NOT NULL COMMENT '系列代碼',
    `base_count` int(11) NOT NULL COMMENT '基本卡數',
    `extra_count` int(11) NOT NULL COMMENT '附加卡數',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡包系列';

BEGIN;
INSERT INTO `series` VALUES (1, '時空激鬥', 'A2',  155, 52, default, default);
INSERT INTO `series` VALUES (2, '超克之光', 'A2a',  75, 21, default, default);
COMMIT;

DROP TABLE IF EXISTS `packages`;
CREATE TABLE `packages` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `series_id` int(11) NOT NULL COMMENT '系列ID',
    `name` varchar(100) NOT NULL COMMENT '卡包名稱',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`series_id`) REFERENCES `series` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡包';

BEGIN;
INSERT INTO `packages` VALUES (1, 1, '帝牙盧卡擴充包', default, default);
INSERT INTO `packages` VALUES (2, 1, '帕路奇亞擴充包', default, default);
INSERT INTO `packages` VALUES (3, 2, '超克之光擴充包', default, default);
COMMIT;

DROP TABLE IF EXISTS `cards`;
CREATE TABLE `cards` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `series_id` int(11) NOT NULL COMMENT '系列ID',
    `number` varchar(100) NOT NULL COMMENT '卡片編號',
    `name` varchar(100) NOT NULL COMMENT '寶可夢名稱',
    `attribute` varchar(100) NOT NULL COMMENT '屬性',
    `rarity` varchar(100) NOT NULL COMMENT '稀有度',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`series_id`) REFERENCES `series` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡片';

BEGIN;
COMMIT;

DROP TABLE IF EXISTS `package_card_relations`;
CREATE TABLE `package_card_relations` (
    `package_id` int(11) NOT NULL,
    `card_id` int(11) NOT NULL,
    PRIMARY KEY (`package_id`, `card_id`),
       FOREIGN KEY (`package_id`) REFERENCES `packages` (`id`)
           ON DELETE CASCADE
           ON UPDATE CASCADE,
       FOREIGN KEY (`card_id`) REFERENCES `cards` (`id`)
           ON DELETE CASCADE
           ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卡片';

BEGIN;
COMMIT;
