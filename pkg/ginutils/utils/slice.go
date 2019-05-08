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
