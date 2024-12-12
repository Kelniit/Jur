package controller

import (
	"github.com/gin-gonic/gin"
)

func MainHallo(c *gin.Context) {
	// Main Page on Sample System
	c.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Hallo, Get All Sample !"})
}

func GetSample(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Sample ID"})
}
