package utils

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	l := NewRandomMap[string, int]()
	n := 5000
	for i := 0; i < n; i++ {
		l.Add(fmt.Sprintf("%d", i), i)
	}
	assert.Equal(t, l.Len(), n)
	for i := 0; i < n; i++ {
		assert.True(t, l.Has(fmt.Sprintf("%d", i)))
	}
}

func TestRemove(t *testing.T) {
	l := NewRandomMap[string, int]()
	n := 5000
	for i := 0; i < n; i++ {
		l.Add(fmt.Sprintf("%d", i), i)
	}
	assert.Equal(t, l.Len(), n)
	for i := 0; i < n; i++ {
		assert.Nil(t, l.Remove(fmt.Sprintf("%d", i)))
	}
	assert.Zero(t, l.Len())
	for i := 0; i < n; i++ {
		assert.False(t, l.Has(fmt.Sprintf("%d", i)))
	}

	for i := 0; i < n; i++ {
		l.Add(fmt.Sprintf("%d", i), i)
	}
	assert.Equal(t, l.Len(), n)
	removedKey := make(map[string]bool)
	// random remove
	for l.Len() > 0 {
		key := fmt.Sprintf("%d", rand.Intn(n))
		_, ok := removedKey[key]
		if !ok {
			assert.Nil(t, l.Remove(key))
			removedKey[key] = true
		}
	}
	assert.Zero(t, l.Len())
}

func TestRandom(t *testing.T) {
	l := NewRandomMap[string, int]()
	n := 10000
	item := l.Random()
	assert.Nil(t, item)
	for i := 0; i < n; i++ {
		l.Add(fmt.Sprintf("%d", i), i)
	}

	for i := 0; i < n; i++ {
		item = l.Random()
		assert.True(t, l.Has(*item.key))
	}
}
