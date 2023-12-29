package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Listening on", li.Addr().String())

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("Connected:", conn.RemoteAddr())

		go telnetConn(conn)
	}
}

func telnetConn(conn net.Conn) {
	defer conn.Close()
	errCh := make(chan error)

	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				errCh <- err
				return
			}

			log.Printf("Message from %v: %v", conn.RemoteAddr(), data)
			fmt.Fprint(conn, data)
		}
	}()

	if err := <-errCh; err == io.EOF {
		log.Println("Disconnected:", conn.RemoteAddr())
	} else {
		log.Println(err)
	}
}
