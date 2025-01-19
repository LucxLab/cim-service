package main

import (
	"github.com/LucxLab/cim-service/internal/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	apiServer := http.New(":8000")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := apiServer.Listen(); err != nil {
			panic(err)
		}
	}()

	<-stop
	if err := apiServer.Close(); err != nil {
		panic(err)
	}
}
