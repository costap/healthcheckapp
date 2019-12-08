package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/costap/healthcheckapp/controllers"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("Staring server listening on port 8080")

	c := controllers.NewHealthController(tpl)

	mux := httprouter.New()
	mux.GET("/", c.Index)
	mux.GET("/info", c.Info)

	http.ListenAndServe(":8080", mux)
}
