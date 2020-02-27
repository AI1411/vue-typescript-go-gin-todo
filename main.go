package main

import (
	"vue-typescript-go-gin/db"
	"vue-typescript-go-gin/router"
)

func main() {
	dbConn := db.Init()

	router.Router(dbConn)
}
