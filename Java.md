# JAVA

java语言特点

- 面向对象
- 强类型
- 可移植
- 编译性

## 面向对象

一种软件开发方法，一种规范，对现实世界理解和抽象的方法

##### 封装

公开对外接口, 隐藏具体实现

##### 继承

子类的对象作为父类的对象来使用

##### 多态

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

[short16 int32 long64] [float32 double64] char16 byte8 boolean1

准动态语言，因为反射

### String

String不是基本数据类型

经常进行数据变更的时候最好不使用string，因为string是不可变的

用StringBuffer（线程安全），StringBuilder

String StringBuffer StringBuilder 都是 char[] 实现的

String中是final

另外两个可变 StringBuffer加了sychronized

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

## 内部类

- 定义在类中的类是内部类

- 静态内部类

  - static修饰的内部类
  - 如果是public的话，可以外部直接像使用一个普通类一样使用

- 实例内部类，没有static修饰

- 局部内部类，方法中的内部类

- 匿名内部类：

  - 匿名的，一个队接口实现的实例

    ```
    interface Compute{
        int sum(int a,int b);
    }
    
    class Mymath{
    	public void SUM(Compute C, int x, int y) {
    		return C.sum(x,y);
    	}
    }
    
    Mymath mymath = new Mymath();
    mymath.SUM(new Compute(){
    	public int sum(int a, int b) {
    		return a+b;
    	}
    },100,100);
    ```

    实际上是现场写了一个接口实现

    如果不用匿名内部类的话，就要现场手写一个局部内部类，来实现这个接口，再实例化

    匿名内部类省略了这个过程

    也是一种传递代码块的手段

## IO流

javaio流：java与外部数据交互的io输送通道，数据传输的管道

上传、下载，都是流的操作

#### 分类：

输入流、输出流：自程序到文件

字节流、字符流：记事本打得开打不开的~

节点流、处理流：节点流是端点，处理流是节点

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
- 泛型是参数化类型的应用，操作的数据类型不限定于特定类型，根据**代码**实际需要，设置不同的数据类型，以实现代码复用。
- 泛型的作用期间是代码期间，编译后擦除，JVM并不支持
  - 只是Java提供的一个语法糖，实际的使用还是通过强制类型转换


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

##### 泛型接口

类似于泛型类，是一类泛型类的一种规约。实现某一个泛型接口，标识这些泛型类有针对一系列类型的一种共同的方法。

```
public interface Generator<T> {

	T next();

}
```

#### 泛型方法

- 调用方法的时候，指定参数类型
- 泛型方法与泛型类没有关系
- 设计上，能使用泛型方法，就不要使用泛型类

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

- ? 代替具体类型实参，类型通配符

- 这玩意是类型实参，不是类型形参

- <? extends A>
- List<? extends A> 没办法添加元素，使用类型通配符的集合只能读取数据，不能写入数据

- extends 泛型变量限定符，限定只能取某类及其子类的泛型变量

- super 泛型变量限定符，限定只能取某类及其父类的泛型变量

- 虚拟机中没有泛型，会被**擦除**掉




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

#### Executor

jdk提供了exceutor框架来实现线程池

## 异常处理

- throw 抛出一个具体的异常

- throws 声明一个方法可能抛出的所有异常信息, 声明但是不处理

##### **try-catch-finally**

- try了普通异常, 必须要catch掉. 
  - 对应的try了RuntimeException，不一定必须要处理
- 多个catch捕获异常时, 子类靠前, 父类靠后
- catch中return了, finally会在return前执行

##### java异常体系结构

java异常体系是Throwable及其子类

- Throwable
  - Error
    - 致命的，无法恢复
    - OutOfMemoryError
    - StackOverflowError
- Exception 可以恢复的异常
  - RuntimeException
  - 其他Exception

##### unchecked异常

- RuntimeException
- Error

##### checked异常

- 需要显示的catch，或者抛出
- 编译阶段就是可以预见的（所以又是编译时异常）

##### 默认异常处理

一个方法如果发生异常，会创建一个异常对象，并转交给JVM。JVM会顺着调用栈去看是否有可以处理异常的代码，如果有，则调用异常处理代码。如果没有找到，会将异常转给默认的异常处理器（JVM的一部分，打印异常信息，终止应用程序）

#### 异常处理原则

延迟捕获，异常发生时，不应该立即捕获，应该考虑当前作用域是否有能力处理，没有的话，应该继续向上抛出



## 字节码

.class就是字节码文件

示例：

```
cafe babe 0000 0033 0015 0a00 0400 1109
0003 0012 0700 1307 0014 0100 0161 0100
```

字节码文件是java 一次编译，到处运行的实现。是虚拟机运行

