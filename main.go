package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-member-api/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/members/new", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/api/members/id/{id}", controllers.GetMemberByID).Methods("GET")
	router.HandleFunc("/api/members/account_id/{account_id}", controllers.GetMembersByAccountID).Methods("GET")
	router.HandleFunc("/api/members/phone_number/{phone_number}", controllers.GetMemberByPhoneNumber).Methods("GET")
	router.HandleFunc("/api/members/client_member_id/{client_member_id}", controllers.GetMemberByClientMemberID).Methods("GET")

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("application running on port:" + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
