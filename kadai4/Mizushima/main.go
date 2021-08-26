// this package implements the Omikuji server.
// It basis on https://qiita.com/ww24/items/7c7863421a1a538c7bc3
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	listener, ch := server(":8080", omikujiHandler)
	fmt.Println("Omikuji Server started at", listener.Addr())

	// ctrl+c signal interrupt
	ctx := context.Background()
	_, cancel := listen(ctx, listener)
	defer cancel()

	log.Println(<-ch)
}

// 
func server(addr string, handler func(w http.ResponseWriter, r *http.Request)) (net.Listener, chan error) {
	ch := make(chan error)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		mux := router(handler)
		ch <- http.Serve(listener, mux)
	}()

	return listener, ch
}

// 
func router(handler func(w http.ResponseWriter, r *http.Request)) *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return mux
}

// 
func listen(ctx context.Context, listener net.Listener) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		<-sig
		if _, err := fmt.Println("\n^Csignal : interrupt."); err != nil {
			cancel()
			log.Fatalf("listen: fmt.Println error: %s", err)
		}
		if err := listener.Close(); err != nil {
			cancel()
			log.Fatalf("listen: listener.Close error: %s", err)
		}
	}()

	return ctx, cancel
}