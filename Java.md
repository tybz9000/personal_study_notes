# JAVA

## 面向对象

##### 封装

公开对外接口, 隐藏具体实现

##### 继承

子类的对象作为父类的对象来使用

#### 多态

允许不同类对象对同一消息做出反应

- 运行时多态
  - 重写overriding实现，一般讲多态是运行时多态
    - 经典多态模型
      - 父类子类
      - 子类重写
      - 调用指向子类
- 编译时多态
  - 重载实现，编译时就能决定运行，静态多态

### 基本数据类型

[short int long] [float double] char byte boolean

准动态语言，因为反射

### String

String不是基本数据类型

经常进行数据变更的时候最好不使用string，因为string是不可变的

用StringBuffer（线程安全），StringBuilder

## 抽象类

使用abstract修饰的类

抽象类不能实例化

abstract修饰的方法叫抽象方法，不需要在抽象类中实现，不一定非得有抽象方法

通过继承实现功能

非继承只能使用静态方法

可以有构造方法，用来封装子类的公共构造调用，子类用时继承

## final

- final 修饰引用

  - 引用是基本数据类型，是常量
  - 引用指向数组、对象等，内容可改，但内存地址不可以改
  - 修饰的类成员变量，必须当场赋值

- final修饰方法

  - 可以被继承，不能被重写

- final修饰类

  - 断子绝孙

  

## 序列化

java对象保存并可恢复的一种手段，将内存中的对象保存下来

比如json

## 克隆

实现方式

Object.clone：浅拷贝，对象中的属性，只复制了内存地址

序列化后反序列化：深拷贝

## IO流

javaio流：java与外部数据交互的io输送通道，数据传输的管道

上传、下载，都是流的操作

#### 分类：

输入流、输出流：自程序到文件

字节流、字符流：记事本打得开打不开的~

节点流、处理流：

```
所有流继承自这四个：
字节输入流：InputStream
字节输出流：OutputStream
字符输入流：Reader
字符输出流：Writer
```

#### 特性

先进先出

顺序读取

只读或只写

#### 常用类与接口

File：面向对象的方式管理文件

OutputStream：二进制字节输出

InputStream：二进制字节输入

Writer：字符输出流

Reader：字符输入流

#### File java.io.File

```
public File(String pathname){}
public File(String parent,String child){}
public File(File parent,String child){}
```

```
public boolean createNewFile()//构造器的文件不存在，则创建.
要求路径存在，否则java.io.IOException 系统找不到指定路径
public boolean mkdir()；//构造器的文件夹不存在，泽创建
不能创建多层文件夹，但不会报异常
public boolean mkdirs();//可以创建多层文件夹
public boolean delete();//删除文件、文件夹，不能递归删
public boolean renameTo(File dest);//改名字、也能移动文件

public boolean isDirectory();//判断是否是目录
public boolean isFile();是否是文件
public boolean exists()
public boolean canRead()
public boolean canWrite()
public boolean isHidden()

public String getAbsolutePath()
public String getPath()
public String getName()
public long length()
public long lastModified()最后修改时间

public String[] list();
public File[] listFiles();相当于ls
```

#### OutputStream

```
FileOutputStream(File file);
FileOutputStream(String path);

FileOutputStream outs = new FileOutputStream("aaa.txt")
String out = "Hello";
byte[] data = out.getBytes();
outs.write(byte[]);
outs.close();//用完了要关!!!!!!!!!!!!!!!!!!!!!!!!!!!!
```

| 分类         | 字节输入    | 字节输出     | 字符输入 | 字符输出 |
| ------------ | ----------- | ------------ | -------- | -------- |
| 基类         | InputStream | OutputStream | Reader   | Writer   |
| 文件(节点流) | File*       | File*        | File*    | File*    |
| 数组         |             |              |          |          |
| 管道         |             |              |          |          |
| 字符串       |             |              |          |          |
| 缓冲流       |             |              |          |          |
| 转换流       |             |              |          |          |
| 对象流       |             |              |          |          |
| 打印         |             |              |          |          |
| 输入输出     |             |              |          |          |
| 特殊         |             |              |          |          |

```
FileReader fr = new FileReader(file1);
int data = fr.read();//读一个byte
fr.read(char[]);读进字符数组

FileWriter fw = new FileWriter(file1);
fw.write(char[])
```

