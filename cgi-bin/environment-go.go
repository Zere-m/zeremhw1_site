package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// CGI header
	fmt.Println("Content-Type: application/json\n")

	// map stores environment variables
	envVars := make(map[string]string)

	// going thru environment variables
	for _, env := range os.Environ() {

	
		var key, value string

		for i := 0; i < len(env); i++ {
			if env[i] == '=' {
				key = env[:i]
				value = env[i+1:]
				break
			}
		}

		envVars[key] = value
	}

	// map to json conversion
	jsonOutput, err := json.MarshalIndent(envVars, "", "  ")
	if err != nil {
		fmt.Println(`{"error": "Could not generate JSON"}`)
		return
	}


	fmt.Println(string(jsonOutput))
}
