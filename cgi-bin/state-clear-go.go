package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Set-Cookie: favorite_color=; Max-Age=0; Path=/\n")
	fmt.Printf("Set-Cookie: favorite_food=; Max-Age=0; Path=/\n")
	fmt.Println("Status: 302 Found")
	fmt.Println("Location: /cgi-bin/state-view-go")
	fmt.Println("Content-Type: text/html\n")
}
