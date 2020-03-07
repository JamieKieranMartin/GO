package main

import "fmt"
import "time"

func main() {
	go heartBeat()
	fmt.Scanln()
}

func heartBeat() {
    for range time.Tick(time.Second * 1) {
        fmt.Println("Foo")
    }
}