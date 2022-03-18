package minimumindexsumoftwolists

func findRestaurant(list1 []string, list2 []string) (ans []string) {
	m := map[string]int{}
	for i := range list1 {
		m[list1[i]] = i
	}
	min_index := 1<<31 - 1

	for j := range list2 {
		if i, ok := m[list2[j]]; ok {
			if i+j > min_index {
				continue
			}
			if i+j < min_index {
				min_index = i + j
				ans = nil
			}
			ans = append(ans, list2[j])
		}
	}
	return ans
}
