package main

import (
	"projek-restoran/connection"
	"projek-restoran/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectToDb()
	rh := router.Handlers{
		DB: db,
		R:  r,
	}
	rh.Routes()

	r.Run(":8081")
}
