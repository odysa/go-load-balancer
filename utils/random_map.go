package utils

import (
	"fmt"
	"math/rand"
)

type MapItem[K comparable, V any] struct {
	key   *K
	value V
}

type RandomMap[K comparable, V any] struct {
	items    []MapItem[K, V]
	indexMap map[K]int
}

func NewRandomMap[K comparable, V any]() *RandomMap[K, V] {
	return &RandomMap[K, V]{
		items:    make([]MapItem[K, V], 0),
		indexMap: make(map[K]int),
	}
}

func (r *RandomMap[K, V]) Add(k K, v V) {
	r.items = append(r.items, MapItem[K, V]{key: &k, value: v})
	r.indexMap[k] = r.Len() - 1
}

func (r *RandomMap[K, V]) Remove(k K) error {
	idx, ok := r.indexMap[k]
	if !ok {
		return fmt.Errorf("key does not exist")
	}
	l := r.Len()
	// swap
	r.items[idx], r.items[l-1] = r.items[l-1], r.items[idx]
	r.items = r.items[:l-1]

	if idx < r.Len() {
		// swap idx
		key := r.items[idx].key
		r.indexMap[*key] = idx
	}

	delete(r.indexMap, k)
	return nil
}

func (r *RandomMap[K, V]) Random() *MapItem[K, V] {
	if r.Len() <= 0 {
		return nil
	}
	l := len(r.items)
	n := rand.Intn(l)
	return &r.items[n]
}

func (r *RandomMap[K, V]) Has(k K) bool {
	_, ok := r.indexMap[k]
	return ok
}

func (r *RandomMap[K, V]) Len() int {
	return len(r.items)
}
