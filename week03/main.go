package main

import (
	"log"
	"net/http"
	"time"

	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

const (
	SERVER_PORT = 8080
)

type server struct {
}

func (s server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello"))
}
func main() {

	log.Println("Starting ...")

	shutdownFunctions := make([]func(context.Context), 0)

	ctx, cancel := context.WithCancel(context.Background())
	shutdownFunctions = append(shutdownFunctions, func(ctx context.Context) {
		cancel()
	})
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		httpServer := &http.Server{
			Addr:         fmt.Sprintf(":%d", SERVER_PORT),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      server{},
		}
		shutdownFunctions = append(shutdownFunctions, func(ctx context.Context) {
			err := httpServer.Shutdown(ctx)
			if err != nil {
				log.Printf("failed to shutdown  server! error: %v", err.Error())
			}
		})

		log.Printf(" server serving at :%d", SERVER_PORT)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to listen: %v", err.Error())
			return err
		}

		return nil
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Printf("received shutdown signal")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	for _, shutdown := range shutdownFunctions {
		shutdown(timeout)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("server returning an error, error: %v", err)
		os.Exit(2)
	}
}
