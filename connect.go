package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func connectToHost(host string, port string, binPath string) error {
	if binPath != "" {
		directory := filepath.Dir(binPath)
		fileName := filepath.Base(binPath)

		copyBinaryTo(directory, fileName)

		args := []string{"connect", "--host", host, "--port", port}
		output := execute(binPath, args)
		fmt.Println(output)
		return nil
	}

	timeout := time.Second * 3
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return fmt.Errorf("failed to connect to %s:%s - %v", host, port, err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Printf("Successfully connected to %s\n", net.JoinHostPort(host, port))
	}
	return nil
}

func downloadFile(fullURLFile string, binPath string) error {
	if binPath != "" {
		directory := filepath.Dir(binPath)
		fileName := filepath.Base(binPath)

		copyBinaryTo(directory, fileName)

		args := []string{"download", "--url", fullURLFile}
		output := execute(binPath, args)
		fmt.Println(output)
		return nil
	}

	file, err := os.CreateTemp("C:/Users/Public/", "mcli.*.dat")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer file.Close()

	client := http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	log.Printf("Downloading file from: %s\n", fullURLFile)
	resp, err := client.Get(fullURLFile)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Printf("Successfully downloaded file to %s (size: %d bytes)\n", file.Name(), size)
	return nil
}
