package main

import (
	"fmt"
)

func main() {
	fmt.Println("Content-Type: text/html\n")

	fmt.Println(`<!DOCTYPE html>
<html>
<head>
    <title>State Demo (Go)</title>
</head>
<body>
    <h1>State Demo (Go Cookies)</h1>

    <p>Enter some data to save in cookies:</p>

    <form action="/cgi-bin/state-save-go" method="POST">
        <label>Your favorite color:</label>
        <input type="text" name="color" required><br><br>

        <label>Your favorite food:</label>
        <input type="text" name="food" required><br><br>

        <button type="submit">Save State</button>
    </form>

    <br>
    <a href="/cgi-bin/state-view-go">View Saved State</a>
</body>
</html>`)
}
