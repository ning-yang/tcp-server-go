package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var addr = flag.String("addr", ":1817", "learn to address:port")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Println("Error on initializing listener:", err)
		os.Exit(1)
	}

	defer listener.Close()

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error on getting hostname:", err)
		os.Exit(1)
	}

	fmt.Println("Listening on - ", *addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn, hostname)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, reply string) {
	fmt.Println(
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		": Handling connection from - ",
		conn.RemoteAddr())
	conn.Write([]byte(reply))
	conn.Close()
}
