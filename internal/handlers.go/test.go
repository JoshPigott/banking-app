package handlers

import (
	"fmt"
	"net/http"
)

func Test(_ http.ResponseWriter, _ *http.Request) {
	fmt.Print("This is a test\n")
}
