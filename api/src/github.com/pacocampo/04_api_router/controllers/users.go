package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/pacocampo/04_api_router/db"
	"github.com/pacocampo/04_api_router/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) List(w http.ResponseWriter, r *http.Request) {
	var uList []models.User

	if err := db.GetAll(uc.session, "users", &uList); err != nil {
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

	if err := db.Get(uc.session, "users", id, &u); err != nil {
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

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	db.Create(uc.session, "users", u)

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

	db.Update(uc.session, "users", u.Id, u)

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

func (uc UserController) TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	fmt.Println("decoded:", decoded)

	var user models.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

func returnNotFound(w http.ResponseWriter, err error) {
	log.Fatal(err)
	w.WriteHeader(404)
}
