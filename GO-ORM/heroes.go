package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.mongodb.org/mongo-driver/bson"
)

func AllHeroes(w http.ResponseWriter, r *http.Request) {
	c = Connections()
	var heroes []*Hero
	heroes = ReturnAllHeroes(c, bson.M{})
	json.NewEncoder(w).Encode(heroes)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New users Endpoint Hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	c = Connections()

	fmt.Fprint(w, "Delete Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
