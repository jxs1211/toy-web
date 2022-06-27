package webv2

import (
	"net/http"
	"sync"
)

// the 2 way is both ok for validating the type wether implement the interface
var _ Handler = (*HandlerBasedMap)(nil)
var _ Handler = &HandlerBasedMap{}

type HandlerBasedMap struct {
	m sync.Map
}

func (h *HandlerBasedMap) Route(method, pattern string, handlerFunc handlerFunc) error {
	key := h.Key(method, pattern)
	h.m.Store(key, handlerFunc)
	return nil
}

func (h *HandlerBasedMap) ServeHTTP(ctx *Context) {
	key := h.Key(ctx.R.Method, ctx.R.URL.Path)
	handler, ok := h.m.Load(key)
	if !ok {
		ctx.W.WriteHeader(http.StatusNotFound)
		return
	}
	handler.(handlerFunc)(ctx)
}

func (h *HandlerBasedMap) Key(method, url string) string {
	return method + "#" + url
}

func NewHandlerBasedMap() *HandlerBasedMap {
	return &HandlerBasedMap{}
}
