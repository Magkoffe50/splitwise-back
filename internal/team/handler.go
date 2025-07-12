package team

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTeamInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserIDs     []uint `json:"user_ids"`
}

func CreateTeamHandler(c *gin.Context) {
	var input CreateTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := CreateTeam(input.Name, input.Description, input.UserIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func GetTeamHandler(c *gin.Context) {
	id := c.Param("id")
	team, err := GetTeamByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}
