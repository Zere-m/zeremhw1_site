package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

func main() {

    //CGI header first
    fmt.Println("Content-Type: application/json\n")


    data := map[string]string{
        "greeting":     "Hello World",
        "language":     "Go",
        "generated_at": time.Now().Format("2006-01-02 15:04:05"),
        "client_ip":    os.Getenv("REMOTE_ADDR"),
    }


    jsonOutput, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        fmt.Println(`{"error": "Could not generate JSON"}`)
        return
    }


    fmt.Println(string(jsonOutput))
}
