package httpserver

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DungeonLord HttpServer")
	})
	http.ListenAndServe(":8080", nil)
}
