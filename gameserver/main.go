package gameserver

import (
	"fmt"
	"net/http"
)

func GameServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DungeonLord GameServer")
	})
	http.ListenAndServe(":8081", nil)
}
