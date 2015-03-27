package main

import (
    "fmt"
)

func main(){
    c := make(chan int)
    go func(){
        c <- 1
        c <- 2
        c <- 3
        c <- 4
        c <- 5
        c <- 6
    }()
    go func(){
        for i:=0; i<6;i++ {
            fmt.Println(<-c)
        }
    }()
}
