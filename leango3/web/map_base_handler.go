package main

import (
	"net/http"
	"sync"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
}

type Handler interface {
	//http.Handler
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers sync.Map
}

// Route 做注册路由
func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handlerFunc(ctx)
	//})
	key := h.key(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func (h *HandlerBasedOnMap) ServeHTTP(ctx *Context) {
	key := h.key(ctx.R.Method, ctx.R.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		ctx.W.WriteHeader(http.StatusNotFound)
		ctx.W.Write([]byte("not found"))
	}

	handler.(func(ctx *Context))(ctx)
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

// 保证HandlerBasedOnMap 一定实现了 Handler 接口
// 如果Handler 发生变动 HandlerBasedOnMap 没有改变的话会报错
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: sync.Map{},
	}
}
