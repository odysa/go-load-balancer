package goloadbalancer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_HOST string = "http://localhost:8000/test"

func TestNewBackend(t *testing.T) {
	b1, err := NewBackend("$6546h")
	assert.Nil(t, b1)
	assert.NotNil(t, err)
	b, err := NewBackend(TEST_HOST)
	assert.Nil(t, err)
	assert.Equal(t, b.proxyUrl, TEST_HOST)
}

func TestBackendActivate(t *testing.T) {
	b, err := NewBackend(TEST_HOST)
	assert.Nil(t, err)
	b.Deactivate()
	n := 5000
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			b.Activate()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.True(t, b.IsActive())
}

func TestBackendDeactivate(t *testing.T) {
	b, err := NewBackend(TEST_HOST)
	assert.Nil(t, err)
	n := 5000
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			b.Deactivate()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.False(t, b.IsActive())
}

func TestIsActive(t *testing.T) {
	b, err := NewBackend(TEST_HOST)
	assert.Nil(t, err)
	assert.True(t, b.IsActive())
	b.Deactivate()
	assert.False(t, b.IsActive())
}

func TestProxyUrl(t *testing.T) {
	b, err := NewBackend(TEST_HOST)
	assert.Nil(t, err)
	assert.Equal(t, b.ProxyUrl(), TEST_HOST)
}
