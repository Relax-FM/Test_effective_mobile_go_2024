package tem2024

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, rtm, wtm time.Duration, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    rtm * time.Second,
		WriteTimeout:   wtm * time.Second,
	}
	logrus.Print(s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	logrus.Print("Try to graceful shutdown server")
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	} else {
		logrus.Print("Server shuted down successfully!")
	}
}