package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber =":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum :=addValues(2,2)
	_ , _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page ans sum of 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	//var sum int
	sum := x + y
	return sum
}

func Divide(w http.ResponseWriter, r *http.Request){
	f, err := divideValues(100.0,0)
	if err!=nil{
		fmt.Fprintf(w, "Can not divide by 0")
		return 
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if(y<=0){
		err:=errors.New("can not divide by Zero")
		return 0, err
	}
	result:= x / y
	return result, nil
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
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting apps on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}