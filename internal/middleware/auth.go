package middleware

import (
	"banking-app/internal/database"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			// I will need to redirct page to login later on
			http.Redirect(w, r, "/login", http.StatusSeeOther) // This will login when I make it
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		sessionID := cookie.Value

		session, err := database.GetSession(sessionID)
		currTime := time.Now().Unix()

		if errors.Is(err, sql.ErrNoRows) {
			http.Redirect(w, r, "/login", http.StatusSeeOther) // This will login when I make it
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		} else if currTime >= session.ExpiryTime {
			err := database.DeleteSession(session.ID)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		} else if session.LoginStatus == false {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
