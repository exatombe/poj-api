package main

import (
	"net/http"
	httphelper "poj/src/http_helper"
)

func main() {
	router := httphelper.Router()
	http.ListenAndServe(":3200", router)
}
