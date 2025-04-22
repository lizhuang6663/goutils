// Package commonutils 提供通用的工具方法
// 包括字符串操作、随机数生成、数值处理等功能
package commonutils

import (
	"math/rand"
	"time"
)

// CommonUtils 包含通用的工具方法
type CommonUtils struct{}

// NewCommonUtils 创建一个 CommonUtils 实例
func NewCommonUtils() *CommonUtils {
	return &CommonUtils{}
}

// IsEmpty 检查字符串是否为空
func (cu *CommonUtils) IsEmpty(s string) bool {
	return len(s) == 0
}

// Reverse 反转字符串，支持 Unicode 字符
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
func (cu *CommonUtils) Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
