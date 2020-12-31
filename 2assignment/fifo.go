package main

import (
    "fmt"
    "sync"
)

func main() {
    queue := make(chan int)

    wg := new(sync.WaitGroup)
    wg.Add(1)
    defer wg.Wait()

    go func(wg *sync.WaitGroup) {
        for {

            r, ok := <-queue
            if !ok {
                wg.Done()
                return
            }

            fmt.Println(r)
        }
    }(wg)

    for i := 2; i <= 10; i++ {
        queue <- i
    }

    close(queue)
}