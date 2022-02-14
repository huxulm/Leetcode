package firstbadversion

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	l, r := 0, n-1

	for l < r {
		mid := int(uint(l+r) >> 1)
		if isBadVersion(mid + 1) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// l == r
	return r + 1
}

func firstBadVersion1(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
