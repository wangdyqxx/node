## go channel

channel 是 goroutine 之间通信的一种方式，可以类比成 Unix 中的进程的通信方式管道。

### CSP模型
Go 语言的并发模型参考了 CSP 理论，其中执行实体对应的是 goroutine， 消息通道对应的就是 channel。

### channel 介绍
channel 提供了一种通信机制，通过它，一个 goroutine 可以向另一 goroutine 发送消息

### channel 创建

channel 使用内置的 make 函数创建，下面声明了一个 chan int 类型的 channel:

```golang
ch := make(chan int)
```

### channel 的读写操作
```golang
ch := make(chan int)

// write to channel
ch <- x

// read from channel
x <- ch

// another way to read
x = <- ch
```

channel 一定要初始化后才能进行读写操作，否则会永久阻塞。

### 关闭 channel

golang 提供了内置的 close 函数对 channel 进行关闭操作。

```golang
ch := make(chan int)

close(ch)
```

> 有关 channel 的关闭，你需要注意以下事项:

* 关闭一个未初始化(nil) 的 channel 会产生 panic
* 重复关闭同一个 channel 会产生 panic
* 向一个已关闭的 channel 中发送消息会产生 panic
* 从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
* 关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息

```golang
ch := make(chan int, 10)
ch <- 11
ch <- 12

close(ch)

for x := range ch {
    fmt.Println(x)
}

x, ok := <- ch
fmt.Println(x, ok)


-----
output:

11
12
0 false
```

### channel 的类型

channel 分为不带缓存的 channel 和带缓存的 channel。

### 无缓存的 channel

从无缓存的 channel 中读取消息会阻塞，直到有 goroutine 向该 channel 中发送消息；同理，向无缓存的 channel 中发送消息也会阻塞，直到有 goroutine 从 channel 中读取消息。

> 通过无缓存的 channel 进行通信时，接收者收到数据 happens before 发送者 goroutine 唤醒

### 有缓存的 channel

有缓存的 channel 的声明方式为指定 make 函数的第二个参数，该参数为 channel 缓存的容量

```golang
ch := make(chan int, 10)
```

有缓存的 channel 类似一个阻塞队列(采用环形数组实现)。当缓存未满时，向 channel 中发送消息时不会阻塞，当缓存满时，发送操作将被阻塞，直到有其他 goroutine 从中读取消息；相应的，当 channel 中消息不为空时，读取消息不会出现阻塞，当 channel 为空时，读取操作会造成阻塞，直到有 goroutine 向 channel 中写入消息。

通过 len 函数可以获得 chan 中的元素个数，通过 cap 函数可以得到 channel 的缓存长度。

### goroutine 通信

```golang
c := make(chan int)  // Allocate a channel.

// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()

doSomethingForAWhile()
<-c
```

主 goroutine 会阻塞，直到执行 sort 的 goroutine 完成。


### range 遍历
channel 也可以使用 range 取值，并且会一直从 channel 中读取数据，直到有 goroutine 对该 channel 执行 close 操作，循环才会结束。
```golang
// consumer worker
ch := make(chan int, 10)
for x := range ch{
    fmt.Println(x)
}
```
等价于
```golang
for {
    x, ok := <- ch
    if !ok {
        break
    }
    
    fmt.Println(x)
}
```

### 配合 select 使用

select 用法类似与 IO 多路复用，可以同时监听多个 channel 的消息状态，看下面的例子

```golang
select {
    case <- ch1:
    ...
    case <- ch2:
    ...
    case ch3 <- 10;
    ...
    default:
    ...
}
```

* select 可以同时监听多个 channel 的写入或读取
* 执行 select 时，若只有一个 case 通过(不阻塞)，则执行这个 case 块
* 若有多个 case 通过，则随机挑选一个 case 执行
* 若所有 case 均阻塞，且定义了 default 模块，则执行 default 模块。若未定义 default 模块，则 select 语句阻塞，直到有 case 被唤醒。
* 使用 break 会跳出 select 块。

#### 1. 设置超时时间

```
ch := make(chan struct{})

// finish task while send msg to ch
go doTask(ch)

timeout := time.After(5 * time.Second)
select {
    case <- ch:
        fmt.Println("task finished.")
    case <- timeout:
        fmt.Println("task timeout.")
}
```

#### 2. quite channel

有一些场景中，一些 worker goroutine 需要一直循环处理信息，直到收到 quit 信号

```golang
msgCh := make(chan struct{})
quitCh := make(chan struct{})
for {
    select {
    case <- msgCh:
        doWork()
    case <- quitCh:
        finish()
        return
}
```

### channel 类结构
```golang
// channel 类型定义
type hchan struct {
    // channel 中的元素数量, len
    qcount   uint           // total data in the queue
    
    // channel 的大小, cap
    dataqsiz uint           // size of the circular queue
    
    // channel 的缓冲区，环形数组实现
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    
    // 单个元素的大小
    elemsize uint16
    
    // closed 标志位
    closed   uint32
    
    // 元素的类型
    elemtype *_type // element type
    
    // send 和 recieve 的索引，用于实现环形数组队列
    sendx    uint   // send index
    recvx    uint   // receive index
    
    // recv goroutine 等待队列
    recvq    waitq  // list of recv waiters
    
    // send goroutine 等待队列
    sendq    waitq  // list of send waiters

    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
    lock mutex
}

// 等待队列的链表实现
type waitq struct {    
    first *sudog       
    last  *sudog       
}

// in src/runtime/runtime2.go
// 对 G 的封装
type sudog struct {
    // The following fields are protected by the hchan.lock of the
    // channel this sudog is blocking on. shrinkstack depends on
    // this for sudogs involved in channel ops.

    g          *g
    selectdone *uint32 // CAS to 1 to win select race (may point to stack)
    next       *sudog
    prev       *sudog
    elem       unsafe.Pointer // data element (may point to stack)

    // The following fields are never accessed concurrently.
    // For channels, waitlink is only accessed by g.
    // For semaphores, all fields (including the ones above)
    // are only accessed when holding a semaRoot lock.

    acquiretime int64
    releasetime int64
    ticket      uint32
    parent      *sudog // semaRoot binary tree
    waitlink    *sudog // g.waiting list or semaRoot
    waittail    *sudog // semaRoot
    c           *hchan // channel
}
```

channel 的主要组成有：一个环形数组实现的队列，用于存储消息元素；两个链表实现的 goroutine 等待队列，用于存储阻塞在 recv 和 send 操作上的 goroutine；一个互斥锁，用于各个属性变动的同步
