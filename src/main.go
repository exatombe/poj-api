package main

import (
	"net/http"
	"poj/src/database"
	httphelper "poj/src/http_helper"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = database.Init()
}

func main() {
	router := httphelper.Router()
	http.ListenAndServe(":3200", router)
}
