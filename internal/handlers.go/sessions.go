package handlers

import (
	"banking-app/internal/sessions"
	"fmt"
	"net/http"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	userId := r.FormValue("userID")
	sessionID, err := sessions.CreateSession(username, userId)

	if err != nil {
		http.Error(w, "Fail to add to database", http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("A session has been made %s\n", sessionID)
	}
	// I need to add expiry here
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
}
