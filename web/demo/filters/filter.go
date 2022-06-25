package filters

import (
	"fmt"

	web "github.com/toy-web/pkg"
)

func init() {
	web.RegisterFilter("my-custom-filter", MyFilter)
}

func MyFilter(next web.Filter) web.Filter {
	return func(c *web.Context) {
		fmt.Println("this is my custom filter")
		next(c)
	}
}
