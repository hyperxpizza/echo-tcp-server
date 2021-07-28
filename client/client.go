package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	ServerHost    = "localhost"
	ServerPort    = "8080"
	RemoteHost    = "localhost"
	PingFrequency = 5 * time.Second
)

var RemotePort string

//manage flags
func init() {
	remotePortPointer := flag.String("port", "7070", "String that specifies remote port value")
	flag.Parse()
	RemotePort = *remotePortPointer
}

func main() {
	serverAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", ServerHost, ServerPort))
	if err != nil {
		log.Fatal(err)
	}

	clientAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", RemoteHost, RemotePort))
	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.DialTCP("tcp", clientAddress, serverAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	fmt.Printf("%s <--> %s\n", clientAddress.String(), serverAddress.String())
	for {
		time.Sleep(PingFrequency)
		err := ping(connection, fmt.Sprintf("Hello from: %s\n", clientAddress.String()))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func ping(connection net.Conn, message string) error {
	_, err := connection.Write([]byte(message))
	if err != nil {
		return err
	}

	log.Printf("Sent: %s\n", message)

	buffer := make([]byte, 512)
	_, err = connection.Read(buffer)
	if err != nil {
		return err
	}

	log.Printf("Recieved: %s\n", message)

	return nil
}
