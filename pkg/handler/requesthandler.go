package handler

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/catchkvs/Coral/pkg/repo"
	"github.com/catchkvs/Coral/pkg/server"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}


func Handle(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	session := server.CreateNewSession(c, "Tag1")
	// Send the session id to the client
	msg := server.ServerMsg{server.CMD_ReceiveSessionId, session.SessionId, session.SessionId}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.WriteMessage(1, msgBytes);

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {

			log.Println("read:", err)
			break
		}
		log.Printf("message type: %s", mt)
		if mt == 2 {
			log.Println("Cannot process binary message right now")
		} else {
			processMessage(message)
		}
	}
}

func processMessage( msg []byte) {
	clientMessage := server.ClientMsg{}
	json.Unmarshal(msg, &clientMessage)
	//log.Println(clientMessage);
	if server.IsSessionExist(clientMessage.SessionId) {
		switch cmd := clientMessage.Command; cmd {
		case server.CMD_Auth:
			log.Println("Auth token: " + clientMessage.Data)
		case server.CMD_CreateFactEntity:
			createFactEntity(clientMessage)
		case server.CMD_UpdateFactEntity:
			updateFactEntity(clientMessage)
		case server.CMD_GetRecentEntities:
			getRecentFactEntities(clientMessage)
		case server.CMD_GetLiveUpdates:
			getLiveUpdates(clientMessage)
		}
	}
}

func createFactEntity(clientMessage server.ClientMsg) {
	decodeFactData, _ := b64.StdEncoding.DecodeString(clientMessage.Data)
	var factEntity model.FactEntity
	json.Unmarshal(decodeFactData, &factEntity)
	repo.SaveFactEntity(&factEntity)
}

func updateFactEntity(clientMessage server.ClientMsg) {

}

func getLiveUpdates(clientMessage server.ClientMsg) {
	session := server.GetSessionStore().GetSession(clientMessage.SessionId)
	dimensionentity := repo.GetDimensionEntity(clientMessage.Data)

}

func getRecentFactEntities(msg server.ClientMsg) {

}