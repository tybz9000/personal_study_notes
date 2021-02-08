# JAVA

### 基本数据类型

[short int long] [float double] char byte boolean

准动态语言，因为反射

## 反射

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

@Retention。在什么级别保存 SOURCE CLASS RUNTIME

@Inherited 被继承

注解用interface定义

类型 命() default

```
int value() default 0;
```

## 泛型

java采用泛型擦除的机制来引入泛型，Java中的泛型仅仅是给编译器javac使用的，确保数据的安全性和免去强制类型转换问题。一旦编译完成，泛型相关数据全部擦除

## JVM

### 类加载过程，类初始化：

##### 加载：

​	class文件载入内存，创建class对象

##### 链接

​	验证：写的对不对，有没有安全问题

​	准备：为静态变量分配内存，设置初始值

​	解析：将类中的二进制数据中的符号引用替换成直接引用（final修改的常量的替换）

##### 初始化

​	执行类初始化，调用clinit，将静态方法合并到一起

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

## 多线程

##### 继承Thread类  

​	Thread类本身就实现了Runnable接口

##### 实现Runnable接口

实现Run方法

start方法执行

##### Thread类通过静态代理

#### 线程状态

创建

就绪：创建，启动后进入就绪

​			阻塞解除后

​			运行释放cpu资源后

阻塞：运行，阻塞后阻塞

运行：就绪，获得cpu资源后运行

死亡：线程执行完成



就绪是运行的必然前态

#### 线程操作

获取当前线程

##### Thread.currentThread()

##### Thread.sleep

阻塞线程特定毫秒

不会释放锁

##### Thread.yield

线程礼让，转为就绪

### 静态代理

```
interface winner {
	void celebrate();
}

class Me inplements winner{
	public winner() {
		balabala;
	}
}

class We inplements winner{
	public winner target;
	public winner() {
		诶嘿
		target.celebrate();
	}
}
```

