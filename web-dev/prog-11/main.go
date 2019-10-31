package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Println("Connection timed out.", err)
	}
	var a string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "bye" {
			conn.Close()

		}
		fmt.Scan(&a)
		fmt.Fprintf(conn, a)
	}
	defer conn.Close()
	fmt.Println("code got here")
}
