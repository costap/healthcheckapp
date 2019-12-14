package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/costap/healthcheckapp/model"

	"github.com/julienschmidt/httprouter"
)

type HealthController struct {
	tpl *template.Template
}

func NewHealthController(tpl *template.Template) *HealthController {
	return &HealthController{tpl: tpl}
}

func (c *HealthController) Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := c.tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func (c *HealthController) Info(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	bs, err := json.Marshal(model.ServiceInfo{Name: "HealthCheck Service", Version: "0.0.1"})
	fmt.Fprintf(w, string(bs))
	HandleError(w, err)
}
