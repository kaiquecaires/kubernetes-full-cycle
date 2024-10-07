package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/config-map", ConfigMap)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	secret := os.Getenv("SECRET")
	w.Write([]byte(fmt.Sprintf("Hello, my name is %s, and I am %s years old, this is my secret: %s", name, age, secret)))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}
	fmt.Fprintf(w, "My Family: %s", string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 || duration.Seconds() > 60 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("duration: %v", duration)))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
