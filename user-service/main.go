package main

import "user-service/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8081")
}
