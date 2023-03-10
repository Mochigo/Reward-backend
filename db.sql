DROP DATABASE IF EXISTS `reward`;

CREATE DATABASE `reward`;

USE `reward`;

CREATE TABLE
    `application` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '申请id',
        `scholarship_item_id` BIGINT NOT NULL COMMENT '奖学金子项id',
        `scholarship_id` BIGINT NOT NUll COMMENT '奖学金id',
        `student_id` BIGINT NOT NULL COMMENT '申请学生id',
        `status` VARCHAR(25) NOT NUll DEFAULT 'PROCESS' COMMENT '申请状态，APPROVE-通过|PROCESS-处理中|FAILURE-驳回',
        PRIMARY KEY(`id`),
        UNIQUE KEY `uniq_idx_stu_scholarship` (
            `student_id`,
            `scholarship_item_id`
        )
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '申请表';

CREATE TABLE
    `certificate` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '荣誉id',
        `application_id` BIGINT NOT NULL COMMENT '申请id',
        -- `student_id` BIGINT NOT NULL COMMENT '所属学生id', 
        `name` VARCHAR(255) NOT NULL COMMENT '荣誉名称',
        `level` CHAR(2) NOT NUll COMMENT '荣誉级别, 校级-01|省级-02|国家级|-03',
        `status` VARCHAR(25) NOT NUll DEFAULT 'PROCESS' COMMENT '申请状态，APPROVED-通过|PROCESS-待处理|REJECTED-驳回',
        `rejected_reason` VARCHAR(255) COMMENT '驳回理由',
        `url` VARCHAR(2048) NOT NULL COMMENT '证明文件连接',
        PRIMARY KEY(`id`),
        KEY `idx_application_id` (`application_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '荣誉表';

CREATE TABLE
    `scholarship` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '奖学金id',
        `name` VARCHAR(255) NOT NULL COMMENT '奖学金名称',
        `college_id` BIGINT NOT NUll COMMENT '学院id',
        `start_time` DATETIME NOT NULL COMMENT '开始时间',
        `end_time` DATETIME NOT NULL COMMENT '结束时间',
        PRIMARY KEY(`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '奖学金表';

CREATE TABLE
    `scholarship_item` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '奖学金子项id',
        `name` VARCHAR(255) NOT NULL COMMENT '奖学金子项名称',
        `scholarship_id` BIGINT NOT NUll COMMENT '奖学金id',
        PRIMARY KEY(`id`),
        KEY `idx_scholarship_id`(`scholarship_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '奖学金子项表';

CREATE TABLE
    `attachment` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '附件子项id',
        `scholarship_id` BIGINT NOT NUll COMMENT '奖学金id',
        `url` VARCHAR(2048) NOT NULL COMMENT '地址',
        PRIMARY KEY(`id`),
        KEY `idx_scholarship_id`(`scholarship_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '附件表';

CREATE TABLE
    `student` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '学生id',
        `score` DOUBLE NOT NUll DEFAULT 0 COMMENT '学生学分绩',
        `uid` VARCHAR(30) NOT NULL COMMENT '学号',
        `password` VARCHAR(255) NOT NULL COMMENT '账号密码，默认为学生学号',
        `college_id` BIGINT NOT NUll COMMENT '学院id',
        PRIMARY KEY(`id`),
        UNIQUE KEY `uniq_idx_uid` (`uid`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '学生表';

INSERT INTO
    `student` (
        `uid`,
        `password`,
        `college_id`
    )
VALUES ('2019213794', '1234', 1);

CREATE TABLE
    `teacher` (
        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '教师id',
        `uid` VARCHAR(30) NOT NULL COMMENT '教师编号',
        `password` VARCHAR(255) NOT NULL COMMENT '账号密码，默认为教师编号',
        `college_id` BIGINT NOT NUll COMMENT '学院id',
        `role` VARCHAR(30) NOT NULL COMMENT '角色',
        PRIMARY KEY(`id`),
        UNIQUE KEY `uniq_idx_uid` (`uid`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '教师表';

CREATE TABLE
    `teacher_student_relationship` (
        `id` BIGINT NOT NULL AUTO_INCREMENT,
        `teacher_id` BIGINT NOT NUll COMMENT '教师id',
        `student_id` BIGINT NOT NUll COMMENT '学生id',
        PRIMARY KEY(`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '教学关系表';

CREATE TABLE
    `college` (
        `id` BIGINT NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(255) NOT NULL COMMENT '学院名称',
        PRIMARY KEY(`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '学院表';

INSERT INTO `college` ( `name` ) VALUES ('计算机学院');