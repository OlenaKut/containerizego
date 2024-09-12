package main

import (
	"net/http"
	"strconv"

	"systementor.se/containerizego/data"
	"systementor.se/containerizego/dto"

	"github.com/gin-gonic/gin"
)

func apiPlayer(c *gin.Context) {
	players := data.FindAll()
	var playerDTOArray [5]dto.PlayerDTO

	for index, item := range players {
		playerDTOArray[index].Id = strconv.Itoa(item.Id)
		playerDTOArray[index].PlayerName = item.PlayerName
		playerDTOArray[index].TeamName = item.TeamName
		playerDTOArray[index].Jersey = item.JerseyNumber
		playerDTOArray[index].BornYear = item.Born
	}
	c.IndentedJSON(http.StatusOK, playerDTOArray)
}

func main() {
	router := gin.Default()
	router.GET("/api/player", apiPlayer)

	router.Run(":8080")
}
