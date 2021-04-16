package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Project string
	Region  string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	project := os.Getenv("PROJECT")
	region := os.Getenv("REGION")

	dat := Data{
		Project: project,
		Region:  region,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, dat)
	})

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Cubeta example image on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
