DROP TABLE IF EXISTS `gov_travel_card_shops`;
CREATE TABLE `gov_travel_card_shops` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL COMMENT '店名',
    `address` varchar(255) NOT NULL COMMENT '地址',
    `latitude` DECIMAL(10, 8) NOT NULL COMMENT '緯度',
    `longitude` DECIMAL(11, 8) NOT NULL COMMENT '經度',
    `phone_number` varchar(20) NOT NULL COMMENT '電話',
    `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '類別ID',
    `note` varchar(255) NOT NULL DEFAULT '' COMMENT '備註',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `name` (`name`),
    KEY `category_id` (`category_id`),
    KEY `latitude` (`latitude`),
    KEY `longitude` (`longitude`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='政府國旅卡店家';

BEGIN;
INSERT INTO `gov_travel_card_shops` VALUES (default, '００Ｂ３寶島眼鏡酒泉街分公司', '台北市大同區酒泉街１３５號１樓', 0, 0, '02-25861956', 15, '', default, default);
INSERT INTO `gov_travel_card_shops` VALUES (default, 'ａｄｉｄａｓ－南西門市', '臺北市大同區光能里南京西路２７號１樓至５樓', 0, 0, '02-25595518', 8, '', default, default);
COMMIT;

DROP TABLE IF EXISTS `gov_travel_card_categories`;
CREATE TABLE `gov_travel_card_categories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL COMMENT '類別名稱',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='政府國旅卡類別';

BEGIN;
INSERT INTO `gov_travel_card_categories` VALUES (default, '旅行業', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '旅宿業', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '交通運輸業', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '觀光遊樂業', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-餐飲', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-農特產及手工藝品', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-加油站', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-體育用品', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-其他觀光服務', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-服飾', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-皮鞋皮件', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-美容護膚', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-商圈及其他', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別-藝文圖書', default, default);
INSERT INTO `gov_travel_card_categories` VALUES (default, '其他業別', default, default);
COMMIT;
