DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `chat_id` int(11) NOT NULL COMMENT 'Chat ID',
    `first_name` varchar(100) NOT NULL COMMENT '暱稱',
    `username` varchar(100) NOT NULL COMMENT '帳號',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '狀態(0: 停用, 1: 啟用)',
    `role` tinyint(1) NOT NULL DEFAULT '1' COMMENT '角色(1: 一般會員, 2: 管理員)',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `chat_id` (`chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='會員';

BEGIN;
COMMIT;

DROP TABLE IF EXISTS `daily_trade_stocks`;
CREATE TABLE `daily_trade_stocks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `code` varchar(50) NOT NULL COMMENT '代號',
    `name` varchar(100) NOT NULL COMMENT '名稱',
    `trade_volume` int NOT NULL COMMENT '成交量',
    `trade_value` bigint NOT NULL COMMENT '成交值',
    `opening_price` float NOT NULL COMMENT '開盤價',
    `highest_price` float NOT NULL COMMENT '最高價',
    `lowest_price` float NOT NULL COMMENT '最低價',
    `closing_price` float NOT NULL COMMENT '收盤價',
    `change` float NOT NULL COMMENT '漲跌',
    `transaction` int NOT NULL COMMENT '筆數',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='每日交易資訊';

BEGIN;
COMMIT;
