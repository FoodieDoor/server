// Package app ties together application resources and handlers.
package app

import (
	"net/http"

	"github.com/FoodieDoor/server/logging"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

type API struct {
	Account *int8
}

func NewAPI() (*API, error) {
	api := &API{}
	return api, nil
}

func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	// r.Mount("/", something)

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