字符流、字节流。搞中文容易出错

#### 缓冲流：提高效率

缓冲流属于处理流

处理流：包裹在原始流外面的一层

提高了效率，原理：降低IO次数开销

关只需要关缓冲流

```
FileReader fr = new FileReader(file1);
BufferedInputStream bis = new BufferedInputStream(fr);
bis.read();



```

## 泛型

- 使用类型参数来解决类型的问题


`在Collection中，不指定泛型通配符的话，默认是Object`

- 编译时类型监测机制
- 参数化类型，操作的数据诶性被指定为一个参数。

#### 泛型类

例子ArrayList<E>

泛型类，在创建对象的时候，决定具体的泛型类型。不指定的话，就是Object

- 泛型类不支持基础数据类型
  - 由于类型擦除机制，泛型底层实际上是Object，而基础数据类型不是继承自Object的
- 泛型类型，在逻辑上可以看做是多个不同的类型，但实际上是同一个

```
class 类名称 <泛型标识,泛型标识> {

}
泛型类型标识，在类名称后面，一行最后

//一个自定义的泛型类，有一个或多个类型变量的类，放在类名后面
public class MyArrayList<E> {
    private Object[] dataArr;
    private int size;
    public MyArrayList (int initialCapacity) {
        dataArr = new Object[initialCapacity];
        size = 0;
    }
    public void push(E e) {
        dataArr[size] = e;
        size++;
    }
}

MyArrayList<Integer> myArrayList = new MyArrayList<Integer> ();
//1.7之后 后边这个<Integer>可以省略
```

#### 泛型方法

调用方法的时候，制定参数类型

```
public <T> List<T> methodName() {

}
类型定义在返回值前边，一行最前
声明了泛型列表的方法才是泛型方法
方法体，参数类型，返回值中可以使用。

//一个类型转换的泛型方法
public static <E, T> List<T> transform(List<E> input, Class<T> clazz) {
        if (CollectionUtils.isEmpty(input)) {
            return Lists.newArrayList();
        }
        List<T> res = Lists.newArrayList();
        try {
            for (E e : input) {
                T t = clazz.newInstance();
                BeanUtils.copyProperties(e, t);
                res.add(t);
            }
        } catch (InstantiationException | IllegalAccessException e) {
            throw Throwables.runtimeException(e, RetCode.CLASS_CAST_FAIL);
        }
        return res;
    }
```

#### 类型通配符

? 代替具体类型实参，类型通配符

这玩意是类型实参，不是类型形参

<? extends A>

List<? extends A> 没办法添加元素，使用类型通配符的集合只能读取数据，不能写入数据

extends 泛型变量限定符，限定只能取某类及其子类的泛型变量

super 泛型变量限定符，限定只能取某类及其父类的泛型变量

虚拟机中没有泛型，会被**擦除**掉



## 反射

程序获取自身状态结构的一种手段

Class c = Class.forName("com.taiyang.User");//反射创建Class对象。一个类只有一个

Object是所有类的源头，Class是所有反射的源头

##### 获取Class

1、object.getClass()

2、Class.forName

3、类名.class *【最高效】*

4、内置对象的type: 

```
integer.type
```

##### 获得父类类型

```
clazz.getSuperClass()
```

##### 获得方法

```
clazz.getMethod("{methodName}",parameterTypes)
clazz.getDeclaredMethods()
```

##### 获得构造器

```
clazz.getConstructors();
clazz.getDeclaredCOnstructor(String.class,int.class)
```

##### 获得属性

```
clazz.getDeclaredFields();//获取全部属性
clazz.getFields();//获取public属性
```

##### 创建对象

```
clazz.newInstance();//调用无参构造器
Constructor constructor = clazz.getDeclaredConsturctor(String.class, int.class);
constructor.newInstance("bala",1);//通过构造器传值调用有参构造器
```

##### 调用方法

```
method.invoke(obj,args);
```

##### 属性赋值

```
field.set(obj,value);
```

##### 获取泛型信息

```
Type[] type = method.getGenericParameterTypes()
```

##### 获取注解信息

```
class.getAnnotations();
```

## Collection

### 集合

数据可以存储基本类型，而集合只能存储对象

数组长度不可变

```
Collections
--List
--Set
--Queue
--SortedSet
Map
--HashMap
--TreeMap
```

### ArrayList

