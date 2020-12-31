package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
rand.Seed(time.Now().Unix())

fmt.Printf(" %d\n", rand.Intn(10))

fmt.Printf("%d\n", rand.Int31n(10))

fmt.Printf(" %d\n", rand.Int63n(10))

fmt.Printf("%d\n", rand.Int())

fmt.Printf("Int31: %d\n", rand.Int31())

fmt.Printf("Int63: %d\n", rand.Int63())
 
}