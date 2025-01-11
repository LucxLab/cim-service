package main

import (
	"github.com/LucxLab/cim-service/internal/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := http.NewServer(":8000")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := server.Listen(); err != nil {
			panic(err)
		}
	}()

	<-stop
	if err := server.Close(); err != nil {
		panic(err)
	}
}
