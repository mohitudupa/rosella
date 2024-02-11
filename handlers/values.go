package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"githib.com/mohitudupa/rosella/data"
	"githib.com/mohitudupa/rosella/utils"
)

type ValuesHandler struct {
	ListURL   string
	GetURL    string
	PostURL   string
	DeleteURL string

	repository data.Repository
}

func NewValuesHandler(r data.Repository) *ValuesHandler {
	return &ValuesHandler{
		ListURL:    "GET /value/{group}",
		GetURL:     "GET /value/{group}/{value}",
		PostURL:    "POST /value/{group}/{value}",
		DeleteURL:  "DELETE /value/{group}/{value}",
		repository: r}
}

func (lh *ValuesHandler) List(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	result, err := lh.repository.ListValues(group)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *ValuesHandler) Get(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	value := req.PathValue("value")
	result, err := lh.repository.GetValue(group, value)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *ValuesHandler) Post(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	value := req.PathValue("value")
	data := ""

	req.Body = http.MaxBytesReader(w, req.Body, 4096)
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

	err = lh.repository.SetValue(group, value, data)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusCreated)
}

func (lh *ValuesHandler) Delete(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	value := req.PathValue("value")
	err := lh.repository.DeleteValue(group, value)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusNoContent)
}
