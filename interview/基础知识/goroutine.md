## goroutine

Goroutine是建立在线程之上的轻量级的抽象。它允许我们以非常低的代价在同一个地址空间中并行地执行多个函数或者方法。相比于线程，
它的创建和销毁的代价要小很多，并且它的调度是独立于线程的。

### Goroutine与线程的区别

许多人认为goroutine比线程运行得更快，这是一个误解。Goroutine并不会更快，它只是增加了更多的并发性。当一个goroutine被阻塞（比如等待IO），
golang的scheduler会调度其他可以执行的goroutine运行。与线程相比，它有以下的几个优点：

* 内存消耗更少：

Goroutine所需要的内存通常只有2kb，而线程则需要1Mb（500倍）

* 创建与销毁的开销更小：

由于线程创建时需要向操作系统申请资源，并且在销毁时将资源归还，因此它的创建和销毁的开销比较大。相比之下，goroutine的创建和销毁是由go语言在运行时自己管理的，因此开销更低。

* 切换开销更小：

只是goroutine之于线程的主要区别，也是golang能够实现高并发的主要原因。线程的调度方式是抢占式的，如果一个线程的执行时间超过了分配给它的时间片，
就会被其他可执行的线程抢占。在线程切换的过程中需要保存/恢复所有的寄存器信息，比如16个通用寄存器，PC（Program Counter）、SP（Stack Pointer）段寄存器。
而goroutine的调度是协同式的，它不会直接地与操作系统内核打交道。当goroutine进行切换的时候，之后很少量的寄存器需要保存和恢复（PC和SP）。
因此goroutine的切换效率更高。

### 如何优雅的退出goroutine

#### 一. 通过Channel传递退出信号

```golang
func run(done chan int) {
        for {
                select {
                case <-done:
                        fmt.Println("exiting...")
                        done <- 1
                        break
                default:
                }
 
                time.Sleep(time.Second * 1)
                fmt.Println("do something")
        }
}
 
func main() {
        c := make(chan int)
 
        go run(c)
 
        fmt.Println("wait")
        time.Sleep(time.Second * 5)
 
        c <- 1
        <-c
 
        fmt.Println("main exited")
}
```

#### 二. 使用waitgroup
通常情况下，我们像下面这样使用waitgroup:

1. 创建一个Waitgroup的实例，假设此处我们叫它wg
2. 在每个goroutine启动的时候，调用wg.Add(1)，这个操作可以在goroutine启动之前调用，也可以在goroutine里面调用。当然，也可以在创建n个goroutine前调用wg.Add(n)
3. 当每个goroutine完成任务后，调用wg.Done()
4. 在等待所有goroutine的地方调用wg.Wait()，它在所有执行了wg.Add(1)的goroutine都调用完wg.Done()前阻塞，当所有goroutine都调用完wg.Done()之后它会返回。

```golang
type Service struct {
        // Other things
 
        ch        chan bool
        waitGroup *sync.WaitGroup
}
 
func NewService() *Service {
    s := &Service{
                // Init Other things
                ch:        make(chan bool),
                waitGroup: &sync.WaitGroup{},
    }
 
    return s
}
 
func (s *Service) Stop() {
        close(s.ch)
        s.waitGroup.Wait()
}
 
func (s *Service) Serve() {
        s.waitGroup.Add(1)
        defer s.waitGroup.Done()
 
        for {
                select {
                case <-s.ch:
                        fmt.Println("stopping...")
                        return
                default:
                }
                s.waitGroup.Add(1)
                go s.anotherServer()
    }
}
func (s *Service) anotherServer() {
        defer s.waitGroup.Done()
        for {
                select {
                case <-s.ch:
                        fmt.Println("stopping...")
                        return
                default:
                }
 
                // Do something
        }
}
 
func main() {
 
        service := NewService()
        go service.Serve()
 
        // Handle SIGINT and SIGTERM.
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        fmt.Println(<-ch)
 
        // Stop the service gracefully.
        service.Stop()
}
```
