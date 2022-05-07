package minimumgeneticmutation

func minMutation(start string, end string, bank []string) (ans int) {
	if start == end {
		return 0
	}
	bankSet := map[string]struct{}{}
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}

	if _, ok := bankSet[end]; !ok {
		return -1
	}

	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		// 每个序列一轮变化，共3x8=24中子序列，如果子序列合法
		// 继续加入队列进行下一轮变化
		for _, cur := range tmp {
			// x 8
			for i, x := range cur {
				// x 3
				for _, y := range "ACGT" {
					if x != y {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}

	return -1
}
