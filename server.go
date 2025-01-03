package todogo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// The hander interface:
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler, // An interface that represents an HTTP handler that handle HTTP queries
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx) // The context manage the shutdown process
}
