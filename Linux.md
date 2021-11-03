# Linux与运维

| 管道符，前面输出的内容丢到后面去处理

printf echo 向屏幕输出，后者带换行

#### Linux 三剑客

#### grep数据查找定位

grep [options] pattern file

- -i 忽略大小写
- -n 显示行号
- -o 精准匹配

file可以是多个

#### awk数据切片

文本分析，可以当做简化的用于处理或统计文本表格，配置的特定C语言来理解

核心是文本行记录处理，行列分割处理

逐行读入，按空格制表符分隔

swk -F '{pattern + action}' {filenames}

#### sed数据修改

## 基础命令

#### shutdown

关机

#### sync

所有数据同步到硬盘

#### reboot

重启

#### top

top 查看进程信息

-H 查看线程信息

-p 指定pid

#### cd

#### ls

#### pwd

显示当前用户所在的目录

#### mkdir

创建目录

#### rmdir

移除目录

#### cp

复制文件或目录

#### rm

rm -rf / 跑路

#### mv

移动

#### clear

清理屏幕

#### df

磁盘系统使用量

#### du

当前磁盘

## 监控与dump

jstack -l 6 > jstack.txt  // jstack dump 看cpu，哪个线程cpu占用高

jmap -dump:format=b,file=jmap.info 6  //看堆，看对象内存占用信息

jmap -histo 6  //也是看堆，看存活实例占用情况，但看的是对象本身大小，往往是char byte等

jstat -gcutil {pid} 5000 // 查看gc

grep java.lang.Thread.State pid.dump| awk '{print $2$3$4$5}' | sort | uniq -c //查看线程状态分布

dump分析

jvisualVM jvm自带

mat 分析工具

jprofiler

远程拷贝

yum install openssh-clients  //安装openssh

scp jstack.txt mi@10.221.65.150:/home/mi/dump  //下载dump文件

scp jmap.info mi@10.221.65.150:/home/mi/dump

scp youpin.log mi@10.221.65.150:/home/mi/dump

远程ssh

apt-get install openssh-server //本地安装ssh

systemctl start sshd.service //本地启动ssh服务

systemctl status sshd.service //本地验证ssh状态

线上下载arthas

wget -c -O out.zip "https://github.com/alibaba/arthas/releases/download/arthas-all-3.3.3/arthas-3.3.3-bin.zip" 

## 常见文件位置

etc配置文件

/etc/profile 环境变量

- source /etc/profile 使环境变量生效

## 权限系统

-rw------- (600)    只有拥有者有读写权限。
-rw-r--r-- (644)    只有拥有者有读写权限；而属组用户和其他用户只有读权限。
-rwx------ (700)    只有拥有者有读、写、执行权限。
-rwxr-xr-x (755)    拥有者有读、写、执行权限；而属组用户和其他用户只有读、执行权限。
-rwx--x--x (711)    拥有者有读、写、执行权限；而属组用户和其他用户只有执行权限。
-rw-rw-rw- (666)    所有用户都有文件读、写权限。
-rwxrwxrwx (777)    所有用户都有读、写、执行权限。

##### 首位

-表文件

d表目录

##### 后面

标识拥有组、群组、其他组

读r:4 写w:2 执行x:1 权限

#### chmod

修改文件权限

#### 添加账号

useradd

#### 切换账号

su

$是普通权限

#是超级权限

## 文件查看

cat 由第一行开始，显示整个文件

tac 逆向显示整个文件

nl命令 展示带行号

more命令，可以翻页，只能向后翻页

less命令，可以翻页，可以往前翻页，q退出

head前面看 -n

tail倒着看 -n

## 链接

##### 硬链接

拷贝性质，原文件删了也没事

##### 软链接

快捷方式

## 进程

1. 每个程序都有自己的一个进程，每个进程都有一个id
2. 每个进程都有一个父进程
3. 进程可前台，可后台

#### ps

查看当前系统中正在执行的各种进程的信息

- -a
- -u 以当前用户的信息显示进程
- -x 显示后台运行程序参数

## 安装

### rpm

rpm -ivh 安装

rpm -e --nodeps 强制删除

解压缩

yum

## Arthas

阿里的java诊断工具

**下载**

wget https://alibaba.github.io/arthas/arthas-boot.jar

**启动**

java -jar arthas-boot.jar

**相关命令**

dashboard 查看监控指标

thread 打印线程堆栈信息

jad 带包的完整类名，反编译线上代码 jad com.mi.xms.youpin.basicdata.express.service.impl.ExpressServiceImpl 

watch 查看函数返回值 watch com.mi.xms.youpin.basicdata.policy.service.impl.SpuPolicyServiceImpl getVoListBySearchVoAndPage {params,returnObj}

stop 退出

cls 清屏

reset 重置增强类

trace 监控方法执行耗时 