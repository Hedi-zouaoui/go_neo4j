package main

import (
    "log"
    "net/http"
	"net"
    "github.com/admin/utils"
)

func main() {
    config, err := utils.LoadConfig(".")
    if err != nil {
		log.Fatal("cannot log config ", err)
	}
	
    // Serve the index.html file
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    // Start the HTTP server
    listener, err := net.Listen("tcp", config.HTTPServerAddress)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Start frontend_microservice at %s", listener.Addr().String())
    err = http.Serve(listener, nil)
    if err != nil {
        log.Fatal("Cannot start frontend_microservice", err)
    }
}
