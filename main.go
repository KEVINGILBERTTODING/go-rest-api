package main

import (
	"github.com/KEVINGILBERTTODING/go-rest-api/controllers/productcontroller"
	"github.com/KEVINGILBERTTODING/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products/create", productcontroller.Create)
	r.POST("/api/products/update", productcontroller.Update)
	r.POST("/api/products/delete", productcontroller.Delete)

	r.Run()
}
