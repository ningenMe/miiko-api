CREATE DATABASE `compro`;
USE `compro`;

CREATE TABLE `health`
(
    `id` integer NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `category`
(
    `category_id`           varchar(255) NOT NULL,
    `category_display_name` varchar(511) NOT NULL,
    `category_system_name`  varchar(511) NOT NULL,
    `category_order`        int(11)      NOT NULL,
    `created_time`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`category_id`),
    UNIQUE KEY `category_system_name_unique` (`category_system_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `topic`
(
    `topic_id`           varchar(255) NOT NULL,
    `category_id`        varchar(255) NOT NULL,
    `topic_display_name` varchar(511) NOT NULL,
    `topic_order`        int(11)      NOT NULL,
    `problem_count`      int(11)      NOT NULL DEFAULT 0,
    `created_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`topic_id`),
    FOREIGN KEY `category_id_foreign` (`category_id`) REFERENCES `category` (`category_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

# ALTER TABLE topic
#     ADD `problem_count` int(11) NOT NULL DEFAULT 0 AFTER `topic_order`;

CREATE TABLE `problem`
(
    `problem_id`           varchar(255)                            NOT NULL,
    `url`                  varchar(511)                            NOT NULL,
    `problem_display_name` varchar(511) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `estimation`           int(10) unsigned                        NOT NULL DEFAULT '0',
    `tag_list`             json                                    NOT NULL,
    `created_time`         timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time`         timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`problem_id`),
    UNIQUE `url` (`url`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `relation_topic_problem`
(
    `topic_id`     varchar(255) NOT NULL,
    `problem_id`   varchar(255) NOT NULL,
    `created_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`topic_id`, `problem_id`),
    KEY `problem_id` (`problem_id`),
    FOREIGN KEY `topic_id_foreign` (`topic_id`) REFERENCES `topic` (`topic_id`),
    FOREIGN KEY `problem_id_foreign` (`problem_id`) REFERENCES `problem` (`problem_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

# CREATE TABLE `atcoder_user`
# (
#     `atcoder_id`             varchar(255) NOT NULL,
#     `topcoder_id`            varchar(255)          DEFAULT NULL,
#     `codeforces_id`          varchar(255)          DEFAULT NULL,
#     `yukicoder_id`           varchar(255)          DEFAULT NULL,
#     `rank`                   int(255)              DEFAULT NULL,
#     `country`                varchar(255)          DEFAULT NULL,
#     `affiliation`            varchar(255)          DEFAULT NULL,
#     `current_rate`           int(255)              DEFAULT NULL,
#     `highest_rate`           int(255)              DEFAULT NULL,
#     `current_performance`    int(255)              DEFAULT NULL,
#     `highest_performance`    int(255)              DEFAULT NULL,
#     `first_participate_time` timestamp    NULL     DEFAULT NULL,
#     `last_participate_time`  timestamp    NULL     DEFAULT NULL,
#     `deleted_time`           timestamp    NULL     DEFAULT NULL,
#     `created_time`           timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     `updated_time`           timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#     PRIMARY KEY (`atcoder_id`),
#     KEY `updated_time` (`updated_time`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8;
#
# CREATE TABLE `atcoder_user_contest`
# (
#     `atcoder_id`          varchar(255) NOT NULL,
#     `contest_name_en`     varchar(255) NOT NULL,
#     `is_rated`            tinyint(1)            DEFAULT NULL,
#     `place`               int(11)               DEFAULT NULL,
#     `old_rating`          int(11)               DEFAULT NULL,
#     `new_rating`          int(11)               DEFAULT NULL,
#     `performance`         int(11)               DEFAULT NULL,
#     `contest_type`        varchar(255)          DEFAULT NULL,
#     `contest_screen_name` varchar(255)          DEFAULT NULL,
#     `contest_name`        varchar(255)          DEFAULT NULL,
#     `end_time`            timestamp    NULL     DEFAULT NULL,
#     `deleted_time`        timestamp    NULL     DEFAULT NULL,
#     `created_time`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     `updated_time`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#     PRIMARY KEY (`atcoder_id`, `contest_name_en`),
#     KEY `contest_name_en` (`contest_name_en`),
#     KEY `end_time` (`end_time`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8;
#
# CREATE TABLE `atcoder_user_history`
# (
#     `atcoder_id`   varchar(256) NOT NULL,
#     `contest_id`   varchar(256) NOT NULL,
#     `is_rated`     tinyint(1)            DEFAULT NULL,
#     `place`        int(11)               DEFAULT NULL,
#     `old_rating`   int(11)               DEFAULT NULL,
#     `new_rating`   int(11)               DEFAULT NULL,
#     `performance`  int(11)               DEFAULT NULL,
#     `deleted_time` timestamp    NULL     DEFAULT NULL,
#     `created_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     `updated_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#     PRIMARY KEY (`atcoder_id`, `contest_id`),
#     KEY `contest_id` (`contest_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8;
