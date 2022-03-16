package alloonedatastructure

import "container/list"

type node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (l *AllOne) Inc(key string) {
	if cur, ok := l.nodes[key]; !ok {
		// Key 不在链表中
		if l.Front() == nil || l.Front().Value.(node).count > 1 {
			l.nodes[key] = l.PushFront(node{map[string]struct{}{key: {}}, 1})
		} else {
			l.Front().Value.(node).keys[key] = struct{}{}
			l.nodes[key] = l.Front()
		}
	} else {
		curNode := cur.Value.(node)
		// key在链表中
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
			l.nodes[key] = l.InsertAfter(node{map[string]struct{}{key: {}}, curNode.count + 1}, cur)
		} else {
			nxt.Value.(node).keys[key] = struct{}{}
			l.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			l.Remove(cur)
		}
	}
}

func (l *AllOne) Dec(key string) {
	cur := l.nodes[key]
	curNode := cur.Value.(node)
	if curNode.count > 1 {
		if prev := cur.Prev(); prev == nil || prev.Value.(node).count < curNode.count-1 {
			l.nodes[key] = l.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
		} else {
			prev.Value.(node).keys[key] = struct{}{}
			l.nodes[key] = prev
		}
	} else {
		// key仅出现一次
		delete(l.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		l.Remove(cur)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Len() > 0 {
		for k := range this.Back().Value.(node).keys {
			return k
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Len() > 0 {
		for k := range this.Front().Value.(node).keys {
			return k
		}
	}
	return ""
}
