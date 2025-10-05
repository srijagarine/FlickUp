package endpoints

import (
	"FlickUp/services"
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

type helloEndpoint struct {
	HelloService services.HelloService
}

func (e *helloEndpoint) DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func (e *helloEndpoint) ProcessRequest(_ context.Context, request interface{}) (interface{}, error) {
	r := request.(*http.Request)
	return e.HelloService.Hello(r)
}

func (e *helloEndpoint) EncodeResponse(_ context.Context, w http.ResponseWriter, request interface{}) error {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := w.Write([]byte(request.(string)))
	return err
}

func NewHelloEndpoint() *kithttp.Server {
	ep := &helloEndpoint{
		HelloService: services.NewHelloServiceManager(),
	}

	return newEndpointServerBuilder(ep).Build()
}
