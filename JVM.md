# JVM

#### Java虚拟机

- Java有一套虚拟机规范，以及基于这套规范的一系列虚拟机实现
- HotSpot：使用范围最广的Java虚拟机
  - Sun JDK、OpenJDK都是用的它
- JRockit：被Oracle收购，很多优秀特性被移植到HotSpot
- IBM J9



![](D:\学习文件3\图片\java虚拟机.png)



#### Java内存空间

- 方法区（1.8之前的永久代，1.8元空间）
  - java规范定义的，永久代与元空间都是方法区的一种实现
  - 类常量、字符串常量、方法代码、类定义数据等
  - 回收永久代需要卸载类
  - 元空间：
    - 存储class matedata
      - Klass结构，Java类在虚拟机中的表示
        - 不同于Class对象，Class对象在堆中
      - method metadata，方法字节码，局部变量表，异常表，参数信息
      - 注解
      - 方法计数器
      - 元空间并不在虚拟机中存贮，直接存储在内存中
- 堆
- 栈
  - 每个线程有自己的虚拟机栈
  - 栈帧
    - 随着方法调用的层层深入，不断压入栈帧，一个方法一个栈帧
    - 局部变量表
    - 操作数栈
    - 动态链接：
      - 指向运行时常量池中该栈帧所属方法的引用
    - 返回地址
- 本地方法栈
- 程序计数器
  - 线程当前执行字节码的行号指示器
  - 不可能OutOfMemory
- 总结
  - 方法区与堆是线程共享的，其他的线程各自独占

# 类加载过程，类初始化：

javac：词法分析、语法分析、语义分析，之后编译成class文件（字节码）

class文件，二进制文件

之后就是类加载进入jvm

##### 加载：

- ​	class文件载入内存
- ​	创建class对象

##### 链接

- 将类的二进制数据合并到JRE中
- ​	验证：写的对不对，有没有安全问题

- ​	准备：为静态变量分配内存，设置初始值


- ​	解析：将类中的二进制数据中的符号引用替换成直接引用（final修改的常量的替换）


##### 初始化

- ​	执行类初始化，调用clinit，将静态方法合并到一起
- 调用构造方法


### 什么时候会导致类加载或者类初始化?

- 创建类的实例(new方法)

- 调用某个类的静态方法

- 访问某个类或者接口的静态属性

- 使用反射机制来强制创建某个类或接口对应的java.lang.Class对象。（Class.forName(“Person”)）

- 初始化某个类的子类

- 直接使用java.exe命令来运行某个主类 main


##### 不会的情况

- 子类访问父类静态变量，不会导致初始化

- 引用常量不会初始化

- static 链接阶段就在常量池了



#### 类加载器

1. **Bootstrap classLoader**引导类加载器，负责java核心库

   无法获取

2. **Ext ClassLoader**扩展类加载器，负责jre/lib/ext 或 -D java.ext.dirs

   systemClassLoader.getParent()

3. **App ClassLoader**系统类加载器，负责 java-classpath 或者 -D java.class.path

   ClassLoader systemClassLoader = ClassLoader.getSystemClasssLoader();

#### 双亲委派机制

保证类加载按顺序进行，不会有人能篡改底层类加载器数据

# 运行时数据区

![](D:\学习文件3\图片\88e60f28a895442d8846ad0e87739184.png)

#### 元空间（方法区、永久代）

常量

静态变量

类元信息：类被类装载子系统加载到这里

1.8之后使用物理内存

#### 堆

##### 年轻代

- 伊甸园区Eden【默认1/10】
- Survivor区 From To
  - 幸存者0
  - 幸存者1
- 老年代

#### 程序计数器

**线程**马上要执行的JVM指令码的内存位置

线程独享

#### 栈

java虚拟机栈

线程私有

每个方法执行时创建栈帧。

##### 栈帧

局部变量表：局部变量的值

操作数栈：运算时的操作数

动态链接：指向运行时常量池中该栈帧所属方法的引用，方法里的符号的指令码，执行过来后，该执行哪些指令码。指向方法区


方法出口：执行完后，应该执行到上一个方法的哪一步

#### 本地方法栈

线程独享，需要调用的本地方法

# GC

#### JVM堆经典布局

- 新生代

  - Eden：新生的对象在此创建
  - S0：Survivor0
  - S1：Survivor1

- 老年代

  - 达到晋升年龄进入老年代 -XX:MaxTenuringThreshold

  - 对象特别大，直接进入老年代 -XX:PretenureSizeThreshold

  - Survivor区内存不够了，最老的一批送入老年代
    - s区满了回直接触发Young GC


  JVM垃圾回收过程

