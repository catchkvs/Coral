package server

import (
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/gorilla/websocket"
	"time"
)

// Stores the live session which are currently running in the server
type SessionStore struct {
	sessions map[string]*Session
	liveUpdateChannelMap map[string] chan *model.FactEntity
	BroadcastChannelMap map[string]  chan *model.FactEntity
	dimensionSessionMap map[string] []*Session
	factTopicSessionMap map[string] []*Session
}


// Session is started when first user connects to it the server
// a unique session Id is given to it.
type Session struct {
	SessionId string
	AuthToken string
	ConnGroup ConnectionGroup
	State string
	Tag string
	CreationTime int64
	LastHeartbeatTime int64
}

// Holds the socket connection and a unique id for it.
type Connection struct {
	Id string
	ClientAddr string
	Conn *websocket.Conn
}

// Connetion Group is to hold multiple connection
type ConnectionGroup struct {
	UserConnection Connection
}


func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

