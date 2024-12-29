package main

import (
	"github.com/mwrks/ticket-app-transactions/initializers"
	"github.com/mwrks/ticket-app-transactions/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}
func main() {
	r := routes.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
