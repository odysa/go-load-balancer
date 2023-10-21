package utils

import "math/rand"

type RandomList[K comparable, V any] struct {
	items    []V
	indexMap map[K]int
}

func (r *RandomList[K, V]) Add(k K, v V) {
	r.items = append(r.items, v)
	r.indexMap[k] = len(r.items) - 1
}

func (r *RandomList[K, V]) Remove(k K, v V) {
	idx := r.indexMap[k]
	l := len(r.items)
	// swap
	r.items[idx], r.items[l-1] = r.items[l-1], r.items[idx]
	r.items = r.items[:len(r.items)-1]
	delete(r.indexMap, k)
}

func (r *RandomList[K, V]) Random() *V {
	l := len(r.items)
	n := rand.Intn(l)
	return &r.items[n]
}
