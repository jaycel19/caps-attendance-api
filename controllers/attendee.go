package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
)

func GetAllAttendees(w http.ResponseWriter, r *http.Request) {
	var attendee services.Attendee
	all, err := attendee.GetAllAttendees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"attendees": all})
}

func GetAttendeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err  := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	attendee, err := models.Attendee.GetAttendeeById(conv_id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendee)
}

func GetAttendeesByCourse(w http.ResponseWriter, r *http.Request) {
	program := chi.URLParam(r, "course")
	attendees, err := models.Attendee.GetByCourses(program)
	if attendees == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendees)
}

func CreateAttendee(w http.ResponseWriter, r *http.Request) {
	var attendeePayload services.Attendee
	err := json.NewDecoder(r.Body).Decode(&attendeePayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	attendeeCreated, err := models.Attendee.CreateAttendee(attendeePayload)
	helpers.WriteJSON(w, http.StatusOK, attendeeCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdateAttendee(w http.ResponseWriter, r *http.Request) {
	var attendeePayload services.Attendee
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	err = json.NewDecoder(r.Body).Decode(&attendeePayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAttendee, err := models.Attendee.UpdateAttendee(conv_id, attendeePayload)
	helpers.WriteJSON(w, http.StatusOK, updatedAttendee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeleteAttendee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {	
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	err = models.Attendee.DeleteAttendee(conv_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
