package main

import (
	"github.com/LucxLab/cim-service/internal/http"
	"github.com/LucxLab/cim-service/internal/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const defaultAddress = ":8000"

func main() {
	mainLogger, loggerErr := logger.New(true)
	if loggerErr != nil {
		log.Fatalf("Failed to create the main logger: %v", loggerErr)
	}
	defer func(mainLogger logger.Logger) {
		_ = mainLogger.Close()
	}(mainLogger)

	apiServerAddress := defaultAddress
	apiServer := http.New(apiServerAddress, mainLogger)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := apiServer.Listen(); err != nil {
			mainLogger.FatalF("Failed to start the API server on address %s: %v", apiServerAddress, err)
		}
	}()

	<-stop
	if err := apiServer.Close(); err != nil {
		mainLogger.FatalF("Failed to close the API server: %v", err)
	}
}
