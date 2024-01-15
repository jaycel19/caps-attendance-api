package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
)

func GetAllAttendance(w http.ResponseWriter, r *http.Request) {
	var attendance services.Attendance
	all, err := attendance.GetAllAttendance()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"attendances": all})
}

func GetAttendanceById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err  := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	attendance, err := models.Attendance.GetAttendanceById(conv_id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendance)
}

func GetAttendanceByCourse(w http.ResponseWriter, r *http.Request) {
	program := chi.URLParam(r, "course")
	attendee, err := models.Attendee.GetByCourses(program)
	if attendee == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendee)
}

func GetAttendanceByEventId(w http.ResponseWriter, r *http.Request) {
	eventId := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(eventId)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	attendee, err := models.Attendance.GetByEventId(conv_id)
	if attendee == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendee)
}

func GetAttendanceByAttendeeId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	
	attendance, err := models.Attendance.GetByAttendeeID(conv_id)
	if attendance == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendance)
}

func GetAttendanceByRange(w http.ResponseWriter, r *http.Request) {
	timeStart := chi.URLParam(r, "timestart")
	timeEnd := chi.URLParam(r, "timeend")
	
	layout := "2006-01-02 15:04:05"
	
	parsedStart, err := time.Parse(layout, timeStart)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	parsedEnd, err := time.Parse(layout, timeEnd)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	attendance, err := models.Attendance.GetByCreatedAtRange(parsedStart, parsedEnd)
	if attendance == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, attendance)
}

func CreateAttendance(w http.ResponseWriter, r *http.Request) {
	var attendancePayload services.Attendance
	err := json.NewDecoder(r.Body).Decode(&attendancePayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	attendanceCreated, err := models.Attendance.CreateAttendance(attendancePayload)
	helpers.WriteJSON(w, http.StatusOK, attendanceCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	var attendancePayload services.Attendance
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	err = json.NewDecoder(r.Body).Decode(&attendancePayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAttendance, err := models.Attendance.UpdateAttendance(conv_id, attendancePayload)
	helpers.WriteJSON(w, http.StatusOK, updatedAttendance)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {	
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	err = models.Attendance.DeleteAttendance(conv_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
