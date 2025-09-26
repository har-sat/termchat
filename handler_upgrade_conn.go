package main

import (
	"fmt"
	"net/http"
)

func (cfg *Config) HandlerUpgradeConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := cfg.Upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		respondWithError(w, 400, err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}
