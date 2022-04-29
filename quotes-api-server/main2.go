package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

//var quotes []quote = make([]quote, 0)

var quotes []quote = []quote{
	{"ABC", "ABCD"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/", getQuotes)
	api.GET("/quotes", getRandomQuotes)

	port := os.Getenv("Port")
	if port == "" {
		port = "32445"
		//localhost:9090
	}

	api.Start(":" + port)

}

func getRandomQuotes(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}
func getQuotes(c echo.Context) error {
	return c.JSON(200, quotes)
}
