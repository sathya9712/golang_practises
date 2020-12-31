package main

import (
"fmt"
"time"
)

func main() {
a := make(chan int, 3)
b := 10
go empl(a, 1)
go empl(a, 2)
go empl(a, 3)
go num(a, b)

time.Sleep(10 * time.Second)
}

func empl(a chan int, id int) {
for i := range a {
fmt.Printf("Worker %d receive value %d\n", id, i)
time.Sleep(1 * time.Second)
}
}

func num(a chan int, b int) {
for i := 0; i < b; i++ {
a <- i
fmt.Printf("Sent value %d\n", i)
time.Sleep(1 * time.Millisecond)
}
close(a)
}