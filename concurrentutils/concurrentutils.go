// Package concurrentutils 提供并发安全的工具方法
// 包括并发字符串拼接和线程安全计数器，优化内存分配
package concurrentutils

import (
	"strings"
	"sync"
)

// ConcurrentUtils 包含并发安全的工具方法
type ConcurrentUtils struct {
	pool sync.Pool
}

// NewConcurrentUtils 创建一个新的 ConcurrentUtils 实例
func NewConcurrentUtils() *ConcurrentUtils {
	return &ConcurrentUtils{
		pool: sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
	}
}

// ConcurrentConcat 使用 sync.Pool 并发安全地拼接字符串
// 参数:
//
//	strs ...string: 要拼接的字符串列表
//
// 返回值:
//
//	string: 拼接后的字符串
func (cu *ConcurrentUtils) ConcurrentConcat(strs ...string) string {
	builder := cu.pool.Get().(*strings.Builder)
	defer cu.pool.Put(builder)
	builder.Reset()
	builder.Grow(estimateLength(strs))
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// estimateLength 估算拼接字符串的总长度，用于预分配内存
// 参数:
//
//	strs []string: 要拼接的字符串列表
//
// 返回值:
//
//	int: 估算的总长度
func estimateLength(strs []string) int {
	total := 0
	for _, s := range strs {
		total += len(s)
	}
	return total
}

// SafeCounter 是一个线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int64
}

// NewSafeCounter 创建一个新的 SafeCounter 实例
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Increment 增加计数器的值
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	sc.count++
	sc.mu.Unlock()
}

// Value 获取当前计数器的值
// 返回值:
//
//	int64: 当前计数
func (sc *SafeCounter) Value() int64 {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}
