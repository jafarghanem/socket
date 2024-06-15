package main

import (
	"fmt"
	"net"
	"os"
	"time"
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
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Message to send
	message := "Hello UDP Server!"

	// Send the message
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		os.Exit(1)
	}

	// Set a read deadline for the response
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// Buffer to hold the response
	buffer := make([]byte, 1024)

	// Read the response
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		os.Exit(1)
	}

	// Print the response
	fmt.Println("Server response:", string(buffer[:n]))
}