- Minor GC = Young gc

  - Eden区没位置的时候，触发Young gc
  - Young GC存活的对象，转移到Survival区
  - STW

- Full GC

  - 老年代空间不够用


#### GC算法

##### 引用计数法

- 比较老了
- 现在的JVM一般不采用这种GC算法
  - 有新引用指向它，那么计数加一
  - 超过生命周期，计数减一
  - 有引用不再指向它，计数减一
- 优点：没有STW
- 缺点：运行时效率低
- 缺点：无法处理循环依赖

##### 可达性分析

引用：有人用

通过对GC Root对象做可达性分析，从GC Root对象开始，**向下搜索**（向内搜索），形成的路径被称为引用链，如果一个对象到GC ROOT没有任何引用，没有形成引用链，那么该对象等待GC回收。

GC Root根：

- Class对象，由BootstrapClassLoader加载的对象是不能被回收的
- Thread，活动的线程
- sychronzied引用的对象
- 虚拟机栈中引用的对象（局部变量表）
  - 比如某段代码块结束，虚拟机栈出栈，GC Root消失

- static，方法区中类静态属性引用的对象
- final，常量引用，方法区中常量引用的对象
- native，本地方法栈的变量

顺着GC Root根的引用去找，没被引用的干掉

#### 具体执行：

被定义为垃圾对象后，并不会马上被回收

定义为不可达对象后，进行一次标记，然后执行一次筛选，执行finalize方法

如果finalize方法后，可达性分析表示还是垃圾对象，那么会被回收

#### 执行引擎

#### 垃圾回收器

- Serial GC：新生代垃圾回收器，单线程，Stop The World，基于**复制**算法
- Parallel Scavenge：新生代垃圾回收器，采用**复制**算法，自适应调节新生代大小、Eden S区比例等。
  - 旨在保证程序到达一个可控制的吞吐量（用于运行用户代码的时间/CPU消耗总时间）
  - 吞吐量低，占用更多内存，更快GC
  - 吞吐量高，降低内存占用，更慢GC
- ParNew: Serial的多线程版本，采用**复制**算法
- Serial Old: Serial的老版本，基于**标记整理算法**
- CMS：采用**标记清除**算法，减少停顿时间，但会内存碎片化
- G1：采用**标记整理**算法，将内存分为若干个Region，每次整理一个Region，各个Region有自己的代。
  - 不使用复制算法，所以不需要设置S1，S2两个Survival

# JVM调优

#### 什么时候需要JVM调优

- 如果JVM参数配置合理的话，大多数场景上不需要调优
  - JVM调优不是常规手段，性能问题第一选择优化程序，最后才需要JVM调优

- JVM调优需要全面的监控，详细分析性能数据，当系统监控指标异常时，需要介入评估
  - gc次数过多
  - gc时间过长
  - 内存占用高
  - full gc频繁
  - GC停顿时间过长
  - 出现OutOfMemory等内存异常

#### JVM调优的目标

- 吞吐量、延迟、内存占用三者不可能三角
  - 延迟：GC低停顿、GC低频率
- 不同的应用，应有各自的JVM调优目标

#### 调整JVM参数

- Eden区太大：Young GC时间长
- Eden区太小：Young GC次数多
- Survival太小：太早进入老年代
- Survival太大：浪费内存

1. 标准参数：-version -help
2. X参数
3. XX参数
   1. 布尔类型XX参数：-XX:+PrintGCDetails +表示开启
   2. KV类型XX参数：-XX:NewSize=256M

- -Xms4G：JVM启动时整个堆的初始化大小

- -Xmx4G：JVM启动时整个堆的最大值

- -Xmn2G：年轻代空间的大小

- -Xss2G：栈大小

- -XX:SurvivorRatio=1：年轻代空间中两个S区同Eden空间大小比例
  - 默认为8，S0:S1:Eden = 1:1:8
    - Eden区比较大，相对YGC次数比较少。但YGC的时候，Eden区占用空间比较大，stw的时间比较长。
  
- -XX:+UseG1GC：使用垃圾回收器类型，1.7以后推荐使用G1GC

- -XX:+UseTlab：使用tlab，默认打开的，在Eden区对各个线程分配一定份额的缓存用以初始化对象

- -XX:+PrintHeapAtGC：GC时打印堆信息

- -XX:+PrintGC：打印GC信息

  - -XX:+PrintGCDetails
  - -XX:+PrintGCDateStamps

  -XX:HeapDumpPath=/usr/local/dump：dump文件路径或者名称

- -XX:+HeapDumpOnOutOfMemoryError：OOM时打印堆信息

