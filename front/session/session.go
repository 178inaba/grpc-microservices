package session

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/xid"
)

const (
	sessionExpires      = 3 * time.Minute
	clearInterval       = 10 * time.Minute
	cookieNameSessionID = "session-id"
)

type Store interface {
	Set(k string, x interface{})
	Delete(k string)
	Get(k string) (interface{}, bool)
}

type StoreOnMemory struct {
	cache *cache.Cache
}

func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{
		cache.New(sessionExpires, clearInterval),
	}
}

func (s *StoreOnMemory) Set(k string, x interface{}) {
	s.cache.Set(k, x, sessionExpires)
}

func (s *StoreOnMemory) Delete(k string) {
	s.cache.Delete(k)
}

func (s *StoreOnMemory) Get(k string) (interface{}, bool) {
	return s.cache.Get(k)
}

func ID() string {
	return xid.New().String()
}

func GetSessionIDFromRequest(r *http.Request) string {
	c, err := r.Cookie(cookieNameSessionID)
	if err != nil {
		return ""
	}

	return c.Value
}

func SetSessionIDToResponse(w http.ResponseWriter, id string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieNameSessionID,
		Value:    id,
		Expires:  time.Now().Add(sessionExpires),
		HttpOnly: true,
	})
}

func DeleteSessionIDFromResponse(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieNameSessionID,
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
	})
}
