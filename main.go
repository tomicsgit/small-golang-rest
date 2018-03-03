package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
)
var people []ManInterface

func main() {
	router := mux.NewRouter()

	faris := Student{Human{1, "Faris Dzibric", 25}, "Hochschule Darmstadt", 1234.5}
	faris.BorrowMoney(450.56)

	marko := Employee{Human{2,"Marko Milojevic", 33}, "AOE GmbH", 4567.8}
	marko.SpendSalary(451)

	milos := Employee{Human{3, "Milos Krsmanovic", 29}, "AOE GmbH", 4567.8}
	milos.SpendSalary(405)

	people = append(people, faris)
	people = append(people, marko)
	people = append(people, milos)

	router.HandleFunc("/", PrintWelcomeMessage).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func PrintWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	var intro = []string{
		"Simple REST API example in Golang",
		"by Dragan Tomic",
		"Thank you for checking my Golang rest API example!",
	}
	message, _ := json.Marshal(intro)
	w.Write(message)
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//params := mux.Vars(r)
	code := 422
	var person Employee

	b1 := []byte(`{"Id":5, "Name":"Dragan Tomic", "Age":33, "Company":"AOE GmbH", "Salary":4322.3}`)
	//body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b1, &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println("========================1")
	fmt.Printf("%+v\n", person)
	fmt.Println("========================2")


	var person1 Person

	b2 := []byte(`{"Emp": {"Id":5, "Name":"Dragan Tomic", "Age":33, "Company":"AOE GmbH", "Salary":4322.3}}`)
	if err := json.Unmarshal(b2, &person1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", person1)
	fmt.Println("========================3")
	/*
		var incomingMessage1 Human
		b2 := []byte(`{"Id":5,"Name":"Dragan Tomic","Age":33}`)
		if err := json.Unmarshal(b2, &incomingMessage1); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", incomingMessage1)
		fmt.Printf("\tcommand: %s\n", incomingMessage1.Name)
		fmt.Println("========================")
	*/
	w.WriteHeader(code)
	return
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if strconv.Itoa(item.GetId()) == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	if p, err := json.Marshal(people); err == nil {
		w.Write(p)
	}
}