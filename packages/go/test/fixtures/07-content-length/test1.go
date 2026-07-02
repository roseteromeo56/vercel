package function

import (
	"html"
	"net/http"
	"os"
	"strconv"
)

// HandlerTest1 function
func HandlerTest1(w http.ResponseWriter, r *http.Request) {
	rdm := os.Getenv("RANDOMNESS_ENV_VAR")
	body := html.EscapeString(rdm) + ":content-length"

	w.WriteHeader(401)
	w.Header().Set("content-length", strconv.Itoa(len(body)))
	w.Write([]byte(body))
}
