package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/{{ cookiecutter.github.org }}/{{ cookiecutter.github.repository }}/api"
)

func main() {
	// Flags
	var (
		host = flag.String("host", "", "listen host")
		port = flag.String("port", "8080", "listen port")
	)
	flag.Parse()

	// Rest server
	server := api.NewServer()

	// HTTP server
	srv := &http.Server{
		Addr:    net.JoinHostPort(*host, *port),
		Handler: server.Router,

		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	run(srv, *host, *port)
}

func run(srv *http.Server, host, port string) {
	// Shutdown the http server when a signal INT, TERM or QUIT is received.
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(
			signals,
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		signal.Ignore(syscall.SIGHUP)

		<-signals
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println("HTTP server shutdown:", err)
		}
	}()

	fmt.Println("[API] Started ", net.JoinHostPort(host, port))
	log.Fatal(srv.ListenAndServe())
}
