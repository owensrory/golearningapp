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
		machine1 := createMachine(*envPtr)
		fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
	} else if *envPtr == "staging" {
		machine1 := createMachine(*envPtr)
		fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
	} else if *envPtr == "prod" {
		machine1 := createMachine(*envPtr)
		fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
	} else {
		fmt.Printf("Error unsupported environment %s", *envPtr)
	}

}

func createMachine(env string) machine {
	var m machine
	m.hostname, _ = os.Hostname()
	m.env = env
	m.os = runtime.GOOS
	m.architecture = runtime.GOARCH
	m.status = fmt.Sprintf("Running in %s mode", env)

	return m
}
