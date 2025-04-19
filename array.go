package lrace

import "strings"

// ArrayCompare 比较数组是否一致, 非严格模式下顺序可不一致
func ArrayCompare[T comparable](arr1, arr2 []T, strict bool) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for k, v := range arr1 {
		if strict {
			if arr2[k] != v {
				return false
			}
		} else {
			if !InArray(arr2, v) {
				return false
			}
		}
	}

	return true
}

// InArray 在数组内
func InArray[T comparable](arr []T, noodle T) bool {
	for _, item := range arr {
		if item == noodle {
			return true
		}
	}
	return false
}

// HasArrayPrefix 数组中的字符以prefix开头
func HasArrayPrefix(prefix string, arr []string) bool {
	if arr == nil {
		return false
	}

	for _, s := range arr {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// ArrayToUpper 字符串数组转为大小
func ArrayToUpper(arr []string) []string {
	for k, v := range arr {
		arr[k] = strings.ToUpper(v)
	}
	return arr
}
