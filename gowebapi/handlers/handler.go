package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Temperature struct {
	ID    string  `json:"id"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

var temperatures = []Temperature{
	{ID: "1", Value: 22.5, Unit: "Celsius"},
	{ID: "2", Value: 72.5, Unit: "Fahrenheit"},
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetTemperatures(c *gin.Context) {
	c.JSON(http.StatusOK, temperatures)
}

func CreateTemperature(c *gin.Context) {
	var newTemp Temperature
	if err := c.BindJSON(&newTemp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	temperatures = append(temperatures, newTemp)
	c.JSON(http.StatusCreated, newTemp)
}

func GetTemperatureByID(c *gin.Context) {
	id := c.Param("id")
	for _, temp := range temperatures {
		if temp.ID == id {
			c.JSON(http.StatusOK, temp)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Temperature not found"})
}

func UpdateTemperature(c *gin.Context) {
	id := c.Param("id")
	var updatedTemp Temperature
	if err := c.BindJSON(&updatedTemp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, temp := range temperatures {
		if temp.ID == id {
			temperatures[i] = updatedTemp
			c.JSON(http.StatusOK, updatedTemp)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Temperature not found"})
}

func DeleteTemperature(c *gin.Context) {
	id := c.Param("id")
	for i, temp := range temperatures {
		if temp.ID == id {
			temperatures = append(temperatures[:i], temperatures[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"message": "Temperature deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Temperature not found"})
}
