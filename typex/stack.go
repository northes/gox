package typex

import "sync"

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

func NewStack() *ItemStack {
	return &ItemStack{items: []Item{}}
}

// Push 入栈
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop 出栈
func (s *ItemStack) Pop() *Item {
	s.lock.RLock()
	defer s.lock.RLocker()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item
}

// Size 栈长度
func (s *ItemStack) Size() int {
	return len(s.items)
}
