
1. ###### users
```sql
-- blog_db.users definition

CREATE TABLE `users` (
                         `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '键',
                         `username` varchar(255) NOT NULL COMMENT '用户名',
                         `password` varchar(255) NOT NULL COMMENT '密码',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
2. ###### blogs_content
```sql
-- blog_db.blogs_content definition

CREATE TABLE `blogs_content` (
                                 `id` int NOT NULL,
                                 `groupId` int DEFAULT NULL COMMENT '分组Id',
                                 `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文章标题',
                                 `content` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文章内容',
                                 `visitors` int DEFAULT '0' COMMENT '浏览量',
                                 `createTime` datetime DEFAULT NULL COMMENT '创建时间',
                                 `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
                                 `isShow` tinyint(1) DEFAULT '1' COMMENT '是否显示',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
3. ###### blog_groups
```sql
-- blog_db.blog_groups definition

CREATE TABLE `blog_groups` (
                               `id` int NOT NULL AUTO_INCREMENT COMMENT '键',
                               `group` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '分组名称',
                               `isShow` tinyint(1) DEFAULT '1' COMMENT '是否显示',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```