package server

import (
	"github.com/catchkvs/Coral/pkg/config"
	"github.com/catchkvs/Coral/pkg/metrics"
	"github.com/catchkvs/Coral/pkg/model"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/speps/go-hashids"
	"log"
	"math/rand"
	"time"
)

const(
	SESSION_STARTED = "STARTED"
	SESSION_INPROGRESS = "INPROGRESS"
	SESSION_ENDED = "ENDED"

)

var store *SessionStore
func init() {
	tmp := SessionStore{
		sessions:             map[string]*Session{},
		liveUpdateChannelMap: map[string] chan *model.FactEntity {},
		dimensionSessionMap:  map[string] []*Session{},
	}
	store = &tmp
}

// Check if session is present in session store
func IsSessionExist(sessionId string) bool {
	if _, ok:= store.sessions[sessionId]; ok {
		return true
	}
	return false
}

func GetSessionStore() *SessionStore {
	return store
}

// Get all sessions.
func (store *SessionStore) GetAllSessions() map[string]*Session {
	return store.sessions
}

func (store *SessionStore) GetSession(sessionId string) *Session {
	return store.sessions[sessionId]
}

// Track a new Dimension Session
func (store *SessionStore) AddDimensionSession(dimensionId string, session *Session) {
	allSessions := store.dimensionSessionMap[dimensionId]
	if len(allSessions) == 10 {
		log.Println("Already reach maximum sessions for this dimension");
	}
	for _, existingSession := range allSessions {
		if existingSession.SessionId == session.SessionId {
			log.Println("Session already present ")
			return
		}
	}

	store.dimensionSessionMap[dimensionId] = append(store.dimensionSessionMap[dimensionId], session)
}

func (store *SessionStore) AddFactChannel(dimensionId string, channel chan *model.FactEntity) bool {
	if _, ok:= store.liveUpdateChannelMap[dimensionId]; ok {
		return false
	}
	store.liveUpdateChannelMap[dimensionId] = channel
	return true
}

func (store *SessionStore) IsFactChannelPresent(dimensionId string) bool {
	if _, ok:= store.liveUpdateChannelMap[dimensionId]; ok {
		return true
	}
	return false;
}

func (store *SessionStore) GetFactChannel(dimensionId string) chan *model.FactEntity {
	return store.liveUpdateChannelMap[dimensionId]
}

func (store *SessionStore) GetDimensionSessions(dimensionId string) []*Session {
	return store.dimensionSessionMap[dimensionId]
}



func (store *SessionStore) CreateNewFactChannel(dimensionId string) chan *model.FactEntity {
	return make(chan *model.FactEntity, 100)
}




// Creates a new session associated with a given connection
func CreateNewSession(conn *websocket.Conn, tag string) *Session {
	metrics.SessionCounter.With(prometheus.Labels{"session":tag}).Inc()
	log.Println("creating a new session...")
	id := newHashId()
	userConnect := Connection{id, conn.RemoteAddr().String(), conn}
	conngroup := ConnectionGroup{userConnect}
	creationTime := time.Now().Unix()
	session := Session { id, "",conngroup , "STARTED", 	tag,creationTime, creationTime	}
	store.sessions[id] = &session
	return &session

}

// write the binary data to the socket
func (s *Session) WriteBinary(data []byte) {
	log.Println("Start Writing to the connection")
	s.ConnGroup.UserConnection.Conn.WriteMessage(2, data)
	log.Println("finished writing to the connection")

}

// Write the text data to the socket
func (s *Session) WriteText(data []byte) {
	s.ConnGroup.UserConnection.Conn.WriteMessage(1, data)
	log.Println("finished writing to the connection")
}


// Cleanup work to remove stale sessions which runs every 5 mins.
func CleanupWorker() {
	for  {
		for sessionId, session := range store.GetAllSessions() {
			timeDiff := time.Now().Unix() - session.LastHeartbeatTime
			if timeDiff > config.GetSessionTimeout() {
				delete(store.sessions, sessionId)
			}
		}
		time.Sleep(300*time.Second)
	}
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
	randomness := rand.Int()
	rand.Seed(time.Now().UnixNano())
	a := []int {year, month, day, hour, minute, second, randomness}
	for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	id, _ := h.Encode(a)
	return id
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}
