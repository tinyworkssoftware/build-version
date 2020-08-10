CREATE SCHEMA if not exists build_version_db;

CREATE USER 'build_version_app'@'%' IDENTIFIED BY 'P@ssw0rd12345';
CREATE DATABASE build_version_db CHARACTER SET utf8 COLLATE utf8_unicode_ci;
GRANT ALL PRIVILEGES ON build_version_db.* TO 'build_version_app'@'%';