- 有序的，基于数组
- 线程不安全
- 核心在于扩容和缩容，扩容1.5倍 10 + 10 >> 1
- 下标查询o(1)
- 删除，把目标节点从后往前挪，不会改size
- modifyCount 用以处理并发的情况，当读时发现modifyCount不一致，直接抛异常

```
private static final int DEFAULT_CAPACITY = 10; //默认长度为10
transient Object[] elementData; //存在这
//transient关键字，序列化，反序列化中被忽略

public ArrayList {
	private static final int DEFAULT_CAPACITY = 10;
	transient Object[] elementData;
	private size = 0;
	
	public ArrayList(){
	}
	
	public void add(Object t) {
		if (elementData == null) {
			elementData = new Object[DEFAULT_CAPACITY];
			size = DEFAULT_CAPACITY;
		}
		if (elementData.length() + 1 > size) {
			int newSize = size + size >> 1;
			transient Object[] newData = new Object[newSize];
			Array.copy(elementData, newData);
		}
		elementData[size++] = T;
	}
}
```

### HashMap

**Hashmap**，允许空值（好理解），允许空键

1. 散列表，数组【桶】+链表【重复值】
2. Hash，任意长度的输入，通过hash算法，变成固定长度的输出。损失数据，但是唯一
3. 不线程安全
4. 外层是node数组，冲突后展开为node链表，1.8引入了红黑树，长度大于8时，桶大于64时，按红黑树存储【解决链化很长的问题】，降到6时，转为链表
5. Node，HashMap的一个静态内部类，实现了Map.Entry，存了hash，key，value，next
6. Node池一定是2n长，方便-1后与哈希值能得到确定的位置。便于路由算法的执行
7. 桶初始16，最大2的30次方，默认负载因子0.75
8. modCount变更次数，结构变更
9. Node桶定位的思路，取int的hashCode，然后mod桶长度
10. 重复插入链表的时候，直接放头部，这样最快。其实就是放桶里，让后把原来桶的沿链表向下移
11. 判重标准，hashCode方法返回值相&&equals

源码

```
//默认table长度
static final int DEFAULT_INITIAL_CAPACITY = 1 << 4; 
//最大table长度
static final int MAXIMUM_CAPACITY = 1 << 30;  
//扩容系数
static final float DEFAULT_LOAD_FACTOR = 0.75f;
//红黑树起点数量，终点数量，起点table长度
static final int TREEIFY_THRESHOLD = 8;
static final int UNTREEIFY_THRESHOLD = 6;
static final int MIN_TREEIFY_CAPACITY = 64;

//Node定义
static class Node<K,V> implements Map.Entry<K,V> {
        final int hash;
        final K key;
        V value;
        Node<K,V> next;
        
//构造方法
判断了一下参数，赋值了负载因子和阈值（长度，一定会是2的次方，小于传入的数的最大的2的次方）

//put方法【懒加载】
这里会判断是否是空，然后初始化

//table Size一定是2的次方实现
static final int tableSizeFor(int cap) {
        int n = cap - 1;
        n |= n >>> 1;
        n |= n >>> 2;
        n |= n >>> 4;
        n |= n >>> 8;
        n |= n >>> 16; //最终整成 0000 0000 1111 1111这样的~ int范围都覆盖那种
        return (n < 0) ? 1 : (n >= MAXIMUM_CAPACITY) ? MAXIMUM_CAPACITY : n + 1;
    }
    9
    1001
    0100  1101
    0011  1111
    16
```

### HashSet

- 基于hashMap实现
- 值一个空Object

### LinkedList

- 双向链表

- 记录起点 终点

- ```
  public class LinkedList<E>{
  	transient Node<E> first;
  	transient Node<E> last;
  	
  	public void add(E e) {
  		final Node<E> l = last;
  		final Node<E> newLast = new Node(e, null, last);
  		last = newLast;
  		if (l == null) {
  			first = newLast;
  		} else {
  			l.next = newLast;
  		}
  		size++;
  		modCount++;
  	}
  	
  	private static class Node<E> {
  		E item;
  		Node<E> next;
  		Node<E> prev;
  	}
  }
  ```

- 查找，折半一次

### Vector

类似于ArrayList的底层结构

线程安全，性能较低

## 注解

元注解，定义注解

@Target。描述什么

@Retention。在什么级别保存 SOURCE CLASS（默认） RUNTIME

