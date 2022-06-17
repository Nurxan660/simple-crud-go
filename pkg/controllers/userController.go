package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testGo/pkg/models"
	"testGo/pkg/repository"

	"github.com/gorilla/mux"
)

type response struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("error when decode body", err)
	}

	userId := repository.InsertUser(user)

	res := response{userId, "Successfully created"}

	json.NewEncoder(w).Encode(res)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert the string into int.  %v", err)
	}

	user, err := repository.GetUser(int64(id))

	if err != nil {
		log.Fatalf("Cannot get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert the string into int.  %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Cannot decode the request body.  %v", err)
	}

	updatedRows := repository.UpdateUser(int64(id), user)

	msg := fmt.Sprintf(" Total rows/record affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert the string into int.  %v", err)
	}

	deletedRows := repository.DeleteUser(int64(id))

	msg := fmt.Sprintf("Total rows/record affected %v", deletedRows)
	res := response{
		Id:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
