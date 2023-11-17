// main.go

package main

import (
	"context"
	"log"

	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	c := oauth.NewComputeEngine()
	meta, err := c.GetRequestMetadata(context.Background())
	if err != nil {
		log.Fatal("GetRequestMetadata failed: ", err)
	}
	log.Println(meta)
	// infoLog := log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// ifaces, err := net.Interfaces()
	// if err != nil {
	// 	infoLog.Println(err)
	// }
	// for _, iface := range ifaces {
	// 	infoLog.Printf("Checking up network interface: |Name: %s, hardware address: %s, flags: %s|", iface.Name, iface.HardwareAddr, iface.Flags)
	// }

}
