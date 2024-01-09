package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID uuid.UUID `json:"id"`
	Attendee uuid.UUID `json:"attendee"`
	AttendeeInstructor uuid.UUID `json:"attendee_instructor"`
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
			&attendance.AttendeeInstructor,
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
		&a.AttendeeInstructor,
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

// TODO: check if all fields is populated
func (a *Attendance) UpdateAttendance(id uuid.UUID, body Attendance) (*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE attendance
		SET
			event = $1,
			attendee = $2,
			attendee_instructor = $3,
			time_in = $4,
			scanned_by = $5,
			updated_at = $6
		WHERE id=$7
	`

	_, err := db.ExecContext(
		ctx,
		query,
		body.Event,
		body.Attendee,
		body.AttendeeInstructor,
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

func (a *Attendance) GetByStudentID(studentID uuid.UUID) ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, event, time_in, scanned_by from attendance where attendee = $1`

	rows, err := db.QueryContext(ctx, query, studentID)
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

func (a *Attendance) GetByInstructorID(instructorID uuid.UUID) ([]*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, event, time_in, scanned_by from attendance where attendee_instructor = $1`

	rows, err := db.QueryContext(ctx, query, instructorID)
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

func (a *Attendance) CreateAttendance(attendance Attendance) (*Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) 
	defer cancel()

	query := `
		INSERT INTO attendance (event, attendee, attendee_instructor, time_in, scanned_by, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		attendance.Event,
		attendance.Attendee,
		attendance.AttendeeInstructor,
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

