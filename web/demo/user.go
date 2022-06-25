package demo

import (
	"fmt"
	"time"

	web "github.com/toy-web/pkg"
)

func SignUp(c *web.Context) {
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		c.BadRequestJson(&commonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid request %v", err),
		})
	}
	c.OkJson(&commonResponse{
		BizCode: 1,
		Msg:     fmt.Sprintf("get request %s ok", c.R.URL.Path),
		Data: map[string]string{
			"result": "success",
		},
	})
}

func SlowService(c *web.Context) {
	time.Sleep(10 * time.Second)
	c.OkJson(&commonResponse{
		BizCode: 1,
		Msg:     "Hi, this is msg from slow service",
	})
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
