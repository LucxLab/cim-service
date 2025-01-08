package main

import "github.com/LucxLab/cim-service/http"

func main() {
	server := &http.Server{
		Address: ":8000",
	}

	if err := server.Listen(); err != nil {
		panic(err)
	}
}
