package endpoints

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

type Endpoint interface {
	DecodeRequest(ctx context.Context, r *http.Request) (interface{}, error)
	ProcessRequest(ctx context.Context, request interface{}) (interface{}, error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, request interface{}) error
}

type EndpointServerBuilder struct {
	ep Endpoint
}

func newEndpointServerBuilder(ep Endpoint) *EndpointServerBuilder {
	return &EndpointServerBuilder{ep: ep}
}

func (b *EndpointServerBuilder) Build() *kithttp.Server {
	return kithttp.NewServer(
		b.ep.ProcessRequest,
		b.ep.DecodeRequest,
		b.ep.EncodeResponse,
	)
}
