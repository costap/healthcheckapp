package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/costap/healthcheckapp/model"
	"github.com/costap/healthcheckapp/repository"
	"github.com/julienschmidt/httprouter"
)

type AdminController struct {
	repo *repository.ServiceInfoRepository
}

func NewAdminController(repo *repository.ServiceInfoRepository) *AdminController {
	return &AdminController{repo: repo}
}

func (c *AdminController) AddService(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		HandleError(w, err)
		return
	}

	var service model.ServiceInfo

	if err := json.Unmarshal(reqBody, &service); err != nil {
		HandleError(w, err)
		return
	}

	c.repo.Save(service)
}

func (c *AdminController) ListServices(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(c.repo.List())
}
