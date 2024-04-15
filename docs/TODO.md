# TODO

## Virifier Transaction Config

### 1. Table Structure

```shell
-- Table structure for v_transaction_config_int
-- ----------------------------
DROP TABLE IF EXISTS `t_app_config_int`;
CREATE TABLE `t_app_config_int` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `k` varchar(64) NOT NULL COMMENT '配置键名',
  `v` bigint(20) NOT NULL COMMENT '配置键值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `k` (`k`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
```

### 2. Table Data
> Example

```shell
-- ----------------------------
-- Records of t_app_config_int
-- ----------------------------
BEGIN;
INSERT INTO `t_app_config_int` VALUES (2, 'opbnb_gas_limit', 90000);
INSERT INTO `t_app_config_int` VALUES (3, 'opbnb_gas_price_max', 100000000);
INSERT INTO `t_app_config_int` VALUES (4, 'opbnb_confirmations_max', 12);
INSERT INTO `t_app_config_int` VALUES (6, 'send_enable', 1);
COMMIT;
```

## Transaction Information 
