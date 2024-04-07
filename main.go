package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/skill-map/database"
	"github.com/muhammedarifp/skill-map/handlers"
	"github.com/muhammedarifp/skill-map/managers"
)

func init() {
	err := database.InitDatabase()
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	app := gin.Default()

	manager := managers.NewUserManager()
	handler := handlers.NewUserHandlerFrom(*manager)
	handler.RegisterApis(app)

	app.Run()
}
