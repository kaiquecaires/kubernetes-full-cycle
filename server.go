package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map", ConfigMap)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Write([]byte(fmt.Sprintf("Hello, my name is %s, and I am %s years old", name, age)))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}
	fmt.Fprintf(w, "My Family: %s", string(data))
}
