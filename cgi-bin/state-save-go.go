package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
)

func main() {
	cl := os.Getenv("CONTENT_LENGTH")
	length := 0
	if cl != "" {
		fmt.Sscanf(cl, "%d", &length)
	}

	data := make([]byte, length)
	if length > 0 {
		_, err := io.ReadFull(os.Stdin, data)
		if err != nil {
			fmt.Println("Content-Type: text/plain\n")
			fmt.Println("Error reading POST data:", err)
			return
		}
	}

	values, _ := url.ParseQuery(string(data))
	color := values.Get("color")
	food := values.Get("food")

	fmt.Printf("Set-Cookie: favorite_color=%s; Max-Age=3600; Path=/\n", color)
	fmt.Printf("Set-Cookie: favorite_food=%s; Max-Age=3600; Path=/\n", food)

	fmt.Println("Status: 302 Found")
	fmt.Println("Location: /cgi-bin/state-view-go")
	fmt.Println("Content-Type: text/html\n")
}
