package demux

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Demux struct {
	Ctx context.Context
	Mux *mux.Router
}

// NewRoute wraps the HTTP transport handler and the endpoint
// with Gorilla route path.
func (demux Demux) NewRoute(
	path string,
	e endpoint.Endpoint,
	method string,
	dec httptransport.DecodeRequestFunc) {
	handler := httptransport.NewServer(
		demux.Ctx,
		e,
		dec,
		encoder,
	)
	route := func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
	demux.Mux.HandleFunc(path, route).Methods(method)
}

func encoder(writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}
