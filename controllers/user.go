package controllers

import (
	"ftm"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session)
{
	return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWritter, r *http.Request,
	p httprouter.Params)
{
	id := p.ByName("id")

	if !bsno.IsObjectHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectHex(id)
	u := models.User{}
	err := uc.Session.DB("keduback").C("users").FindIn(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		ftm.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	ftm.Println(w, "%s\n", uj)
}

func (uc UserController) CreateUser (w http.ResponseWritter, r *http.Request,
	_ httprouter.Params)
{
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("keduback").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(w, "%s\n", uj)
}


DeleteUser