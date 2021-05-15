package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/istherepie/testing-actions/internal/service"
)

func App(instance string, host string, port int) int {

	addr := fmt.Sprintf("%v:%d", host, port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println(err)
		return 1
	}

	mux := service.New(instance)

	log.Printf("Server starting on %v", addr)
	serveErr := http.Serve(listener, mux)

	if serveErr != nil {
		log.Println(serveErr)
		return 1
	}

	return 0
}

func main() {
	instance := flag.String("instance", "unknown", "Instance name.")
	host := flag.String("host", "localhost", "Service host.")
	port := flag.Int("port", 8080, "Service port.")
	flag.Parse()

	retCode := App(*instance, *host, *port)
	os.Exit(retCode)

	// Change
}
