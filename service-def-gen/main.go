package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type ServiceDefWrapper struct {
	Services []*ServiceDef `json:"services"`
}

type ServiceDef struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// {"service": {
// 	"name": "web",
// 	"tags": ["rails"],
// 	"port": 80
//  "checks": [{
//     "script": "/usr/local/bin/check_redis.py",
//     "interval": "10s"
//  }]
// }}

func main() {
	name := flag.String("name", "", "service name")
	ip := flag.String("ip", "", "docker bridge ip")
	port := flag.Int("port", 0, "service port")
	flag.Parse()

	svcDef := &ServiceDefWrapper{
		Services: []*ServiceDef{&ServiceDef{
			Name:    *name,
			Address: *ip,
			Port:    *port,
		}},
	}

	svcDefStr, err := json.MarshalIndent(svcDef, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(svcDefStr))
}
