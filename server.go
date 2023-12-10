package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/elanutta/video-api/internal/video"
)

func main() {
	e := echo.New()

	e.POST("/", StartConsumingCamera)

	e.Logger.Fatal(e.Start(":1323"))
}

func StartConsumingCamera(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	url := c.FormValue("url")

	video.GetVideoStream("rtsp://" + login + ":" + password + "@" + url)

	return c.String(http.StatusOK, "ok")
}
