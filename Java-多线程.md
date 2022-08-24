## 多线程

并发编程三要素

- 原子性
- 可见性：
  - 线程的变更，应该对其他线程可见
  - 常见的实现方法：锁
- 有序性

### 进程：

运行的程序，例如一个运行中的Java虚拟机实例

### 线程：

操作系统能够进行运算调度的最小单位。它被包含在进程之中，是进程中的实际运作单位

### 并发：

一个时间段内执行多个任务，可能是切换执行

### 并行：

同一个时刻，多个任务同时

### 继承Thread类  

Thread类本身就实现了Runnable接口

### 实现Runnable接口

- 实现Run方法
- Runnable接口是典型的函数式接口，只有一个run方法
- 执行方法要这样：new Thread(实现runnable的实例).start()
  - Thread的构造参数里有runnable
  - start是Thread的一个方法
  - 实现runnable接口，实际上是实现了Thread的具体执行部分
  - 调用start方法，关键是调用start0方法，这是个通过jni调用的native方法
    - 该方法，实现调用了系统线程
    - 回调了runable方法

### 实现Callable接口

- 实现Callable可以实现多线程返回值
- 实现call方法
- 可以抛出异常
- 使用ThreadPoolExecutor.submit(Callable c)来执行，将Callable对象推给线程池执行，返回Feature对象
- feature对象通过get()方法获取返回值

# Thread类

#### Thread类通过静态代理

- Thread是Runnable中执行方法的静态代理

### 守护线程

子线程结束时，守护线程自动结束

setDeamon(true)

#### 线程状态

```
public enum State {
    /**
     * Thread state for a thread which has not yet started.
     * 创建状态
     */
    NEW,

    /**
     * Thread state for a runnable thread.  A thread in the runnable
     * state is executing in the Java virtual machine but it may
     * be waiting for other resources from the operating system
     * such as processor.
     * 可执行，不一定在执行
     */
    RUNNABLE,

    /**
     * Thread state for a thread blocked waiting for a monitor lock.
     * A thread in the blocked state is waiting for a monitor lock
     * to enter a synchronized block/method or
     * reenter a synchronized block/method after calling
     * {@link Object#wait() Object.wait}.
     * 阻塞状态 等待monitor lock
     */
    BLOCKED,

    /**
     * Thread state for a waiting thread.
     * A thread is in the waiting state due to calling one of the
     * following methods:
     * <ul>
     *   <li>{@link Object#wait() Object.wait} with no timeout</li>
     *   <li>{@link #join() Thread.join} with no timeout</li>
     *   <li>{@link LockSupport#park() LockSupport.park}</li>
     * </ul>
     *
     * <p>A thread in the waiting state is waiting for another thread to
     * perform a particular action.
     *
     * For example, a thread that has called <tt>Object.wait()</tt>
     * on an object is waiting for another thread to call
     * <tt>Object.notify()</tt> or <tt>Object.notifyAll()</tt> on
     * that object. A thread that has called <tt>Thread.join()</tt>
     * is waiting for a specified thread to terminate.
     * 等待状态 等待其他线程
     */
    WAITING,

    /**
     * Thread state for a waiting thread with a specified waiting time.
     * A thread is in the timed waiting state due to calling one of
     * the following methods with a specified positive waiting time:
     * <ul>
     *   <li>{@link #sleep Thread.sleep}</li>
     *   <li>{@link Object#wait(long) Object.wait} with timeout</li>
     *   <li>{@link #join(long) Thread.join} with timeout</li>
     *   <li>{@link LockSupport#parkNanos LockSupport.parkNanos}</li>
     *   <li>{@link LockSupport#parkUntil LockSupport.parkUntil}</li>
     * </ul>
     * 定时等待状态
     */
    TIMED_WAITING,

    /**
     * Thread state for a terminated thread.
     * The thread has completed execution.
     * 结束状态
     */
    TERMINATED;
}
```

- 创建

