package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irwNd2/gin-rest-api/controllers/productcontroller"
	"github.com/irwNd2/gin-rest-api/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Fetch)
	r.GET("/api/products/:id", productcontroller.Detail)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products/:id", productcontroller.Delete)

	r.Run()
}
