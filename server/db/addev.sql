

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

drop table if exists `user_admins`;
create table `user_admins` (
    `id` int(10) unsigned not null auto_increment,
    `create_at` timestamp(0) null default null,
    `update_at` timestamp(0) null default null,
    `deleted_at` timestamp(0) null default null,
    `uuid` varbinary(255) null  default null comment 'uuid',
    `phone_number` varchar(20) not null comment '登陆手机号',
    `password` varchar(255) character set utf8 collate utf8_general_ci null default null comment '登陆密码',
    `real_name` varchar(255) character set utf8 collate utf8_general_ci null default null comment '用户真名',
    `nick_name` varchar(255) character set utf8 collate utf8_general_ci null default null comment '用户昵称',
    `avatar` varchar(255) character set utf8 collate utf8_general_ci null default null comment '用户头像',
    `role_id` int(10) unsigned not null comment '角色id',
    `role_name` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '角色名称',
    `company_id` int(10) unsigned not null comment '用户所属公司id',
    `company_name` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '用户所属公司名称',
    `address` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '用户住址',
    primary key (`id`) using btree
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact ;

