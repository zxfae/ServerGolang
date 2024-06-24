package server

import (
	"time"
)

/*
Invalidates existing session for the user,
locks the sessions map,
generates a new session ID,
sets the expiry time,
adds new session to the map,
and session ID.
*/
func createSession(username string) string {
	invalidateSession(username)
	mutex.Lock()
	defer mutex.Unlock()

	sessionID := generateSessionID()
	expiry := time.Now().Add(24 * time.Hour)

	sessions[sessionID] = Session{
		Username: username,
		Expiry:   expiry,
	}
	return sessionID
}

/*
Locks the sessions map, deletes the session with the given session ID.
*/
func deleteSession(sessionID string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(sessions, sessionID)
}

/*
getSession: Locks the sessions map, check session by session ID, checks if session is expired.
If not found or expired, it returns false.
*/
func getSession(sessionID string) (Session, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	session, exists := sessions[sessionID]
	if !exists || session.Expiry.Before(time.Now()) {
		return Session{}, false
	}

	return session, true
}

// Locks the sessions map,
// iterates through the sessions to find and delete any session for the given username
func invalidateSession(username string) {
	mutex.Lock()
	defer mutex.Unlock()

	for id, session := range sessions {
		if session.Username == username {
			delete(sessions, id)
			break
		}
	}
}
