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
    `address` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '用户住址',
    primary key (`id`) using btree
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '用户表';

drop table if exists `lifts`;
create table `lifts` (
    `id` int(10) unsigned not null auto_increment,
    `create_at` timestamp(0) null default null,
    `update_at` timestamp(0) null default null,
    `deleted_at` timestamp(0) null default null,
    `nick_name` varchar(255) character set utf8 collate utf8_general_ci null default null comment '电梯别名',
    `code` varchar(255) character set utf8 collate utf8_general_ci null default null comment '电梯编号',
    `installer_id` int(10) unsigned null default null comment '电梯安装公司id',
    `maintainer_id` int(10) unsigned null default null comment '电梯维保公司id',
    `checker_id` int(10) unsigned null default null comment '电梯年检公司id',
    `owner_id` int(10) unsigned null default null comment '电梯使用公司id',
    `factory_time` timestamp(0) null default null comment '电梯出厂时间',
    `install_time` timestamp(0) null default null comment '电梯安装时间',
    `check_time` timestamp(0) null default null comment '电梯年检时间',
    `lift_model_id` int(10) unsigned null default null comment '电梯型号',
    `category_id` int(10) not null comment '电梯类别',
    `floor_count` int not null comment '总楼层',
    `latitude` decimal(10,7) null default null comment '电梯位置地理经度',
    `longitude` decimal(10,7) null default null comment '电梯位置地理纬度',
    `address_id` int(10) null default null comment '地址id',
    `region_id` int null default null comment '区域id',
    `building` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '楼栋',
    `cell` int null default null comment '单元',
    `ad_device_id` int(10) unsigned null default null comment '广告机设备id'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯表';

drop table if exists `user_lifts`;
create table `user_lifts` (
   `id`         int(10) unsigned not null auto_increment,
   `create_at`  timestamp(0)     null default null,
   `update_at`  timestamp(0)     null default null,
   `deleted_at` timestamp(0)     null default null,
   `user_id` int(10) unsigned not null comment '用户',
   `lift_id` int(10) unsigned not null comment '电梯',
   `category_id` int(10) not null comment '用户电梯关系类型'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '用户电梯关系表';

drop table if exists `addresss`;
create table `addresss` (
    `id`         int(10) unsigned not null auto_increment,
    `create_at`  timestamp(0)     null default null,
    `update_at`  timestamp(0)     null default null,
    `deleted_at` timestamp(0)     null default null,
    `address_name` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '地址，精确到类似小区级别'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '地址表';


drop table if exists `regions`;
create table `regions` (
    `id`         int(10) unsigned not null auto_increment,
    `create_at`  timestamp(0)     null default null,
    `update_at`  timestamp(0)     null default null,
    `deleted_at` timestamp(0)     null default null,
    `province` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '省份',
    `city` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '城市',
    `district` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '行政区'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '区域表';

drop table if exists `lift_models`;
create table `lift_models` (
   `id`         int(10) unsigned not null auto_increment,
   `create_at`  timestamp(0)     null default null,
   `update_at`  timestamp(0)     null default null,
   `deleted_at` timestamp(0)     null default null,
   `factory_id` int(10) unsigned null default null comment '电梯生产商id',
   `brand` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '电梯品牌',
   `modal` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '电梯型号',
   `load` int null default null comment '电梯载重'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯型号表';


drop table if exists `companys`;
create table `companys` (
   `id`         int(10) unsigned not null auto_increment,
   `create_at`  timestamp(0)     null default null,
   `update_at`  timestamp(0)     null default null,
   `deleted_at` timestamp(0)     null default null,
   `full_name` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '公司全称',
   `alias` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '公司简称',
   `legal_person` varchar(50) character set utf8 collate utf8_general_ci null default null comment '法人信息',
   `phone_number` varchar(20) character set utf8 collate utf8_general_ci null default null comment '联系方式',
   `status` varchar(20) character set utf8 collate utf8_general_ci null default null comment '公司状态',
   `reg_code` varchar(20) character set utf8 collate utf8_general_ci null default null comment '工商注册号',
   `org_code` varchar(20) character set utf8 collate utf8_general_ci null default null comment '组织机构号',
   `credit_code` varchar(20) character set utf8 collate utf8_general_ci null default null comment '统一信用代码',
   `tax_code` varchar(20) character set utf8 collate utf8_general_ci null default null comment '纳税人识别号',
   `address` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '注册地址',
   `category_id` int(10) not null comment '电梯类别'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '公司表';


drop table if exists `category_subject`;
create table `category_subject` (
    `id`         int(10) unsigned not null auto_increment,
    `create_at`  timestamp(0)     null default null,
    `update_at`  timestamp(0)     null default null,
    `deleted_at` timestamp(0)     null default null,
    `subject_name` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '类别主体名'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '类别主体表，电梯/公司/记录/......';


drop table if exists `categorys`;
create table `categorys` (
  `id`         int(10) unsigned not null auto_increment,
  `create_at`  timestamp(0)     null default null,
  `update_at`  timestamp(0)     null default null,
  `deleted_at` timestamp(0)     null default null,
  `category_subject_id` int(10) not null comment '类别主体id',
  `category_name` varchar(50)  character set utf8 collate utf8_general_ci null default null comment '公司类别名'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯类别表';

drop table if exists `lift_changes`;
create table `lift_changes` (
   `id`         int(10) unsigned not null auto_increment,
   `create_at`  timestamp(0)     null default null,
   `update_at`  timestamp(0)     null default null,
   `deleted_at` timestamp(0)     null default null,
   `lift_id` int(10) unsigned not null comment '电梯id',
   `content` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '变更内容'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯变更表';

drop table if exists `lift_records`;
create table `lift_records` (
    `id`         int(10) unsigned not null auto_increment,
    `create_at`  timestamp(0)     null default null,
    `update_at`  timestamp(0)     null default null,
    `deleted_at` timestamp(0)     null default null,
    `lift_id` int(10) unsigned not null comment '电梯id',
    `category_id` int(10) unsigned not null comment '记录类别',
    `images` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '图片记录',
    `content` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '文字记录',
    `start_time` timestamp(0) null default null comment '开始时间',
    `end_time` timestamp(0) null default null comment '结束时间',
    `worker_id` int(10) unsigned not null comment '操作人员',
    `recorder_id` int(10) unsigned not null comment '记录人员'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯操作记录';

drop table if exists `lift_troubles`;
create table `lift_troubles` (
    `id`         int(10) unsigned not null auto_increment,
    `create_at`  timestamp(0)     null default null,
    `update_at`  timestamp(0)     null default null,
    `deleted_at` timestamp(0)     null default null,
    `lift_id` int(10) unsigned not null comment '电梯id',
    `from_category_id` int(10) unsigned not null comment '故障来源类别',
    `start_time` timestamp(0) null default null comment '故障开始时间',
    `start_user_id` int(10) null default null comment '发起故障人员',
    `response_time` timestamp(0) null default null comment '故障响应时间',
    `response_user_id` int(10) null default null comment '故障响应人员',
    `scene_time` timestamp(0) null default null comment '达到现场时间',
    `scene_user_id` int(10) null default null comment '达到现场人员',
    `fix_time` timestamp(0) null default null comment '解除故障时间',
    `fix_user_id` int(10) null default null comment '解除故障人员',
    `fix_category_id` int(10) unsigned null comment '解除故障方式类别',
    `reason_category_id` int(10) unsigned null comment '故障原因类别',
    `content` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '故障详情',
    `progress` int unsigned not null default 1 comment '故障进度',
    `recorder_id` int(10) unsigned not null comment '记录人员',
    `feedback_content` varchar(255)  character set utf8 collate utf8_general_ci null default null comment '反馈内容',
    `feedback_rate` int unsigned not null default 100 comment '反馈评分'
) engine = InnoDB auto_increment = 100 character set = utf8 collate = utf8_general_ci row_format = compact comment '电梯故障记录';

