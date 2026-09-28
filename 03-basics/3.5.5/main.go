package homework

func toFrequencyMap(s []string) map[string]int {
	mp := make(map[string]int)
	for _, v := range s {
		mp[v]++
	}
	return mp
}
