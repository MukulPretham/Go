package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	router := gin.Default()

	//
	db_string := "host=localhost user=postgres password=9059015626 dbname=postgres port=5432"
	db,err := gorm.Open(postgres.Open(db_string))
	if err!=nil{
		log.Fatal("server was not able to connect to db")
	}
	db.AutoMigrate(&Book{})

	router.POST("/book",func(ctx *gin.Context) {
		var currBook Book
		if err := ctx.BindJSON(&currBook); err!=nil{
			ctx.IndentedJSON(http.StatusBadRequest,gin.H{"error":"invalid request"})
			return
		}
		result := db.Create(&currBook)
		if result.Error != nil{
			ctx.IndentedJSON(http.StatusBadGateway,gin.H{"error":"db error"})
			return
		}
		ctx.JSON(200,gin.H{ "message" : "entry created"})
	})

	router.GET("/books",func(ctx *gin.Context) {
		var allBooks []Book
		result:= db.Find(&allBooks)
		if result.Error != nil{
			ctx.IndentedJSON(http.StatusBadGateway,gin.H{"error":"db error"})
			return
		}
		ctx.IndentedJSON(http.StatusOK,allBooks)
	})

	router.GET("/book/:author",func(ctx *gin.Context) {
		authorName := ctx.Param("author")

		var currBooks []Book
		result:= db.Where("author = ?",authorName).Find(&currBooks)
		if result.Error != nil{
			ctx.IndentedJSON(http.StatusBadGateway,gin.H{"error":"db error"})
			return
		}
		ctx.IndentedJSON(http.StatusOK,currBooks)
	})

	router.POST("/update",func(ctx *gin.Context) {
		var updateReq UpdateReq
		err := ctx.BindQuery(&updateReq)
		fmt.Print(updateReq)
		if err !=nil{
			ctx.IndentedJSON(http.StatusBadRequest,gin.H{"error":"bad request"})
			return
		}
		result := db.Model(&Book{}).Where("author = ?",updateReq.Value).Update("author","removed")
		
		if result.Error != nil{	
			ctx.IndentedJSON(http.StatusBadRequest,gin.H{"error":"bad query"})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted,gin.H{"message":"done"})
	})

	router.Run(":3000")
}