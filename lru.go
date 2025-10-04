package main

import (
	"container/list"
	"fmt"
)

func TestLRUCache() {
	c := NewLRUCache[int](2)
	_, a := c.Get(1)
	fmt.Println(a == false)
	c.Put(1, 'a')
	t, a := c.Get(1)
	fmt.Println(a == true)
	fmt.Println(t == 'a')
	c.Put(2, 'a')
	c.Put(3, 'a')
	_, a = c.Get(1)
	fmt.Println(a == false)

	c1 := NewLRUCache[string](2)
	c1.Put("aa", 1)
	c1.Put("ab", struct{}{})
	c1.Get("aa")
	c1.Put("ac", 2)
	t1, a1 := c1.Get("aa")
	fmt.Println(a1 == true)
	fmt.Println(t1 == 1)
	t1, a1 = c1.Get("ab")
	fmt.Println(a1 == false)
	t1, a1 = c1.Get("ac")
	fmt.Println(a1 == true)
	fmt.Println(t1 == 2)
}

type LRUCacheElement[T comparable] struct {
	key   T
	value any
}

type LRUCache[T comparable] struct {
	l   list.List
	m   map[T]*list.Element
	len int
}

func NewLRUCache[T comparable](len int) *LRUCache[T] {
	c := &LRUCache[T]{}
	c.l = list.List{}
	c.m = make(map[T]*list.Element)
	c.len = len
	return c
}

func (c *LRUCache[T]) Put(key T, value any) {
	if e, ok := c.m[key]; ok {
		c.l.MoveToFront(e)
		e.Value.(*LRUCacheElement[T]).value = value
		return
	}

	if len(c.m) >= c.len {
		e := c.l.Back()
		delete(c.m, e.Value.(LRUCacheElement[T]).key)
	}

	c.l.PushFront(LRUCacheElement[T]{key, value})
	c.m[key] = c.l.Front()
}

func (c *LRUCache[T]) Get(key T) (any, bool) {
	if e, ok := c.m[key]; ok {
		c.l.MoveToFront(e)
		return e.Value.(LRUCacheElement[T]).value, true
	}
	return nil, false
}
