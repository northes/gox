package typex

import "sync"

type QueueItem struct {
	items []Item
	lock  sync.RWMutex
}

// NewQueue 创建队列
func NewQueue() *QueueItem {
	return &QueueItem{
		items: []Item{},
	}
}

// Enqueue 入队列
func (q *QueueItem) Enqueue(t Item) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, t)
}

// Dequeue 出队列
func (q *QueueItem) Dequeue() *Item {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.IsEmpty() {
		return &Item{}
	}
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &item
}

func (q *QueueItem) List() []Item {
	return q.items
}

// Front 获取第一个元素，不移除
func (q *QueueItem) Front() *Item {
	q.lock.RLock()
	defer q.lock.RUnlock()
	item := q.items[0]
	return &item
}

// IsEmpty 判空
func (q *QueueItem) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 队列长度
func (q *QueueItem) Size() int {
	return len(q.items)
}
