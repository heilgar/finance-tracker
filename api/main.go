package main

import (
	"heilgar/finance-tracker/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
