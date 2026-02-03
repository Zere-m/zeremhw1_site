package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Content-Type: text/html\n")

	cookieHeader := os.Getenv("HTTP_COOKIE")
	cookies := map[string]string{}

	for _, part := range strings.Split(cookieHeader, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			cookies[kv[0]] = kv[1]
		}
	}

	color := cookies["favorite_color"]
	if color == "" {
		color = "(none saved)"
	}

	food := cookies["favorite_food"]
	if food == "" {
		food = "(none saved)"
	}

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
    <title>View State (Go)</title>
</head>
<body>
    <h1>Saved State (Go Cookies)</h1>

    <p><b>Favorite Color:</b> %s</p>
    <p><b>Favorite Food:</b> %s</p>

    <br>
    <a href="/cgi-bin/state-form-go">Go Back</a>

    <form action="/cgi-bin/state-clear-go" method="POST" style="margin-top:20px;">
        <button type="submit">Clear Saved State</button>
    </form>
</body>
</html>`, color, food)
}
