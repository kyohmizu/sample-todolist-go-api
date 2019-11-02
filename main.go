package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todoList = []todo{
	{
		ID:   1,
		Text: "勉強",
		Done: false,
	},
	{
		ID:   2,
		Text: "掃除",
		Done: true,
	},
}

var maxID = 2

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newTodo)

	maxID++
	newTodo.ID = maxID
	todoList = append(todoList, newTodo)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTodo)
}

func getOneTodo(w http.ResponseWriter, r *http.Request) {
	id := getID(r)

	for _, t := range todoList {
		if t.ID == id {
			json.NewEncoder(w).Encode(t)
		}
	}
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todoList)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := getID(r)
	var updatedTodo todo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedTodo)

	for i, t := range todoList {
		if t.ID == id {
			todoList[i].Text = updatedTodo.Text
			todoList[i].Done = updatedTodo.Done
			json.NewEncoder(w).Encode(t)
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := getID(r)

	for i, t := range todoList {
		if t.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}
}

func getID(r *http.Request) int {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	return int(id)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", createTodo).Methods("POST")
	r.HandleFunc("/", getTodoList).Methods("GET")
	r.HandleFunc("/{id}", getOneTodo).Methods("GET")
	r.HandleFunc("/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/{id}", deleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9999", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(10),
		handlers.AllowCredentials(),
	)(r)))
}
