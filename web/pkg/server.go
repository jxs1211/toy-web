package web

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Router
	Run(addr string) error
	Shutdown(ctx context.Context) error
}

type sdkHttpServer struct {
	name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method, pattern string, handleFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Run(addr string) error {
	// http.Handle("/", s.handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		s.root(c)
	})
	return http.ListenAndServe(addr, nil)
}

func (s *sdkHttpServer) Shutdown(ctx context.Context) error {
	fmt.Println("server shutdown start")
	time.Sleep(time.Second)
	fmt.Println("server shutdown end")
	return nil
}

func NewSdkHttpServer(name string, builders ...FilterBuilder) *sdkHttpServer {
	handler := NewHandlerBasedMap()
	root := handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		builder := builders[i]
		root = builder(root) // Filter func
	}
	s := &sdkHttpServer{
		name:    name,
		handler: handler,
		root:    root,
	}

	return s
}

func Test(ctx *Context) {
	fmt.Fprintf(ctx.W, "Hi, %s\n", ctx.R.URL.Path)
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Printf("read err: %v\n", err)
		return
	}
	// fmt.Printf("req: %v\n", req)
	resp := commonResponse{
		Msg: fmt.Sprintf("path: %s, email: %s", ctx.R.URL.Path, req.Email),
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("write err: %v\n", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
