package cmd

import (
	"flux/app/middleware"
	"flux/app/routes"
	"flux/app/routes/api"
	"github.com/foolin/goview"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strconv"
)

var Serve = cli.Command{
	Name:        "serve",
	Usage:       "Starts the web server",
	Description: "Flux web server is all you need to run.",
	Action:      serve,
	Flags: []cli.Flag{
		intFlag("port, p", 8080, "Web server port"),
		stringFlag("config, c", "", "Custom config path"),
	},
}

func serve(c *cli.Context) error {
	var partials []string
	err := filepath.Walk("views/partials", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := filepath.Rel("views/", path)
		if err != nil {
			return nil
		}

		partials = append(partials, file)
		return nil
	})
	if err != nil {
		return err
	}

	gv := goview.New(goview.Config{
		Root:         "views",
		Extension:    ".tmpl.html",
		Master:       "layouts/master",
		Partials:     partials,
		DisableCache: true,
		Funcs:        template.FuncMap{},
	})

	goview.Use(gv)

	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	err = routes.Init(c, r)
	if err != nil {
		return err
	}

	err = api.Init(c, r)
	if err != nil {
		return err
	}

	err = middleware.Init(c)
	if err != nil {
		return err
	}

	r.Use(middleware.HeaderMiddleware)

	return http.ListenAndServe(":"+strconv.Itoa(c.Int("port")), r)
}
