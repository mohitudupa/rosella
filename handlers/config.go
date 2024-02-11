package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"githib.com/mohitudupa/rosella/data"
	"githib.com/mohitudupa/rosella/utils"
)

type ConfigsHandler struct {
	ListURL   string
	GetURL    string
	PostURL   string
	DeleteURL string

	repository data.Repository
}

func NewConfigsHandler(r data.Repository) *ConfigsHandler {
	return &ConfigsHandler{
		ListURL:    "GET /config/{group}",
		GetURL:     "GET /config/{group}/{config}",
		PostURL:    "POST /config/{group}/{config}",
		DeleteURL:  "DELETE /config/{group}/{config}",
		repository: r}
}

func (lh *ConfigsHandler) List(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	result, err := lh.repository.ListConfigs(group)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *ConfigsHandler) Get(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	config := req.PathValue("config")
	result, err := lh.repository.GetConfig(group, config)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (lh *ConfigsHandler) Post(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	config := req.PathValue("config")
	data := data.Config{}

	req.Body = http.MaxBytesReader(w, req.Body, 1048576)
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

	err = lh.repository.SetConfig(group, config, data)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusCreated)
}

func (lh *ConfigsHandler) Delete(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	config := req.PathValue("config")
	err := lh.repository.DeleteConfig(group, config)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusNoContent)
}
