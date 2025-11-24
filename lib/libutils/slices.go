package libutils

import "strconv"

type slicesUtil struct{}

var SlicesUtil = &slicesUtil{}

func (sl *slicesUtil) IsInSliceInt(value int, s []int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func (sl *slicesUtil) IsInSliceStr(value string, s []string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func (sl *slicesUtil) IntSliceToStringSlice(s []int) []string {
	strS := make([]string, len(s), len(s))
	for idx, v := range s {
		strS[idx] = strconv.Itoa(v)
	}

	return strS
}

func (sl *slicesUtil) StringSliceToIntSlice(s []string) ([]int, error) {
	intS := make([]int, len(s), len(s))
	var err error
	for idx, v := range s {
		intV, err1 := strconv.Atoi(v)
		if err == nil && err1 != nil {
			err = err1
		}
		intS[idx] = intV
	}

	return intS, err
}

// IntersectionIntSlice 取两个 int slice 的交集
// 返回同时存在于两个 slice 中的元素（去重）
func (sl *slicesUtil) IntersectionIntSlice(s1, s2 []int) []int {
	if len(s1) == 0 || len(s2) == 0 {
		return []int{}
	}

	// 使用 map 记录 s2 中的元素，提高查找效率
	s2Map := make(map[int]bool)
	for _, v := range s2 {
		s2Map[v] = true
	}

	// 使用 map 去重结果
	resultMap := make(map[int]bool)
	var result []int

	// 遍历 s1，如果元素在 s2 中存在，则加入结果
	for _, v := range s1 {
		if s2Map[v] && !resultMap[v] {
			resultMap[v] = true
			result = append(result, v)
		}
	}

	return result
}

// UnionIntSlice 取两个 int slice 的并集
// 返回两个 slice 中所有不重复的元素
func (s *slicesUtil) UnionIntSlice(s1, s2 []int) []int {
	// 使用 map 去重
	resultMap := make(map[int]bool)
	var result []int

	// 添加 s1 中的所有元素
	for _, v := range s1 {
		if !resultMap[v] {
			resultMap[v] = true
			result = append(result, v)
		}
	}

	// 添加 s2 中的所有元素
	for _, v := range s2 {
		if !resultMap[v] {
			resultMap[v] = true
			result = append(result, v)
		}
	}

	return result
}
