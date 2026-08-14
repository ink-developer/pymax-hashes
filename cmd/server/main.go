package main

import (
	"net"
	"net/http"
	"pymax-hashes/internal/config"
	"pymax-hashes/internal/httpapi"
	"pymax-hashes/internal/storage"
	"strconv"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	storage := storage.New(config.DataPath)
	server := httpapi.New(storage, config)

	http.HandleFunc("GET /versions.json", server.GetVersions)
	http.HandleFunc("PUT /versions/{version}", server.AddVersion)
	http.HandleFunc("/", server.NotFound)

	addr := net.JoinHostPort(config.Host, strconv.Itoa(config.Port))

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
