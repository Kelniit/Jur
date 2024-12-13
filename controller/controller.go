package controller

import (
	"Jur/config"
	"Jur/entities"
	"Jur/utilities"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func MainHallo(c *gin.Context) {
	// Main Page on Sample System
	c.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAll(c *gin.Context) {
	// Get All Sample Table
	database, errtable := config.GalaSetup()

	if errtable != nil {
		utilities.FailMess(c, 500, "Fail ! Connection Unavailable !")
		return
	}

	var samplesresult []entities.SampleTabler

	result := database.Find(&samplesresult)

	if result.Error != nil {
		utilities.FailMess(c, 500, "Fail to Get Sample Data", result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		utilities.FailMess(c, 404, "Data Unavailable !")
		return
	}

	c.JSON(200, samplesresult)
}

func GetSample(c *gin.Context) {
	SampleID := c.Param("SampleID")
	result := fmt.Sprintf("Sample %s is Available !", SampleID)
	c.JSON(200, gin.H{"result": result})
}

func CreateSample(c *gin.Context) {
	// Create Sample Data
	var samples []entities.SampleTabler

	if json_error := c.ShouldBindJSON(&samples); json_error != nil {
		utilities.FailMess(c, 400, "Fail to Bind JSON", json_error.Error())
		return
	}

	database, err := config.GalaSetup()

	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	// if create_err := database.Create(&sample).Error; create_err != nil {
	// 	utilities.FailMess(c, 500, "Fail to Create Sample !")
	// 	return
	// }

	if create_err := database.CreateInBatches(samples, len(samples)).Error; create_err != nil {
		utilities.FailMess(c, 500, "Fail to Create Multiple Sample !")
	}

	c.JSON(200, gin.H{"result": "Create Sample ID !"})
}

func DeleteSample(c *gin.Context) {
	SampleID := c.Param("SampleID")
	// database.Delete(&, SampleID)
	result := fmt.Sprintf("Sample %s is Deleted !", SampleID)
	c.JSON(200, gin.H{"result": result})
}
