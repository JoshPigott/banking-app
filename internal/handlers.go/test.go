package handlers

import (
	"fmt"
	"net/http"
)

func Test(_ http.ResponseWriter, _ *http.Request) {
	fmt.Print("This is a test")
}

func Login(_ http.ResponseWriter, _ *http.Request) {
	fmt.Print("You are in the login page")
}
