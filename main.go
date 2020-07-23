/* SimpleEchoServer
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	initi()
	var dial, port string
	flag.StringVar(&dial, "d", "no message", "send a msg")
	flag.StringVar(&port, "p", "no message", "send a msg")
	flag.Parse()

	//connect(dial)
	go startListen(port)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		s := strings.Split(text, " ")
		//dial = s[1]
		var command = s[0]
		if strings.Compare(command, "conn") == 0 {
			connect(dial)
		}
}

// TODO desconect a peer should not kill a host connected
func initi() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func startListen(port string) {
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkError(err)
		// convert CRLF to LF
		go handleClient(conn)
	}
}

func connect(dial string) {
	fmt.Println("connecting to... " + dial)

	var caller net.Conn
	caller, errx := net.Dial("tcp", ":"+ dial)
	// accepts multi thread
	for errx != nil {
		//fmt.Fprintf(os.Stderr, "Fatal error: %s", errx)
		caller, errx = net.Dial("tcp", ":"+ dial)
	}
	fmt.Println("connected to... " + dial)

	for {
		handleWrite(caller)
	}
}

func handleWrite(caller net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	caller.Write([]byte(text))
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
		log.Print(err.Error())
		//fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
