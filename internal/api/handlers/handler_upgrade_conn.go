package handlers

import (
	"fmt"
	"net/http"

	"github.com/har-sat/termchat/internal/config"
	"github.com/har-sat/termchat/utils"
)

func  HandlerUpgradeConnection(w http.ResponseWriter, r *http.Request) {
	cfg := config.Cfg

	conn, err :=cfg.Upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error upgrading connection: %v\n", err))
		return
	}
	defer conn.Close()
	cfg.Clients[conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		for client := range cfg.Clients {
			err := client.WriteMessage(messageType, p); 
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
