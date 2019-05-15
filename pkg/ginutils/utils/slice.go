package utils

// UniqueUintSlice uint slice 去重
func UniqueUintSlice(s []uint) []uint {
	result := []uint{}
	m := make(map[uint]bool)

	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}

	return result
}

// InStringSlice 是否存在于 slice 中
func InStringSlice(s []string, o string) bool {
	for _, v := range s {
		if v == o {
			return true
		}
	}

	return false
}

// InIntSlice 是否存在于 slice 中
func InIntSlice(s []int, o int) bool {
	for _, v := range s {
		if v == o {
			return true
		}
	}

	return false
}
