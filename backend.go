package goloadbalancer

import (
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/google/uuid"
)

type Backend struct {
	uuid     uuid.UUID
	proxy    httputil.ReverseProxy
	active   atomic.Bool
	proxyUrl string
}

func NewBackend(proxy string) (*Backend, error) {
	p, err := url.ParseRequestURI(proxy)
	if err != nil {
		return nil, err
	}
	b := &Backend{
		uuid:     uuid.New(),
		proxy:    *httputil.NewSingleHostReverseProxy(p),
		active:   atomic.Bool{},
		proxyUrl: proxy,
	}
	b.Activate()
	return b, nil
}

func (b *Backend) Activate() {
	b.active.Store(true)
}

func (b *Backend) Deactivate() {
	b.active.Store(false)
}

func (b *Backend) IsActive() bool {
	return b.active.Load()
}

func (b *Backend) Uuid() *uuid.UUID {
	return &b.uuid
}

func (b *Backend) ProxyUrl() string {
	return b.proxyUrl
}
