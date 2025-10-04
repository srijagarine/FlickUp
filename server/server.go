package server

import (
	"FlickUp/endpoints"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.NewRoute().Path("/hello").Methods(http.MethodGet).Handler(endpoints.NewHelloEndpoint())

	return muxRouter
}
