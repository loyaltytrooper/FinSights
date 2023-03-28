package controllers

import (
	helpers "FinSights/helpers"
	"FinSights/models"
	"FinSights/services"
	"github.com/gin-gonic/gin"
	"net/http"
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

	data := services.GetFundTrail(r.FileUrl, r.Password)

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
