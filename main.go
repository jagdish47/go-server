package main

import (
	"fmt"
	"log"
	"net/http"
)




func helloHandler(w http.ResponseWriter,r *http.Request){

	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
	}

	fmt.Fprint(w, "Hello!")

	
}


func formHandler(w http.ResponseWriter,r *http.Request){

	err:= r.ParseForm();

	if err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	if r.URL.Path != "/form"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}


	fmt.Fprintf(w, "POST request successfull")

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w,name)
	fmt.Fprintf(w,email)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting server a port 8080\n")

	err:= http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}