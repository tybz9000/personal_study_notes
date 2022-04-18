# Thread类

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
