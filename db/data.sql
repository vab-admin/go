BEGIN;
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (11, 3);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (14, 4);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (15, 6);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (15, 7);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (15, 8);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (16, 5);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (16, 8);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (17, 13);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (18, 14);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (19, 16);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (19, 17);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (19, 18);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (20, 15);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (20, 18);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (21, 8);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (22, 9);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (23, 11);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (23, 12);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (23, 13);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (24, 10);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (24, 13);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (25, 19);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (26, 20);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (27, 22);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (27, 23);
INSERT INTO `admin_rule_apis` (`rule_id`, `api_id`) VALUES (28, 21);
COMMIT;

BEGIN;
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (1, 0, '/system', 'System', 'Layout', '', '系统管理', 0, 0, 'settings-2-line', 0, 0, 0, '', 0, '', 0, 2, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (2, 1, 'admin/users', 'AdminUser', '@/views/auth/admin/index', '', '管理员管理', 0, 0, 'user-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (3, 1, 'auth/rules', 'AdminRule', '@/views/auth/rule/index', '', '菜单管理', 0, 0, 'menu-fill', 0, 0, 0, '', 0, '', 0, 2, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (4, 1, 'auth/roles', 'AdminRole', '@/views/auth/role/index', '', '角色管理', 0, 0, 'group-line', 0, 0, 0, '', 0, '', 0, 1, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (5, 1, 'auth/api', 'SystemApi', '@/views/auth/api/index', '', '接口管理', 0, 0, 'align-justify', 0, 0, 0, '', 0, '', 0, 3, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (11, 2, '', '', '', '', '查询', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminUser.query');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (14, 2, '', '', '', '', '删除', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminUser.del');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (15, 2, '', '', '', '', '编辑', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminUser.edit');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (16, 2, '', '', '', '', '添加', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminUser.add');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (17, 3, '', '', '', '', '查询', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRule.query');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (18, 3, '', '', '', '', '删除', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRule.del');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (19, 3, '', '', '', '', '编辑', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRule.edit');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (20, 3, '', '', '', '', '添加', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRule.add');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (21, 4, '', '', '', '', '查询', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRole.query');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (22, 4, '', '', '', '', '删除', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRole.del');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (23, 4, '', '', '', '', '编辑', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRole.edit');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (24, 4, '', '', '', '', '添加', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'adminRole.add');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (25, 5, '', '', '', '', '查询', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'systemApi.query');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (26, 5, '', '', '', '', '删除', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'systemApi.del');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (27, 5, '', '', '', '', '编辑', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'systemApi.edit');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (28, 5, '', '', '', '', '添加', 0, 0, '', 0, 0, 0, '', 0, '', 0, 0, 1, 2, 'systemApi.add');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (30, 0, '/cp', 'cp', '', '', 'CP管理', 0, 0, 'user-3-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (31, 0, '/user', 'User', '', '', '客户管理', 0, 0, 'user-star-line', 0, 0, 0, '', 0, '', 0, 1, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (32, 30, '', 'CpMobile', '', '', '手机号码管理', 0, 0, 'smartphone-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (33, 30, '', 'CpMobileTotal', '', '', '号码统计', 0, 0, 'list-unordered', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (34, 30, '', 'CpSms', '', '', '短信管理', 0, 0, 'mail-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (35, 30, '', 'Cp', '', '', 'CP管理', 0, 0, 'list-ordered', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (36, 31, '', 'UserMobile', '', '', '手机号码管理', 0, 0, 'smartphone-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (38, 31, '', 'UserSms', '', '', '短信管理', 0, 0, 'mail-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (39, 31, '', 'User', '', '', '客户管理', 0, 0, 'list-ordered', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (40, 0, '/log', 'Log', '', '', '日志管理', 0, 0, 'clipboard-line', 0, 0, 0, '', 0, '', 0, 3, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (41, 0, '', 'Total', '', '', '数据统计', 0, 0, 'line-chart-line', 0, 0, 0, '', 0, '', 0, 4, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (42, 30, '', 'CpResponse', '', '', 'CP响应', 0, 0, 'database-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (43, 31, '', 'UserResponse', '', '', '客户响应', 0, 0, 'database-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (44, 40, 'oper', 'OperLog', '', '', '操作日志', 0, 0, 'lightbulb-flash-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (45, 40, 'login', 'LoginLog', '', '', '登录日志', 0, 0, 'article-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
INSERT INTO `admin_rules` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `hidden`, `level_hidden`, `icon`, `no_keep_alive`, `no_closable`, `no_column`, `badge`, `tab_hidden`, `target`, `dot`, `sort`, `status`, `type`, `code`) VALUES (46, 31, 'total', 'UserTotal', '', '', '客户数据统计', 0, 0, 'article-line', 0, 0, 0, '', 0, '', 0, 0, 1, 1, '');
COMMIT;

BEGIN;
INSERT INTO `admin_users` (`id`, `mobile`, `account`, `nickname`, `password`, `created_at`, `updated_at`) VALUES (1, '17606518462', 'admin', '', '24326124313024617159783744597a5634385953557173476836456165487456424c4f507944673048524a4e41626b4e485159446843717562643269', 1683822535, 1685027513);
COMMIT;

BEGIN;
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (3, '查询管理员', 'GET', '/api/v1/admin/users', 1685027076);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (4, '删除管理员', 'DELETE', '/api/v1/admin/users/:id', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (5, '添加管理员', 'POST', '/api/v1/admin/users', 1685032171);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (6, '更新管理员', 'PUT', '/api/v1/admin/users/:id', 1685032188);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (7, '编辑管理员', 'GET', '/api/v1/admin/users/:id/edit', 1685032241);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (8, '查询角色', 'GET', '/api/v1/admin/roles', 1685032424);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (9, '删除角色', 'DELETE', '/api/v1/admin/roles/:id', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (10, '添加角色', 'POST', '/api/v1/admin/roles', 1685032171);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (11, '更新角色', 'PUT', '/api/v1/admin/roles/:id', 1685032188);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (12, '编辑角色', 'GET', '/api/v1/admin/roles/:id/edit', 1685032241);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (13, '查询规则', 'GET', '/api/v1/admin/rules', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (14, '删除规则', 'DELETE', '/api/v1/admin/rules/:id', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (15, '添加规则', 'POST', '/api/v1/admin/rules', 1685032171);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (16, '更新规则', 'PUT', '/api/v1/admin/rules/:id', 1685032188);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (17, '编辑规则', 'GET', '/api/v1/admin/rules/:id/edit', 1685032241);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (18, '查询所有规则', 'GET', '/api/v1/admin/api/all', 1685034830);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (19, '查询接口', 'GET', '/api/v1/admin/api', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (20, '删除接口', 'DELETE', '/api/v1/admin/api/:id', 1685031710);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (21, '添加接口', 'POST', '/api/v1/admin/api', 1685032171);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (22, '更新接口', 'PUT', '/api/v1/admin/api/:id', 1685032188);
INSERT INTO `system_api` (`id`, `name`, `method`, `path`, `created_at`) VALUES (23, '编辑接口', 'GET', '/api/v1/admin/api/:id/edit', 1685032241);
COMMIT;
