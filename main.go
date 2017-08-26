package main

import "github.com/mtdx/case-api/routes"

func main() {
	// TODO:connect
	//	db := db.Init()
	//	defer db.Close()

	router := routes.SetupRouter()
	router.Run()
}
