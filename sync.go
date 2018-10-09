package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func parallel(wg *sync.WaitGroup) {
    fmt.Println("A")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("B")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("C")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("D")
    wg.Done()
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    wg := new(sync.WaitGroup)
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go parallel(wg)
    }
    wg.Wait()
}