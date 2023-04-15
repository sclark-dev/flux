package api

import (
	"flux/app/utils"
	"net/http"
	"strings"
)

func Base(w http.ResponseWriter, r *http.Request) {
	json.WriteJsonHTTP(w, 200, map[string]any{
		"api":     strings.ToLower(Context.App.Name),
		"version": Context.App.Version,
	})
}
