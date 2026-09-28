package homework

func mapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	var res []int
	for k := range m1 {
		if _, ok := m2[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
