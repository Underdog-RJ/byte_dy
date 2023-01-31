DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
     `user_id` bigint(20) NOT NULL COMMENT '点赞用户id',
     `video_id` bigint(20) NOT NULL COMMENT '被点赞的视频id',
     `create_time` datetime NOT NULL COMMENT '创建时间',
     `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认点赞为0，取消赞为1',
     PRIMARY KEY (`id`),
     UNIQUE KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
     KEY `userIdIdx` (`user_id`) USING BTREE,
     KEY `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1229 DEFAULT CHARSET=utf8 COMMENT='点赞表';