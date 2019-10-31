package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
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
	defer conn.Close()
	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("****METHOD", m)
		}
		if ln == " " {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html><head><title>TCP Server Response</title></head>
	<body><p>Hey there</p></body></html>`

	fmt.Fprintf(conn, "HTTP/1/1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
