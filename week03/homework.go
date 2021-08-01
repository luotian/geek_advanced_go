package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

var hs http.Server

func init() {
	fmt.Println("Homework For week03 : errgroup")

	hs.Addr = ":80"
}

func innerServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Recevied Request")
}

func startServer() error {
	http.HandleFunc("/req", innerServer)
	err := hs.ListenAndServe()

	return err
}

func main() {
	bgctx := context.Background()
	cctx, cancel := context.WithCancel(bgctx)

	g, ctx := errgroup.WithContext(cctx)

	g.Go(func() error {
		return startServer()
	})

	g.Go(func() error {
		<-ctx.Done()
		return hs.Shutdown(ctx)
	})

	chanel := make(chan os.Signal)
	signal.Notify(chanel)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-chanel:
				cancel()
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
}
