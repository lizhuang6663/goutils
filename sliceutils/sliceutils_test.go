// Package sliceutils 提供切片工具方法的单元测试
package sliceutils

import (
	"reflect"
	"strings"
	"testing"
)

// TestSliceUtils_Deduplicate 测试 Deduplicate 方法
func TestSliceUtils_Deduplicate(t *testing.T) {
	su := NewSliceUtils()
	input := []string{"a", "b", "a", "c", "b"}
	expected := []string{"a", "b", "c"}
	result := su.Deduplicate(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Deduplicate failed, expected %v, got %v", expected, result)
	}
}

// TestSliceUtils_Filter 测试 Filter 方法
func TestSliceUtils_Filter(t *testing.T) {
	su := NewSliceUtils()
	input := []string{"apple", "banana", "app"}
	expected := []string{"apple", "app"}
	result := su.Filter(input, func(s string) bool {
		return strings.HasPrefix(s, "app")
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter failed, expected %v, got %v", expected, result)
	}
}

// TestSliceUtils_Map 测试 Map 方法
func TestSliceUtils_Map(t *testing.T) {
	su := NewSliceUtils()
	input := []string{"a", "b", "c"}
	expected := []string{"A", "B", "C"}
	result := su.Map(input, strings.ToUpper)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map failed, expected %v, got %v", expected, result)
	}
}

// TestSliceUtils_Reverse 测试 Reverse 方法
func TestSliceUtils_Reverse(t *testing.T) {
	su := NewSliceUtils()
	input := []string{"a", "b", "c"}
	expected := []string{"c", "b", "a"}
	result := su.Reverse(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse failed, expected %v, got %v", expected, result)
	}
}
