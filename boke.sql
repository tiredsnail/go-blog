-- 创建WYcartoon库
create database if not exists blog_baiwuya;

-- 选择数据库
use blog_baiwuya;
--
-- 表的结构 `cat`
--
create table if not exists `blog_type`(
  `type_id` int unsigned auto_increment primary key,
  `name` varchar(20) unique NOT NULL COMMENT '分类名',
  `order` tinyint(2) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0-隐藏 1-显示 2-header',
  `title` varchar(50) comment '网站title',
  `keywords` varchar(100) comment '网站关键词',
  `description` varchar(255) comment '网站说明'
)engine=innodb default charset=utf8 comment '分类表';

--
-- 表的结构 `blog_article`
--
CREATE TABLE IF NOT EXISTS `blog_article` (
  `article_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(20) comment '分类名',
  `headline` varchar(45) DEFAULT '' comment '标题',
  `summary` varchar(255) comment '摘要',
  `content` text,
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned DEFAULT '0',
  `comm` smallint(5) unsigned NOT NULL DEFAULT '0' comment '评论数量',
  `uv` int(11) NOT NULL DEFAULT 0 COMMENT '访客量',
  `pv` int(11) NOT NULL DEFAULT 0 COMMENT '点击量',
  `state` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-隐藏 1-显示 2-未提交',
  PRIMARY KEY (`article_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='文章表';

--
-- 表的结构 `tag`
--

CREATE TABLE IF NOT EXISTS `tag` (
  `tag_id` int unsigned auto_increment primary key,
  `art_id` int(10) unsigned NOT NULL DEFAULT '0',
  `tag` char(10) NOT NULL DEFAULT '0',
  KEY `at` (`art_id`,`tag`),
  KEY `ta` (`tag`,`art_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `comment`
--
CREATE TABLE IF NOT EXISTS `blog_comment` (
  `comment_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned NOT NULL,
  `nick` varchar(20) NOT NULL  COLLATE utf8_unicode_ci DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '' comment '邮箱',
  `url` varchar(100) NOT NULL DEFAULT '' comment '用户主页',
  `content` varchar(1000) NOT NULL COLLATE utf8_unicode_ci DEFAULT '',
  `superior` varchar(20) DEFAULT ''  COLLATE utf8_unicode_ci comment '父级(回复)昵称',
  `ip` varchar(20) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`comment_id`),
  KEY `article_id` (`article_id`),
  KEY `ip` (`ip`),
  KEY `nick` (`nick`),
  KEY `superior` (`superior`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL DEFAULT '',
  `nick` char(20) NOT NULL DEFAULT '',
  `email` char(30) NOT NULL DEFAULT '',
  `password` char(32) NOT NULL DEFAULT '',
  `lastlogin` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;