package main

import "fmt"

var score = 99.5
var points = []int{10, 20, 30}

func main(){
    sayHello("mario")
    showScore()

    for _, v := range points{
        fmt.Println(v)
    }
}

func sayHello(name string) {
    fmt.Printf("Hello %s\n", name)
}

func showScore() {
    fmt.Printf("score: %.1f\n", score)
}