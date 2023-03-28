package main

import (
	"FinSights/router"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	r := router.RouterInit()
	err := r.Run(":" + PORT)
	if err != nil {
		panic("Error opening server on port")
		return
	}
}
