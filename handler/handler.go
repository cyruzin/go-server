package handler

import (
	"net/http"
	"server/pkg/rest"
)

// Root is the main entry point of this api.
func Root(w http.ResponseWriter, req *http.Request) {
	rest.ToJSON(w, http.StatusOK, &rest.APIMessage{Message: "Hello there! Welcome!"})
}

// HelloWorld is a basic example. Here we are capturing the username
// in the query params. If the user doesn't exist, we assume it's anonymous.
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	userParams := req.URL.Query()["user"]

	if len(userParams) == 0 {
		rest.ToJSON(w, http.StatusOK, &rest.APIMessage{Message: "Hello anonymous!"})
		return
	}

	userName := userParams[0]

	rest.ToJSON(w, http.StatusOK, &rest.APIMessage{Message: "Hello " + userName})
}
