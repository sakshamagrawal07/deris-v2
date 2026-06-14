package main

import (
	"flag"
	"log"

	"github.com/sakshamsharma/deris-v2/server"
	"github.com/sakshamsharma/deris-v2/config"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host for the deris server")
	flag.IntVar(&config.Port, "port", 7379, "Port for the deris server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Deris server rolling...")
	server.RunAsyncTCPServer();
}