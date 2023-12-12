package logic

// 判断元素是否在数组中
func InArray(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
