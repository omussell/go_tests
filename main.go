package main

import "github.com/brianvoe/gofakeit"
import "fmt"

func main() {
        gofakeit.Seed(0)
        fmt.Println(gofakeit.Name())
}
