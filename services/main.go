package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

func New (dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}

type Models struct {
	Attendee Attendee
	Admin Admin
	Event Event
	Attendance Attendance
	Personnel Personnel
	JsonResponse JsonResponse
}
