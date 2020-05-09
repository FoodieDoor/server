// Package app ties together application resources and handlers.
package app

import (
	"net/http"

	"github.com/FoodieDoor/server/logging"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// Router provides application routes.
// func () Router() *chi.Mux {
// 	r := chi.NewRouter()
// 	return r
// }

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
