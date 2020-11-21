DROP TABLE IF EXISTS tb_users;
CREATE TABLE tb_users (
	`account_id` INT NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    `real_name` VARCHAR(50) NOT NULL DEFAULT '',
    `phone_number` VARCHAR(50) NULL DEFAULT '',
    `address` VARCHAR(200) NULL DEFAULT '',
    `major` VARCHAR(200) NULL DEFAULT '',
    `gender` int not null default 1,
    `age` int not null default 1,
    `status` tinyint default 0,
    `active_str` VARCHAR(30) NULL DEFAULT '',
    `user_type` int not null default 1,
    PRIMARY KEY(`account_id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
INSERT INTO tb_users (account_id, user_name, `password`, `status`, user_type, active_str) VALUES (1,'admin','35cf546482e8847c27fe85a2b375f1fb', 1, 0,'PIARiJrJ');

DROP TABLE IF EXISTS tb_users_roles;
CREATE TABLE tb_users_roles (
	`row_id` INT NOT NULL AUTO_INCREMENT,
    `account_id` INT NOT NULL,
    `role_id` INT NOT NULL,
    PRIMARY KEY(`row_id`)
)DEFAULT CHARSET=utf8;

INSERT INTO tb_users_roles (`account_id`, `role_id`) VALUES (1,1);

DROP TABLE IF EXISTS tb_functions;
CREATE TABLE tb_functions (
	`function_id` INT NOT NULL AUTO_INCREMENT,
	`number` INT NOT NULL,
	`order` int NOT NULL,
	`name` VARCHAR(200) NULL,
	`path` VARCHAR(200) NULL,
	`type` TINYINT NOT NULL DEFAULT 1,
	`parent_function_id` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`function_id`)
)DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_functions_items;
CREATE TABLE tb_functions_items (
	`function_item_id` INT NOT NULL AUTO_INCREMENT,
	`function_id` INT NOT NULL DEFAULT 0,
	`item_number` INT NOT NULL DEFAULT 0,
	`item_name` VARCHAR(200) NULL,
  PRIMARY KEY (`function_item_id`)
)DEFAULT CHARSET=utf8;

# ---增加默认的
INSERT INTO tb_functions (`function_id`, `name`, `order`, `path`, `type`, `parent_function_id`) VALUES (10,'首页', 0, '', 0, 0);

INSERT INTO tb_functions (`function_id`, `name`, `order`, `path`, `type`, `parent_function_id`) VALUES (1000,'权限管理', 0, '', 0, 0);
INSERT INTO tb_functions (`function_id`, `name`, `order`, `path`, `type`, `parent_function_id`) VALUES (1001,'用户管理', 1, '', 0, 1000);
INSERT INTO tb_functions (`function_id`, `name`, `order`, `path`, `type`, `parent_function_id`) VALUES (1002,'角色管理', 2, '', 0, 1000);
INSERT INTO tb_functions (`function_id`, `name`, `order`, `path`, `type`, `parent_function_id`) VALUES (1003,'功能点管理',3, '', -1, 1000);

INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10011,1001, '查看用户');
INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10012,1001, '新增用户');
INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10013,1001, '编辑用户');
INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10014,1001, '删除用户');

-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10021,1001, '查看角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10022,1001, '新增角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10023,1001, '编辑角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10024,1001, '删除角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10025,1001, '查看角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10026,1001, '新增角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10027,1001, '编辑角色');
-- INSERT INTO tb_functions_items (`function_item_id`, `function_id`, `item_name`) VALUES (10028,1001, '删除角色');

DROP TABLE IF EXISTS tb_roles;
CREATE TABLE tb_roles (
	`role_id` INT NOT NULL AUTO_INCREMENT,
	`code` VARCHAR(10) NOT NULL,
	`name` VARCHAR(200) NULL,
	`type` TINYINT NOT NULL DEFAULT 1,
	`status` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`role_id`)
)DEFAULT CHARSET=utf8;

INSERT INTO tb_roles (`role_id`, `code`, `name`, `type`, `status` ) VALUES (1, 'admin', 'admin', 0, 1);

DROP TABLE IF EXISTS tb_roles_functions;
CREATE TABLE tb_roles_functions (
    `row_id` INT NOT NULL AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `function_id` INT NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`row_id`)
)DEFAULT CHARSET=utf8;

INSERT INTO tb_roles_functions (`role_id`, `function_id`, `status`) VALUES (1,10,1);

INSERT INTO tb_roles_functions (`role_id`, `function_id`, `status`) VALUES (1,1000,1);
INSERT INTO tb_roles_functions (`role_id`, `function_id`, `status`) VALUES (1,1001,1);
INSERT INTO tb_roles_functions (`role_id`, `function_id`, `status`) VALUES (1,1002,1);
INSERT INTO tb_roles_functions (`role_id`, `function_id`, `status`) VALUES (1,1003,1);

DROP TABLE IF EXISTS tb_roles_items;
CREATE TABLE tb_roles_items (
    `row_id` INT NOT NULL AUTO_INCREMENT,
    `role_id` INT NOT NULL,
    `item_id` INT NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`row_id`)
)DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS tb_works_records;
CREATE TABLE tb_works_records(
	`record_id` INT NOT NULL AUTO_INCREMENT,
    `record_name` VARCHAR(50) NOT NULL,
    `account_id` INT(11) NOT NULL,
    `record_type` VARCHAR(50) NOT NULL DEFAULT '',
    `record_body` VARCHAR(1000) NULL DEFAULT '',
    `level` VARCHAR(50) NULL DEFAULT '',
    `record_address` VARCHAR(200) NULL DEFAULT '',
    `create_time` DATETIME NULL DEFAULT '2020-10-10',
    `status` tinyint default 0,
    PRIMARY KEY(`record_id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

DROP TABLE IF EXISTS tb_tasks_records;
CREATE TABLE tb_tasks_records(
	`task_id` INT NOT NULL AUTO_INCREMENT,
    `task_name` VARCHAR(50) NOT NULL,
    `account_id` INT(11) NOT NULL,
    `task_address` VARCHAR(200) NOT NULL DEFAULT '',
    `task_body` VARCHAR(1000) NULL DEFAULT '',
    `level` VARCHAR(50) NULL DEFAULT '',
    `create_time` DATETIME NULL DEFAULT '2020-10-10',
    `status` tinyint default 0,
    PRIMARY KEY(`task_id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;