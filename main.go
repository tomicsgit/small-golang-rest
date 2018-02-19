package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)
var people []ManInterface

func main() {
	router := mux.NewRouter()

	people = append(people, Student{Human{1, "Faris Dzibric", 25}, "Hochschule Darmstadt", 1234.5})
	people = append(people, Employee{Human{2,"Marko Milojevic", 33}, "AOE GmbH", 4567.8})
	people = append(people, Employee{Human{3, "Milos Krsmanovic", 29}, "AOE GmbH", 4567.8})

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func GetPeople (w http.ResponseWriter, r *http.Request) {
	if p, err := json.Marshal(people); err == nil {
		w.Write(p)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if strconv.Itoa(item.GetId()) == params["id"] {
			if p, err := json.Marshal(item); err == nil {
				w.Write(p)
			}
		}
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person ManInterface
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.SetId(strconv.Atoi(params["id"]))
	people = append(people, person)
	if p, err := json.Marshal(people); err == nil {
		w.Write(p)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if strconv.Itoa(item.GetId()) == params["id"] {
			people = people[:index+copy(people[index:], people[index+1:])]
			break
		}
		if p, err := json.Marshal(people); err == nil {
			w.Write(p)
		}
	}
}