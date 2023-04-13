CREATE DATABASE /*!32312 IF NOT EXISTS*/`casbin_demo` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `casbin_demo`;

/*Table structure for table `casbin_rule` */

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) NOT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8;

/*Data for the table `casbin_rule` */

insert  into `casbin_rule`(`id`,`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) values 

(3,'p','role::admin','admin::resource','read-write','','',''),

(5,'p','role::user','user::resource','read-write','','',''),

(57,'g','test1','role::user','','','',''),

(59,'g','role::admin','role::user','','','',''),

(63,'g','test2','role::admin',NULL,NULL,NULL,NULL);

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;

/*Data for the table `user` */

insert  into `user`(`id`,`username`,`password`,`email`,`phone`) values 

(36,'test1','123',NULL,NULL),

(38,'test2','123',NULL,NULL);