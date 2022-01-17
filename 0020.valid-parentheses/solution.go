package validparentheses

import "container/list"

func isValid(s string) bool {

	stack := list.New()

	for i := range s {
		// 检查栈顶元素是否和 s[i] 匹配
		// 匹配规则: [] ()
		if stack.Len() > 0 {
			top := stack.Front().Value.(byte)
			if (top == '(' && s[i] == ')') || (top == '[' && s[i] == ']') || (top == '{' && s[i] == '}') {
				// 弹出栈顶
				stack.Remove(stack.Front())
				continue
			}
		}

		// push front
		stack.PushFront(s[i])
	}

	return stack.Len() == 0
}
