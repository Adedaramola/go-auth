package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(port uint32) *Server {
	return &Server{
		&http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (srv *Server) Listen() {
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to serve:", err)
		}
	}()

	srv.gracefulShutdown()
}

func (srv *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Server exiting")

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("failed to shutdown properly:", err)
	}

	log.Println("Server exited gracefully")
}

func (srv *Server) SetupRoutes(handler http.Handler) {
	srv.Handler = handler
}
