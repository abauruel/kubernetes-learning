package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)

	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello %s, I have %s years old", name, age)
}
func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}
func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv(("USER"))
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "My USER: %s. PASSWORD: %s", user, password)
}
