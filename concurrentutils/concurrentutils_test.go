// Package concurrentutils 提供并发工具方法的单元测试
package concurrentutils

import (
	"sync"
	"testing"
)

// TestConcurrentUtils_ConcurrentConcat 测试 ConcurrentConcat 方法
func TestConcurrentUtils_ConcurrentConcat(t *testing.T) {
	cu := NewConcurrentUtils()
	input := []string{"hello", "world", "!"}
	expected := "helloworld!"
	result := cu.ConcurrentConcat(input...)
	if result != expected {
		t.Errorf("ConcurrentConcat failed, expected %s, got %s", expected, result)
	}
}

// TestSafeCounter 测试 SafeCounter 的并发安全性
func TestSafeCounter(t *testing.T) {
	sc := NewSafeCounter()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Increment()
		}()
	}
	wg.Wait()
	if sc.Value() != 100 {
		t.Errorf("SafeCounter failed, expected 100, got %d", sc.Value())
	}
}
