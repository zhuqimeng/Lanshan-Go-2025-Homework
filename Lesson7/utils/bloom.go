package utils

import (
	"context"
	"math"

	"github.com/redis/go-redis/v9"
	"github.com/spaolacci/murmur3"
)

// BloomFilter Redis布隆过滤器实现
type BloomFilter struct {
	client     *redis.Client
	key        string   // Redis中的bitmap key
	m          uint64   // 位数组长度
	k          uint64   // 哈希函数数量
	hashSeeds  []uint32 // 哈希种子
	bitOffsets []int64  // 位偏移量缓存
}

// NewBloomFilter 创建布隆过滤器
func NewBloomFilter(client *redis.Client, key string, expectedItems uint64, falsePositiveRate float64) *BloomFilter {
	// 计算最优参数
	m := calculateM(expectedItems, falsePositiveRate)
	k := calculateK(expectedItems, m)

	// 生成哈希种子
	hashSeeds := make([]uint32, k)
	for i := uint64(0); i < k; i++ {
		hashSeeds[i] = uint32(i*1000 + 131) // 使用不同的种子
	}

	return &BloomFilter{
		client:    client,
		key:       key,
		m:         m,
		k:         k,
		hashSeeds: hashSeeds,
	}
}

// calculateM 计算位数组长度
func calculateM(n uint64, p float64) uint64 {
	return uint64(math.Ceil(-float64(n) * math.Log(p) / (math.Log(2) * math.Log(2))))
}

// calculateK 计算哈希函数数量
func calculateK(n, m uint64) uint64 {
	return uint64(math.Ceil(float64(m) / float64(n) * math.Log(2)))
}

// getBitPositions 获取元素的所有位位置
func (bf *BloomFilter) getBitPositions(data []byte) []int64 {
	positions := make([]int64, bf.k)

	// 使用Murmur3哈希
	hash1 := murmur3.Sum32WithSeed(data, bf.hashSeeds[0])
	hash2 := murmur3.Sum32WithSeed(data, bf.hashSeeds[1])

	// 使用双重哈希法生成多个哈希值
	for i := uint64(0); i < bf.k; i++ {
		combined := uint64(hash1) + i*uint64(hash2)
		position := combined % bf.m
		positions[i] = int64(position)
	}

	return positions
}

// Add 添加元素到布隆过滤器
func (bf *BloomFilter) Add(ctx context.Context, item string) error {
	data := []byte(item)
	positions := bf.getBitPositions(data)

	// 批量设置位
	pipe := bf.client.Pipeline()
	for _, pos := range positions {
		pipe.SetBit(ctx, bf.key, pos, 1)
	}

	_, err := pipe.Exec(ctx)
	return err
}

// Exists 检查元素是否可能存在
func (bf *BloomFilter) Exists(ctx context.Context, item string) (bool, error) {
	data := []byte(item)
	positions := bf.getBitPositions(data)

	// 批量获取位值
	pipe := bf.client.Pipeline()
	cmdS := make([]*redis.IntCmd, len(positions))
	for i, pos := range positions {
		cmdS[i] = pipe.GetBit(ctx, bf.key, pos)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// 检查所有位是否为1
	for _, cmd := range cmdS {
		if cmd.Val() != 1 {
			return false, nil
		}
	}
	return true, nil
}

// Clear 清空布隆过滤器
func (bf *BloomFilter) Clear(ctx context.Context) error {
	return bf.client.Del(ctx, bf.key).Err()
}