@Inherited 被继承

注解用interface定义

类型 命() default

```
int value() default 0;
```

## lambda表达式

### 原生lambda

把一个代码块赋给一个变量

```
a = public void sayHello(int i) {

​	print("say hello")

}
```

省略public

省略void【编译器能判断】

省略名字【赋值给a了】

省略参数类型

省略大括号

```
a = (i) -> print("say hello")
```

a的类型是一个接口 void a(int i)

它是个接口的实现，这个接口只能有一个方法。函数式接口

甚至可以作为一个参数传递给方法

```
interface SayInterface{
	void say(String s);
}

public static void saySomething(SayInterface sayMethod, String s) {
	sayMethod.say(s);
}

class SayInterfaceImpl implement SayInterface {
	public void say(String s) {
		print(s)
	}
}

saySomething(new SayInterfaceImpl(), "hellow")
```

saySomething((s) -> print(s), "hellow")

### 函数式接口包

例子，以特定条件从List中筛选元素

```
public void Filter(List l) {
	for(Object object : l) {
		if(o.isXXX()) {
			o.doXXX()
		}
	}
}

l.stream().filter(o -> o.isXXX()).foreach(o -> o.doXXX())
```

链式编程，每步返回的结果类型不变。

filter、foreach 参数都是一个函数式方法

#### 四大函数式接口

Consumer  参数T，返回void

```
(x) -> System.printLn(x)
```

Supplier 无，返回R

```
() -> new R()
```

Function 参数T，返回R

Predicate 参数T，返回Boolean

### Optional

擅长处理空指针问题，是一个容器，用来装载元素

```
Optional<Person> personOpt = Optional.ofNullable(person); //创建实例，可以为空，可以不为空

personOpt.get() //把值取出来，但为空可能抛异常

if(personOpt.isPresent()){}

personOpt.isPresent(() -> sayHello());//这个可以传入一个consumer

personOpt.orElse(new Person());//有值返回，无值返回参数

personOpt.orElseGet(() -> return new Person())//有值返回，无值执行参数

personOpt.orElseThrow(() -> throw new YmsException())//有值返回，无值抛异常

personOpt.map(Person::getName)// 返回name的包装optional

personOpt.flatMap(Person::getName)//直接返回name值

personOpt.filter(p -> person.getAge() > 18)//返回过滤后的optional

```

### Stream

将要处理的元素集合看作一种流，内部没有元素，方法和optional很像

## 多线程

##### 进程：

运行的程序，例如一个运行中的Java虚拟机实例

##### 线程：

操作系统能够进行运算调度的最小单位。它被包含在进程之中，是进程中的实际运作单位

##### 并发：

一个时间段内执行多个任务，可能是切换执行

##### 并行：

同一个时刻，多个任务同时

##### 继承Thread类  

Thread类本身就实现了Runnable接口

##### 实现Runnable接口

实现Run方法

start方法执行

##### Thread类通过静态代理

##### 守护线程

子线程结束时，守护线程自动结束

setDeamon(true)

#### 线程状态

- 创建

- 就绪：
  - 创建，启动后进入就绪
  - 阻塞解除后
  - 运行释放cpu资源后
- 阻塞：运行，阻塞后阻塞

- 运行：就绪，获得cpu资源后运行

- 死亡：线程执行完成




就绪是运行的必然前态

**java线程状态**

- NEW：还没启动的线程
- RUNNABLE：正在执行，或者等待cpu资源中，就绪
- BLOCKED：被阻塞,等待获取监视器锁进入synchronized代码块或者在调用Object.wait之后重新进入synchronized代码块
- WAITING：无限期等待另一个线程执行特定动作后唤醒它,也就是调用Object.wait后会等待拥有同一个监视器锁的线程调用notify/notifyAll来进行唤醒
- TIMED_WAITING：有时限的等待另一个线程执行特定动作
- TERMINATED：已经完成了执行

### 线程同步

多个线程访问同一个变量时, 保持数据一致的手段是为线程同步

(对于多进程而言, 进程有独立内存空间, 独立程序计数器, 很少产生数据冲突, 不会有同步问题)

#### 线程操作

获取当前线程

**Thread.currentThread()**

**start()**

**Thread.sleep()**

阻塞线程特定毫秒

不会释放锁、不会让出监视器

**yield()**

