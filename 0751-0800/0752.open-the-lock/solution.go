package openthelock

// 方法一: BFS
func openLock(deadends []string, target string) int {
	ds := map[string]bool{}
	for _, v := range deadends {
		ds[v] = true
	}
	vis := map[string]bool{}
	q := []string{"0000"}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, s := range tmp {
			for i, x := range s {
				next := []string{}
				if x-'0' == 0 {
					next = append(next, s[:i]+"9"+s[i+1:])
				} else {
					next = append(next, s[:i]+string(x-1)+s[i+1:])
				}
				next = append(next, s[:i]+string('0'+(x-'0'+1)%10)+s[i+1:])
				for _, v := range next {
					if ds[v] {
						continue
					}
					if v == target {
						return step + 1
					}
					if !vis[v] {
						q = append(q, v)
						vis[v] = true
					}
				}
			}
		}
	}
	return -1
}

// 方法二: 双向 BFS 优化
func openLock1(deadends []string, target string) int {
	type void struct{}
	deads := map[string]bool{}
	for _, v := range deadends {
		deads[v] = true
	}
	// 用集合不用队列，可以快速判断元素是否存在
	q1, q2 := map[string]void{}, map[string]void{}
	visited := map[string]bool{}

	// step := 0
	q1["0000"] = void{}
	q2[target] = void{}

	for step := 0; len(q1) > 0 && len(q2) > 0; step++ {
		// 哈希集合在遍历的过程中不能修改，用 temp 存储扩散结果
		tmp := make(map[string]void)

		/* 将 q1 中的所有节点向周围扩散 */
		for s := range q1 {
			/* 判断是否到达终点 */
			if deads[s] {
				continue
			}

			if _, ok := q2[s]; ok {
				return step
			}

			visited[s] = true

			/* 将一个节点的未遍历相邻节点加入集合 */
			for i, x := range s {
				next := ""
				// 向下转动
				if x-'0' == 0 {
					next = s[:i] + "9" + s[i+1:]
				} else {
					next = s[:i] + string(x-1) + s[i+1:]
				}
				if !visited[next] {
					// 加入队列
					tmp[next] = void{}
				}
				// 向上转动
				next = s[:i] + string('0'+(x-'0'+1)%10) + s[i+1:]
				if !visited[next] {
					// 加入队列
					tmp[next] = void{}
				}
			}
		}
		// 这里交换 q1 q2，下一轮 for 就是扩散 q2
		q1, q2 = q2, tmp
	}
	return -1
}
