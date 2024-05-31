package main

import "heilgar/finance-tracker/config"

func main() {
	r := config.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
