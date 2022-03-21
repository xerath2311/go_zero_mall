CREATE table `user`(
`id` bigint unsigned not null auto_increment,
`name` varchar(255) not null default '' comment '用户姓名',
`gender` tinyint(3) unsigned not null default '0' comment '用户性别',
`mobile` varchar(255) not null default '' comment '用户电话',
`password` varchar(255) not null default '' comment '用户密码',
`create_time` timestamp null default current_timestamp,
`update_time` timestamp null default current_timestamp on update current_timestamp,
primary key (`id`),
unique key `idx_mobile_unique` (`mobile`)
) engine=innodb default charset=utf8mb4;