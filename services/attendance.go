package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID uuid.UUID `json:"id"`
	Attendee uuid.UUID `json:"attendee"`
	Event uuid.UUID `json:"event"`
	TimeIn time.Time `json:"time_in"`
	ScannedBy uuid.UUID `json:"scanned_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Attendance) GetAllAttendance() ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from attendance`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var attendances []*Attendance
	for rows.Next() {
		var attendance Attendance 
		err := rows.Scan(
			&attendance.ID,
			&attendance.Attendee,
			&attendance.Event,
			&attendance.TimeIn,
			&attendance.ScannedBy,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, &attendance)
	}
	return attendances, nil
}

func (a *Attendance) GetAttendanceById(id uuid.UUID) (*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM attendance WHERE id = $1
	`
	var attendance Attendance

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&attendance.ID,
		&attendance.Attendee,
		&attendance.Event,
		&attendance.TimeIn,
		&attendance.ScannedBy,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}


func (a *Attendance) GetByCreatedAtRange(startTime, endTime time.Time) ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, attendee, event, time_in, scanned_by, created_at, updated_at
		FROM attendance
		WHERE created_at BETWEEN $1 AND $2
	`

	rows, err := db.QueryContext(ctx, query, startTime, endTime)
	if err != nil {
		return nil, err
	}

	var attendances []*Attendance
	for rows.Next() {
		var attendance Attendance
		err := rows.Scan(
			&attendance.ID,
			&attendance.Attendee,
			&attendance.Event,
			&attendance.TimeIn,
			&attendance.ScannedBy,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, &attendance)
	}
	return attendances, nil
}

// TODO: check if all fields is populated
func (a *Attendance) UpdateAttendance(id uuid.UUID, body Attendance) (*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE attendance
		SET
			event = $1,
			attendee = $2,
			time_in = $3,
			scanned_by =$4,
			updated_at = $5
		WHERE id=$6
	`

	_, err := db.ExecContext(
		ctx,
		query,
		body.Event,
		body.Attendee,
		body.TimeIn,
		body.ScannedBy,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (a *Attendance) DeleteAttendance(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM attendance WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Attendance) GetByAttendeeID(attendeeID uuid.UUID) ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, event, time_in, scanned_by from attendance where attendee = $1`

	rows, err := db.QueryContext(ctx, query, attendeeID)
	if err != nil {
		return nil, err
	}
	var attendances []*Attendance
	for rows.Next() {
		var attendance Attendance 
		err := rows.Scan(
			&attendance.ID,
			&attendance.Event,
			&attendance.TimeIn,
			&attendance.ScannedBy,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, &attendance)
	}
	return attendances, nil
}

func (a *Attendance) GetByEventId(eventId uuid.UUID) ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, attendee, event, time_in, scanned_by from attendance where event = $1`

	rows, err := db.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}
	var attendances []*Attendance
	for rows.Next() {
		var attendance Attendance 
		err := rows.Scan(
			&attendance.ID,
			&attendance.Attendee,
			&attendance.Event,
			&attendance.TimeIn,
			&attendance.ScannedBy,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, &attendance)
	}
	return attendances, nil
}

func (a *Attendance) CreateAttendance(attendance Attendance) (*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) 
	defer cancel()
	
	query := `
		INSERT INTO attendance (event, attendee, time_in, scanned_by, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning *
	`
	_, err := db.ExecContext(
		ctx,
		query,
		attendance.Event,
		attendance.Attendee,
		attendance.TimeIn,
		attendance.ScannedBy,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

