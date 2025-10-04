package main

import (
	"container/list"
	"golang.org/x/exp/constraints"
)

type leaseMonotonic[T constraints.Ordered] struct {
	l list.List
}

func newLeaseMonotonic[T constraints.Ordered]() *leaseMonotonic[T] {
	return &leaseMonotonic[T]{l: list.List{}}
}

func (ms *leaseMonotonic[T]) push(value T) (T, bool) {
	// push data finally
	defer ms.l.PushBack(value)

	// pop the lease values
	for ms.l.Len() > 0 && ms.l.Back().Value.(T) > value {
		ms.l.Remove(ms.l.Back())
	}
	// calculate return value
	if ms.l.Len() > 0 {
		return ms.l.Back().Value.(T), true
	} else {
		var zero T
		return zero, false
	}
}

func (ms *leaseMonotonic[T]) top() (T, bool) {
	if ms.l.Len() == 0 {
		var zero T
		return zero, false
	}
	return ms.l.Front().Value.(T), true
}

func (ms *leaseMonotonic[T]) pop() {
	if ms.l.Len() > 0 {
		ms.l.Remove(ms.l.Front())
	}
}

type greaterMonotonic[T constraints.Ordered] struct {
	l list.List
}

func newGreaterMonotonic[T constraints.Ordered]() *greaterMonotonic[T] {
	return &greaterMonotonic[T]{l: list.List{}}
}

func (ms *greaterMonotonic[T]) push(value T) (T, bool) {
	// push data finally
	defer ms.l.PushBack(value)

	// pop the lease values
	for ms.l.Len() > 0 && ms.l.Back().Value.(T) < value {
		ms.l.Remove(ms.l.Back())
	}
	// calculate return value
	if ms.l.Len() > 0 {
		return ms.l.Back().Value.(T), true
	} else {
		var zero T
		return zero, false
	}
}

func (ms *greaterMonotonic[T]) top() (T, bool) {
	if ms.l.Len() == 0 {
		var zero T
		return zero, false
	}
	return ms.l.Front().Value.(T), true
}

func (ms *greaterMonotonic[T]) pop() {
	if ms.l.Len() > 0 {
		ms.l.Remove(ms.l.Front())
	}
}
