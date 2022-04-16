package miniparser

import "unicode"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val  int
	list []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.list == nil
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.list
}

func deserialize(s string) *NestedInteger {
	index := 0
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[index] == '[' {
			index++
			for s[index] != ']' {
				ni.Add(*dfs())
				if s[index] == ',' {
					index++
				}
			}
			index++
			return ni
		}

		negative := s[index] == '-'
		if negative {
			index++
		}
		num := 0
		for ; index < len(s) && unicode.IsDigit(rune(s[index])); index++ {
			num = num*10 + int(s[index]-'0')
		}
		if negative {
			num = -num
		}
		ni.SetInteger(num)
		return ni
	}
	return dfs()
}
