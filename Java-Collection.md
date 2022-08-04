## Collection

### 集合

- 数组可以存储基本类型，而集合只能存储对象
  - 集合存储的都是包装的基本数据类型

数组长度不可变

```
Collection接口
--List接口
--Set接口
--Queue接口
--AbstractCollection抽象类
----AbstractList抽象类（实现List接口）
------Vector
------ArrayList
------LinkedList（同时也实现了Queue）
----AbstractQueue抽象类（实现Queue接口）
------PriorityQueue
----AbstractSet抽象类（实现Set接口）
------HashSet
--------LinkedHashSet
Map
--HashMap
--TreeMap
```

## Collection

```
size()
isEmpty()
contains()
iterator()//Collection实现了Iterable
Object[] toArray()
add()
remove()
containsAll()
addAll()
removeAll()
removeIf()//根据条件删除特定的
retainAll()
clear()
```

### List

有序的collection，此接口的用户可以对列表中每个元素的插入位置进行精确的控制。用户可以根据元素的整数索引访问元素，并搜索列表中的元素。

```
addAll()
sort()
get(int)
set(int, E)
add(int, E)
remove(int)
indexOf(Object)
lastIndexOf(Object)
subList(int,int)
```

#### ArrayList

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

#### LinkedList

- 同时也实现了Queue

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

#### Vector

- 类似于ArrayList的底层结构
  - 动态数组
- 线程安全，性能较低

#### Stack

继承vector

### Set

不包含重复元素的collection

#### HashSet

- 基于hashMap实现
- 值一个空Object

Queue：队列通常已FIFO的方式排序各个元素，并按FIFO的方式插入元素

#### LinkedHashSet

继承HashSet，用LinkedHashMap实现

#### TreeSet

用TreeMap来实现

## Map

### HashMap

**Hashmap**，允许空值（好理解），允许空键

1. 散列表，数组【桶】+链表【重复值】
2. Hash，任意长度的输入，通过hash算法，变成固定长度的输出。损失数据
3. 不线程安全
4. modCount变更次数，记录变更
5. 桶
   1. 外层是node数组，冲突后展开为node链表，1.8引入了红黑树，长度大于8时，桶大于64时，按红黑树存储【解决链化很长的问题】，降到6时，转为链表
   2. Node，HashMap的一个静态内部类，实现了Map.Entry，存了hash，key，value，next
   3. Node池一定是2n长，方便-1后与哈希值能得到确定的位置。便于路由算法的执行
   4. 桶初始16，最大2的30次方，默认负载因子0.75
   5. Node桶定位的思路，取int的hashCode，然后mod桶长度
   6. 重复插入链表的时候，直接放头部，这样最快。其实就是放桶里，让后把原来桶的沿链表向下移
   7. 判重标准，hashCode方法返回值相&&equals

6. 扩容
   1. 负载因子，默认0.75
      1. 当前桶中，有75%的元素有值的时候，触发扩容。

   2. 扩容桶，重做hash
      1. 实际上桶都是*2扩容，看下最高位是0还是1，就好判断位置了

7. 数据机构
   1. 动态数组 + 链表 （红黑树）


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
```

### Hashtable

- Hashtable是synchronized的，适用于多线程环境；hashmap不是同步的，适用于单线程环境
- 也是**数组加链表**
  - HashTable未启用红黑树
- 不允许键为null
- 直接使用key的hashCode，而hashMap的哈希算法是自己的
- HashMap继承AbstractMap；Hashtable继承Dictionary

### LinkedHashMap

- 继承自Hashmap
- 在hashmap的基础上，使用了一对双向链表来记录元素添加的顺序

### TreeMap

- 有序集合
- 整体通过红黑树实现
