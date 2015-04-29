package main

import (
	"net/http"
	"os"

	"github.com/benschw/opin-go/ophttp"
	"github.com/benschw/opin-go/rando"
	"github.com/benschw/opin-go/rest"
)

func DemoHandler(resp http.ResponseWriter, req *http.Request) {
	host, _ := os.Hostname()

	client := NewGreetingClient()
	greeting, _ := client.GetGreeting()

	rest.SetOKResponse(resp, &DemoGreeting{
		Message:  "hello from demo on " + host + "/" + rando.MyIp(),
		Greeting: greeting,
	})
}

// Resource Handler for `/greeting`
func GreetingHandler(resp http.ResponseWriter, req *http.Request) {
	host, _ := os.Hostname()

	rest.SetOKResponse(resp, Greeting{Message: "hello from greeting on " + host + "/" + rando.MyIp()})
}

// Wire and start http server
func RunServer(server *ophttp.Server) {
	http.Handle("/greeting", http.HandlerFunc(GreetingHandler))
	http.Handle("/demo", http.HandlerFunc(DemoHandler))
	server.Start()
}
