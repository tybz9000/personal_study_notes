# JVM

### 类装载子系统

### 运行时数据区

#### 元空间（方法区、永久代）

常量

静态变量

类元信息：类被类装载子系统加载到这里

1.8之后使用物理内存

#### 堆

##### 年轻代

伊甸园区Eden【默认1/10】

Survivor区 From To

##### 老年代【默认2/3】

#### 程序计数器

线程马上要执行的JVM指令码的内存位置

线程独享

#### 栈

java虚拟机栈

线程私有

每个方法执行时创建栈帧。

##### 栈帧

局部变量表：局部变量的值

操作数栈：运算时的操作数

动态链接：方法里的符号的指令码，执行过来后，该执行哪些指令码。指向方法区

方法出口：执行完后，应该执行到上一个方法的哪一步

##### 本地方法栈

线程独享，需要调用的本地方法

### GC

minor gc，伊甸园区满了

#### GC算法

##### 可达性分析

GC Root根：

- 类加载器
- Thread
- 虚拟机栈的本地变量
- static成员
- 常量引用
- 本地方法栈的变量

顺着GC Root根的引用去找，没被引用的干掉

#### 执行引擎

### 类加载过程，类初始化：

##### 加载：

​	class文件载入内存，创建class对象

##### 链接

​	验证：写的对不对，有没有安全问题

​	准备：为静态变量分配内存，设置初始值

- [ ] ​	解析：将类中的二进制数据中的符号引用替换成直接引用（final修改的常量的替换）


##### 初始化

- [ ] ​	执行类初始化，调用clinit，将静态方法合并到一起


### 什么时候会导致类加载或者类初始化?

创建类的实例(new方法)

调用某个类的静态方法

访问某个类或者接口的静态属性

使用反射机制来强制创建某个类或接口对应的java.lang.Class对象。（Class.forName(“Person”)）

初始化某个类的子类

直接使用java.exe命令来运行某个主类 main

##### 不会的情况

子类访问父类静态变量，不会导致初始化

引用常量不会初始化

​	static 链接阶段就在常量池了

#### 类加载器

1. **Bootstrap classLoader**引导类加载器，负责java核心库

   无法获取

2. **Ext ClassLoader**扩展类加载器，负责jre/lib/ext 或 -D java.ext.dirs

   systemClassLoader.getParent()

3. **App ClassLoader**系统类加载器，负责 java-classpath 或者 -D java.class.path

   ClassLoader systemClassLoader = ClassLoader.getSystemClasssLoader();

#### 双亲委派机制

保证类加载按顺序进行，不会有人能篡改底层类加载器数据