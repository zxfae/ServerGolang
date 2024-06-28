package server

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// Structure of session with Username and Expiry.
type Session struct {
	Username string
	Expiry   time.Time
}

/*
Mutex is used to lock and unlock access
to the sessions map to manage concurrency.
*/
var (
	sessions = map[string]Session{}
	mutex    = sync.Mutex{}
)

func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
