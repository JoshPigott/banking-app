package handlers

import (
	"banking-app/internal/sessions"
	"fmt"
	"net/http"
	"time"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	userId := r.FormValue("userID")
	expiryTime := time.Now().Add(time.Hour)
	fakeExpiryTime := time.Now().Add(30 * time.Second)
	sessionID, err := sessions.CreateSession(username, userId, fakeExpiryTime)

	if err != nil {
		http.Error(w, "Fail to add to database", http.StatusInternalServerError)
		return
	} else {
		fmt.Printf("A session has been made %s\n", sessionID)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Expires:  expiryTime,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
}
