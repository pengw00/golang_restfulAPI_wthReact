package controllers

import (
	"net/http"
	u "goapi/utils"
	"goapi/models"
	"encoding/json"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request){
	account := &models.Account{}
	
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error accor
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
	}
	resp := account.Create() // create account
	u.Respond(w, resp)
}
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
		account := &models.Account{}
		err := json.NewDecoder(r.Body).Decode(account)
		if err != nil {
			u.Respond(w, u.Message(false, "Invalid request"))
			return
		}

		resp := models.Login(account.Email, account.Password)
		u.Respond(w, resp)
}