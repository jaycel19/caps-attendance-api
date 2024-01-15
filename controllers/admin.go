package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
	"github.com/jaycel19/capstone-api/util"
)

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var loginPayload services.Admin
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	adminResp, err := models.Admin.AdminLogin(loginPayload)
	if err != nil {
		helpers.WriteJSON(w, http.StatusForbidden, helpers.Envelope{"Error": "Wrong username"})
		return
	}
	err = util.CheckPassword(loginPayload.Password, adminResp.Password)
	if err != nil {

		helpers.WriteJSON(w, http.StatusForbidden, helpers.Envelope{"Error": "Password not match"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"Message": "Logged in", "admin": adminResp})
}

func GetAllAdmin(w http.ResponseWriter, r *http.Request) {
	var admin services.Admin
	all, err := admin.GetAllAdmins()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"personnel": all})
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminPayload services.Admin
	err := json.NewDecoder(r.Body).Decode(&adminPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	adminCreated, err := models.Admin.CreateAdmin(adminPayload)
	helpers.WriteJSON(w, http.StatusOK, adminCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminPayload services.Admin
	username := chi.URLParam(r, "username")

	err := json.NewDecoder(r.Body).Decode(&adminPayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAdmin, err := models.Admin.UpdateAdmin(username, adminPayload)
	helpers.WriteJSON(w, http.StatusOK, updatedAdmin)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	err := models.Attendee.DeleteAdmin(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
