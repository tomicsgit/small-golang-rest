package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
	"io/ioutil"
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := 422
	var dat Employee

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	w.WriteHeader(code)
	return
/*	params := mux.Vars(r)

	var person ManInterface
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Println(json.Unmarshal(body, &person))
	}

//	_ = json.NewDecoder(r.Body).Decode(&person)
fmt.Println(person)
fmt.Println(json.Marshal(people))

	person.SetId(strconv.Atoi(params["id"]))
	people = append(people, person)
	if p, err := json.Marshal(people); err == nil {
		w.Write(p)
	}
*/
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