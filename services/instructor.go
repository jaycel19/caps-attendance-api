package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Instructor struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (i *Instructor) GetAllInstructor() ([]*Instructor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from instructors`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var instructors []*Instructor
	for rows.Next() {
		var instructor Instructor
		err := rows.Scan(
			&instructor.ID,
			&instructor.Name,
			&instructor.UpdatedAt,
			&instructor.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		instructors = append(instructors, &instructor)
	}
	return instructors, nil
}


func (i *Instructor) GetInstructorById(id uuid.UUID) (*Instructor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM instructors WHERE id = $1
	`
	var instructor Instructor

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&instructor.ID,
		&instructor.Name,
		&instructor.CreatedAt,
		&instructor.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &instructor, nil
}

// TODO: check if all fields is populated
func (i *Instructor) UpdateInstructor(id uuid.UUID, body Instructor) (*Instructor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE instructors 
		SET
			name = $1,
			updated_at = $2
		WHERE id=$3
	`

	_, err := db.ExecContext(
		ctx,
		query,
		body.Name,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (i *Instructor) DeleteInstructor(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM instructors WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (i *Instructor) CreatePersonnel(instructor Instructor) (*Instructor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO instructors (name, created_at, updated_at)
		values ($1, $2, $3) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		instructor.Name,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &instructor, nil
}
