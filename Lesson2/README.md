# 朱启梦的第二次go组作业

对之前的部分代码编写了单元测试（任务2的函数工厂不太容易test，没写）

### 思考题回答

**Q1：如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么？**

A：切片之间会互相干扰，一个切片修改会导致其他切片被连锁修改；若想要规避这种问题，可以新开一个空的切片，对这个切片append其他切片，这样操作时不会改变其他切片和原数组。

**Q2：怎样沿用“扩容”的思想对切片进行“缩容”？**

A：沿用"扩容"的思想对切片进行"缩容"，主要是通过重新分配更小的底层数组来释放未使用的内存。
也就是说，新建一个切片，人为调整它的容量，实例代码如下：
```go
package main

import "fmt"

func shrinkSlice(s []int, newCap int) []int {
    if newCap >= cap(s) {
        return s // 不需要缩容
    }
    
    // 创建新的切片，容量为 newCap
    newSlice := make([]int, len(s), newCap)
    copy(newSlice, s)
    return newSlice
}

func main() {
    s := make([]int, 10, 20)
    fmt.Printf("原始 - 长度: %d, 容量: %d\n", len(s), cap(s))
    
    // 缩容到容量15
    s = shrinkSlice(s, 15)
    fmt.Printf("缩容后 - 长度: %d, 容量: %d\n", len(s), cap(s))
}
```

**Q3：函数真正拿到的参数值其实只是它们的副本，那么函数返回给调用方的结果值也会被复制吗？**

A：基本类型和数组会被复制，而切片，映射和指针的地址会被复制，但是指向的底层数据结构是共享的（指针是相同的）。