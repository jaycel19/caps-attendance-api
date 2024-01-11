package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jaycel19/capstone-api/helpers"
	"github.com/jaycel19/capstone-api/services"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var student services.Student
	all, err := student.GetAllStudents()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"students": all})
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err  := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	student, err := models.Student.GetStudentById(conv_id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, student)
}

func GetStudentsByCourse(w http.ResponseWriter, r *http.Request) {
	program := chi.URLParam(r, "course")
	students, err := models.Student.GetByCourses(program)
	if students == nil {
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
	}
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// TODO: handle error
	helpers.WriteJSON(w, http.StatusOK, students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var studentPayload services.Student
	err := json.NewDecoder(r.Body).Decode(&studentPayload)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	studentCreated, err := models.Student.CreateStudent(studentPayload)
	helpers.WriteJSON(w, http.StatusOK, studentCreated)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var studentPayload services.Student
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	err = json.NewDecoder(r.Body).Decode(&studentPayload)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedStudent, err := models.Student.UpdateStudent(conv_id, studentPayload)
	helpers.WriteJSON(w, http.StatusOK, updatedStudent)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	conv_id, err := uuid.Parse(id)
	if err != nil {	
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	err = models.Student.DeleteStudent(conv_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "Deleted!")
}
