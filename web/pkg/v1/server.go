package webv1

import (
	"fmt"
	"log"
	"net/http"
)

type Server interface {
	Router
	Run(addr string) error
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

func NewSdkHttpServer(name string, builders ...FilterBuilder) *sdkHttpServer {
	handler := NewHandlerBasedMap()
	root := handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		builder := builders[i]
		root = builder(root)
	}
	s := &sdkHttpServer{
		name:    name,
		handler: handler,
		root:    root,
	}

	return s
}

func Test(ctx *Context) {
	log.Println("Test start")
	log.Println("call Test")
	fmt.Fprintf(ctx.W, "Hi, %s\n", ctx.R.URL.Path)
	log.Println("Test end")
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
