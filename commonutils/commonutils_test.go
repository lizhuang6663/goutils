// Package commonutils 提供通用工具方法的单元测试
package commonutils

import (
	"strings"
	"testing"
)

// TestCommonUtils_Reverse 测试 Reverse 方法
func TestCommonUtils_Reverse(t *testing.T) {
	cu := NewCommonUtils()
	input := "hello"
	expected := "olleh"
	result := cu.Reverse(input)
	if result != expected {
		t.Errorf("Reverse failed, expected %s, got %s", expected, result)
	}
}

// TestCommonUtils_RandomString 测试 RandomString 方法
func TestCommonUtils_RandomString(t *testing.T) {
	cu := NewCommonUtils()
	result := cu.RandomString(10)
	if len(result) != 10 {
		t.Errorf("RandomString failed, expected length 10, got %d", len(result))
	}
	if strings.ContainsAny(result, " !@#$%^&*()") {
		t.Errorf("RandomString contains invalid characters")
	}
}

// TestCommonUtils_Clamp 测试 Clamp 方法
func TestCommonUtils_Clamp(t *testing.T) {
	cu := NewCommonUtils()
	tests := []struct {
		value, min, max, expected int
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{15, 0, 10, 10},
	}
	for _, test := range tests {
		result := cu.Clamp(test.value, test.min, test.max)
		if result != test.expected {
			t.Errorf("Clamp failed, value=%d, min=%d, max=%d, expected %d, got %d",
				test.value, test.min, test.max, test.expected, result)
		}
	}
}
