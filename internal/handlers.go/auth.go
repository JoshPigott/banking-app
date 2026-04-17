package handlers

import (
	"banking-app/internal/services"
	"net/http"
)

// Set reponse with a cookie containing session id
func setSessionCookie(w http.ResponseWriter, userID string) error {
	sessionID, expiryTime, err := services.CreateSession(userID)
	if err != nil {
		return err
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
	return nil
}

// Create an account and session if valid username and password
func SignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !services.IsValidPassword(password) || !services.IsValidUsername(username) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`<div>Invalid password or username</div>`))
		return
	}

	userID, err := services.CreateUserAccount(username, password)
	if err != nil {
		http.Error(w, "Interal server error", http.StatusInternalServerError)
		return
	}

	err = setSessionCookie(w, userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// If username and password are valid create a session
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	valid, userID := services.ValidLoginCredentials(username, password)

	if !valid {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`<div>Invalid password or username</div>`))
		return
	}

	err = setSessionCookie(w, userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