- 就绪：

  - 创建，启动后进入就绪
  - 阻塞解除后
  - 运行释放cpu资源后

- 阻塞：运行，阻塞后阻塞

- 运行：就绪，获得cpu资源后运行

- 死亡：线程执行完成

  

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

随机唤醒一个在**一样的对象监视器**上等待的线程

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

### 成员变量

- name：线程名
- priority：优先级，表示cpu调度线程的可能性，值的范围是0到10
- daemon：是否是守护线程
  - 守护线程：始终运行的线程，jvm中无任何用户线程时，守护线程才销毁，jvm退出
- target：要执行的代码，runnable
- threadLocals
- inheritableThreadLocals：可继承的线程变量，子线程可以拿到父线程的这个
- tid：Thread Id
- threadStatus：线程状态

### 方法

- static currentThread()
  - 获取当前线程
- static yield()
  - 线程礼让
  - 执行到待执行
  - 如果cpu认为接下来还要执行这个，那么就会执行这个
- static sleep(long millis)
  - 当前线程睡眠
- interrupt()
  - 设置线程中断位
  - 处于阻塞状态的线程（被wait()，join()，sleep()阻塞的线程），会抛出异常，把中断状态掰回去
    - 其实是唤醒了阻塞中的线程
    - 不会中断运行中的线程
    - 一个线程不该由其他线程来强制中断或停止。这个方法只是通知线程应该中断了
- join()
  - 外部线程等待主线程完成
    - 调用后外部线程开始等待，状态是TIMED_WAITING
    - 内部线程完成后，外部线程继续

- dumpStack()
  - 输出栈信息
  
- #### Synchronized
  
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

# 线程池

*线程池*是一种多线程处理形式，处理过程中将任务添加到队列，然后在创建线程后自动启动这些任务

#### 线程池的优势

- 降低系统资源开销，复用线程，减少线程创建和销毁的开销
- 提高响应速度，想要有线程用，就有线程用，不用新创建
- 提高系统可靠性，不会无限制创建线程，被打挂

#### 线程池创建方式

- 通过ThreadPoolExecutor手动创建线程池

```java
ThreadPoolExecutor threadPool = new ThreadPoolExecutor(5, 10, 100, TimeUnit.SECONDS, new LinkedBlockingQueue<>(10));
threadPool.execute(() -> say("hello"))
```

- 通过Executors类创建线程池
  - 封装好了ThreadPoolExecutor构造器的参数
  - newFixedThreadPool(int nThreads)
    - 创建固定大小线程池
  - newSingleThreadExecutor()
    - 创建只有一个线程大小的线程池
  - newCachedThreadPool()
    - 创建不限大小的线程池

ThreadPoolExecutor

```
// Java线程池的完整构造函数
public ThreadPoolExecutor(
  int corePoolSize, 
  // 线程池长期维持的线程数，即使线程处于Idle状态，也不会回收。
  int maximumPoolSize, 
  // 线程数的上限
  long keepAliveTime, 
  TimeUnit unit, 
  // 超过corePoolSize的线程的idle时长，
  // 超过这个时间，多余的线程会被回收。
  BlockingQueue<Runnable> workQueue, 
  // 任务的排队队列
  ThreadFactory threadFactory, 
  // 新线程的产生方式
  RejectedExecutionHandler handler
  // 拒绝策略
  ) 
```

ExecutorService接口

- 包含submit方法

### 线程池状态

| 状态       | 含义                                                         |
| ---------- | ------------------------------------------------------------ |
| RUNNING    | 允许提交并处理任务                                           |
| SHUTDOWN   | 不允许提交新的任务，但是会处理完已提交的任务                 |
| STOP       | 不允许提交新的任务，也不会处理阻塞队列中未执行的任务，并设置正在执行的线程的中断标志位 |
| TIDYING    | 所有任务执行完毕，池中工作的线程数为0，等待执行terminated()勾子方法 |
| TERMINATED | terminated()勾子方法执行完毕                                 |

