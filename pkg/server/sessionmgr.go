package server

import (
	"github.com/gorilla/websocket"
	"github.com/speps/go-hashids"
	"log"
	"time"
)

const(
	SESSION_STARTED = "STARTED"
	SESSION_INPROGRESS = "INPROGRESS"
	SESSION_ENDED = "ENDED"

)

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


// Creates a new session
func CreateNewSession(conn *websocket.Conn, tag string) *Session {
	id := newHashId()
	userConnect := Connection{id, conn.RemoteAddr().String(), conn}
	conngroup := ConnectionGroup{userConnect}
	creationTime := time.Now().Unix()
	session := Session { id, "",conngroup , "STARTED", 	tag,creationTime, creationTime	}
	SessionStore[id] = &session
	return &session

}

// write the binary data to the socket
func (s *Session) WriteBinary(data []byte, connType int) {
	log.Println("Start Writing to the connection")
	s.ConnGroup.UserConnection.Conn.WriteMessage(2, data)
	log.Println("finished writing to the connection")

}

// Write the text data to the socket
func (s *Session) WriteText(data []byte, connType int) {
	s.ConnGroup.UserConnection.Conn.WriteMessage(1, data)
	log.Println("finished writing to the connection")
}

func newHashId() string {
	var hd = hashids.NewData()
	hd.Salt = "Coral Server"
	h, err := hashids.NewWithData(hd)
	handleError(err)
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	id, _ := h.Encode([]int{year, month, day, hour, minute, second})
	return id
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}
