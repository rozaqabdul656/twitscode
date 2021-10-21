package main

import (
	"testkerjatwits/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	authorized := r.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	authorized.Use()
	{
		authorized.GET("/binarytodecimal", controller.BinaryToDecimal) // Nomer 1 a
		authorized.GET("/decimaltobinary", controller.DecimalToBinary) // Nomer 1 b
	}

	controller.Polyndrome("di rumah saya ada kasur rusak")
	controller.Polyndrome("aku suka makan nasi")
	controller.Polyndrome("abcde edcbza")

	r.Run()

}
