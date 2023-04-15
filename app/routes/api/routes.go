package api

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var Context *cli.Context

func Init(c *cli.Context, router *mux.Router) error {
	Context = c

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("", Base).Methods("GET")
	apiRouter.HandleFunc("/", Base).Methods("GET")

	return nil
}
