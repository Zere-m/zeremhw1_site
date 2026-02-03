package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {


	fmt.Println("Content-Type: application/json")
	fmt.Println("Cache-Control: no-cache")
	fmt.Println()


	method := os.Getenv("REQUEST_METHOD")
	contentType := os.Getenv("CONTENT_TYPE")

	clientIP := os.Getenv("REMOTE_ADDR")
	userAgent := os.Getenv("HTTP_USER_AGENT")
	host := os.Getenv("HTTP_HOST")

	timestamp := time.Now().Format("2006-01-02 15:04:05")


	dataReceived := make(map[string]string)


	if method == "GET" {
		queryString := os.Getenv("QUERY_STRING")

		values, _ := url.ParseQuery(queryString)
		for key, val := range values {
			if len(val) > 0 {
				dataReceived[key] = val[0]
			}
		}

	} else {
		
		bodyBytes, _ := io.ReadAll(os.Stdin)
		body := string(bodyBytes)

		// json body
		if strings.Contains(contentType, "application/json") {
			var jsonData map[string]string
			err := json.Unmarshal(bodyBytes, &jsonData)

			if err != nil {
				dataReceived["error"] = "Invalid JSON received"
			} else {
				dataReceived = jsonData
			}

		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {

			//parsing d body
			values, _ := url.ParseQuery(body)
			for key, val := range values {
				if len(val) > 0 {
					dataReceived[key] = val[0]
				}
			}

		} else {
			// raw body if encoding is unknown
			dataReceived["raw_body"] = body
		}
	}


	response := map[string]any{
		"message":       "Echo Endpoint (Go)",
		"language":      "Go",
		"method":        method,
		"hostname":      host,
		"timestamp":     timestamp,
		"client_ip":     clientIP,
		"user_agent":    userAgent,
		"data_received": dataReceived,
	}


	jsonOutput, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(`{"error": "Could not generate JSON output"}`)
		return
	}

	fmt.Println(string(jsonOutput))
}
