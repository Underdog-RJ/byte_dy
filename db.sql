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

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键，视频唯一id',
      `user_id` bigint(20) NOT NULL COMMENT '视频作者id',
      `play_url` varchar(255) NOT NULL COMMENT '播放url',
      `cover_url` varchar(255) NOT NULL COMMENT '封面url',
      `favorite_count` bigint(20) NOT NULL COMMENT '点赞数量',
      `comment_count` bigint(20) NOT NULL COMMENT '评论数量',
      `publish_time` datetime NOT NULL COMMENT '发布时间戳',
      `title` varchar(255) DEFAULT NULL COMMENT '视频名称',
      PRIMARY KEY (`id`),
      KEY `time` (`publish_time`) USING BTREE,
      KEY `author` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8 COMMENT='视频表';

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
        `user_id` bigint(20) NOT NULL COMMENT '评论发布用户id',
        `video_id` bigint(20) NOT NULL COMMENT '评论视频id',
        `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
        `create_time` datetime NOT NULL COMMENT '评论发布时间',
        `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认评论发布为0，取消后为1',
        PRIMARY KEY (`id`),
        KEY `videoIdIdx` (`video_id`) USING BTREE COMMENT '评论列表使用视频id作为索引-方便查看视频下的评论列表'
) ENGINE=InnoDB AUTO_INCREMENT=1206 DEFAULT CHARSET=utf8 COMMENT='评论表';