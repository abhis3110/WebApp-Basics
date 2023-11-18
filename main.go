package main

import (
	//"errors"
	"fmt"
	"net/http"
	"html/template"
)

const portNumber =":8080"

func Home(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "This is the home page")
	renderTemplate(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
	
}

func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _:=template.ParseFiles("./templates/"+ tmpl)
	err:=parsedTemplate.Execute(w, nil)
	if err!=nil{
		fmt.Println("error parsing template", err)
		return
	}
}






func main() {
	//fmt.Println("Hello world !")
	http.HandleFunc("/default", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello world")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of Bytes written: %d", n))
	})

	http.HandleFunc("/",Home)
	http.HandleFunc("/about", About)
	

	fmt.Println(fmt.Sprintf("Starting apps on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}