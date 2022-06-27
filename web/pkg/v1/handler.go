package webv1

import (
	"net/http"
)

type Router interface {
	Route(method, pattern string, handlerFunc handlerFunc) error
}

type Handler interface {
	ServeHTTP(c *Context)
	Router
}

type handlerFunc func(c *Context)

type HandlerBasedMap struct {
	// Handler
	// map's key is method + url
	m map[string]func(ctx *Context)
}

func (h *HandlerBasedMap) Route(method, pattern string, handleFunc handlerFunc) error {
	key := h.Key(method, pattern)
	h.m[key] = handleFunc
	return nil
}

func (h *HandlerBasedMap) ServeHTTP(ctx *Context) {
	key := h.Key(ctx.R.Method, ctx.R.URL.Path)
	if f, ok := h.m[key]; ok {
		f(NewContext(ctx.W, ctx.R))
	} else {
		ctx.W.WriteHeader(http.StatusNotFound)
		ctx.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedMap) Key(method, url string) string {
	return method + "#" + url
}

// 接口断言
// 1
var _ Handler = (*HandlerBasedMap)(nil)

// 2
// var _ Handler = &HandlerBasedMap{}

func NewHandlerBasedMap() Handler {
	return &HandlerBasedMap{
		m: make(map[string]func(ctx *Context), 200),
	}
}
