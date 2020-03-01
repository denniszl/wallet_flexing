package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/denniszl/wallet_flexing/internal/endpoints"
	"github.com/denniszl/wallet_flexing/internal/flex"
	httptransport "github.com/denniszl/wallet_flexing/internal/http"
	"github.com/denniszl/wallet_flexing/internal/repository"
	"github.com/denniszl/wallet_flexing/internal/repository/disk"
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
	p := os.Getenv("PORT")
	if p == "" {
		p = "6969"
	}
	var (
		addr     = p
		httpAddr = flag.String("http.addr", ":"+p, "HTTP listen address")
	)

	var err error
	var repo repository.Repository
	{
		repo, err = disk.NewRepository()
		if err != nil {
			panic(err)
		}
	}

	var service flex.Service
	{
		service = flex.NewService(repo)
	}

	var handlerEndpoints endpoints.Endpoints
	{
		handlerEndpoints = endpoints.NewEndpoints(service)
	}

	var handler http.Handler
	{
		handler = httptransport.MakeHTTPHandler(handlerEndpoints)
	}

	log.Infof("Listening on port: %s", addr)
	err = http.ListenAndServe(*httpAddr, handler)
	if err != nil {
		panic("exited")
	}

	os.Exit(1)
}
