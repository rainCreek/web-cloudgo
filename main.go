// learn from Mr.PanMaolin's "main.go"

package main

import (
	"os"
	"./service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	// default port : 8080
	if len(port) == 0 {
		port = PORT
	}

	// port for httpd listening
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// server.go receive para: port num
	service.NewServer(port)
}
