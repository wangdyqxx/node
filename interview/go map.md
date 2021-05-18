## go map
Map 是一种无序的键值对的集合。
Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的

## go map 是否是线程安全的？否
### 安全的使用map
* 方式一：sync.Map
```golang
var ma sync.Map// 该类型是开箱即用，只需要声明既可
ma.Store("key", "value") // 存储值
ma.Delete("key") //删除值
ma.LoadOrStore("key", "value")// 获取值，如果没有则存储
fmt.Println(ma.Load("key"))//获取值

//遍历
ma.Range(func(key, value interface{}) bool {
    fmt.Printf("key:%s ,value:%s \n", key, value)
    //如果返回：false，则退出循环，
    return true
})
```
* 方式二：增加同步机制
```golang
// 通过匿名结构体声明了一个变量counter，变量中包含了map和sync.RWMutex
var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}
// 读取数据的时候使用读锁
counter.RLock()
n := counter.m["Tony"]
counter.RUnlock()
// 写数据的使用使用写锁
counter.Lock()
counter.m["Tony"]++
counter.Unlock()
```

## 怎么在golang中释放map内存
golang释放map内存的方法：首先删除map中的所有key，map占用内存仍处于【使用状态】；然后map置为nil，map占用的内存处于【空闲状态】；最后处于空闲状态内存，一定时间内在下次申请的可重复被使用，不必再向操作系统申请即可。
golang的map在key被删除之后，并不会立即释放内存，所以随着程序的运行，实际上map占用的内存只会越来越大。此外，GC会在标记阶段访问map中的每一个元素，当map非常大时这会对程序性能带来非常大的开销。不过go 1.5版本之后，如果map的key和value中都不包含指针，那么GC会忽略这个map。