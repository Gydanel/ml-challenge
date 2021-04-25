package main

import (
	"flag"
	"fmt"
	"os"

	"ml-challenge/app/server"
	"ml-challenge/config"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
