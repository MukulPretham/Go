package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ReqBody struct{
	Message string
}

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/album",func(ctx *gin.Context) {
		ctx.IndentedJSON(200,albums)
	})

	router.POST("/album",func(ctx *gin.Context) {
		var body album
		if err:= ctx.BindJSON(&body);err != nil{
			ctx.JSON(400,gin.H{"errro":"error while reading body"})
		}
		albums = append(albums, body)
		fmt.Print(albums)
		ctx.JSON(200,gin.H{"message":"album added"})
	})

	router.GET("/album/:id",func(ctx *gin.Context) {
		param := ctx.Param("id")
		for _,val := range albums{
			if val.ID == param{
				ctx.IndentedJSON(200,val)
				return
			}
		}
		ctx.IndentedJSON(404,gin.H{"error":"album not found"})
	})

	router.Run(":3000")
}