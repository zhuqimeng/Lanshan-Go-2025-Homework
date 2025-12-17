package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

const DefaultSize = 256

// 设置种子，保证不同哈希函数有不同的计算方式
var seeds = []uint{7, 11, 13, 31, 37, 61}

// 布隆过滤器结构，包括二进制数组和多个哈希函数
type BloomFilter struct {
	set       *bitset.BitSet
	hashFuncs [6]func(seed uint, value string) uint
}

// 构造一个布隆过滤器，包括数组和哈希函数的初始化
func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	bf.set = bitset.New(DefaultSize)
	for i := 0; i < len(bf.hashFuncs); i++ {
		bf.hashFuncs[i] = createHash()
	}
	return bf
}

// 构造6个哈希函数，每个哈希函数有参数seed保证计算方式的不同
func createHash() func(seed uint, value string) uint {
	return func(seed uint, value string) uint {
		var result uint = 0
		for i := 0; i < len(value); i++ {
			result = result*seed + uint(value[i])
		}
		return result & (DefaultSize - 1)
	}
}

// 添加元素
func (b *BloomFilter) add(value string) {
	for i, f := range b.hashFuncs {
		b.set.Set(f(seeds[i], value))
	}
}

// 判断元素是否存在
func (b *BloomFilter) contains(value string) bool {
	for i, f := range b.hashFuncs {
		if !b.set.Test(f(seeds[i], value)) {
			return false
		}
	}
	return true
}
func main() {
	filter := NewBloomFilter()
	filter.add("asd")
	fmt.Println(filter.contains("asd"))    // 输出: true
	fmt.Println(filter.contains("2222"))   // 输出: false
	fmt.Println(filter.contains("155343")) // 输出: false
}
