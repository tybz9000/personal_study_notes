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


