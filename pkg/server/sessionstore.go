package server



// Stores the live session which are currently running in the server
type SessionStore struct {
	sessionMap map[string]*Session
}

func (store *SessionStore) IsSessionExist(sessionId string) bool {
	if _, ok:= store.sessionMap[sessionId]; ok {
		return true
	}
	return false
}
