package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Define the server address and port
	serverAddr := "127.0.0.1:12345"

	// Resolve the server address
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", serverAddr)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read incoming data
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Print received message
		fmt.Println("Received message from client:", string(buffer[:n]))

		// Respond to the client
		response := "Message received"
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
		}
	}
}
