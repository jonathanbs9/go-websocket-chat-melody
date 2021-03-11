package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/olahol/melody.v1"
	"net/http"
)

func main() {
	// Creo una instancia de Echo
	e := echo.New()

	// Creo una instancia de Melody
	m := melody.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Le sirvo el archivo index.html para visualizar el chat
	e.GET("/", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "index.html")
		return nil
	})

	e.GET("/ws", func(c echo.Context) error {
		m.HandleRequest(c.Response().Writer, c.Request())
		return nil
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	e.Logger.Fatal(e.Start(":5000"))
}
