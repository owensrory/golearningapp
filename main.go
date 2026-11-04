package main

import (
	"errors"
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

	envs := []string{"dev", "staging", "prod"}

	for i := 0; i < len(envs); i++ {
		if *envPtr == envs[i] {
			machine1 := createMachine(*envPtr)
			fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status)
			return
		}
	}
	err := errors.New("Invalid environment specified. Please use 'dev', 'staging', or 'prod'.")
	fmt.Printf("Error: %v", err)
	os.Exit(1)

}

func createMachine(env string) machine {
	var m machine
	result, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error retrieving hostname: %v", err)
		os.Exit(1)
	}
	m.hostname = result
	m.env = env
	m.os = runtime.GOOS
	m.architecture = runtime.GOARCH
	m.status = fmt.Sprintf("Running in %s mode", env)

	return m
}
