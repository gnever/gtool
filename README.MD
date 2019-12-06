
## 集成了部分自己常用的工具


## 安装

推荐使用 `go.mod`:
```
require github.com/gnever/gtool latest
```

### package 

* `list` 并发安全的双向链表 

```
l := list.New()

l.PushBack(1)
l.PushBack(2)

fmt.Println(l.PopFront())

```

* `queue` 并发安全队列

```
q := queue.New()
q.Push(1)

for {
    if v := q.Pop(); v != nil {
        fmt.Println("for got - ", v)
    }
}

//或

for {
    select {
    case v := <-q.C:
        fmt.Println("select got - ", v)
    }
}

```

* workerPool 用于管理一定数量的 goroutine .用来限制 worker 的数量并且将消费和生产分离

```
w := New(20) //20 为启动的 worker 数量,默认为 10 

w.Add(func() {
        time.Sleep(time.Millisecond * 1)
        })

```

* `godaemon` 使程序在后台执行

除了可以 ``` nohup commond > /dev/null 2>&1 & ``` 这种方式启动后台程序外
还可以直接引入下面的 package 使程序直接支持 daemon 模式。

	
```
import _ "github.com/gnever/gtool/godaemon"
```

假设编译后程序包为 command ，直接执行 ``` command -d=true ``` 程序就进入后台模式了
