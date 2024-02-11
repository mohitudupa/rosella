package handlers

import (
	"net/http"

	"githib.com/mohitudupa/rosella/utils"
)

const (
	HealthyUrl = "GET /healthy"
)

func Healthy(w http.ResponseWriter, req *http.Request) {
	utils.JsonResponse(w, []byte("true"), http.StatusOK)
}
