// Package stringutils 提供字符串相关的工具方法，注重性能优化。
// 包括高效字符串拼接、转换、裁剪等功能，适用于高性能场景。
package stringutils

import (
	"strings"
	"unicode"
	"unsafe"
)

// StringUtils 包含字符串相关的工具方法
type StringUtils struct{}

// NewStringUtils 创建一个 StringUtils 实例
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// Concat 使用 strings.Builder 高效拼接多个字符串
func (su *StringUtils) Concat(strs ...string) string {
	var builder strings.Builder
	builder.Grow(estimateLength(strs))
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// estimateLength 估算拼接字符串的总长度，用于预分配内存
func estimateLength(strs []string) int {
	total := 0
	for _, s := range strs {
		total += len(s)
	}
	return total
}

// FastStringToBytes 使用 unsafe 包将字符串转换为字节切片，无内存拷贝
// 注意: 使用 unsafe 操作需谨慎，确保字节切片生命周期不超过字符串
func (su *StringUtils) FastStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// FastBytesToString 使用 unsafe 包将字节切片转换为字符串，无内存拷贝
// 注意: 使用 unsafe 操作需谨慎，确保字符串生命周期不超过字节切片
func (su *StringUtils) FastBytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// TrimWhitespace 移除字符串首尾的空白字符，支持 Unicode
func (su *StringUtils) TrimWhitespace(s string) string {
	return strings.TrimFunc(s, unicode.IsSpace)
}

// Capitalize 将字符串的首字母转换为大写
func (su *StringUtils) Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ContainsAny 检查字符串是否包含任意指定的子字符串
func (su *StringUtils) ContainsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// SplitByLength 按指定长度分割字符串
func (su *StringUtils) SplitByLength(s string, length int) []string {
	if length <= 0 {
		return []string{s}
	}
	result := make([]string, 0, (len(s)+length-1)/length)
	for i := 0; i < len(s); i += length {
		end := i + length
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}
