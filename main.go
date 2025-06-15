package main

import (
	"log"
	"os"
	"solicitudPrestamo/src/view/controller"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {
	println("Hello, World!")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY"),
	})

	if err != nil {
		log.Println("Error al iniciar sentry: ", err)
	}

	defer sentry.Flush(2 * time.Second)

	func() {
		defer func() {
			err := recover()
			if err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
			}
		}()

		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Logger(), gin.Recovery())

		r.GET("/mutant", controller.TestApp)
		r.GET("/mutant-sentry", controller.TestSentry)

		r.Run(":8087")

	}()
}
