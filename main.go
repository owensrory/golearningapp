package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

type machine struct {
	hostname     string
	os           string
	architecture string
	env          string
	status       string
}

func main() {
	appId := "InfraProbe v0.1"
	textPtr := flag.String("text", "", "Text to parse.")
	flag.Parse()

	if *textPtr == "dev" {
		var machine1 machine
		machine1.hostname, _ = os.Hostname()
		machine1.env = *textPtr
		machine1.os = runtime.GOOS
		machine1.architecture = "ARM64"
		machine1.status = "Running in development mode"
		fmt.Printf("%s\n Hostname:%s\n Env:%s\n OS:%s\n Architecture:%s\n Status:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
	} else {
		fmt.Println("Invalid input")
	}

}
