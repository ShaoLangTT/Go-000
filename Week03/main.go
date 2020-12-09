package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 作业:基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return HttpServer(ctx)
	})

	g.Go(func() error {
		return SignalServer(ctx)
	})

	if err := g.Wait(); err == nil {
		fmt.Println("Success")
	}
}

func HttpServer(ctx context.Context) error {
	srv := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	go func() {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = srv.Shutdown(ctx)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	return srv.ListenAndServe()
}

func SignalServer(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	select {
	case <-ctx.Done(): // 监听退出
		os.Exit(0)
		return nil
	case call := <-c:
		return errors.Errorf("os exit syscall : %v", call)
	}
}
