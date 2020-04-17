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

var GetMembersByAccountID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account_id, err := strconv.Atoi(params["account_id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data, err := models.GetMembersByAccountID(int(account_id))
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	unmarshalledMembers, err := json.Marshal(data)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	// members := make([]MemberWithID)
	// for _, member := range data {
	// 	unmarshalledMember, err := json.Marshal(*member)
	// 	if err != nil {
	// 		u.Respond(w, u.Message(false, err.Error()))
	// 		return
	// 	}
	// 	members = append(members, unmarshalledMember)
	// }

	resp := u.Message(true, "success")
	resp["data"] = string(unmarshalledMembers)
	u.Respond(w, resp)
}

var GetMemberByPhoneNumber = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	phone_number := params["phone_number"]

	data, err := models.GetMemberByPhoneNumber(phone_number)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	unmarshalledData, err := json.Marshal(*data)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error marshalling returned object"))
	}
	resp := u.Message(true, "success")
	resp["data"] = string(unmarshalledData)
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

	data, err := models.GetMemberByClientMemberID(int(client_member_id))
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	unmarshalledData, err := json.Marshal(*data)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error marshalling returned object"))
	}
	resp := u.Message(true, "success")
	resp["data"] = string(unmarshalledData)
	u.Respond(w, resp)
}

var GetMemberByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data, err := models.GetMemberByID(int(id))
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	unmarshalledData, err := json.Marshal(*data)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error marshalling returned object"))
	}
	resp := u.Message(true, "success")
	resp["data"] = string(unmarshalledData)
	u.Respond(w, resp)
}