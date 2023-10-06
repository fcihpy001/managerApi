-- 创建gin_demo数据库
CREATE DATABASE `defi` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建gin_demo数据库的用户
CREATE USER 'defi'@'%' IDENTIFIED WITH mysql_native_password BY 'defi1688' PASSWORD EXPIRE NEVER;
-- 授权gin_demo数据库的用户
GRANT ALL PRIVILEGES ON defi.* TO 'defi'@'%';
-- 刷新权限
FLUSH PRIVILEGES;
