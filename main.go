package main

import "fmt"

func main() {

	token := Authenticate()
    fmt.Println(token)
    Requests(token)
}
