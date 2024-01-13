package websocket

import (
	"net/http"
  "fmt"
)

var SessionMap = InitLockingSessionMap()

func NewSession(id string, email string) {
	// Create a session
	var session = Session{
		Id:    id,
		Email: email,
		Cookie: http.Cookie{
			Name:     "registration_cookie",
			Value:    fmt.Sprintf(`id="%s"; Email="%s"`, id, email),
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		},
	}

	SessionMap.Set(id, session)
}
