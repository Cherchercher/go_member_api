package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-member-api/models"
	u "go-member-api/utils"
	"net/http"
	"strconv"
)

type createBody struct {
	Name string
}

var CreateMember = func(w http.ResponseWriter, r *http.Request) {

	var bd createBody
	if r.Body == nil {
		u.Respond(w, u.Message(false, "Please send a request body"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&bd); err != nil {
		u.Respond(w, u.Message(false, "Error while reading request body"))
	}

	err := models.BatchCreateMethodII(bd.Name)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, u.Message(true, "success"))
}

var GetMemberByAccountID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account_id, err := strconv.Atoi(params["account_id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetMemberByAccountID(int(account_id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetMemberByPhoneNumber = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phone_number := params["phone_number"]

	data := models.GetMemberByPhoneNumber(phone_number)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetMemberByClientMemberID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client_member_id, err := strconv.Atoi(params["client_member_id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetMemberByClientMemberID(int(client_member_id))
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error marshalling returned object"))
	}
	resp := u.Message(true, "success")
	resp["data"] = string(b)
	u.Respond(w, resp)
}
