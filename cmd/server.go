package main

import (
	"context"
	"net"
	"net/http"

	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"
)

var server *http.Server

func start(context.Context) error {
	router.Handle(http.MethodGet, "/_health", func(c *router.Context) { c.OK("ok") })
	router.Handle(http.MethodGet, "/", handler)
	router.Handle(http.MethodPost, "/", handler)

	addr := config.GetString("addr")

	server = &http.Server{Addr: addr, Handler: router.Handler()}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Info().Msgf("server start success on %s", addr)

	return server.Serve(ln)
}

func stop() {
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Error().Msgf("server shutdown error: %v", err)
	} else {
		log.Info().Msgf("server stop success")
	}
}
