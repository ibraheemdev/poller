package renderer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"path"
	"path/filepath"

	"github.com/ibraheemdev/poller/pkg/authboss"
)

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

type templateRenderer struct {
	// url mount path
	mountPath string

	templates map[string]*template.Template
}

// NewTemplateRenderer creates a new setup to render layout based go templates
func newTemplateRenderer(mountPath string, layoutsDir string, templatesDir string) *templateRenderer {
	r := &templateRenderer{
		templates: make(map[string]*template.Template),
		mountPath: mountPath,
	}
	r.Load(layoutsDir, templatesDir)
	return r
}

func (t *templateRenderer) Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	tmpl, ok := t.templates[page]
	if !ok {
		return nil, "", fmt.Errorf("the template %s does not exist", page)
	}
	buf := &bytes.Buffer{}
	err = tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return nil, "", fmt.Errorf("failed to render template for page %s: %w", page, err)
	}
	return buf.Bytes(), "text/html", nil
}

func (t *templateRenderer) Load(layoutsDir, templatesDir string) {
	layouts, err := filepath.Glob(layoutsDir)
	if err != nil {
		log.Fatal(err)
	}

	templates, err := filepath.Glob(templatesDir)
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"mountpathed": func(location string) string {
			if t.mountPath == "/" {
				return location
			}
			return path.Join(t.mountPath, location)
		},
	}

	mainTemplate, err := template.New("main").Funcs(funcMap).Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range templates {
		fileName := filepath.Base(tpl)
		files := append(layouts, tpl)

		t.templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}

		t.templates[fileName] = template.Must(t.templates[fileName].ParseFiles(files...))
	}
}
