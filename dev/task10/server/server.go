package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	const host = "127.0.0.1"
	const port = "8090"

	log.Println("Запуск")
	defer log.Println("остановка")

	ln, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	done := make(chan interface{})

	in := make(chan string)
	go func() {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-done:
				close(in)
				return
			default:
				msg, err := reader.ReadString('\n')
				if err == nil {
					in <- msg
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case msg := <-in:
				fmt.Fprint(conn, strings.ToUpper(msg))
			}
		}
	}()

	<-sig
	close(done)
}
