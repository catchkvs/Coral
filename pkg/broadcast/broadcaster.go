package broadcast

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/catchkvs/Coral/pkg/server"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)
var upgrader = websocket.Upgrader{}
var mux sync.Mutex


func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle connection")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	session := server.CreateNewSession(c, "Tag1")
	// Send the session id to the client
	msg := server.ServerMsg{server.CMD_ReceiveSessionId, session.SessionId, session.SessionId}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	c.WriteMessage(1, msgBytes)

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
	log.Println(clientMessage)
	if server.IsSessionExist(clientMessage.SessionId) {
		switch cmd := clientMessage.Command; cmd {
		case server.CMD_Auth:
			log.Println("Auth token: " + clientMessage.Data)
		case server.CMD_BroadcastFact:
			broadcastFact(clientMessage)
		case server.CMD_GetLiveUpdates:
			getLiveUpdates(clientMessage)
		}
	}
}

func broadcastFact(msg server.ClientMsg) {
	decodeFactData, _ := b64.StdEncoding.DecodeString(msg.Data)
	var factEntity model.FactEntity
	json.Unmarshal(decodeFactData, &factEntity)
	store := server.GetSessionStore()
	channel := store.GetTopicChannel(factEntity.Name)
	channel <- &factEntity
}

func getLiveUpdates(msg server.ClientMsg) {
	mux.Lock()
	session := server.GetSessionStore().GetSession(msg.SessionId)
	decodeFactData, _ := b64.StdEncoding.DecodeString(msg.Data)
	var subscription model.DeviceSubscription
	json.Unmarshal(decodeFactData, &subscription)

	store := server.GetSessionStore()
	if !store.IsTopicChannelPresent(subscription.Topic) {
		log.Println("Creating a topic channel..." , subscription.Topic)
		channel := store.CreateNewTopicChannel(subscription.Topic)
		store.AddFactChannel(subscription.Topic, channel)
		go BroadcastUpdator(subscription.Topic, channel)
	}

	// Add the session to dimension session mapping
	store.AddTopicSubscription(subscription.Topic, subscription.DeviceId, session)
	mux.Unlock()
}
func BroadcastUpdator(topic string, topicChannel chan *model.FactEntity) {
	log.Println("Starting Order Updator....")

	for {
		newFact := <-topicChannel
		data, _ := json.Marshal(newFact)
		store := server.GetSessionStore()
		sessions := store.GetSessionsByFactTopic(topic)
		for _, session := range sessions {
			log.Println("Updating the session with id: " , session.SessionId)
			msg := server.ServerMsg{
				Command:   server.CMD_NewFactData,
				Data:      string(data),
				SessionId: session.SessionId,
			}
			msgData, err := json.Marshal(msg)
			HandleError(err)
			session.WriteText(msgData)

		}
		time.Sleep(10*time.Millisecond)
	}
}

func HandleError(err error){
	log.Println("Error while processing", err)
}
