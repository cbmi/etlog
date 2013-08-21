package main

import (
	"etlog/services"
	"flag"
	"fmt"
	"os"
)

var defaultServerPorts = map[string]int{
	"http": 4000,
	"tcp":  4001,
	"udp":  4002,
}

func main() {
	serverTypeFlag := flag.String("type", "http", "server type: http, tcp, or udp")
	serverHostFlag := flag.String("host", "", "server host")
	serverPortFlag := flag.Int("port", 0, "server port")

	flag.Parse()

	//
	serverType := *serverTypeFlag
	serverHost := *serverHostFlag
	serverPort := *serverPortFlag

	if serverPort == 0 {
		serverPort = defaultServerPorts[serverType]
	}

	switch serverType {
	case "udp":
		services.Udp(serverHost, serverPort)
	case "tcp":
		services.Tcp(serverHost, serverPort)
	case "http":
		services.Http(serverHost, serverPort)
	default:
		fmt.Println("server type must be http, tcp, or udp")
		os.Exit(1)
	}
}
