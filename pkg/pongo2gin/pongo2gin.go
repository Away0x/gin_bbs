// Package pongo2gin is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the Pongo2 template
// library https://github.com/flosch/pongo2
// https://godoc.org/gitlab.com/go-box/pongo2gin
package pongo2gin

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// RenderOptions is used to configure the renderer.
type RenderOptions struct {
	TemplateDir string
	ContentType string
}

// Pongo2Render is a custom Gin template renderer using Pongo2.
type Pongo2Render struct {
	Options  *RenderOptions
	Template *pongo2.Template
	Context  pongo2.Context
}

// New creates a new Pongo2Render instance with custom Options.
func New(options RenderOptions) *Pongo2Render {
	return &Pongo2Render{
		Options: &options,
	}
}

// Default creates a Pongo2Render instance with default options.
func Default() *Pongo2Render {
	return New(RenderOptions{
		TemplateDir: "templates",
		ContentType: "text/html; charset=utf-8",
	})
}

// Instance should return a new Pongo2Render struct per request and prepare
// the template by either loading it from disk or using pongo2's cache.
func (p Pongo2Render) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	filename := path.Join(p.Options.TemplateDir, name)

	// always read template files from disk if in debug mode, use cache otherwise.
	if gin.Mode() == "debug" {
		template = pongo2.Must(pongo2.FromFile(filename))
	} else {
		template = pongo2.Must(pongo2.FromCache(filename))
	}

	return Pongo2Render{
		Template: template,
		Context:  data.(pongo2.Context),
		Options:  p.Options,
	}
}

// Render should render the template to the response.
func (p Pongo2Render) Render(w http.ResponseWriter) error {
	p.WriteContentType(w)
	err := p.Template.ExecuteWriter(p.Context, w)
	return err
}

// WriteContentType should add the Content-Type header to the response
// when not set yet.
func (p Pongo2Render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{p.Options.ContentType}
	}
}
