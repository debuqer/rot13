package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--server" {
		li, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Panic(err)
		}
		defer li.Close()

		fmt.Println("Starting in port 8080:")
		for {
			conn, err := li.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			go handle(conn)
		}

	} else if os.Args[1] == "--version" {
		fmt.Printf("GOROT13 version 0.0.1\n")
	} else {
		fmt.Printf("%s\n", rot13([]byte(os.Args[1])))
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()

		fmt.Fprintf(conn, "Goodbye!\n")
		if text == "exit" {
			break
		}

		encrypted := rot13([]byte(text))

		fmt.Fprintf(conn, string(encrypted)+"\n")
	}

	defer conn.Close()
}

func rot13(s []byte) []byte {
	var chars = make([]byte, len(s))

	for index, element := range s {
		if element == 32 {
			chars[index] = 32
		} else if element <= 109 {
			chars[index] = element + 13
		} else {
			chars[index] = element - 13
		}
	}

	return chars
}
