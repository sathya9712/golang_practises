package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func enq(q []string, value string) []string {
	q = append(q, value)
	fmt.Println("Enqueued:", value)
	return q
}

func dq(q []string) []string {
	value := q[0]
	fmt.Println("Dequeued:", value)
	return q[0:]
}

func main() {

	var q []string

	data := os.Args
	if len(data) == 1 {
		fmt.Println("Hey,Please provide port number for server")
		return
	}

	input := ":" + data[1]
	a, err := net.Listen("tcp", input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer a.Close()

	c, err := a.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		infodata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		q = enq(q, infodata)
		fmt.Println(q)

		q = dq(q)
		fmt.Println(q)
		if strings.TrimSpace(string(infodata)) == "QUIT" {
			fmt.Println("Your TCP server is Exit!")
			return
		}

		fmt.Print("=> ", string(infodata))
		b := time.Now()
		currentime := b.Format("2006-01-02 15:04:05 Monday") + "\n"
		c.Write([]byte(currentime))
	}
}
