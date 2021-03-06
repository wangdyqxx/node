## go调度

https://blog.csdn.net/yizhou35/article/details/108977664

### 3个缩写，请一定记住
G: goroutine

M: 工作线程

P: 处理器，它包含了运行Go代码的资源，M必须和一个P关联才能运行G。

### 调度器的有两大思想：

![调度器](../img/81f5f8dc9e928b041.png)

 * 复用线程：协程本身就是运行在一组线程之上，不需要频繁的创建、销毁线程，而是对线程的复用。
在调度器中复用线程还有2个体现：
 	1）work stealing，当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。
 	2）hand off，当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

 * 利用并行：GOMAXPROCS设置P的数量，当GOMAXPROCS大于1时，就最多有GOMAXPROCS个线程处于运行状态，这些线程可能分布在多个CPU核上同时运行，使得并发利用并行。另外，GOMAXPROCS也限制了并发的程度，比如GOMAXPROCS = 核数/2，则最多利用了一半的CPU核进行并行。

### 调度器的两小策略：

  * 抢占：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU`10ms`，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

  * 全局G队列：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G。

Go是并发，不是并行，因为Go做的是在一段时间内完成几十万、甚至几百万的工作，而不是同一时间同时在做大量的工作。`并发可以利用并行提高效率，调度器是有并行设计的`

并行依赖多核技术，每个核上在某个时间只能执行一个线程，当我们的CPU有8个核时，我们能同时执行8个线程，这就是并行。

