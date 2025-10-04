package services

import "net/http"

type HelloService interface {
	Hello(r *http.Request) (string, error)
}

type HelloServiceManager struct{}

func (h *HelloServiceManager) Hello(r *http.Request) (string, error) {
	st := r.URL.Query().Get("name")
	if st != "" {
		return "Hello, " + st + "!", nil
	}
	return "Hello, World!", nil
}

func NewHelloServiceManager() *HelloServiceManager {
	return &HelloServiceManager{}
}
