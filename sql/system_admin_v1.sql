/*
 Navicat Premium Data Transfer

 Source Server         : aliyun
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 47.106.206.78:3306
 Source Schema         : system_admin

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 06/03/2023 11:08:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access_register
-- ----------------------------
DROP TABLE IF EXISTS `access_register`;
CREATE TABLE `access_register` (
                                   `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                                   `name` varchar(50) DEFAULT NULL COMMENT '姓名',
                                   `phone` varchar(50) DEFAULT NULL COMMENT '手机号码',
                                   `type` int(1) DEFAULT NULL COMMENT '出入类型（1：出校 0：入校）',
                                   `card` varchar(80) DEFAULT NULL COMMENT '身份证号',
                                   `remark` varchar(255) DEFAULT NULL COMMENT '备注',
                                   `dept` varchar(50) DEFAULT NULL COMMENT '部门',
                                   `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
                                   `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                   `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                                   `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='出入登记表';

-- ----------------------------
-- Records of access_register
-- ----------------------------
BEGIN;
INSERT INTO `access_register` (`id`, `name`, `phone`, `type`, `card`, `remark`, `dept`, `create_by`, `create_time`, `update_time`, `is_delete`) VALUES (1, '范晨', '17730312784', 1, '342623199906214418', '正常返校', '软件186', 'admin', '2022-12-17 14:39:37', '2023-02-23 06:31:39', 0);
INSERT INTO `access_register` (`id`, `name`, `phone`, `type`, `card`, `remark`, `dept`, `create_by`, `create_time`, `update_time`, `is_delete`) VALUES (2, '范晨', '17730312784', 0, '342623199906214418', '正常返校', '软件186', 'admin', '2022-12-17 18:04:10', '2021-12-17 18:04:10', 0);
INSERT INTO `access_register` (`id`, `name`, `phone`, `type`, `card`, `remark`, `dept`, `create_by`, `create_time`, `update_time`, `is_delete`) VALUES (3, '杜东亮', '19851920126', 1, '34262319990901159x', '正常请假出校', '软件186', 'admin', '2022-12-17 18:20:27', '2021-12-17 18:20:27', 0);
INSERT INTO `access_register` (`id`, `name`, `phone`, `type`, `card`, `remark`, `dept`, `create_by`, `create_time`, `update_time`, `is_delete`) VALUES (5, '李文', '13226948870', 1, '342623199906214412', '正常返校', '软件186', '', '2023-02-23 06:41:53', '2023-02-23 06:41:53', 0);
INSERT INTO `access_register` (`id`, `name`, `phone`, `type`, `card`, `remark`, `dept`, `create_by`, `create_time`, `update_time`, `is_delete`) VALUES (6, '李文节', '13227878838', 1, '342623199906214413', '无', '软件186', 'admin', '2023-02-28 14:35:52', '2023-02-28 14:35:52', 0);
COMMIT;

-- ----------------------------
-- Table structure for access_return
-- ----------------------------
DROP TABLE IF EXISTS `access_return`;
CREATE TABLE `access_return` (
                                 `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                                 `name` varchar(50) DEFAULT NULL COMMENT '姓名',
                                 `phone` varchar(60) DEFAULT NULL COMMENT '手机号码',
                                 `card` varchar(80) DEFAULT NULL COMMENT '身份证号码',
                                 `remark` varchar(255) DEFAULT NULL COMMENT '备注',
                                 `dept` varchar(50) DEFAULT NULL COMMENT '部门',
                                 `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='未归用户表';

-- ----------------------------
-- Records of access_return
-- ----------------------------
BEGIN;
INSERT INTO `access_return` (`id`, `name`, `phone`, `card`, `remark`, `dept`, `create_time`) VALUES (1, '杜东亮', '19851920126', '34262319990901159x', '正常请假出校', '软件186', '2022-12-17 18:20:27');
INSERT INTO `access_return` (`id`, `name`, `phone`, `card`, `remark`, `dept`, `create_time`) VALUES (3, '李文节', '13227878838', '342623199906214413', '无', '软件186', '2023-02-28 14:35:53');
COMMIT;

-- ----------------------------
-- Table structure for another_for_unit_test
-- ----------------------------
DROP TABLE IF EXISTS `another_for_unit_test`;
CREATE TABLE `another_for_unit_test` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of another_for_unit_test
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for good_info
-- ----------------------------
DROP TABLE IF EXISTS `good_info`;
CREATE TABLE `good_info` (
                             `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '物资信息id',
                             `type_id` bigint(20) DEFAULT NULL COMMENT '类型id',
                             `name` varchar(50) DEFAULT NULL COMMENT '物资名称',
                             `img` varchar(255) DEFAULT NULL COMMENT '图片链接',
                             `size` varchar(50) DEFAULT NULL COMMENT '物资规格',
                             `unit` varchar(50) DEFAULT NULL COMMENT '物资单位',
                             `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
                             `remark` varchar(50) DEFAULT NULL COMMENT '备注',
                             `total` int(11) DEFAULT NULL COMMENT '库存',
                             `status` int(1) DEFAULT '1' COMMENT '是否启用',
                             `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                             `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                             `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
                             `version` int(11) DEFAULT '1' COMMENT '乐观锁',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='物资信息表';

-- ----------------------------
-- Records of good_info
-- ----------------------------
BEGIN;
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (1, 1, '冰露饮用水', 'https://geektutu.com/post/geecache/geecache.jpg', '打', '20瓶/打', 'admin', '必不可少的饮用水22', 42, 1, '2022-12-10 09:31:46', '2023-03-01 11:38:52', 0, 8);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (2, 4, '医用口罩', 'https://geektutu.com/post/geecache/geecache.jpg', '包', '10只/包', 'admin', '普通的医用口罩', 37, 1, '2022-12-10 15:39:32', '2023-03-01 11:49:41', 0, 3);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (3, 4, 'N95口罩', 'https://geektutu.com/post/geecache/geecache.jpg', '盒', '15只/盒', 'admin', '防护性很强的口罩 安全', 53, 1, '2022-12-10 18:13:04', '2023-03-01 11:38:52', 0, 2);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (4, 1, '葡萄糖', 'https://geektutu.com/post/geecache/geecache.jpg', '盒', '10只/盒', 'admin', '点滴必备', 60, 1, '2022-12-10 18:23:26', '2021-12-18 20:11:39', 0, 3);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (5, 3, '酒精棉', 'https://geektutu.com/post/geecache/geecache.jpg', '箱', '20包/箱', 'admin', '消毒用', 30, 1, '2022-12-11 09:18:15', '2021-12-11 09:29:44', 0, 3);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (6, 1, '可乐', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/37591v2-7ac995b29de01e9985a33538d6c3bee8_r.jpg', '20瓶/箱', '瓶', 'admin', '可乐好喝111', NULL, 1, '2023-03-01 09:29:32', '2023-03-01 10:14:08', 1, 2);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (7, 0, '香蕉', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/12936v2-7ac995b29de01e9985a33538d6c3bee8_r.jpg', '10条', '条', 'admin', '好吃', 0, 1, '2023-03-01 09:38:32', '2023-03-01 09:38:32', 0, 1);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (8, 0, 'aaa', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/55416v2-7ac995b29de01e9985a33538d6c3bee8_r.jpg', '1', '1', 'admin', '11', 0, 1, '2023-03-01 10:31:31', '2023-03-01 10:31:31', 1, 1);
INSERT INTO `good_info` (`id`, `type_id`, `name`, `img`, `size`, `unit`, `create_by`, `remark`, `total`, `status`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (9, 0, 'bb', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/44034v2-7ac995b29de01e9985a33538d6c3bee8_r.jpg', '11', '1', 'admin', '22', 0, 1, '2023-03-01 10:31:45', '2023-03-01 10:31:45', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for good_stock
-- ----------------------------
DROP TABLE IF EXISTS `good_stock`;
CREATE TABLE `good_stock` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '出入库信息id',
                              `accept` varchar(60) DEFAULT NULL COMMENT '去向',
                              `create_by` varchar(50) DEFAULT NULL COMMENT '操作人',
                              `good_num` int(11) DEFAULT NULL COMMENT '物资数量',
                              `good_size` varchar(50) DEFAULT NULL COMMENT '物资规格',
                              `good_name` varchar(50) DEFAULT NULL COMMENT '物资名',
                              `people_name` varchar(50) DEFAULT NULL COMMENT '联系人',
                              `people_phone` varchar(50) DEFAULT NULL COMMENT '联系人电话',
                              `operate_type` int(1) DEFAULT '1' COMMENT '操作类型',
                              `remark` varchar(255) DEFAULT NULL COMMENT '备注',
                              `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                              `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                              `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
                              `version` int(11) DEFAULT '1' COMMENT '乐观锁',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='物资出入库表';

-- ----------------------------
-- Records of good_stock
-- ----------------------------
BEGIN;
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (1, '大数据学院仓库', 'admin', 20, '打', '冰露饮用水', '范晨', '17730312784', 0, '第一次物品入库', '2022-12-11 19:12:44', '2022-12-11 19:12:44', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (2, '大数据学院仓库', 'admin', 10, '打', '冰露饮用水', '范晨', '17730312784', 0, '第二次物资入库', '2022-12-11 19:18:23', '2022-12-11 19:18:23', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (3, '大数据学院仓库', 'admin', 10, '包', '医用口罩', '范晨', '17730312784', 0, '第二次物资入库', '2022-12-11 19:18:23', '2022-12-11 19:18:23', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (4, '大数据学院仓库', 'admin', 10, '盒', '葡萄糖', '范晨', '17730312784', 0, '第二次物资入库', '2022-12-11 19:18:23', '2022-12-11 19:18:23', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (5, '西太湖保安厅', 'admin', 5, '打', '冰露饮用水', '杜东亮', '19851920613', 1, '给保安搞点水喝喝', '2022-12-11 20:52:31', '2022-12-11 20:52:31', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (6, '刘国钧学院仓库', '赵丹妮', 30, '盒', 'N95口罩', '赵丹妮', '18855331293', 0, '入库刘国钧学院仓库', '2022-12-13 14:12:48', '2022-12-13 14:12:48', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (7, '大数据学院仓库', 'admin', 5, '打', '冰露饮用水', '范晨', '17730312784', 0, '入库', '2022-12-17 12:47:25', '2022-12-17 12:47:25', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (8, '大数据学院仓库', 'admin', 10, '包', '医用口罩', '范晨', '17730312784', 0, '入库', '2022-12-17 12:47:25', '2022-12-17 12:47:25', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (9, '华师', 'admin', 2, '打', '冰露饮用水', '', '', 1, '我', '2023-03-01 11:38:52', '2023-03-01 11:38:52', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (10, '华师', 'admin', 3, '盒', 'N95口罩', '', '', 1, '我', '2023-03-01 11:38:52', '2023-03-01 11:38:52', 0, 1);
INSERT INTO `good_stock` (`id`, `accept`, `create_by`, `good_num`, `good_size`, `good_name`, `people_name`, `people_phone`, `operate_type`, `remark`, `create_time`, `update_time`, `is_delete`, `version`) VALUES (11, '华工', 'admin', 3, '包', '医用口罩', '丰富1', '13223445543', 1, '111', '2023-03-01 11:49:41', '2023-03-01 11:49:41', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for good_type
-- ----------------------------
DROP TABLE IF EXISTS `good_type`;
CREATE TABLE `good_type` (
                             `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '物资类型id',
                             `type` varchar(30) DEFAULT NULL COMMENT '物资类型',
                             `order_num` int(4) DEFAULT NULL COMMENT '排序',
                             `status` int(1) DEFAULT '1' COMMENT '状态',
                             `create_by` varchar(30) DEFAULT NULL COMMENT '创建人',
                             `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                             `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                             `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
                             `version` int(11) DEFAULT '1' COMMENT '乐观锁',
                             `remark` varchar(80) DEFAULT NULL COMMENT '备注',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='物资类型表';

-- ----------------------------
-- Records of good_type
-- ----------------------------
BEGIN;
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (1, '食物', 1, 1, 'admin', '2022-12-09 16:22:33', '2022-12-09 16:22:36', 0, 1, '可食用的食物');
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (2, '衣物', 2, 0, 'admin', '2022-12-09 19:04:58', '2022-12-18 20:12:20', 0, 5, '穿的衣服');
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (3, '医疗用具', 3, 1, 'admin', '2022-12-09 19:06:21', '2022-12-11 09:18:55', 0, 4, '医疗必备');
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (4, '防护用具', 4, 1, 'admin', '2022-12-09 19:08:08', '2022-12-10 15:10:26', 0, 5, '医护人员穿戴的防护用具');
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (5, '玩具', 0, 1, 'admin', '2023-03-01 12:10:21', '2023-03-01 12:12:02', 1, 1, '玩具解闷');
INSERT INTO `good_type` (`id`, `type`, `order_num`, `status`, `create_by`, `create_time`, `update_time`, `is_delete`, `version`, `remark`) VALUES (6, '玩具', 0, 0, 'admin', '2023-03-01 12:12:37', '2023-03-01 12:19:20', 0, 1, '玩具解闷');
COMMIT;

-- ----------------------------
-- Table structure for health_clock
-- ----------------------------
DROP TABLE IF EXISTS `health_clock`;
CREATE TABLE `health_clock` (
                                `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '打卡id',
                                `username` varchar(50) DEFAULT NULL COMMENT '姓名',
                                `health_type` int(1) DEFAULT NULL COMMENT '健康状况',
                                `temperature` float(6,1) DEFAULT NULL COMMENT '温度',
  `middle_high` int(1) DEFAULT NULL COMMENT '中高风险',
  `diagnosis` int(1) DEFAULT NULL COMMENT '确诊',
  `return_info` int(1) DEFAULT NULL COMMENT '境外返回',
  `address` varchar(60) DEFAULT NULL COMMENT '地址',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门id',
  `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='打卡健康表';

-- ----------------------------
-- Records of health_clock
-- ----------------------------
BEGIN;
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (1, 'admin', 2, 38.0, 1, 1, 1, '江苏省-苏州市-昆山市', '2022-12-15 14:24:57', '2022-12-15 14:24:57', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (2, 'admin', 2, 38.0, 1, 1, 1, '北京市-市辖区-海淀区', '2022-12-15 14:28:15', '2022-12-15 14:28:15', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (3, 'student', 2, 38.0, 1, 1, 1, '江苏省-南京市-玄武区', '2022-12-15 14:45:49', '2022-12-15 14:45:49', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (4, 'student', 2, 36.9, 1, 1, 1, '江苏省-南京市-建邺区', '2022-12-16 12:30:04', '2022-12-16 12:30:04', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (5, 'admin', 2, 37.1, 1, 1, 1, '江苏省-南京市-秦淮区', '2022-12-16 15:26:04', '2022-12-16 15:26:04', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (6, 'admin', 2, 37.3, 1, 1, 1, '安徽省-芜湖市-无为县', '2022-12-18 20:14:05', '2022-12-18 20:14:05', 103, 0);
INSERT INTO `health_clock` (`id`, `username`, `health_type`, `temperature`, `middle_high`, `diagnosis`, `return_info`, `address`, `create_time`, `update_time`, `dept_id`, `is_delete`) VALUES (9, 'admin', 0, 37.1, 0, 1, 0, '北京市-县-延庆县', '2023-03-01 15:49:56', '2023-03-01 15:49:56', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for health_report
-- ----------------------------
DROP TABLE IF EXISTS `health_report`;
CREATE TABLE `health_report` (
                                 `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                                 `username` varchar(50) DEFAULT NULL COMMENT '用户名',
                                 `dept_id` bigint(11) DEFAULT NULL COMMENT '部门id',
                                 `phone_number` varchar(50) DEFAULT NULL COMMENT '手机号',
                                 `img3` varchar(255) DEFAULT NULL COMMENT '核酸报告',
                                 `img2` varchar(255) DEFAULT NULL COMMENT '行程码',
                                 `img1` varchar(255) DEFAULT NULL COMMENT '健康码',
                                 `type` int(1) DEFAULT NULL COMMENT '返校情况',
                                 `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                 `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='二码一报告表';

-- ----------------------------
-- Records of health_report
-- ----------------------------
BEGIN;
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (1, 'admin', 103, '17730312784', 'http://localhost:8081/img/2021/12/09a5984e9a1448a5989396ae02ee2218.jpg', 'http://localhost:8081/img/2021/12/70b927c1e2474706bb7d96e5bf515af7.jpg', 'http://localhost:8081/img/2021/12/ca5ed73630a94883864215d2b4d5ba07.jpg', 1, '2022-12-16 15:05:52', '2022-12-16 15:05:52');
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (2, 'admin', 103, '17730312784', 'http://localhost:8081/img/2021/12/73054455739a4e4a90999af1badef193.jpg', 'http://localhost:8081/img/2021/12/1ea4667c7734445ebcee993c8166830f.jpg', 'http://localhost:8081/img/2021/12/2402dbf4d1a34efca35abec5d7288770.jpg', 1, '2022-12-16 15:08:29', '2022-12-16 15:08:29');
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (3, 'admin', 103, '17730312784', 'http://localhost:8081/img/2021/12/3bd20328127a4913880a103997789fa4.jpg', 'http://localhost:8081/img/2021/12/efb111d5de62419ab771d6c310a1e64c.jpg', 'http://localhost:8081/img/2021/12/fe54d296d493481abf39d5b5a197390f.jpg', 0, '2022-12-16 15:13:04', '2022-12-16 15:13:04');
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (4, 'student', 103, '17730312784', 'http://localhost:8081/img/2021/12/b1fb6d36d7d446d282efa082c4d18887.jpg', 'http://localhost:8081/img/2021/12/d37c7ab0149c407a9487b91926d6a8f9.jpg', 'http://localhost:8081/img/2021/12/73caed944b514424a2f22c32f26701f5.jpg', 0, '2022-12-16 15:15:17', '2022-12-16 15:15:17');
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (5, 'admin', 103, '17730312784', 'http://localhost:8081/img/2021/12/d8b797f1567c471eae9176232d1d97ad.jpg', 'http://localhost:8081/img/2021/12/41d62d3b8e314ee4b575b238cd9c8bdd.jpg', 'http://localhost:8081/img/2021/12/351a3a4da719497097f7781c5d1e33e1.jpg', 1, '2022-12-18 20:15:04', '2022-12-18 20:15:04');
INSERT INTO `health_report` (`id`, `username`, `dept_id`, `phone_number`, `img3`, `img2`, `img1`, `type`, `create_time`, `update_time`) VALUES (6, 'admin', 103, '17730312784', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/28710Snipaste_2023-03-01_16-59-37.png', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/28710Snipaste_2023-03-01_16-59-27.png', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/1/28710Snipaste_2023-03-01_16-59-01.png', 0, '2023-03-01 17:05:30', '2023-03-01 17:05:30');
COMMIT;

-- ----------------------------
-- Table structure for leave_apply
-- ----------------------------
DROP TABLE IF EXISTS `leave_apply`;
CREATE TABLE `leave_apply` (
                               `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
                               `username` varchar(50) DEFAULT NULL COMMENT '用户名',
                               `dept_id` bigint(11) DEFAULT NULL COMMENT '部门id',
                               `reason` varchar(255) DEFAULT NULL COMMENT '请假原因',
                               `leave_type` int(1) DEFAULT NULL COMMENT '请假类型（1：事假 2：病假）',
                               `status` int(1) DEFAULT NULL COMMENT '状态（0：撤销 1：待审核 2：审核通过 3：审核不通过）',
                               `student_type` int(1) DEFAULT NULL COMMENT '学生类型（1：本科生 2：研究生 3：博士生）',
                               `nickname` varchar(64) DEFAULT NULL COMMENT '学生姓名',
                               `time` varchar(64) DEFAULT NULL COMMENT '请假时间区间',
                               `day` varchar(50) DEFAULT NULL COMMENT '请假天数',
                               `address` varchar(100) DEFAULT NULL COMMENT '目的地',
                               `traffic` varchar(50) DEFAULT NULL COMMENT '交通工具',
                               `clazz` int(1) DEFAULT NULL COMMENT '是否有课程（1：没有 0：有）',
                               `dormitory` varchar(50) DEFAULT NULL COMMENT '宿舍',
                               `phone_number` varchar(50) DEFAULT NULL COMMENT '手机号码',
                               `exam` int(1) DEFAULT NULL COMMENT '考试（1：没有 0：有）',
                               `opinion` varchar(120) DEFAULT NULL COMMENT '审核意见',
                               `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                               `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                               `is_delete` int(1) DEFAULT '0' COMMENT '逻辑删除',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='请假审批表';

-- ----------------------------
-- Records of leave_apply
-- ----------------------------
BEGIN;
INSERT INTO `leave_apply` (`id`, `username`, `dept_id`, `reason`, `leave_type`, `status`, `student_type`, `nickname`, `time`, `day`, `address`, `traffic`, `clazz`, `dormitory`, `phone_number`, `exam`, `opinion`, `create_time`, `update_time`, `is_delete`) VALUES (1, 'student', 103, '元旦回家和女朋友约会', 1, 2, 3, '学生用户', '2022-01-01 至 2022-01-03', '3', '安徽芜湖', '高铁G7774', 1, '7-540', '17730312784', 1, '合情合理 允许你请假', '2022-12-18 13:30:12', '2022-12-18 20:20:53', 0);
INSERT INTO `leave_apply` (`id`, `username`, `dept_id`, `reason`, `leave_type`, `status`, `student_type`, `nickname`, `time`, `day`, `address`, `traffic`, `clazz`, `dormitory`, `phone_number`, `exam`, `opinion`, `create_time`, `update_time`, `is_delete`) VALUES (2, 'student', 103, '啊啊啊', 2, 1, 3, '学生用户', '2023-03-02 至 2023-04-09', '1', '北京', '地铁', 0, '222', '132234545555', 0, NULL, '2023-03-02 11:19:14', '2023-03-02 11:52:28', 0);
INSERT INTO `leave_apply` (`id`, `username`, `dept_id`, `reason`, `leave_type`, `status`, `student_type`, `nickname`, `time`, `day`, `address`, `traffic`, `clazz`, `dormitory`, `phone_number`, `exam`, `opinion`, `create_time`, `update_time`, `is_delete`) VALUES (3, 'student', 103, '学习', 2, 1, 2, '学生用户', '2023-03-17 至 2023-04-23', '1', '广州', '飞机', 1, '333', '13223333333', 0, '', '2023-03-02 11:53:16', '2023-03-02 11:53:16', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
                            `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
                            `parent_id` bigint(20) DEFAULT '0' COMMENT '父部门id',
                            `dept_name` varchar(30) DEFAULT '' COMMENT '部门名称',
                            `order_num` int(4) DEFAULT '0' COMMENT '显示顺序',
                            `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
                            `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
                            `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
                            `status` int(1) DEFAULT '1' COMMENT '部门状态（1正常 0停用）',
                            `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
                            `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                            `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
                            `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                            `is_delete` int(1) DEFAULT '0',
                            `version` int(11) DEFAULT '1',
                            PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8mb4 COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (100, 0, '常州大学', 0, '校长', '15888888888', 'asdasaa@qq.com', 1, 'admin', '2022-11-29 18:52:24', '', '2023-03-03 09:11:00', 0, 5);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (101, 100, '西太湖校区', 1, '范晨', '16888888888', '1571025887@qq.com', 1, 'admin', '2022-11-29 18:52:24', '', '2022-12-03 18:51:45', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (102, 100, '科教城校区', 2, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-03 18:52:18', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (103, 101, '阿里云大数据学院', 1, '杜东亮', '19851922596', '2473786752@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:35:27', 0, 3);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (104, 101, '刘国钧学院', 2, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:35:40', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (105, 101, '商学院', 3, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:35:58', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (106, 101, '机械学院', 4, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:36:09', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (107, 101, '运维部门', 5, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', NULL, 1, 1);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (108, 102, '艺术学院', 1, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:36:18', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (109, 102, '工学院', 2, '范晨', '15888888888', 'ry@qq.com', 1, 'admin', '2022-11-29 18:52:25', '', '2022-12-17 19:36:38', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (200, 102, '石油化工学院', 3, '倪兴林', '19851920126', 'xigongnei942@163.com', 1, '', '2022-12-03 18:57:54', '', '2022-12-17 19:36:55', 0, 2);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (201, 100, '大学城校区', 1, '李文瑞', '13228783390', 'fefa@qq.com', 0, NULL, '2023-03-03 09:12:47', 'admin', '2023-03-03 09:30:23', 1, 1);
INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_delete`, `version`) VALUES (202, 201, 'test', 0, '111', '13223333344', '111@qq.com', 1, 'admin', '2023-03-03 09:31:19', '', '2023-03-03 09:31:19', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_login_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_info`;
CREATE TABLE `sys_login_info` (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
                                  `username` varchar(50) DEFAULT '' COMMENT '用户账号',
                                  `ip` varchar(50) DEFAULT '' COMMENT '登录IP地址',
                                  `login_location` varchar(100) DEFAULT '' COMMENT '登录地点',
                                  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
                                  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
                                  `status` int(1) DEFAULT '1' COMMENT '登录状态（0成功 1失败）',
                                  `msg` varchar(100) DEFAULT '' COMMENT '提示消息',
                                  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8mb4 COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_login_info
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_info` (`id`, `username`, `ip`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (1, 'student', '127.0.0.1', ' 本机地址', 'Firefox 110.0', 'Windows 10', 1, '登录成功', '2023-03-06 11:06:19');
INSERT INTO `sys_login_info` (`id`, `username`, `ip`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`) VALUES (2, 'admin', '127.0.0.1', ' 本机地址', 'Firefox 110.0', 'Windows 10', 1, '登录成功', '2023-03-06 11:06:30');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `parent_id` bigint(20) DEFAULT NULL COMMENT '父菜单ID，一级菜单为0',
                            `name` varchar(64) NOT NULL COMMENT '菜单名',
                            `path` varchar(255) DEFAULT NULL COMMENT '菜单URL',
                            `perms` varchar(255) DEFAULT NULL COMMENT '授权(多个用逗号分隔，如：user:list,user:create)',
                            `component` varchar(255) DEFAULT NULL,
                            `type` int(5) NOT NULL COMMENT '类型     0：目录   1：菜单   2：按钮',
                            `icon` varchar(32) DEFAULT NULL COMMENT '菜单图标',
                            `order_num` int(11) DEFAULT NULL COMMENT '排序',
                            `create_time` datetime DEFAULT NULL,
                            `update_time` datetime DEFAULT NULL,
                            `status` int(5) NOT NULL,
                            `is_delete` int(1) NOT NULL DEFAULT '0',
                            `version` int(11) NOT NULL DEFAULT '1',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (1, 0, '系统管理', '', 'sys:manage', '', 0, 'el-icon-s-operation', 1, '2022-01-15 18:58:18', '2022-01-15 18:58:20', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (2, 1, '用户管理', '/sys/user', 'sys:user:list', 'sys/User', 1, 'el-icon-s-custom', 1, '2022-01-15 19:03:45', '2022-01-15 19:03:48', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (3, 1, '角色管理', '/sys/role', 'sys:role:list', 'sys/Role', 1, 'el-icon-s-promotion', 2, '2022-01-15 19:03:45', '2022-12-01 18:12:43', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (4, 1, '菜单管理', '/sys/menu', 'sys:menu:list', 'sys/Menu', 1, 'el-icon-menu', 3, '2022-01-15 19:03:45', '2022-01-15 19:03:48', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (5, 0, '系统工具', '', 'sys:tools', NULL, 0, 'el-icon-s-tools', 2, '2022-01-15 19:06:11', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (6, 5, '登录日志', '/sys/loginInfo', 'sys:login:list', 'sys/LoginInfo', 1, 'el-icon-s-order', 1, '2022-01-15 19:07:18', '2022-12-04 19:03:11', 1, 0, 6);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (7, 3, '添加角色', '', 'sys:role:save', '', 2, '', 1, '2022-01-15 23:02:25', '2022-01-17 21:53:14', 0, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (9, 2, '添加用户', NULL, 'sys:user:save', NULL, 2, NULL, 1, '2022-01-17 21:48:32', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (10, 2, '修改用户', NULL, 'sys:user:update', NULL, 2, NULL, 2, '2022-01-17 21:49:03', '2022-01-17 21:53:04', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (11, 2, '删除用户', NULL, 'sys:user:delete', NULL, 2, NULL, 3, '2022-01-17 21:49:21', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (12, 2, '分配角色', NULL, 'sys:user:role', NULL, 2, NULL, 4, '2022-01-17 21:49:58', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (13, 2, '重置密码', NULL, 'sys:user:repass', NULL, 2, NULL, 5, '2022-01-17 21:50:36', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (14, 3, '修改角色', NULL, 'sys:role:update', NULL, 2, NULL, 2, '2022-01-17 21:51:14', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (15, 3, '删除角色', NULL, 'sys:role:delete', NULL, 2, NULL, 3, '2022-01-17 21:51:39', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (16, 3, '分配权限', NULL, 'sys:role:perm', NULL, 2, NULL, 4, '2022-01-17 21:52:02', '2022-12-03 14:54:45', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (17, 4, '添加菜单', NULL, 'sys:menu:save', NULL, 2, NULL, 1, '2022-01-17 21:53:53', '2022-01-17 21:55:28', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (18, 4, '修改菜单', NULL, 'sys:menu:update', NULL, 2, NULL, 2, '2022-01-17 21:56:12', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (19, 4, '删除菜单', NULL, 'sys:menu:delete', NULL, 2, NULL, 3, '2022-01-17 21:56:36', NULL, 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (21, 1, '部门管理', '/sys/dept', 'sys:dept:list', 'sys/Dept', 1, 'el-icon-s-data', 4, '2022-12-03 14:50:45', '2022-12-03 14:51:08', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (22, 21, '添加部门', NULL, 'sys:dept:save', NULL, 2, '', 1, '2022-12-03 14:52:52', '2022-12-03 14:52:52', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (23, 21, '修改部门', NULL, 'sys:dept:update', NULL, 2, '', 2, '2022-12-03 14:53:25', '2022-12-03 14:53:25', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (24, 21, '删除部门', NULL, 'sys:dept:delete', NULL, 2, '', 3, '2022-12-03 14:53:52', '2022-12-03 14:53:52', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (25, 5, 'mysql数据监控', '/monitor/druid', 'monitor:druid:list', 'monitor/Druid', 1, 'el-icon-set-up', 3, '2022-12-03 21:00:35', '2023-03-06 10:39:02', 0, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (26, 5, '系统接口', '/monitor/swagger', 'monitor:swagger:list', 'monitor/Swagger', 1, 'el-icon-more', 5, '2022-12-03 22:01:41', '2022-12-07 19:35:18', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (27, 5, '服务监控', '/monitor/server', 'monitor:server:list', 'monitor/Server', 1, 'el-icon-cpu', 4, '2022-12-03 22:59:56', '2022-12-03 22:59:56', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (28, 5, '操作日志', '/monitor/operate', 'monitor:operate:list', 'monitor/Operate', 1, 'el-icon-s-platform', 2, '2022-12-05 14:45:08', '2022-12-07 19:35:01', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (29, 6, '删除登录日志', NULL, 'sys:login:delete', NULL, 2, '', 1, '2022-12-07 18:31:50', '2022-12-07 18:34:56', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (30, 6, '清空登录日志', NULL, 'sys:login:clear', NULL, 2, '', 2, '2022-12-07 18:32:10', '2022-12-07 18:35:03', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (32, 28, '删除操作日志', NULL, 'monitor:operate:delete', NULL, 2, '', 1, '2022-12-07 18:35:20', '2022-12-07 18:35:20', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (33, 28, '清空操作日志', NULL, 'monitor:operate:clear', NULL, 2, '', 1, '2022-12-07 18:35:36', '2022-12-07 18:35:36', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (34, 0, '物资管理', NULL, 'good:manage', NULL, 0, 'el-icon-goods', 3, '2022-12-09 14:42:21', '2022-12-09 14:42:21', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (35, 34, '物资资料', '/good/info', 'good:info:list', 'good/Info', 1, 'el-icon-odometer', 1, '2022-12-09 14:52:35', '2022-12-09 14:52:35', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (36, 34, '物资分类', '/good/type', 'good:type:list', 'good/Type', 1, 'el-icon-s-open', 2, '2022-12-09 15:07:51', '2022-12-09 15:07:51', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (37, 34, '物资库存', '/good/total', 'good:total:list', 'good/Total', 1, 'el-icon-set-up', 3, '2022-12-10 18:31:27', '2022-12-10 18:31:27', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (38, 34, '物资出入库', '/good/stock', 'good:stock:list', 'good/Stock', 1, 'el-icon-s-data', 4, '2022-12-11 09:39:26', '2022-12-11 09:39:26', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (39, 35, '添加物资', NULL, 'good:info:save', NULL, 2, '', 1, '2022-12-12 15:19:21', '2022-12-12 15:19:21', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (40, 35, '更新物资', NULL, 'good:info:update', NULL, 2, '', 2, '2022-12-12 15:20:10', '2022-12-12 15:20:57', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (41, 35, '删除物资', NULL, 'good:info:delete', NULL, 2, '', 3, '2022-12-12 15:20:51', '2022-12-12 15:20:51', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (42, 36, '添加分类', NULL, 'good:type:save', NULL, 2, '', 1, '2022-12-12 15:21:48', '2022-12-12 15:21:48', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (43, 36, '更新分类', NULL, 'good:type:update', NULL, 2, '', 2, '2022-12-12 15:22:08', '2022-12-12 15:22:35', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (44, 36, '删除分类', NULL, 'good:type:delete', NULL, 2, '', 3, '2022-12-12 15:22:30', '2022-12-12 15:22:30', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (45, 38, '出入库', NULL, 'good:stock:operate', NULL, 2, '', 1, '2022-12-12 15:23:35', '2022-12-12 15:23:35', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (46, 0, '健康管理', NULL, 'health:manage', NULL, 0, 'el-icon-location', 4, '2022-12-13 20:44:35', '2022-12-13 20:44:35', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (47, 46, '健康打卡', '/health/clock', 'health:clock:save', 'health/Clock', 1, 'el-icon-s-promotion', 1, '2022-12-14 09:17:41', '2022-12-14 09:17:41', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (48, 46, '打卡信息', '/health/clockInfo', 'health:clock:list', 'health/ClockInfo', 1, 'el-icon-s-opportunity', 2, '2022-12-14 09:19:59', '2022-12-14 09:22:09', 1, 0, 2);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (49, 46, '二码一报告', '/health/report', 'health:report:save', 'health/Report', 1, 'el-icon-chat-round', 3, '2022-12-14 09:21:57', '2022-12-14 09:21:57', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (50, 46, '报告信息', '/health/reportInfo', 'health:report:list', 'health/ReportInfo', 1, 'el-icon-s-data', 4, '2022-12-14 09:23:48', '2022-12-14 09:23:48', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (51, 0, '出行管理', NULL, 'access:manage', NULL, 0, 'el-icon-user-solid', 5, '2022-12-16 18:45:38', '2022-12-16 18:45:38', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (52, 51, '出入登记', '/access/register', 'access:register:list', 'access/Register', 1, 'el-icon-s-platform', 1, '2022-12-16 18:50:55', '2022-12-16 18:50:55', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (53, 51, '未归人员', '/access/return', 'access:return:list', 'access/Return', 1, 'el-icon-phone', 2, '2022-12-17 15:17:30', '2022-12-17 15:17:30', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (54, 0, '请假管理', '', 'leave:manage', NULL, 0, 'el-icon-paperclip', 6, '2022-12-17 18:25:58', '2022-12-17 18:25:58', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (55, 54, '请假申请', '/leave/apply', 'leave:apply:list', 'leave/Apply', 1, 'el-icon-s-promotion', 1, '2022-12-17 18:28:13', '2022-12-17 18:28:13', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (56, 54, '请假记录', '/leave/record', 'leave:record:list', 'leave/Record', 1, 'el-icon-s-platform', 2, '2022-12-17 18:29:23', '2022-12-17 18:29:23', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (58, 52, '添加记录', NULL, 'access:register:save', NULL, 2, '', 1, '2022-12-18 18:57:18', '2022-12-18 18:57:18', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (59, 55, '新增请假', NULL, 'leave:apply:save', NULL, 2, '', 1, '2022-12-18 18:58:10', '2022-12-18 18:58:10', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (60, 56, '审核请假', NULL, 'leave:record:examine', NULL, 2, '', 1, '2022-12-18 18:59:24', '2022-12-18 18:59:24', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (61, 55, '修改请假', NULL, 'leave:apply:update', NULL, 2, '', 2, '2022-12-18 19:08:52', '2022-12-18 19:08:52', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (62, 5, 'redis监控', '/monitor/redis', 'monitor:redis:list', 'monitor/Redis', 1, 'el-icon-s-data', 1, '2023-03-02 16:41:32', '2023-03-02 16:41:32', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (66, 0, '测试菜单', '', 'test', '', 0, 'el-icon-delete', 0, '2023-03-03 12:05:45', '2023-03-03 12:05:45', 0, 1, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (67, 5, '公告管理', '/monitor/notice', 'monitor:notice:list', 'monitor/Notice', 1, 'el-icon-star-on', 1, '2023-03-03 12:11:48', '2023-03-03 12:11:48', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (68, 67, '设置公告', NULL, 'monitor:notice:set', NULL, 2, '', 1, '2023-03-03 14:29:03', '2023-03-03 14:29:03', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (69, 67, '修改公告', NULL, 'monitor:notice:update', NULL, 2, '', 1, '2023-03-03 14:29:20', '2023-03-03 14:29:20', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (70, 67, '删除公告', NULL, 'monitor:notice:delete', NULL, 2, '', 1, '2023-03-03 14:29:35', '2023-03-03 14:29:35', 1, 0, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `perms`, `component`, `type`, `icon`, `order_num`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (71, 67, '创建公告', NULL, 'monitor:notice:save', NULL, 2, '', 1, '2023-03-03 14:29:57', '2023-03-03 14:29:57', 1, 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `title` varchar(64) DEFAULT NULL,
                              `content` varchar(64) DEFAULT NULL,
                              `status` int(5) NOT NULL,
                              `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
                              `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
                              `create_time` datetime DEFAULT NULL,
                              `update_time` datetime DEFAULT NULL,
                              `is_delete` int(1) DEFAULT '0',
                              `remark` varchar(50) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
BEGIN;
INSERT INTO `sys_notice` (`id`, `title`, `content`, `status`, `create_by`, `update_by`, `create_time`, `update_time`, `is_delete`, `remark`) VALUES (5, 't1', 't2c', 0, '', '', '2023-02-28 02:33:59', '2023-02-28 02:33:59', 1, '');
INSERT INTO `sys_notice` (`id`, `title`, `content`, `status`, `create_by`, `update_by`, `create_time`, `update_time`, `is_delete`, `remark`) VALUES (6, '更新通知', '<p>系统将更新</p>', 1, 'admin', '', '2023-03-03 14:33:25', '2023-03-03 14:33:55', 0, 'test');
COMMIT;

-- ----------------------------
-- Table structure for sys_operate_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_operate_log`;
CREATE TABLE `sys_operate_log` (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
                                   `title` varchar(50) DEFAULT '' COMMENT '模块标题',
                                   `business_type` varchar(50) DEFAULT '' COMMENT '业务类型',
                                   `method` varchar(100) DEFAULT '' COMMENT '方法名称',
                                   `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
                                   `oper_url` varchar(255) DEFAULT '' COMMENT '请求URL',
                                   `oper_ip` varchar(128) DEFAULT '' COMMENT '主机地址',
                                   `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
                                   `oper_param` varchar(2000) DEFAULT '' COMMENT '请求参数',
                                   `oper_name` varchar(50) DEFAULT NULL COMMENT '操作人',
                                   `json_result` varchar(2000) DEFAULT '' COMMENT '返回参数',
                                   `status` int(1) DEFAULT '1' COMMENT '操作状态（1正常 0异常）',
                                   `error_msg` varchar(2000) DEFAULT '' COMMENT '错误消息',
                                   `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_operate_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_operate_log` (`id`, `title`, `business_type`, `method`, `request_method`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `oper_name`, `json_result`, `status`, `error_msg`, `oper_time`) VALUES (1, '清除日志', '', 'ClearOpLog', 'POST', '/sys/operateLog', '127.0.0.1', ' 本机地址', '', 'admin', '{\"code\":200,\"msg\":\"操作成功\",\"data\":\"清空成功\"}', 1, '', '2023-03-06 11:01:38');
INSERT INTO `sys_operate_log` (`id`, `title`, `business_type`, `method`, `request_method`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `oper_name`, `json_result`, `status`, `error_msg`, `oper_time`) VALUES (2, '登录日志管理', '清空登录日志', 'ClearLoginInfo', 'POST', '/sys/loginInfo', '127.0.0.1', ' 本机地址', '', 'admin', '{\"code\":200,\"msg\":\"操作成功\",\"data\":\"所有数据已清除\"}', 1, '', '2023-03-06 11:06:06');
INSERT INTO `sys_operate_log` (`id`, `title`, `business_type`, `method`, `request_method`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `oper_name`, `json_result`, `status`, `error_msg`, `oper_time`) VALUES (3, '', '', '', 'POST', '/logout', '127.0.0.1', ' 本机地址', '', '', '{\"code\":200,\"msg\":\"操作成功\",\"data\":\"退出成功\"}', 1, '', '2023-03-06 11:06:12');
INSERT INTO `sys_operate_log` (`id`, `title`, `business_type`, `method`, `request_method`, `oper_url`, `oper_ip`, `oper_location`, `oper_param`, `oper_name`, `json_result`, `status`, `error_msg`, `oper_time`) VALUES (4, '', '', '', 'POST', '/logout', '127.0.0.1', ' 本机地址', '', '', '{\"code\":200,\"msg\":\"操作成功\",\"data\":\"退出成功\"}', 1, '', '2023-03-06 11:06:24');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `name` varchar(64) NOT NULL,
                            `code` varchar(64) NOT NULL,
                            `remark` varchar(64) DEFAULT NULL COMMENT '备注',
                            `create_time` datetime DEFAULT NULL,
                            `update_time` datetime DEFAULT NULL,
                            `status` int(5) NOT NULL,
                            `is_delete` int(1) DEFAULT '0',
                            `version` int(11) DEFAULT '1',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `name` (`name`) USING BTREE,
                            UNIQUE KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `name`, `code`, `remark`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (3, '学生用户', 'student', '只有健康管理功能', '2022-01-04 10:09:14', '2023-03-06 09:28:48', 1, 0, 17);
INSERT INTO `sys_role` (`id`, `name`, `code`, `remark`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (6, '超级管理员', 'admin', '系统默认最高权限，不可以编辑和任意修改', '2022-01-16 13:29:03', '2023-03-03 14:30:17', 1, 0, 25);
INSERT INTO `sys_role` (`id`, `name`, `code`, `remark`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (8, '老师用户', 'teacher', '老师用户 查看健康管理和审批功能', '2022-12-02 13:15:34', '2022-12-18 18:49:55', 1, 0, 18);
INSERT INTO `sys_role` (`id`, `name`, `code`, `remark`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (10, '后勤用户', 'service', '后勤用户 负责管理物资', '2022-12-12 14:49:17', '2022-12-18 18:49:03', 1, 0, 4);
INSERT INTO `sys_role` (`id`, `name`, `code`, `remark`, `create_time`, `update_time`, `status`, `is_delete`, `version`) VALUES (11, '开发者', 'developer', '开发者', '2023-03-03 15:52:28', '2023-03-03 15:52:55', 1, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `role_id` bigint(20) NOT NULL,
                                 `menu_id` bigint(20) NOT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=921 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (898, 10, 34);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (899, 10, 35);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (900, 10, 39);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (901, 10, 40);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (902, 10, 41);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (903, 10, 36);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (904, 10, 42);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (905, 10, 43);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (906, 10, 44);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (907, 10, 37);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (908, 10, 38);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (909, 10, 45);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (910, 10, 51);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (911, 10, 52);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (912, 10, 53);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (913, 8, 6);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (914, 8, 46);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (915, 8, 48);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (916, 8, 50);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (917, 8, 51);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (918, 8, 53);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (919, 8, 54);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (920, 8, 56);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1056, 6, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1057, 6, 2);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1058, 6, 9);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1059, 6, 10);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1060, 6, 11);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1061, 6, 12);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1062, 6, 13);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1063, 6, 3);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1064, 6, 7);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1065, 6, 14);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1066, 6, 15);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1067, 6, 16);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1068, 6, 4);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1069, 6, 17);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1070, 6, 18);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1071, 6, 19);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1072, 6, 21);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1073, 6, 22);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1074, 6, 23);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1075, 6, 24);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1076, 6, 5);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1077, 6, 6);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1078, 6, 29);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1079, 6, 30);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1080, 6, 62);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1081, 6, 67);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1082, 6, 68);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1083, 6, 69);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1084, 6, 70);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1085, 6, 71);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1086, 6, 28);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1087, 6, 32);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1088, 6, 33);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1089, 6, 25);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1090, 6, 27);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1091, 6, 26);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1092, 6, 34);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1093, 6, 35);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1094, 6, 39);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1095, 6, 40);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1096, 6, 41);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1097, 6, 36);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1098, 6, 42);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1099, 6, 43);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1100, 6, 44);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1101, 6, 37);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1102, 6, 38);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1103, 6, 45);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1104, 6, 46);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1105, 6, 47);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1106, 6, 48);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1107, 6, 49);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1108, 6, 50);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1109, 6, 51);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1110, 6, 52);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1111, 6, 53);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1112, 6, 54);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1113, 6, 55);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1114, 6, 56);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1173, 3, 2);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1174, 3, 10);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1175, 3, 3);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1176, 3, 4);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1177, 3, 21);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1178, 3, 46);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1179, 3, 47);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1180, 3, 48);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1181, 3, 49);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1182, 3, 50);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1183, 3, 54);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1184, 3, 55);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1185, 3, 59);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1186, 3, 61);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1187, 3, 56);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `username` varchar(64) DEFAULT NULL,
                            `password` varchar(64) DEFAULT NULL,
                            `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
                            `avatar` varchar(255) DEFAULT NULL,
                            `phone_number` varchar(64) DEFAULT NULL,
                            `city` varchar(64) DEFAULT NULL,
                            `dept_id` bigint(20) DEFAULT NULL,
                            `create_time` datetime DEFAULT NULL,
                            `update_time` datetime DEFAULT NULL,
                            `remark` varchar(50) DEFAULT NULL,
                            `status` int(5) NOT NULL,
                            `is_delete` int(1) DEFAULT '0',
                            `version` int(11) DEFAULT '1',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `UK_USERNAME` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/6/56639Snipaste_2023-03-01_16-59-01.png', '17730312781', '广东广州', 103, '2022-01-12 22:13:53', '2023-03-06 11:00:57', '超级管理员', 1, 0, 53);
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (2, 'student', 'e10adc3949ba59abbe56e057f20f883e', '学生用户', 'https://go-pj-oss.oss-cn-guangzhou.aliyuncs.com/uploadfile/2023/March/6/20628v2-7ac995b29de01e9985a33538d6c3bee8_r.jpg', '17730312784', '广州', 103, '2022-12-13 14:07:44', '2023-03-06 09:24:59', '学生用户', 1, 0, 4);
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (3, 'service', 'e10adc3949ba59abbe56e057f20f883e', '后勤用户', 'http://localhost:8081/img/2021/12/59b880040d6443b28956519118c2d507.jpg', '18855331293', '芜湖', 104, '2022-12-13 14:11:13', '2022-12-17 19:23:06', '后勤用户', 1, 0, 6);
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (4, 'teacher', 'e10adc3949ba59abbe56e057f20f883e', '老师用户', 'http://localhost:8081/img/2021/12/400bba2995ec46fc97249cece09f9ca4.jpg', '13966346765', '常州', 200, '2022-12-16 12:27:41', '2022-12-16 18:03:52', '老师用户', 1, 0, 3);
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (5, 'user1', 'e10adc3949ba59abbe56e057f20f883e', '', 'https://geektutu.com/post/geecache/geecache.jpg', '13229875590', '伊拉克', 103, '2023-03-02 17:52:59', '2023-03-04 14:33:19', 'user for test', 1, 0, 1);
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `avatar`, `phone_number`, `city`, `dept_id`, `create_time`, `update_time`, `remark`, `status`, `is_delete`, `version`) VALUES (7, 'test', 'e10adc3949ba59abbe56e057f20f883e', '', '', '13223334456', 'test', 103, '2023-03-03 17:57:04', '2023-03-03 18:02:36', 'test', 1, 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `user_id` bigint(20) NOT NULL,
                                 `role_id` bigint(20) NOT NULL,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (19, 2, 3);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (20, 3, 10);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (21, 1, 3);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (22, 1, 6);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (23, 1, 8);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (24, 1, 10);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (25, 4, 8);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (26, 5, 3);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (27, 7, 3);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
