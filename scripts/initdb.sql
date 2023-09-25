-- 创建gin_demo数据库
CREATE DATABASE `managerdb` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建gin_demo数据库的用户
CREATE USER 'manager'@'%' IDENTIFIED WITH mysql_native_password BY 'manager' PASSWORD EXPIRE NEVER;
-- 授权gin_demo数据库的用户
GRANT ALL PRIVILEGES ON managerdb.* TO 'manager'@'%';
-- 刷新权限
FLUSH PRIVILEGES;
