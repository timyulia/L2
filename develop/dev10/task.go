package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type config struct {
	host    string
	port    string
	timeout int
}

func parse() *config {
	conf := config{}
	conf.timeout = *flag.Int("timeout", 10, "timeout to connect")
	args := flag.Args()
	if len(args) > 0 {
		conf.host = args[0]
	} else {
		conf.host = "localhost"
		conf.port = "3000"
		return &conf
	}
	if len(args) > 1 {
		conf.port = args[1]
	} else {
		conf.port = "3000"
	}
	return &conf
}

type Client struct {
	Conn net.Conn
}

func newClient(conf config) (*Client, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(conf.host, conf.port), time.Duration(conf.timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected")

	return &Client{Conn: conn}, nil
}

func (c *Client) Start() {
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go handle(c.Conn, errChan)

	for {
		select {
		case <-sigChan:
			fmt.Println("\nCtrl+C: closing connection")
			return
		case err := <-errChan:
			if err == io.EOF {
				fmt.Println("\nCtrl+D: closing connection")
				return
			}
			log.Printf("got error: %s\n", err)
			return
		}
	}
}

func handle(conn net.Conn, errChan chan error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			errChan <- err
			return
		}

		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			errChan <- err
			return
		}

		fmt.Print("->: " + message)
	}
}

func main() {
	conf := parse()
	client, err := newClient(*conf)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Conn.Close()

	client.Start()
}
