package sessionmanager

import (
	"net/http"
	"sandbox/internal/structs"
	"time"
)

const SESSION_EXPIRY = 30 * 24 * 60 * 60

// Entry point for handling sessions, if a session is not found
func Handle(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		// make a new guest cookie
		newSession := newGuestSession()
		newCookie := http.Cookie{
			Name:     "session",
			Value:    newSession.Token,
			Expires:  time.Unix(newSession.CreationTime+SESSION_EXPIRY, 0),
			HttpOnly: true, // to prevent XSS
			Path:     "/",
		}
		http.SetCookie(w, &newCookie)
		return
	}
	existingSession := getSession_stub(sessionCookie.Value)
	// make a new session
	if existingSession == nil {
		return
	}
	// handle expired sessions

}
func getSession_stub(token string) *structs.Session {
	// waiting for ruqaya to write the database.GetSession() - already done!
	panic("Get Session not implemented")
}

func newGuestSession() *structs.Session {
	panic("newGuestSession not implemented")
}

func renewSession() *structs.Session {
	panic("newGuestSession not implemented")
}
