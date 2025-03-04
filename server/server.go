package server

import (
	"context"
	"net/http"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

func (s *HTTPServer) Run(handler http.Handler, port string) error {
	s.server = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
