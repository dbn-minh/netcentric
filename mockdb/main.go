package main

import (
	"mockdb/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
