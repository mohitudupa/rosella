package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"githib.com/mohitudupa/rosella/data"
	"githib.com/mohitudupa/rosella/utils"
)

type FlagsHandler struct {
	ListURL   string
	GetURL    string
	PostURL   string
	DeleteURL string

	repository data.Repository
}

func NewFlagsHandler(r data.Repository) *FlagsHandler {
	return &FlagsHandler{
		ListURL:    "GET /flag/{group}",
		GetURL:     "GET /flag/{group}/{flag}",
		PostURL:    "POST /flag/{group}/{flag}",
		DeleteURL:  "DELETE /flag/{group}/{flag}",
		repository: r,
	}
}

func (fh *FlagsHandler) List(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	result, err := fh.repository.ListFlags(group)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (fh *FlagsHandler) Get(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	flag := req.PathValue("flag")
	result, err := fh.repository.GetFlag(group, flag)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(result)
	utils.JsonResponse(w, jsonResponse, http.StatusOK)
}

func (fh *FlagsHandler) Post(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	flag := req.PathValue("flag")
	data := false

	req.Body = http.MaxBytesReader(w, req.Body, 8)
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

	err = fh.repository.SetFlag(group, flag, data)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusCreated)
}

func (fh *FlagsHandler) Delete(w http.ResponseWriter, req *http.Request) {
	group := req.PathValue("group")
	flag := req.PathValue("flag")
	err := fh.repository.DeleteFlag(group, flag)
	if err != nil {
		utils.ErrorJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, nil, http.StatusNoContent)
}
