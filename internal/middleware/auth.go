package middleware

import (
	"banking-app/internal/database"
	"errors"
	"net/http"
	"time"
)

// Check if session is valid
func validSession(cookie *http.Cookie) error {
	var err error = nil
	session, err := database.GetSession(cookie.Value)
	if err != nil {
		return err
	}
	// Check session expiry
	if time.Now().Unix() >= session.ExpiryTime {
		database.DeleteSession(session.ID)
		err = errors.New("Expiried session")
	}
	return err
}

// Wraps request to check if session and user in login
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectToLogin := func() {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		cookie, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			redirectToLogin()
			return
		}
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		err = validSession(cookie)
		if err != nil {
			redirectToLogin()
			return
		}

		next(w, r)
	}
}
