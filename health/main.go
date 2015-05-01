package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/benschw/opin-go/rest"
)

type HealthStatus struct {
	Status string `json:"status"`
}

func main() {
	flag.Parse()

	address := flag.Arg(0)
	log.Printf("Testing Status endpoint: '%s'", address)
	r, err := rest.MakeRequest("GET", address, nil)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	var status HealthStatus
	if err = rest.ProcessResponseEntity(r, &status, http.StatusOK); err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println(status.Status)
	switch status.Status {
	case "UP":
		os.Exit(0)
	case "OK":
		os.Exit(0)
	case "WARN":
		os.Exit(1)
	default:
		os.Exit(2)
	}
}
