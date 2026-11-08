package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
)

type machine struct {
	hostname          string
	os                string
	architecture      string
	env               string
	status            string
	networkinterfaces []networkInterface
}

type networkInterface struct {
	name string
	mac  string
	ips  []string
}

func main() {
	appId := "InfraProbe v0.2"
	envPtr := flag.String("env", "", "Environment to parse.")
	flag.Parse()

	envs := []string{"dev", "staging", "prod"}

	for i := 0; i < len(envs); i++ {
		if *envPtr == envs[i] {
			machine1, err := createMachine(*envPtr)
			if err != nil {
				fmt.Printf("Error: %v", err)
				os.Exit(1)
			} else {
				fmt.Printf("%s\nHostname:%s\nEnv:%s\nOS:%s\nArchitecture:%s\nStatus:%s\nLogical CPUs:%d\nNetwork Interfaces:%d\n", appId, machine1.hostname, machine1.env, machine1.os, machine1.architecture, machine1.status, runtime.NumCPU(), len(machine1.networkinterfaces))
				for _, iface := range machine1.networkinterfaces {
					fmt.Printf("Interface:%v\n    MAC:%s\n    IPs:\n", iface.name, iface.mac)
					for _, ip := range iface.ips {
						fmt.Printf("        %s\n", ip)
					}

				}
				return
			}
		}
	}
	err := fmt.Sprint("Invalid environment specified. Please use dev, staging, or prod.")
	fmt.Printf("Error: %v", err)
	os.Exit(1)

}

func createMachine(env string) (machine, error) {
	var m machine
	result, err := os.Hostname()
	if err != nil {
		m.hostname = fmt.Sprintf("Error retrieving hostname: %v", err)
	} else {
		m.hostname = result
	}
	m.env = env
	m.os = runtime.GOOS
	m.architecture = runtime.GOARCH
	m.status = fmt.Sprintf("Running in %s mode", env)
	m.networkinterfaces, err = getInterfaces()

	return m, err
}

func getInterfaces() ([]networkInterface, error) {

	listofinterfaces := []networkInterface{}
	interfaces, err := net.Interfaces()
	if err != nil {
		return listofinterfaces, fmt.Errorf("Error retrieving network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		var ni networkInterface
		ni.name = iface.Name
		ni.mac = iface.HardwareAddr.String()
		if ni.mac == "" {
			ni.mac = "N/A"
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return listofinterfaces, fmt.Errorf("Error retrieving addresses for interface %s: %v", iface.Name, err)
		}
		for _, addr := range addrs {
			ni.ips = append(ni.ips, addr.String())

		}
		if len(ni.ips) < 1 {
			ni.ips = append(ni.ips, "N/A")

		}

		listofinterfaces = append(listofinterfaces, ni)
	}

	return listofinterfaces, err

}
