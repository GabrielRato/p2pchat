/* SimpleEchoServer
 */
package main

import (
	"bufio"
	"flag"
	"net"
	"os"
	"fmt"
)

func main() {
	var dial, port string
	flag.StringVar(&dial, "d", "no message", "send a msg")
	flag.StringVar(&port, "p", "no message", "send a msg")
	flag.Parse()
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	var caller net.Conn
	caller, errx := net.Dial("tcp", ":"+ dial)

	conn, err := listener.Accept()
	if err != nil {
	}
	for {
		// accepts multi thread
		for errx != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
			caller, errx = net.Dial("tcp", ":"+ dial)
		}
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		go caller.Write([]byte(text))
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		checkError(err2)
		fmt.Print("-> ")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
