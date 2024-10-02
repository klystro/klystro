package main

import (
	"fmt"
	"log"
	"net/http"

	"klystro/config"
	"klystro/pkg/db"

	"github.com/joho/godotenv"
	"github.com/quic-go/quic-go/http3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.InitConfig()
	db.InitDatabases()
	defer db.CloseDatabases()

	// Define your handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Klystro server! You're using %s\n", r.Proto)
	})

	// Set up HTTP/1.1 and HTTP/2 server
	h2s := &http2.Server{}
	h1h2Server := &http.Server{
		Addr:    ":8090",
		Handler: h2c.NewHandler(handler, h2s),
	}

	// Set up HTTP/3 server
	h3Server := &http3.Server{
		Addr:    ":443",
		Handler: handler,
	}

	// Start HTTP/1.1 and HTTP/2 server
	go func() {
		log.Println("Starting HTTP/1.1 and HTTP/2 server on :8080")
		if err := h1h2Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Start HTTP/3 server
	log.Println("Starting HTTP/3 server on :443")
	if err := h3Server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}
}
