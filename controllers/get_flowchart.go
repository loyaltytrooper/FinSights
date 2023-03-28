package controllers

import (
	helpers "FinSights/helpers"
	"FinSights/models"
	"FinSights/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetFundFlow(c *gin.Context) {
	var r models.GetFlowRequest

	// validate the request
	jsonError := c.ShouldBindJSON(&r)
	if jsonError != nil {
		response := helpers.CreateResponse(nil, jsonError)
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	fileUrl := strings.Split(r.FileUrl, "/")
	fileName := fileUrl[len(fileUrl)-1]

	fmt.Println(fileUrl)
	fmt.Println(fileName)

	err := helpers.DownloadFile(fileName, r.FileUrl)
	if err != nil {
		return
	}

	data := services.GetFundTrail(fileName, r.Password)

	//if err != nil {
	//	response := responsehandler.CreateResponse(nil, err, message)
	//	c.JSON(http.StatusInternalServerError, response)
	//	c.Abort()
	//	return
	//}
	response := helpers.CreateResponse(data, nil)
	c.JSON(http.StatusOK, response)
	return
}
