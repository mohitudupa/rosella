package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"githib.com/mohitudupa/rosella/data"
	"githib.com/mohitudupa/rosella/utils"
)

type LimitsHandler struct {
	ListURL   string
	GetURL    string
	PostURL   string
	DeleteURL string

	repository data.Repository
}

func NewLimitsHandler(r data.Repository) *LimitsHandler {
	return &LimitsHandler{
		ListURL:    "GET /limit/{group}",
		GetURL:     "GET /limit/{group}/{limit}",
		PostURL:    "POST /limit/{group}/{limit}",
		DeleteURL:  "DELETE /limit/{group}/{limit}",
		repository: r}
}

func (lh *LimitsHandler) List(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	result, err := lh.repository.ListLimits(group)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *LimitsHandler) Get(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	limit := req.PathValue("limit")
	result, err := lh.repository.GetLimit(group, limit)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *LimitsHandler) Post(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	limit := req.PathValue("limit")
	data := float64(0)

	req.Body = http.MaxBytesReader(w, req.Body, 256)
	responseBody, err := io.ReadAll(req.Body)
	if err != nil {
		utils.ErrorJsonResponse(w, "could not read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		utils.ErrorJsonResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = lh.repository.SetLimit(group, limit, data)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusCreated)
}

func (lh *LimitsHandler) Delete(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	limit := req.PathValue("limit")
	err := lh.repository.DeleteLimit(group, limit)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusNoContent)
}
