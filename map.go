package lrace

// MergeMaps 合并两个map
func MergeMaps(m1, m2 map[string]any) map[string]any {
	len1 := len(m1)
	len2 := len(m2)

	merged := make(map[string]any, Ternary(len1 > len2, len1, len2))
	for k, v := range m1 {
		merged[k] = v
	}

	for k, v := range m2 {
		merged[k] = v
	}

	return merged
}
