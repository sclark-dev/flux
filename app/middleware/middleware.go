package middleware

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
)

var Context *cli.Context

func Init(c *cli.Context) error {
	Context = c

	return nil
}

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fmt.Sprintf("%s %s", Context.App.Name, Context.App.Version))
		next.ServeHTTP(w, r)
	})
}
