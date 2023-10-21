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
	p, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}
	return &Backend{
		uuid:     uuid.New(),
		proxy:    *httputil.NewSingleHostReverseProxy(p),
		active:   atomic.Bool{},
		proxyUrl: proxy,
	}, nil
}

func (b *Backend) Activate() {
	b.active.Store(true)
}

func (b *Backend) Deactivate() {
	b.active.Store(false)
}

func (b *Backend) Uuid() *uuid.UUID {
	return &b.uuid
}

func (b *Backend) ProxyUrl() string {
	return b.proxyUrl
}
