package main

import "task-service/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
