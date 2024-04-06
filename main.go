package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"errors"

	"example/kedubak-yanisdolivet/controllers"
)

func getSession() *mgo.Session
{
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}

func main()
{
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/me", uc.GetUser)
	r.POST("/user/edit", uc.CreateUser)
	r.DELETE("/user/remove", uc.DeleteUser)
	http.ListenAndServe("localhost: 8080")
}
