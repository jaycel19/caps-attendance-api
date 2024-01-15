package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	var event services.Event
	all, err := event.GetAllEvents()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"events": all})
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err  := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	event, err := models.Event.GetEventById(conv_id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var eventPayload services.Event
	err := json.NewDecoder(r.Body).Decode(&eventPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	eventCreated, err := models.Event.CreateEvent(eventPayload)
	helpers.WriteJSON(w, http.StatusOK, eventCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var eventPayload services.Event
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	err = json.NewDecoder(r.Body).Decode(&eventPayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedEvent, err := models.Event.UpdateEvent(conv_id, eventPayload)
	helpers.WriteJSON(w, http.StatusOK, updatedEvent)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {	
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	err = models.Event.DeleteEvent(conv_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
