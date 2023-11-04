package main

import (
	"api/db"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Initialize your database thingy
	databaseThingy, err := db.NewDatabaseThingy()
	if err != nil {
		panic(err) // For simplicity, we just panic if we can't connect to the database.
	}

	// Set up Gin
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Define your routes and handlers
	// PostCard godoc
	// @Summary Create card
	// @Description Create card
	// @ID post-card
	// @Produce  json
	// @Success 200 CardRequest
	// @Router /card [post]
	r.POST("/card", func(c *gin.Context) {
		var cardRequest models.CardRequest
		if err := c.ShouldBindJSON(&cardRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		key, err := databaseThingy.Insert("Card", &cardRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cardRequest.ID = key.Name // Assume the ID is the key name
		c.JSON(http.StatusOK, cardRequest)
	})

	// GetCards godoc
	// @Summary Retrieve cards
	// @Description get cards
	// @ID get-cards
	// @Produce  json
	// @Success 200 {array} CardRequest
	// @Router /cards [get]
	r.GET("/cards", func(c *gin.Context) {
		var cards []models.CardRequest
		err := databaseThingy.Select("Card", &cards)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cards)
	})

	// Start the server
	r.Run() // By default, it listens on :8080
}
