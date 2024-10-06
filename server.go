package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
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
