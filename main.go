package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/costap/healthcheckapp/controllers"
	"github.com/costap/healthcheckapp/repository"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("Staring server listening on port 8080")

	repo := repository.NewServiceInfoRepository()

	hc := controllers.NewHealthController(tpl)
	ac := controllers.NewAdminController(repo)

	mux := httprouter.New()
	mux.GET("/", hc.Index)
	mux.GET("/info", hc.Info)
	mux.POST("/api/services", ac.AddService)
	mux.GET("/api/services", ac.ListServices)
	mux.ServeFiles("/assets/*filepath", http.Dir("./public"))

	http.ListenAndServe(":8080", mux)
}
