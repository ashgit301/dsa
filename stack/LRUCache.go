
package stack


type DNode struct {
	Key   int
	Value int
	Prev  *DNode
	Next  *DNode
}

type LRUCache struct {
	Cache    map[int]*DNode
	Size     int
	Capacity int
	Head     *DNode
	Tail     *DNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Head:     &DNode{},
		Tail:     &DNode{},
		Cache:    make(map[int]*DNode, capacity),
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (this *LRUCache) AddToHead(node *DNode) {
	temp := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	node.Next = temp
	temp.Prev = node
}

func (this *LRUCache) Remove(node *DNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	//node.prev.next = node.next
	//node.next.prev = node.prev
}

// Get get: if exist, move to head, return key, else return -1
func (this *LRUCache) Get(key int) int {
	if this.Cache[key] != nil {
		node := this.Cache[key]
		this.Remove(node)
		this.AddToHead(node)
		return node.Value
	}
	return -1
}

// Put put: check if key exist, if exist then delete
// if not full add to head, else evict tail then add to head
func (this *LRUCache) Put(key int, value int) {
	// key in cache
	if node := this.Cache[key]; node != nil {
		delete(this.Cache, node.Key)
		this.Remove(node)
	}
	// create new node from key value
	newNode := &DNode{
		Key:   key,
		Value: value,
	}
	// not full
	if len(this.Cache) < this.Capacity {
		this.Cache[key] = newNode
		this.AddToHead(newNode)
		return
	}
	// cache full, evict curr tail
	tailNode := this.Tail.Prev
	delete(this.Cache, tailNode.Key)
	this.Remove(tailNode)
	// append new node to head
	this.Cache[key] = newNode
	this.AddToHead(newNode)
}