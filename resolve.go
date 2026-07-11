package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"time"
)

func resolve(hostname string, binPath string) error {
	if hostname == "" {
		return fmt.Errorf("hostname cannot be empty")
	}

	if binPath != "" {
		directory := filepath.Dir(binPath)
		fileName := filepath.Base(binPath)

		copyBinaryTo(directory, fileName)

		args := []string{"dnsquery", "--host", hostname}
		output := execute(binPath, args)
		fmt.Println(output)
		return nil
	}

	log.Printf("Resolving DNS for: %s\n", hostname)

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	ips, err := r.LookupHost(context.Background(), hostname)
	if err != nil {
		return fmt.Errorf("failed to resolve %s: %v", hostname, err)
	}

	if len(ips) == 0 {
		fmt.Printf("No IP addresses found for %s\n", hostname)
		return nil
	}

	fmt.Printf("Resolved %s to:\n", hostname)
	for _, ip := range ips {
		fmt.Printf("  %s\n", ip)
	}

	return nil
}
