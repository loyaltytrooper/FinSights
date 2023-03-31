package controllers

import (
	helpers "FinSights/helpers"
	"FinSights/models"
	"FinSights/services"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
		response := helpers.CreateResponse(nil, errors.New("Error downloading file"))
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var data models.TransactionsJSON
	if fileName[len(fileName)-3:] == "csv" {
		data = services.ParseCSVFile(fileName)
		fmt.Println(data)
		client := &http.Client{
			Timeout: time.Second * 300,
		}

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormFile("file", fileName)
		if err != nil {
			fmt.Println(err.Error() + "from Get_flowchart")
		}
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error() + "from Get_flowchart")
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			fmt.Println(err.Error() + "from Get_flowchart")
		}
		// Close multipart writer.
		writer.Close()
		req, err := http.NewRequest("POST", "https://kv-py.onrender.com", bytes.NewReader(body.Bytes()))
		if err != nil {
			fmt.Println(err.Error() + "from Get_flowchart")
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rsp, _ := client.Do(req)
		if rsp.StatusCode != http.StatusOK {
			log.Printf("Request failed with response code: %d", rsp.StatusCode)
		} else {
			defer rsp.Body.Close()

			var cResp models.ML_Response

			if err := json.NewDecoder(rsp.Body).Decode(&cResp); err != nil {
				log.Fatal("ooopsss! an error occurred, please try again")
			}
			fmt.Println(cResp.Data)

			for _, v := range data.MetaData.Children {
				for _, val := range cResp.Data {
					if v.TxnId == strconv.Itoa(val.TxnId) {
						v.MetaData.Suspicious = true
						v.MetaData.Score = val.Score
					}
				}
			}
		}
	} else {
		data = services.GetFundTrail(fileName, r.Password)

		err = os.RemoveAll("act_" + fileName[0:len(fileName)-4])
		if err != nil {
			response := helpers.CreateResponse(nil, errors.New("Error removing directory"))
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		err = os.Remove(fileName[0:(len(fileName)-4)] + "transactions.csv")
		if err != nil {
			response := helpers.CreateResponse(nil, errors.New("Error removing csv"))
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	err = os.Remove(fileName)
	if err != nil {
		response := helpers.CreateResponse(nil, errors.New("Error removing "))
		c.JSON(http.StatusInternalServerError, response)
		return
	}

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
