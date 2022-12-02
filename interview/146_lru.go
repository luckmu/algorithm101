package interview

// Get() && Put() time-complexity be O(1)
// so -> map[key]what?
// LRU need seq, arr or linklist -> move ele in arr exceed O(1) -> linklist -> bi-direction linklist is always easier
// add Header && Tailer to bi-direction linklist is easier to operate (head -> linklist -> tail)
// find in linklist exceed O(1) -> map[key]*listnode

type LRUNode struct {
	Key, Val   int
	Prev, Next *LRUNode
}

type LRUCache struct {
	C int
	M map[int]*LRUNode
	H *LRUNode
	T *LRUNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		M: make(map[int]*LRUNode),
		H: &LRUNode{},
		T: &LRUNode{},
		C: capacity,
	}
	l.H.Next = l.T
	l.T.Prev = l.H
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.M[key]; !ok {
		return -1
	}
	node := this.M[key]
	this.MoveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.M[key]; ok {
		node.Val = value
		this.MoveToHead(node)
		return
	}
	node := &LRUNode{}
	node.Key, node.Val = key, value
	this.M[key] = node
	if len(this.M) > this.C {
		node := this.RemoveTail()
		delete(this.M, node.Key)
	}
	this.AddToHead(node)
}

func (this *LRUCache) AddToHead(node *LRUNode) {
	node.Prev = this.H
	node.Next = this.H.Next
	this.H.Next.Prev = node
	this.H.Next = node
}
func (this *LRUCache) MoveToHead(node *LRUNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}
func (this *LRUCache) RemoveNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev, node.Next = nil, nil
}
func (this *LRUCache) RemoveTail() *LRUNode {
	node := this.T.Prev
	this.RemoveNode(node)
	return node
}

func Q146() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
