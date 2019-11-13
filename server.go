package main

import (
	"io"
	"log"
	"net/http"
	"test/queroquitar/giphy"
	"text/template"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		qname := c.QueryParam("name")

		img, err := giphy.GetRandomGif("hello")
		if err != nil {
			log.Println("Giphy:", err)
		}

		return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
			"name":  qname,
			"image": img,
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
