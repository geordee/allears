package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		slog.Info(fmt.Sprintf("\n%s\n---\n", string(dump)))
		w.WriteHeader(http.StatusOK)
	})

	slog.Info(fmt.Sprintf("Server running on port %s\n", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
