package helpers

import (
	"CYAMod/models"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type MyHandler struct{}


var(
	stories models.Story
	tmpl *template.Template
	err error
)

func init(){
	stories = MapJSONToStories()
	tmpl,err = template.ParseFiles("pages/layout.html")
	if err != nil {
		log.Fatal("Parse Files: ",err)
	}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint:= strings.Trim(r.URL.Path,"/")
	tmpl.Execute(w,stories[endpoint])
}