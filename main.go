package main

import (
	"FinSights/router"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	r := router.RouterInit()
	err := r.Run(":" + PORT)
	if err != nil {
		panic("Error opening server on port")
		return
	}
}
