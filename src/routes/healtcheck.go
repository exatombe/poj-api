package routes

import (
	"net/http"
	utils "poj/src/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, map[string]string{"status": "ok"})
}
