package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jesseobrien/fake-iot/internal/http"
)

var httpListenHost string
var httpListenPort string
var certPath string
var keyPath string

func init() {
	flag.StringVar(&httpListenHost, "host", "127.0.0.1", "The host IP to bind the HTTP server to.")
	flag.StringVar(&httpListenPort, "port", "8080", "The host port to bind the HTTP server to.")
	flag.StringVar(&certPath, "cert", "certs/server.crt", "Path to the server's certificate for HTTPS.")
	flag.StringVar(&keyPath, "key", "certs/server.key", "Path to the server's signing key for HTTPS.")
	flag.Parse()
}

func main() {
	apiToken := os.Getenv("FAKEIOT_API_TOKEN")
	if apiToken == "" {
		log.Fatal("error: environment variable FAKEIOT_API_TOKEN is not set")
	}

	listenAddress := fmt.Sprintf("%s:%s", httpListenHost, httpListenPort)
	log.Fatal(http.Run(listenAddress, certPath, keyPath, apiToken))
}
