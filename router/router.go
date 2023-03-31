package router

import (
	"FinSights/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterInit() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(gin.CustomRecovery(ErrorHandler))

	publicRoutes := router.Group("/api")
	{
		trailAnalysis := publicRoutes.Group("/trail_analysis")
		{
			v1 := trailAnalysis.Group("/v1")
			{
				v1.POST("/fund_flow", controllers.GetFundFlow)
			}
		}
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ErrorHandler(c *gin.Context, err any) {
	c.Next()

	fmt.Println(err)

	for _, err := range c.Errors {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   true,
		"message": "Internal Server Error",
	})
}
