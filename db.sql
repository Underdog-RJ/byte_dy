DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
     `user_id` bigint(20) NOT NULL COMMENT '点赞用户id',
     `video_id` bigint(20) NOT NULL COMMENT '被点赞的视频id',
     `create_time` datetime NOT NULL COMMENT '创建时间',
     `update_time` datetime NOT NULL COMMENT '更新时间',
     `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认点赞为0，取消赞为1',
     PRIMARY KEY (`id`),
     UNIQUE KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
     KEY `userIdIdx` (`user_id`) USING BTREE,
     KEY `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1229 DEFAULT CHARSET=utf8 COMMENT='点赞表';

drop TABLE if exists `videos`;
create table videos
(
    id                 bigint auto_increment comment '自增主键，视频唯一id'
        primary key,
    user_id            bigint           not null comment '视频作者id',
    play_url           varchar(255)     null comment '播放url',
    cover_url          varchar(255)     null comment '封面url',
    favorite_count     bigint default 0 not null comment '点赞数量',
    comment_count      bigint default 0 not null comment '评论数量',
    publish_time       datetime         not null comment '发布时间戳',
    title              varchar(255)     null comment '视频名称',
    video_status       int    default 0 null comment '视频状态',
    video_size         bigint           null,
    video_md5          varchar(255)     null,
    video_ext          varchar(255)     null,
    original_file_path varchar(255)     null
)
    comment '视频表' charset = utf8mb3;

create index author
    on videos (user_id);

create index time
    on videos (publish_time);

create table users
(
    id              int unsigned auto_increment
        primary key,
    created_at      datetime     null,
    updated_at      datetime     null,
    deleted_at      datetime     null,
    user_name       varchar(255) null,
    password_digest varchar(255) null,
    follow_count    int          null,
    follower_count  int          null,
    constraint user_name
        unique (user_name)
);

create index idx_users_deleted_at
    on users (deleted_at);


create table follower
(
    id          int auto_increment
        primary key,
    follower_id int      null comment '被关注的人的id',
    followee_id int      null comment '关注的人的id',
    create_time datetime null comment '创建时间',
    update_time datetime null comment '更新时间'
);


