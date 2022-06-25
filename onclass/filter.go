package onclass

import (
	"fmt"
	"time"
)

type Filter func(ctx *Context)

type FilterBuilder func(next Filter) Filter

// var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		fmt.Printf("cost  %d\n", end-start)
	}
}
