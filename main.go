package main

import (
	"Jur/controller"
	"Jur/router"

	"github.com/gin-gonic/gin"
)

func main() {
	rute := gin.Default()

	// database, err := config.GalaSetup()
	// if err != nil {
	// log.Fatalf("Fail : %v", err)
	// }

	// database.AutoMigrate(&entities.SampleTabler{})

	rute.GET("/", controller.MainHallo)

	router.SampleRouter(rute)

	rute.Run()
}
