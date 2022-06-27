package webv2

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorInvalidRouterPattern = errors.New("invalid router pattern")

var supportMethods = [4]string{http.MethodPost, http.MethodGet,
	http.MethodDelete, http.MethodPut}

type node struct {
	path     string
	children []*node
	// 如果这是叶子节点，
	// 那么匹配上之后就可以调用该方法
	handler handlerFunc
}

type HandlerBasedOnTree struct {
	root *node
}

func NewHandlerBasedOnTree() Handler {
	root := &node{}
	return &HandlerBasedOnTree{
		root: root,
	}
}

func (h *HandlerBasedOnTree) Route(method, pattern string, handlerFunc handlerFunc) error {
	err := h.validatePattern(pattern)
	if err != nil {
		return err
	}
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root
	for i, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			h.createSubTree(cur, paths[i:], handlerFunc)
			return nil
		}
		cur = matchChild
	}
	// 离开了循环，说明我们加入的是短路径，
	// 比如说我们先加入了 /order/detail
	// 再加入/order，那么会走到这里
	cur.handler = handlerFunc
	return nil
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findMatchHandler(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) findMatchHandler(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range root.children {
		// 并不是 * 的节点命中了，直接返回
		// != * 是为了防止用户乱输入
		if child.path == path && child.path != "*" {
			return child, true
		}
		// 命中了通配符的，我们看看后面还有没有更加详细的
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFn handlerFunc) {
	cur := root
	for _, p := range paths {
		nn := newNode(p)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
	// validate "*", it must be the suffix and it previous must be "/" if it exists
	// so pattern "/*" is legal, but "abc*" is illegal
	pos := strings.Index(pattern, "*")
	if pos > 0 {
		if pos != len(pattern)-1 {
			return ErrorInvalidRouterPattern
		}
		if pattern[pos-1] != '/' {
			return ErrorInvalidRouterPattern
		}
	}
	return nil
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 2),
	}
}
