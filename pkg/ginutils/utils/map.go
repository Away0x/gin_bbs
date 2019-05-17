package utils

// MergeMap 合并 map
func MergeMap(o1, o2 map[string]interface{}) map[string]interface{} {
	o3 := map[string]interface{}{}

	for k, v := range o1 {
		o3[k] = v
	}
	for k, v := range o2 {
		o3[k] = v
	}

	return o3
}
