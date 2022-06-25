package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// there is no any further extension, so define as a struct, not an interface
type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	PathParams map[string]string
}

func (c *Context) ReadJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:          w,
		R:          r,
		PathParams: make(map[string]string, 1),
	}
}

func newContext() *Context {
	fmt.Println("create new context")
	return &Context{}
}

func (c *Context) Reset(w http.ResponseWriter, r *http.Request) {
	c.W = w
	c.R = r
	c.PathParams = make(map[string]string, 1)
}
