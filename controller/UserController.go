package controller

import (
	"fmt"
	"lenslocked/M"
	"log"
	"net/http"
)

type UserController struct {
	Template struct {
		New   Template
		Login Template
	}
	US M.UserService
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	//TODO: get the userName and password from request
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	name := "yyChat"
	_, err := u.US.Create(M.CreateUserParms{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "create user error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("sucessful to sign in !!!"))
	//TODO: use the model to add user data
	//TODO: then base on the result to render some page or return error to the responseWriter
}

func (u UserController) Find(name string) {

}

func (u UserController) New(w http.ResponseWriter, r *http.Request) {
	u.Template.New.Execute(w, nil)
}
func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	u.Template.Login.Execute(w, nil)
}
func (u UserController) ProcessLogin(w http.ResponseWriter, r *http.Request) {
	parms := M.AuthenticateParms{}
	parms.Email = r.PostFormValue("email")
	parms.Password = r.PostFormValue("password")
	user, err := u.US.Authenticate(parms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "user %v authencate successful", user)
}
