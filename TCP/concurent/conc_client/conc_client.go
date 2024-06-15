package main

import (
	"fmt"
	"net"
)
var Conn net.Conn

func write_data(id int)(string){
    defer Conn.Close()
	// Connect to the server
	Conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
        fmt.Println(err)
        return ""
    }
	// Send some data to the server
    simp_msg:= fmt.Sprintf("Hello %d, server!",id)
	_, err = Conn.Write([]byte(simp_msg))
    if err != nil {
        fmt.Println(err)
        return ""
    }

    // Close the connection
    
	return simp_msg
}

func main() {
	for i:=0;i<10;i++{
		go func(j int){
			write_data(j)
		}(i)
	}

}