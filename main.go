package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("requests.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file")
			return
		}
		defer file.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			return
		}

		_, err = io.WriteString(file, fmt.Sprintf("\nMethod: %s, URL: %s\nBody: %s\n", r.Method, r.URL, string(body)))
		if err != nil {
			fmt.Println("Error writing to file")
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
