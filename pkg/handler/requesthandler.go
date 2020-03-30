package handler

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/catchkvs/Coral/pkg/repo"
	"github.com/catchkvs/Coral/pkg/server"
	"github.com/gorilla/websocket"
	"github.com/speps/go-hashids"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

type DimensionConnInput struct {
	Id string
	Name string
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle connection")
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
	log.Println(clientMessage);
	if server.IsSessionExist(clientMessage.SessionId) {
		switch cmd := clientMessage.Command; cmd {
		case server.CMD_Auth:
			log.Println("Auth token: " + clientMessage.Data)
		case server.CMD_CreateFactEntity:
			createFactEntity(clientMessage)
		case server.CMD_GetLiveUpdates:
			getLiveUpdates(clientMessage)
		}
	}
}

func createFactEntity(clientMessage server.ClientMsg) {
	decodeFactData, _ := b64.StdEncoding.DecodeString(clientMessage.Data)
	var factEntity model.FactEntity
	json.Unmarshal(decodeFactData, &factEntity)
	factEntity.Id = newFactId()
	log.Println("creating a fact entity...", factEntity)
	repo.SaveFactEntity(&factEntity)
	store := server.GetSessionStore()

	// update the channel with fact entity
	if store.IsFactChannelPresent(factEntity.DimensionId) {
		log.Println("Updating the channel...")
		channel := store.GetFactChannel(factEntity.DimensionId)
		channel <- &factEntity
	}
}

func getLiveUpdates(clientMessage server.ClientMsg) {
	session := server.GetSessionStore().GetSession(clientMessage.SessionId)
	decodeFactData, _ := b64.StdEncoding.DecodeString(clientMessage.Data)
	var dimensionConnInput DimensionConnInput
	json.Unmarshal(decodeFactData, &dimensionConnInput)
	log.Println("Live updates for dimensionConnInput", dimensionConnInput)
	dimensionentity := repo.GetDimensionEntity(dimensionConnInput.Name, dimensionConnInput.Id)
	store := server.GetSessionStore()
	if !store.IsFactChannelPresent(dimensionentity.Id) {
		log.Println("Creating a fact channel...")
		channel := store.CreateNewFactChannel(dimensionentity.Id)
		store.AddFactChannel(dimensionentity.Id, channel)
		go factUpdator(dimensionentity.Id, channel)
	}

	// Add the session to dimension session mapping
	store.AddDimensionSession(dimensionentity.Id, session)

}

func factUpdator(dimensionId string, factChannel chan *model.FactEntity) {
	log.Println("Starting Order Updator....")
	for {
		newFact := <-factChannel
		data, _ := json.Marshal(newFact)
		store := server.GetSessionStore()
		sessions := store.GetDimensionSessions(dimensionId)
		for _, session := range sessions {
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

func newFactId() string {
	var hd = hashids.NewData()
	hd.Salt = "fact entity"
	h, err := hashids.NewWithData(hd)
	HandleError(err)
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	rand.Seed(time.Now().UnixNano())
	id, _ := h.Encode([]int{year, month, day, hour, minute, second, rand.Intn(1000)})
	return "O_"+ id
}

func HandleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}
