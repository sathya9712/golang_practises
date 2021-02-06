package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	data := os.Args
	if len(data) == 1 {
		fmt.Println("Hey,Please provide localhost:portnum")
		return
	}

	CONNECT := data[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		command, _ := reader.ReadString('\n')
		fmt.Fprintf(c, command+"\n")

		response, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + response)
		if strings.TrimSpace(string(command)) == "QUIT" {
			fmt.Println("Your TCP client is Exit...")
			return
		}
	}
}
