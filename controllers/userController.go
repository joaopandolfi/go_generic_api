package controllers

import (
	"net/http"
	"strconv"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
	"../models"
)

// UserController -
type UserController struct {
}

// NewClientUser - This endpoint create a new user, based on institution
// @rest
func (cc UserController) NewClientUser(w http.ResponseWriter, r *http.Request) {
	userService := NewUserService()
	form, _ := handlers.GetForm(r)

	instInt, _ := strconv.Atoi(form.Get("institution"))
	//cpf, _ := strconv.Atoi(utils.OnlyNumbers(form.Get("cpf")))
	cpf := utils.OnlyNumbers(form.Get("cpf"))
	user := models.User{
		People: models.People{
			Name: form.Get("name"),
			CPF:  cpf,
		},
		Email:     "",
		Username:  form.Get("username"),
		Picture:   "",
		Password:  form.Get("password"),
		Instution: instInt,
	}

	result, err := userService.NewUserClient(user)
	if err != nil {
		utils.Debug("Error on create new user",err.Error())
		handlers.RESTResponseError(w, "Error on create new user")
	} else {
		handlers.RESTResponse(w, result)
	}
}

// SetEspecialty Set in session the speciality is in the question
// @rest
// Store in session with çabel `specialty`
func (cc UserController) SetEspecialty(w http.ResponseWriter, r *http.Request) {
	sess, _ := handlers.GetSession(r)
	vars := handlers.GetVars(r)

	specialty, _ := strconv.Atoi(vars["specialty"])

	if specialty > 0 {
		sess.Values[models.SESSION_VALUE_SPECIALTY] = specialty
		err := sess.Save(r, w)
		if err == nil {
			handlers.RESTResponse(w, true)
			return
		}
	}

	handlers.RESTResponseError(w, "Invalid Specialty")
}
