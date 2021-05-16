use go;
create table `user` (
    `id` bigint(20) not null auto_increment,
    `user_id` bigint(20) not null ,
    `username` varchar(64) not null ,
    `password` varchar(64) not null ,
    `email` varchar(64) ,
    `gender` tinyint(4) not null default '0',
    `create_time` timestamp null default CURRENT_TIMESTAMP ,
    `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP ,
    primary key (`id`),
    unique key `idx_username` (`username`) using btree,
    unique key `idx_user_id` (`user_id`) using btree
)ENGINE=InnoDB,charset = utf8mb4