// main.go

package main

import (
	"log"
	"net"
	"os"
)

func main() {
	infoLog := log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ifaces, err := net.Interfaces()
	if err != nil {
		infoLog.Println(err)
	}
	for _, iface := range ifaces {
		infoLog.Printf("Checking up network interface: |Name: %s, hardware address: %s, flags: %s|", iface.Name, iface.HardwareAddr, iface.Flags)
	}

}
