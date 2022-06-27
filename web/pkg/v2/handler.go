package webv2

type Router interface {
	Route(method, pattern string, handlerFunc handlerFunc) error
}

type Handler interface {
	ServeHTTP(c *Context)
	Router
}

type handlerFunc func(c *Context)
