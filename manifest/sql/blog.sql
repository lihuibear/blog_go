create table admin
(
    id         int unsigned auto_increment
        primary key,
    username   varchar(100) not null comment '用户名',
    password   varchar(100) not null comment '密码',
    nickname   varchar(100) null comment '昵称',
    avatar     longtext     null comment '头像，base64',
    register   datetime     not null comment '注册时间',
    salt       char(32)     not null comment '随机盐值',
    last_login datetime     not null comment '最后登录时间',
    constraint username唯一
        unique (username)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table article
(
    id          int unsigned auto_increment
        primary key,
    grp_id      int unsigned                   not null comment '分组id',
    title       varchar(100)                   not null comment '标题',
    author      varchar(30)                    null comment '作者',
    thumb       varchar(200)                   null comment '图片地址',
    tags        varchar(255)                   null comment '标签，依英文逗号隔开',
    description varchar(255)                   null comment '简介',
    content     longtext                       null comment '内容',
    `order`     smallint           default 0   not null comment '排序，越大越靠前',
    ontop       tinyint unsigned   default '0' not null comment '是否置顶',
    onshow      tinyint unsigned   default '1' not null comment '是否显示',
    hist        mediumint unsigned default '0' not null comment '点击数',
    post        mediumint unsigned default '0' not null comment '评论数',
    created_at  datetime                       not null comment '创建时间',
    updated_at  datetime                       not null comment '更新时间',
    deleted_at  datetime                       null comment '删除时间',
    lasted_at   datetime                       null comment '最后浏览时间'
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table article_grp
(
    id          mediumint unsigned auto_increment
        primary key,
    name        varchar(50)                  not null comment '名称',
    tags        varchar(255)                 null comment '标签，依英文逗号隔开',
    description varchar(255)                 null comment '简介',
    onshow      tinyint unsigned default '1' not null comment '是否显示',
    `order`     smallint         default 0   not null comment '排序，越大越靠前'
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table link
(
    id          int unsigned auto_increment
        primary key,
    name        varchar(20)  not null comment '名称',
    description varchar(100) null comment '描述',
    link        varchar(255) not null comment '链接'
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table reading
(
    id          int unsigned auto_increment
        primary key,
    name        varchar(100) not null comment '书名',
    author      varchar(100) not null comment '作者',
    status      int unsigned not null comment '状态: 1弃读 2完结 9在读',
    finished_at date         null comment '读完时间'
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table reply
(
    id         int unsigned auto_increment
        primary key,
    aid        int          not null comment '文章id',
    rid        int          not null comment '根回复id，一个根可视为一个楼层',
    pid        int          not null comment '回复的id',
    p_name     varchar(20)  null comment '回复的名称',
    email      varchar(50)  not null comment '回复人邮箱',
    name       varchar(20)  not null comment '回复人名称',
    site       varchar(100) null comment '回复人网站',
    content    longtext     not null comment '回复内容',
    status     tinyint(1)   not null comment '状态： 1待审核 2审核通过 3审核失败',
    reason     varchar(50)  null comment '审核失败原因',
    created_at datetime     not null comment '创建时间',
    updated_at datetime     not null comment '更新时间',
    deleted_at datetime     null comment '删除时间'
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table sentence
(
    id       int unsigned auto_increment
        primary key,
    book_id  int unsigned not null,
    sentence longtext     null
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table sentence_tag
(
    s_id int unsigned not null comment '句子id',
    t_id int unsigned not null comment 'tag id',
    constraint s_id
        unique (s_id, t_id)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table tag
(
    id     int unsigned auto_increment
        primary key,
    grp_id int unsigned not null comment '分组id',
    name   varchar(30)  not null comment '标签名称',
    constraint name
        unique (name)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

create table tag_grp
(
    id   int unsigned auto_increment
        primary key,
    name varchar(30) null comment '标签分组名称',
    constraint name
        unique (name)
)
    collate = utf8mb4_general_ci
    row_format = DYNAMIC;

