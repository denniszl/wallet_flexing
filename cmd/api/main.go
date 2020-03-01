package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/common/log"
)

func main() {
	// Handle panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "PANIC: %+v", r)
			os.Exit(2)
		}
	}()

	// use one if passed in, otherwise default it
	p := os.Getenv("port")
	if p == "" {
		p = "6969"
	}
	var (
		addr     = p
		httpAddr = flag.String("http.addr", ":"+p, "HTTP listen address")
	)

	log.Infof("Listening on port: %s", addr)
	err := http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		panic("exited")
	}

	os.Exit(1)
}