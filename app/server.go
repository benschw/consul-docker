package main

import (
	"net/http"

	"github.com/benschw/opin-go/ophttp"
	"github.com/benschw/opin-go/rest"
)

func DemoHandler(resp http.ResponseWriter, req *http.Request) {
	client := NewGreetingClient()
	greeting, _ := client.GetGreeting()

	rest.SetOKResponse(resp, string(greeting[:]))
}

// Resource Handler for `/greeting`
func GreetingHandler(resp http.ResponseWriter, req *http.Request) {
	rest.SetOKResponse(resp, "hello world")
}

// Wire and start http server
func RunServer(server *ophttp.Server) {
	http.Handle("/greeting", http.HandlerFunc(GreetingHandler))
	http.Handle("/demo", http.HandlerFunc(DemoHandler))
	server.Start()
}
