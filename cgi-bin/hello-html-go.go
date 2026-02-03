

package main

import (
	"fmt"
	"os"
	"time"
)

func main(){

    fmt.Println("Content-Type: text/html\n")

    fmt.Println("<!DOCTYPE html>")
    fmt.Println("<html>")
    fmt.Println("<head><title>Hello HTML Go</title></head>")
    fmt.Println("<body>")

    fmt.Println("<h1 align='center'>Hello HTML World</h1><hr>")
    fmt.Println("<p>Language: Go</p>")
    fmt.Printf("<p>Generated at: %s</p>\n", time.Now().Format("2006-01-02 15:04:05"))

    clientIP := os.Getenv("REMOTE_ADDR")
    fmt.Printf("<p>Your IP Address: %s</p>\n", clientIP)

    fmt.Println("</body></html>")
}
