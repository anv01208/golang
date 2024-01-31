package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/putnug1122/snippetbox/internal/models"
	"github.com/putnug1122/snippetbox/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var function = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.go.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		pattern := []string{
			"html/base.go.tmpl",
			"html/partials/*.go.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(function).ParseFS(ui.Files, pattern...)
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob("./ui/html/partials/*.go.tmpl")
		// if err != nil {
		// 	return nil, err
		// }

		// ts, err = ts.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}
	return cache, nil
}
