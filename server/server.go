// An example of echo server
// @author Wojciech Frackowski

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const PORT = "8080"

func main() {
	runServer()
}

func runServer() {
	log.Printf("Starting the echo server on port: %s", PORT)
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("net.Listen error: %s\n", err)
	}

	connections := clientConnections(server)
	for {
		go handleConnections(<-connections)
	}
}

func clientConnections(listener net.Listener) chan net.Conn {
	c := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil && client == nil {
				log.Printf("Cound not accept connection: %v\n", err)
				continue
			}
			i++
			log.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			c <- client
		}
	}()
	return c
}

func handleConnections(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // End Of Line or error
			break
		}
		client.Write(line)
	}
}
