// Package sliceutils 提供切片相关的工具方法
// 包括切片元素删除、去重、过滤、映射等功能，注重内存效率
package sliceutils

// SliceUtils 包含切片相关的工具方法
type SliceUtils struct{}

// NewSliceUtils 创建一个 SliceUtils 实例
func NewSliceUtils() *SliceUtils {
	return &SliceUtils{}
}

// Remove 从字符串切片中删除指定索引的元素
func (su *SliceUtils) Remove(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

// Deduplicate 移除字符串切片中的重复元素
func (su *SliceUtils) Deduplicate(slice []string) []string {
	seen := make(map[string]struct{}, len(slice))
	result := make([]string, 0, len(slice))
	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Filter 根据谓词函数过滤切片元素
func (su *SliceUtils) Filter(slice []string, predicate func(string) bool) []string {
	result := make([]string, 0, len(slice))
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 对切片中的每个元素应用转换函数
func (su *SliceUtils) Map(slice []string, transform func(string) string) []string {
	result := make([]string, len(slice))
	for i, item := range slice {
		result[i] = transform(item)
	}
	return result
}

// Reverse 反转字符串切片
func (su *SliceUtils) Reverse(slice []string) []string {
	result := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = slice[len(slice)-1-i]
	}
	return result
}
