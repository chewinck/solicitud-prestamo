package controller

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func TestApp(c *gin.Context) {
	c.String(200, "Sollcitud prestamo up\n")
}

func TestSentry(c *gin.Context) {
	_, err := strconv.Atoi("TEST")
	if err != nil {
		sentry.CaptureException(err)
		c.String(200, "Sentry up\n")
		return
	}
	c.String(200, "Sentry down\n")
}
