// Package stringutils 提供字符串工具方法的单元测试
package stringutils

import (
	"reflect"
	"testing"
)

// TestStringUtils_Concat 测试 Concat 方法
func TestStringUtils_Concat(t *testing.T) {
	su := NewStringUtils()
	input := []string{"hello", "world", "!"}
	expected := "helloworld!"
	result := su.Concat(input...)
	if result != expected {
		t.Errorf("Concat failed, expected %s, got %s", expected, result)
	}
}

// TestStringUtils_TrimWhitespace 测试 TrimWhitespace 方法
func TestStringUtils_TrimWhitespace(t *testing.T) {
	su := NewStringUtils()
	input := "  hello world  "
	expected := "hello world"
	result := su.TrimWhitespace(input)
	if result != expected {
		t.Errorf("TrimWhitespace failed, expected %s, got %s", expected, result)
	}
}

// TestStringUtils_Capitalize 测试 Capitalize 方法
func TestStringUtils_Capitalize(t *testing.T) {
	su := NewStringUtils()
	input := "hello"
	expected := "Hello"
	result := su.Capitalize(input)
	if result != expected {
		t.Errorf("Capitalize failed, expected %s, got %s", expected, result)
	}
}

// TestStringUtils_SplitByLength 测试 SplitByLength 方法
func TestStringUtils_SplitByLength(t *testing.T) {
	su := NewStringUtils()
	input := "abcdefgh"
	expected := []string{"abc", "def", "gh"}
	result := su.SplitByLength(input, 3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SplitByLength failed, expected %v, got %v", expected, result)
	}
}
