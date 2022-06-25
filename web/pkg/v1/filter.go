package webv1

import (
	"log"
	"time"
)

type Filter func(ctx *Context)

type FilterBuilder func(next Filter) Filter

// var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		log.Println("MetricFilterBuilder start")
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		log.Printf("cost  %d\n", end-start)
		log.Println("MetricFilterBuilder end")
	}
}

func LogFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		log.Println("LogFilterBuilder start")
		next(ctx)
		log.Printf("[%s] %s\n", ctx.R.Method, ctx.R.URL.Path)
		log.Println("LogFilterBuilder end")
	}
}
