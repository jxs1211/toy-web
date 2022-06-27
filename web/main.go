package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/toy-web/demo"
	_ "github.com/toy-web/demo/filters"
	web "github.com/toy-web/pkg"
	// 可行
	// webv1 "github.com/toy-web/pkg/v1"
	// 可行
	// "github.com/toy-web/pkg"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

// func h(ctx context.Context) {
// 	fmt.Println("h start")
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("time out ")
// 	}
// 	fmt.Println("h end")
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	h(ctx)
// 	cancel()
// }

func main() {
	shutdown := web.NewGracefulShutdown()
	s := web.NewSdkHttpServer("test",
		web.MetricFilterBuilder,
		shutdown.ShutdownFilterBuilder,
	)
	// adminServer := web.NewSdkHttpServer("admin-test-server",
	// 	// 注意，如果你真实环境里面，使用的是多个 server监听不同端口，
	// 	// 那么这个 shutdown最好也是多个。互相之间就不会有竞争
	// 	// MetricFilterBuilder 是无状态的，所以不存在这种问题
	// 	web.MetricFilterBuilder, shutdown.ShutdownFilterBuilder)

	s.Route(http.MethodPost, "/user/create/", demo.SignUp)
	s.Route(http.MethodPost, "/slowService", demo.SlowService)

	// go func() {
	// 	if err := adminServer.Run(":12345"); err != nil {
	// 		panic(err)
	// 	}
	// }()

	go func() {
		if err := s.Run(":1234"); err != nil {
			// 快速失败，因为服务器都没启动成功，啥也做不了
			panic(err)
		}
		// 假设我们后面还有很多动作
	}()

	web.WaitForShutdown(
		// notify gateway we are going to be close
		func(ctx context.Context) error {
			fmt.Println("mock notify gateway")
			time.Sleep(time.Second * 2)
			return nil
		},
		// Reject new and waiting request
		shutdown.RejectNewAndWaiting,
		// 全部请求处理完了我们就可以关闭 server了
		web.BuildCloseServerHook(s),
		// release resources
		func(ctx context.Context) error {
			// 假设这里我要清理一些执行过程中生成的临时资源
			fmt.Println("mock release resources")
			time.Sleep(time.Second * 2)
			return nil
		},
	)
	// filterNames := ReadFromConfig
	// 匿名引入之后，就可以在这里按名索引 filter
	//web.NewSdkHttpServerWithFilterNames("my-server", filterNames...)

}
