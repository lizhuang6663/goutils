// Package commonutils 提供通用的工具方法
// 包括字符串操作、随机数生成、数值处理等功能
package commonutils

import (
	"math/rand"
	"time"
)

// CommonUtils 包含通用的工具方法
type CommonUtils struct{}

// NewCommonUtils 创建一个新的 CommonUtils 实例
func NewCommonUtils() *CommonUtils {
	return &CommonUtils{}
}

// IsEmpty 检查字符串是否为空
// 参数:
//
//	s string: 输入字符串
//
// 返回值:
//
//	bool: 如果字符串为空返回 true，否则返回 false
func (cu *CommonUtils) IsEmpty(s string) bool {
	return len(s) == 0
}

// Reverse 反转字符串，支持 Unicode 字符
// 参数:
//
//	s string: 输入字符串
//
// 返回值:
//
//	string: 反转后的字符串
func (cu *CommonUtils) Reverse(s string) string {
	runes := []rune(s)
	length := len(runes)
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = runes[length-1-i]
	}
	return string(result)
}

// RandomString 生成指定长度的随机字符串
// 参数:
//
//	length int: 随机字符串的长度
//
// 返回值:
//
//	string: 生成的随机字符串
func (cu *CommonUtils) RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

// Clamp 限制数值在指定范围内
// 参数:
//
//	value int: 输入值
//	min int: 最小值
//	max int: 最大值
//
// 返回值:
//
//	int: 限制后的值
func (cu *CommonUtils) Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
