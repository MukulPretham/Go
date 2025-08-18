package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReqBody struct {
	Message string
}

type album struct {
	ID     uint  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	var dsn = "host=localhost user=postgres password=9059015626 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	db.AutoMigrate(&album{})

	router := gin.Default()

	router.GET("/album", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, albums)
	})

	router.POST("/album", func(ctx *gin.Context) {
		var body album
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(400, gin.H{"errro": "error while reading body"})
		}
		db.Create(&body)
		fmt.Print(albums)
		ctx.JSON(200, gin.H{"message": "album added"})
	})

	router.GET("/album/:name", func(ctx *gin.Context) {
		param := ctx.Param("id")
		for _, val := range albums {
			if val.Artist == param {
				ctx.IndentedJSON(200, val)
				return
			}
		}
		ctx.IndentedJSON(404, gin.H{"error": "album not found"})
	})

	router.Run(":3000")
}
