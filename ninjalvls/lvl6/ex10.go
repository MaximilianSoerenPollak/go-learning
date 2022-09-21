package main

import "fmt"

func main() {
    for i:=0; i < 10; i++ {
        fmt.Println(enclosed())
    }
}


func enclosed() func() int {
    x := 0
    return func() int {
         x++
         return x
    }
}

