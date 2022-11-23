package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var times string

func main() {

	flag.StringVar(&times, "timeout", "10s", "время ожидания вышло")
	flag.Parse()
	const host = "127.0.0.1"
	const port = "8090"

	ints, _ := strconv.Atoi(times[:len(times)-1])
	to := time.Duration(ints) * time.Second

	var conn net.Conn
	var err error

	start := time.Now()
	for time.Since(start) < to {
		conn, err = net.Dial("tcp", host+":"+port)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("Таймаут %v", to)
	}
	defer conn.Close()
	log.Printf("Соеденение к: %s:%s", host, port)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			sms, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Cмс", sms)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		_, err := fmt.Fprintf(conn, in+"\n")
		if err != nil {
			log.Fatal("Exit")
		}
	}
}
