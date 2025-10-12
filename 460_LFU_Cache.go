package main

import (
	"container/list"
	"fmt"
)

func TestLFUCache() {
	//LFUCache lfu = new LFUCache(2);
	lfu := ConstructorLFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(len(lfu.dataMap) == 2)
	fmt.Println(len(lfu.countMap) == 1)
	fmt.Println(lfu.countList.Len() == 1)
	fmt.Println(len(lfu.listMap) == 1)
	fmt.Println(lfu.dataMap[1].Value.(*TListLeafElement).Count == 1)

	fmt.Println(lfu.Get(1) == 1)
	fmt.Println(lfu.dataMap[1].Value.(*TListLeafElement).Count == 2)
	fmt.Println(len(lfu.countMap) == 2)
	fmt.Println(lfu.countList.Len() == 2)
	fmt.Println(len(lfu.listMap) == 2)
	//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
	//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
	//lfu.get(1);      // return 1
	//                 // cache=[1,2], cnt(2)=1, cnt(1)=2

	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2) == -1)
	fmt.Println(lfu.dataMap[3].Value.(*TListLeafElement).Count == 1)
	fmt.Println(lfu.Get(3) == 3)
	fmt.Println(lfu.dataMap[3].Value.(*TListLeafElement).Count == 2)
	fmt.Println(lfu.dataMap[1].Value.(*TListLeafElement).Count == 2)
	//lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest,
	//invalidate 2.
	//                 // cache=[3,1], cnt(3)=1, cnt(1)=2
	//lfu.get(2);      // return -1 (not found)
	//lfu.get(3);      // return 3
	//                 // cache=[3,1], cnt(3)=2, cnt(1)=2

	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1) == -1)
	fmt.Println(lfu.dataMap[3].Value.(*TListLeafElement).Count == 2)
	fmt.Println(lfu.dataMap[4].Value.(*TListLeafElement).Count == 1)
	fmt.Println(lfu.Get(3) == 3)
	fmt.Println(lfu.dataMap[3].Value.(*TListLeafElement).Count == 3)
	fmt.Println(lfu.Get(4) == 4)
	fmt.Println(lfu.dataMap[4].Value.(*TListLeafElement).Count == 2)
	//lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1
	//.
	//                 // cache=[4,3], cnt(4)=1, cnt(3)=2
	//lfu.get(1);      // return -1 (not found)
	//lfu.get(3);      // return 3
	//                 // cache=[3,4], cnt(4)=1, cnt(3)=3
	//lfu.get(4);      // return 4
	//                 // cache=[4,3], cnt(4)=2, cnt(3)=3
	lfu.Put(4, 5)
	fmt.Println(lfu.Get(4) == 5)
}

type TListLeafElement struct {
	Value int
	Count int
	Key   int
}

type LFUCache struct {
	capacity  int
	listMap   map[int]*list.List
	dataMap   map[int]*list.Element
	countList *list.List
	countMap  map[int]*list.Element
}

func ConstructorLFU(capacity int) LFUCache {
	lfu := LFUCache{capacity: capacity}
	lfu.listMap = make(map[int]*list.List)
	lfu.dataMap = make(map[int]*list.Element)
	lfu.countList = list.New()
	lfu.countMap = make(map[int]*list.Element)
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if vv, ok := this.dataMap[key]; ok {
		return this.updateCount(vv)
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if v, ok := this.dataMap[key]; ok {
		this.updateCount(v)
		this.dataMap[key].Value.(*TListLeafElement).Value = value
		return
	}

	if len(this.dataMap) >= this.capacity {
		this.removeLastOne()
	}

	this.insertNewOne(key, value)
}

func (this *LFUCache) removeLastOne() {
	mc := this.countList.Back().Value.(int) // 访问次数少的在后
	l := this.listMap[mc]
	v := l.Front() //最先使用的在后

	delete(this.dataMap, v.Value.(*TListLeafElement).Key)
	l.Remove(v)
	if l.Len() == 0 {
		delete(this.listMap, mc)
		this.countList.Remove(this.countMap[mc])
		delete(this.countMap, mc)
	}
}

func (this *LFUCache) insertNewOne(key, value int) {
	e := &TListLeafElement{value, 1, key}
	if _, ok := this.listMap[e.Count]; !ok {
		this.listMap[e.Count] = list.New()
		this.countList.PushBack(e.Count)
		this.countMap[e.Count] = this.countList.Back()
	}
	this.listMap[e.Count].PushBack(e)
	this.dataMap[key] = this.listMap[e.Count].Back()
}

func (this *LFUCache) updateCount(vv *list.Element) int {
	v := vv.Value.(*TListLeafElement)

	nn := &TListLeafElement{Value: v.Value, Count: v.Count + 1, Key: v.Key}
	if _, ok := this.listMap[nn.Count]; !ok {
		this.listMap[nn.Count] = list.New()
		this.countMap[nn.Count] = this.countList.InsertBefore(nn.Count, this.countMap[v.Count])
	}
	this.listMap[nn.Count].PushBack(nn)
	this.dataMap[nn.Key] = this.listMap[nn.Count].Back()

	l := this.listMap[v.Count]
	l.Remove(vv)
	if l.Len() == 0 {
		delete(this.listMap, v.Count)
		this.countList.Remove(this.countMap[v.Count])
		delete(this.countMap, v.Count)
	}
	return v.Value
}
