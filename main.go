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
	envPtr := flag.String("env", "", "Environment to parse.")
	flag.Parse()

	if *envPtr == "dev" {
		var machine1 machine
		machine1.hostname, _ = os.Hostname()
		machine1.env = *envPtr
		machine1.os = runtime.GOOS
		machine1.architecture = runtime.GOARCH
		machine1.status = "Running in development mode"
		fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
	} else {
		fmt.Println("Invalid input")
	}

}
