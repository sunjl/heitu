DROP USER 'dev'@'localhost';

CREATE USER 'dev'@'localhost'
  IDENTIFIED BY 'dev';

DROP DATABASE IF EXISTS `heitu`;

CREATE DATABASE IF NOT EXISTS `heitu`
  DEFAULT CHARACTER SET `utf8`
  COLLATE `utf8_unicode_ci`;

USE `heitu`;

GRANT ALL PRIVILEGES ON heitu.* TO 'dev'@'localhost'
IDENTIFIED BY 'dev'
WITH GRANT OPTION;

FLUSH PRIVILEGES;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id`           INT(11)      NOT NULL        AUTO_INCREMENT,
  `username`     VARCHAR(255) NULL,
  `password`     VARCHAR(255) NULL,
  `display_name` VARCHAR(255) NULL,
  `phone`        VARCHAR(255) NULL,
  `user_type`    VARCHAR(255) NULL,
  `token`        VARCHAR(255) NULL,
  `create_time`  DATETIME     NULL,
  `update_time`  DATETIME     NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_users_username` (`username`),
  UNIQUE KEY `uk_users_token` (`token`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `hosts`;

CREATE TABLE `hosts` (
  `id`          INT(11)      NOT NULL        AUTO_INCREMENT,
  `group_name`  VARCHAR(255) NULL,
  `name`        VARCHAR(255) NULL,
  `ip`          VARCHAR(255) NULL,
  `processor`   VARCHAR(255) NULL,
  `memory`      BIGINT(20)   NULL,
  `disk`        BIGINT(20)   NULL,
  `platform`    VARCHAR(255) NULL,
  `create_time` DATETIME     NULL,
  `update_time` DATETIME     NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_hosts_name` (`name`),
  UNIQUE KEY `uk_hosts_ip` (`ip`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `project`;

CREATE TABLE `projects` (
  `id`          INT(11)      NOT NULL        AUTO_INCREMENT,
  `group_name`  VARCHAR(255) NULL,
  `name`        VARCHAR(255) NULL,
  `git`         VARCHAR(255) NULL,
  `create_time` DATETIME     NULL,
  `update_time` DATETIME     NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_projects_name` (`name`),
  UNIQUE KEY `uk_projects_git` (`git`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `configs`;

CREATE TABLE `configs` (
  `id`          INT(11)      NOT NULL        AUTO_INCREMENT,
  `project_id`  INT(11)      NULL,
  `version`     VARCHAR(255) NULL,
  `environment` VARCHAR(255) NULL,
  `file_name`   VARCHAR(255) NULL,
  `content`     TEXT         NULL,
  `create_time` DATETIME     NULL,
  `update_time` DATETIME     NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_configs` (`project_id`, `version`, `environment`),
  CONSTRAINT `fk_configs_project_id` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `tasks`;

CREATE TABLE `tasks` (
  `id`          INT(11)      NOT NULL        AUTO_INCREMENT,
  `host_name`   VARCHAR(255) NULL,
  `ip`          VARCHAR(255) NULL,
  `name`        VARCHAR(255) NULL,
  `content`     TEXT         NULL,
  `status`      VARCHAR(255) NULL,
  `create_time` DATETIME     NULL,
  `update_time` DATETIME     NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `services`;

CREATE TABLE `services` (
  `id`          INT(11)      NOT NULL        AUTO_INCREMENT,
  `group_name`  VARCHAR(255) NULL,
  `project_id`  INT(11)      NULL,
  `name`        VARCHAR(255) NULL,
  `protocol`    VARCHAR(255) NULL,
  `host_name`   VARCHAR(255) NULL,
  `ip`          VARCHAR(255) NULL,
  `port`        INT(11)      NULL,
  `create_time` DATETIME     NULL,
  `update_time` DATETIME     NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_services` (`ip`, `port`),
  CONSTRAINT `fk_services_project_id` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
