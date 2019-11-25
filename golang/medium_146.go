package golang

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (44.32%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    23.9K
 * Total Submissions: 53.7K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
*/

// @lc code=start

func Constructor(capacity int) LRUCache {
	m := make(map[int]*CacheNode)
	c := LRUCache{Cap: capacity, Map: m}
	return c
}

func (this *LRUCache) Get(key int) int {
	found, ok := this.Map[key]
	if !ok {
		return -1
	}
	if this.Head == found {
		return found.Val
	}
	if this.Tail == found {
		this.Tail = found.Prev
	}
	// move found to head
	if found.Next != nil {
		found.Next.Prev = found.Prev
	}
	if found.Prev != nil {
		found.Prev.Next = found.Next
	}
	this.Head.Prev, found.Next = found, this.Head
	this.Head = found
	return found.Val
}

func (this *LRUCache) Put(key int, value int) {
	found, ok := this.Map[key]
	if ok {
		found.Val = value
		_ = this.Get(found.Key)
		return
	}

	// add to head
	n := &CacheNode{Key: key, Val: value}

	if len(this.Map) == 0 {
		this.Tail = n
	} else {
		this.Head.Prev, n.Next = n, this.Head
	}
	this.Map[key], this.Head = n, n
	if this.Cap == this.Len {
		delete(this.Map, this.Tail.Key)
		this.Len, this.Tail = this.Len-1, this.Tail.Prev
	}
	this.Len++
}

type CacheNode struct {
	Next *CacheNode
	Prev *CacheNode
	Key  int
	Val  int
}

type LRUCache struct {
	Cap  int
	Len  int
	Head *CacheNode
	Tail *CacheNode
	Map  map[int]*CacheNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
