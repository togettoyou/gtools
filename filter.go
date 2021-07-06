package gtools

// RemoveDuplicateStr 字符串数组去重
func RemoveDuplicateStr(slice []string) []string {
	temp := map[string]struct{}{}
	result := make([]string, 0, len(slice))
	for _, e := range slice {
		l := len(temp)
		temp[e] = struct{}{}
		// 加入map后，map长度变化，则元素不重复
		if len(temp) != l {
			result = append(result, e)
		}
	}
	return result
}

// RemoveDuplicateInt 整型数组去重
func RemoveDuplicateInt(slice []int) []int {
	temp := map[int]struct{}{}
	result := make([]int, 0, len(slice))
	for _, e := range slice {
		l := len(temp)
		temp[e] = struct{}{}
		// 加入map后，map长度变化，则元素不重复
		if len(temp) != l {
			result = append(result, e)
		}
	}
	return result
}
