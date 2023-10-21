package goloadbalancer

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

type Scheduler interface {
	Next() *uuid.UUID
	Add(proxy string) error
	Remove(proxy string)
}

type RoundRobinScheduler struct {
	backends []*Backend
	sync.RWMutex
}

func (r *RoundRobinScheduler) Add(proxy string) error {
	r.Lock()
	defer r.Unlock()
	b, err := NewBackend(proxy)
	if err != nil {
		return err
	}
	r.backends = append(r.backends, b)
	return nil
}

func (r *RoundRobinScheduler) Remove(proxy string) error {
	r.Lock()
	defer r.Unlock()
	toRemove := -1
	for i, b := range r.backends {
		if b.ProxyUrl() == proxy {
			toRemove = i
			break
		}
	}
	if toRemove == -1 {
		return fmt.Errorf("proxy %s does not exist", proxy)
	}
	// remove from backends
	r.backends = append(r.backends[:toRemove], r.backends[toRemove+1:]...)
	return nil
}

func (r *RoundRobinScheduler) Next() *Backend {
	l := len(r.backends)
	n := rand.Intn(l)
	return r.backends[n]
}
