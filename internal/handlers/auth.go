package handlers

import (
	"banking-app/internal/database"
	"banking-app/internal/helpers"
	"net/http"
)

// Create an account and session if valid username and password
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Unpacks form data to get username and password
	err := r.ParseForm()
	if err != nil {
		writeServerError(w)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	// Check if username and password are valid
	if !helpers.IsValidCredentials(username, password) {
		writeInvalidCredentials(w)
		return
	}
	// Create account in database
	userID, err := helpers.CreateUserAccount(username, password)
	if err != nil {
		writeServerError(w)
		return
	}
	// Adds new session for the account
	err = setSessionCookie(w, userID)
	if err != nil {
		writeServerError(w)
		return
	}
	writeAuthSuccess(w)
}

// If username and password are valid create a session
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	// Unpacks form data to get username and password
	err := r.ParseForm()
	if err != nil {
		writeServerError(w)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if password and username and username belong to an account
	valid, userID := helpers.ValidLoginCredentials(username, password)
	if !valid {
		writeInvalidCredentials(w)
		return
	}
	// Adds new session for the account
	err = setSessionCookie(w, userID)
	if err != nil {
		writeServerError(w)
		return
	}
	writeAuthSuccess(w)
}

// Deletes user session
func Logout(w http.ResponseWriter, r *http.Request) {
	// Gets cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		logoutError(w)
		return
	}
	if err = database.DeleteSession(cookie.Value); err != nil {
		logoutError(w)
		return
	}
	w.Header().Set("HX-Redirect", "/login")
}

// Set reponse with a cookie containing session id
func setSessionCookie(w http.ResponseWriter, userID string) error {
	// Creates session in database
	sessionID, expiryTime, err := helpers.CreateSession(userID)
	if err != nil {
		return err
	}

	// Sets response to contain a cookie containing the session
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

func writeInvalidCredentials(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(`<div class="login__message login__message--show">
	<p>Invalid password or username</p></div>`))
}

func writeServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(`<div class="login__message login__message--show">
	<p>Unable to process request</p></div>`))
}

func writeAuthSuccess(w http.ResponseWriter) {
	w.Header().Set("HX-Redirect", "/online-banking")
	w.WriteHeader(http.StatusCreated)
}

func logoutError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte(`Unable to login out`))
}
