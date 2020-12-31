package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
file, err := os.Open("sa.txt")

if err != nil {
log.Fatalf("failed opening file: %s", err)
}

scanner := bufio.NewScanner(file)
scanner.Split(bufio.ScanLines)
var a []string

for scanner.Scan() {
a = append(a, scanner.Text())
}

file.Close()

for _, b := range a {
fmt.Println(b)
}
}
