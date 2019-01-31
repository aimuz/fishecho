package router

import (
	"fishecho/config"
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"path/filepath"
)

type Template struct {
}

var templates map[string]*template.Template

func Init(reset bool) (err error) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	if !reset {
		for range templates {
			return
		}
	}

	templatesDir := "./template/"

	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		return err
	}

	widgets, err := filepath.Glob(templatesDir + "widgets/*.html")
	if err != nil {
		return err
	}
	for _, layout := range layouts {
		files := append(widgets, layout)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}

	return
}

func (tpl *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := Init(config.C.Debug)
	if err != nil {
		return err
	}

	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist. ", name)
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}
