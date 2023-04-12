package main

import "fmt"

func main() {

    token := Authenticate()
	fmt.Println(token.Access_Token)
//	Requests(token)
}
