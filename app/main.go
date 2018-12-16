package main

import (
	. "../app/database"
)

func main() {
	defer SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}
