package http

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	ID        string
	Data      map[string]interface{}
	ExpiresAt time.Time
}

var (
	sessionStore   = make(map[string]*Session)
	sessionTTL     = 30 * time.Minute
	sessionIDLen   = 32
	sessionIDChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sessionIDMutex sync.Mutex
	sessionMutex   sync.Mutex
)

const sessionCookieName = "expressgo_session_id"

func generateSessionID() string {
	sessionIDMutex.Lock()
	defer sessionIDMutex.Unlock()

	b := make([]byte, sessionIDLen)
	for i := range b {
		b[i] = sessionIDChars[rand.Intn(len(sessionIDChars))]
	}
	return string(b)
}

func (ctx *Context) GetSession() *Session {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	var sessionId string
	cookie, err := ctx.Request.r.Cookie(sessionCookieName)
	if err != nil || cookie.Value == "" {
		sessionId = generateSessionID()
		http.SetCookie(ctx.Response.Writer, &http.Cookie{
			Name:     sessionCookieName,
			Value:    sessionId,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // Set to true if using HTTPS
			MaxAge:   int(sessionTTL.Seconds()),
			Expires:  time.Now().Add(sessionTTL),
		})
	} else {
		sessionId = cookie.Value
	}

	session, exists := sessionStore[sessionId]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		session = &Session{
			ID:        sessionId,
			Data:      make(map[string]interface{}),
			ExpiresAt: time.Now().Add(sessionTTL),
		}
		sessionStore[sessionId] = session
	}
	session.ExpiresAt = time.Now().Add(sessionTTL)
	ctx.Response.Writer.Header().Set("X-Session-ID", session.ID)
	return session
}

func (ctx *Context) SetSessionData(key string, value interface{}) {
	session := ctx.GetSession()
	session.Data[key] = value
	sessionStore[session.ID] = session
}

func (ctx *Context) GetSessionData(key string) (interface{}, bool) {
	session := ctx.GetSession()
	value, exists := session.Data[key]
	return value, exists
}

func (ctx *Context) DeleteSessionData(key string) {
	session := ctx.GetSession()
	delete(session.Data, key)
	sessionStore[session.ID] = session
}
