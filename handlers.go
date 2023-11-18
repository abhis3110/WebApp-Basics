package main 

import 
(
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "This is the home page")
	renderTemplate(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"about.page.tmpl")
}
