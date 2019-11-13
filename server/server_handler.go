package server

import (
	"log"
	"net/http"
	"test/queroquitar/giphy"

	"github.com/labstack/echo"
)

// WelcomeHTTPHandler return echo render
func WelcomeHTTPHandler(c echo.Context) error {
	qname := c.QueryParam("name")

	img, err := giphy.GetRandom("hello")
	if err != nil {
		log.Println("Giphy:", err)
	}

	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"name":  qname,
		"image": img,
	})
}
