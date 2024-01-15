package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
	"github.com/jaycel19/capstone-api/util"
)

func LoginPersonnel(w http.ResponseWriter, r *http.Request) {
	var loginPayload services.Personnel
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	adminResp, err := models.Personnel.PersonnelLogin(loginPayload)
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

func GetAllPersonnel(w http.ResponseWriter, r *http.Request) {
	var personnel services.Personnel
	all, err := personnel.GetAllPersonnel()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"personnel": all})
}

func GetPersonnelById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err  := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	personnel, err := models.Personnel.GetPersonnelById(conv_id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, personnel)
}

func CreatePersonnel(w http.ResponseWriter, r *http.Request) {
	var personnelPayload services.Personnel
	err := json.NewDecoder(r.Body).Decode(&personnelPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	personnelCreated, err := models.Personnel.CreatePersonnel(personnelPayload)
	helpers.WriteJSON(w, http.StatusOK, personnelCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdatePersonnel(w http.ResponseWriter, r *http.Request) {
	var personnelPayload services.Personnel
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	err = json.NewDecoder(r.Body).Decode(&personnelPayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedPersonnel, err := models.Personnel.UpdatePersonnel(conv_id, personnelPayload)
	helpers.WriteJSON(w, http.StatusOK, updatedPersonnel)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeletePersonnel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {	
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	err = models.Personnel.DeletePersonnel(conv_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
