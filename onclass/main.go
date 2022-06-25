package onclass

import "net/http"

func main() {

	s := NewSdkHttpServer("test", MetricFilterBuilder)
	s.Route(http.MethodGet, "/test", test)
	s.Route(http.MethodPost, "/signUp", SignUp)
	s.Run(":1234")
}
