package main

import (
	"context"
	"github.com/VlaDkoR-280/VPN-Checker/internal/adapters/vpn"
	"log"
	"time"
)

func main() {
	log.Println("Starting VPN Checker")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	adapter, errInit := vpn.InitAdapter(ctx, "test")
	if errInit != nil {
		log.Fatal(errInit)
	}

	log.Println("VPN adapter initialized")

	ip, errGetIP := adapter.GetIPAddress(ctx)
	if errGetIP != nil {
		log.Fatal(errGetIP)
	}

	log.Println("ip address: ", ip)
}
