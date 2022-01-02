package main

import (
	"YamadaKESHE/model"
	"YamadaKESHE/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