线程礼让，转为就绪

**join()**

插队的肯定先走

setName

getName

setPriority

getPriority

interrupt()

中断，如果是阻塞状态就停止

否则改变中断标识位

**wait()**

- 线程通信，告诉当前线程，释放锁，然后开始睡眠等待。

- 直到有线程进入监视器调用notify() notifyAll()启用它
- 必须要作用于同步代码块，否则 Exception in thread "main" java.lang.IllegalMonitorStateException
- wait、notify、notifyAll是用在同步代码块上的，产生了数据竞争

**notify()**

随机唤醒一个在一样的对象监视器上等待的线程

**notifyAll()**

唤醒所有的在一样对象监视器上等待的线程



```
public class WaitNotify {

  public static void main(String[] args) {

    Object lock = new Object();
    
    // thread1
    new Thread(() -> {

        System.out.println("thread1 is ready");

        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {

        }
        synchronized (lock) {
            lock.notify();
            System.out.println("thread1 is notify,but not exit synchronized");
            System.out.println("thread1 is exit synchronized");
        }


    }).start();

    // thread2
    new Thread(() -> {
        System.out.println("thread2 is ready");

        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {

        }
        synchronized (lock) {
            try {
                System.out.println("thread2 is waiting");
                lock.wait();
                System.out.println("thread2 is awake");
            } catch (InterruptedException e) {
            }
        }
    }).start();
  }
}
//线程一二同时开始执行，线程二sleep一秒后等待，
//线程一sleep两秒后唤醒
//线程一二在同一个同步代码块，在对象监视器上
thread1 is ready
thread2 is ready
thread2 is waiting
thread1 is notify,but not exit synchronized
thread 1 is exit synchronized
thread2 is awake

```

**对象监视器**

java语言中，对象监视器和锁概念上是接近的

Java允许任何对象都可以成为一个锁也叫做对象监视器

#### Synchronized

**java内置锁**

- 代码块
- 方法：作用于实例，不参与继承，父类方法有，也没法继承到子类，
- 静态方法：作用于类，作用到所有访问到这个类的对象
- 类：作用于类，类似于静态方法，也是给整个类加锁

进入同步代码块、同步类时获取**对象**内置锁

是互斥锁，如果另外一个线程要获得这个锁，必须阻塞（Blocked）等待

#### 锁

##### 乐观vs悲观

- 乐观锁：要加锁的不一定会被改，改了我再说。记录更新id之类的
  - 多读、乐观锁
  - CAS：比较并交换，操作前拿出值，我现在要写的时候，原值拿出来再比较，不变则操作，否则失败。
- 悲观锁：要加锁的一定是有人要改的，都锁住，谁也不能动。（synchronized、Lock）
  - 多写、悲观锁

##### 独享锁vs共享锁

- 独享锁：锁同时只能被一个线程持有

  - 往往用作写锁
- 共享锁：锁同时可以被多个线程持有

  - 往往用作读锁
- 互斥锁、读写锁：独享、共享锁的一种实现

  - ReadWriteLock：读写锁
  - ReenTrantLock：

### 线程池

线程池可以看做是线程的**集合**，在没有任务时，线程处于空闲状态。请求到来时，线程池给请求分配一个空闲的线程，请求完成后，线程回到线程池中，这样就完成了线程的**重用**

​	减少了线程总的生命周期开销

**核心线程**：永不回收的线程

**阻塞队列**：干不完的活儿

**非核心线程**：忙碌的时候使用的线程

**空闲时间**：非核心线程空闲到一定时间后，回收

**饱和策略**：

- AbortPolicy(抛出一个异常，默认的)
- DiscardPolicy(新提交的任务直接被抛弃)
- DiscardOldestPolicy（丢弃队列里最老的任务，将当前这个任务继续提交给线程池）
- CallerRunsPolicy（交给线程池调用所在的线程进行处理，即将某些任务回退到调用者)

![image-20210705204424134](C:\Users\taiyang\Documents\image-20210705204424134.png)

#### Executor

jdk提供了exceutor框架来实现线程池

## 异常处理

throw 抛出一个具体的异常

throws 声明一个方法可能抛出的所有异常信息, 声明但是不处理

**try-catch-finally**

try了普通异常, 必须要catch掉. 

多个catch捕获异常时, 子类靠前, 父类靠后

catch中return了, finally会在return前执行
