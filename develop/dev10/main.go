package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleServ(conn)
	}
}

// handle представляет собой простейший обработчик tcp соединения, который выводит данные, записанные клиентом в сокет
func handleServ(conn net.Conn) {
	defer conn.Close()
	notify := make(chan error)
	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				notify <- err
				return
			}

			if strings.TrimSpace(data) == "STOP" {
				fmt.Println("Exiting TCP server")
				notify <- nil
				return
			}

			fmt.Fprintf(conn, "I heard you said: %s", data)
		}
	}()

	for {
		select {
		case err := <-notify:
			if io.EOF == err {
				fmt.Println("connection dropped", err)
			}
		}
	}
}
