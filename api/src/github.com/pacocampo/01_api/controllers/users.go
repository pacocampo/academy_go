package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pacocampo/01_api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) List(w http.ResponseWriter, r *http.Request) {
	var uList []models.User
	if err := uc.session.DB("db-api").C("users").Find(nil).All(&uList); err != nil {
		returnNotFound(w, err)
		return
	}

	data, _ := json.Marshal(uList)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(data)

}

func (uc UserController) Get(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		returnNotFound(w, nil)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("db-api").C("users").FindId(oid).One(&u); err != nil {
		returnNotFound(w, err)
		return
	}

	data, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(data)
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return
	}

	u.Id = bson.NewObjectId()

	uc.session.DB("db-api").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	w.Write(uj)
}

func (uc UserController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.ObjectIdHex(params["id"])

	if err := uc.session.DB("db-api").C("users").UpdateId(u.Id, u); err != nil {
		fmt.Println("damn")
		returnNotFound(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("db-api").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}

func returnNotFound(w http.ResponseWriter, err error) {
	log.Fatal(err)
	w.WriteHeader(404)
}
