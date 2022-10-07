# Redis

基于key-value键值对的存储系统

#### 五种基本数据结构：

String、

List：底层实际上是个链表

Set：是通过HashTable实现的

Hash：类似于存了一个Map<String, Object>

Zset：类似于Set，不过每个元素都有一个分数

#### 其他功能

Geospatial地理位置

Hyperloglog数量统计，占内存少，有一定容错需求

bitmap，统计一定时序的bool很好用

#### 事务

redis单条保证一致性，事务不保证原子性，没有隔离性

有编译错误，都不运行。有运行时异常，抛出异常，其他的正常执行

一次性、顺序性、排他性

更像一个代码块

#### 锁

##### 悲观锁

不管做什么都加锁

##### 乐观锁

不加锁，改的时候看看别人改了没

watch（命令）

##### jedis

redis java中间件

new jedis(host,port)

### Spring-boot 整合REDIS

#### spring-data 

数据相关操作都封装在这里了

spring-redis 2.0以上不用jedis，jedis直连，线程不安全。

使用lettuce使用netty

#### redisTemplate

redisTemplate.opsForValue() 后面的操作和jedis一摸一样的。可以练练

常用操作提出来了

存对象必须序列化

## Redis应用，延时调用

- **存储**：数据使用ZSet存储，对特定的key，存储一个以**时间**为score的**有序集合**
- 入列：入列数据，score是期望执行的时间
- 出列：通过定时调用，通过**lua脚本**调用，使用redis：zrangebyscore方法筛选特定score区间的数据，分数区间是0到当前时间

# Redis集群

# 缓存其他问题

### 缓存穿透

- 查询一个值为null，我们通常会再查一次数据库
- 高并发场景下，缓存穿透会导致性能问题
  - 保存null value
  - 单独逻辑处理

### 缓存击穿

- 查的时候刚好过期了
  - 定时刷新

### 缓存雪崩

- 大量的数据在同一个时间过期了，同时请求数据库刷新。
  - 随机刷新时间

### 缓存数据库一致性问题

- 并发修改缓存和数据库，无法保证最后缓存与数据库一致性

线程1更新X DB = 1

线程2 更新X DB = 2

线程2 更新X 缓存 = 2

线程1 更新X 缓存 = 1

最后 DB:2 缓存:1

- 加操作锁，编辑操作事务化
- 推荐：先删除缓存，再更新数据库【在并发读+写操作时也会有问题】
- 缓存订阅数据库binLog，动态更新缓存

