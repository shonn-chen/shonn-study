# 数据库锁
## 基于数据库表记录插入删除实现分布式锁
### 实现方式
#### 数据库表
```
CREATE TABLE `lock_tab` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `lock_name` varchar(64) NOT NULL DEFAULT '',
    `lock_remark` varchar(255) NOT NULL DEFAULT '' COMMENT 'description of lock',
    `create_time` int(10) unsigned NOT NULL DEFAULT 0,
    `update_time` int(10) unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_lock_name` (`lock_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
#### 获取锁
```
INSERT INTO lock_tab(lock_name, lock_remark, create_time, update_time) VALUES('lock_name', 'lock_remark', REPLACE(unix_timestamp(now(3)),'.',''), REPLACE(unix_timestamp(now(3)),'.',''))
```
因为有uidx_lock_name的唯一索引保证，可以认为插入成功即获取相应的锁
#### 释放锁
```
DELETE FROM lock_tab WHERE lock_name='lock_name';
```
#### code example
[基于数据库表记录插入删除实现分布式锁]()
