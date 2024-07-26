package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Attendee struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Program string `json:"program"`
	YearLevel string `json:"year_level"`
	Type string `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Attendee) GetAllAttendees() ([]*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from attendees`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var attendees []*Attendee
	for rows.Next() {
		var attendee Attendee
		err := rows.Scan(
			&attendee.ID,
			&attendee.Name,
			&attendee.Program,
			&attendee.YearLevel,
			&attendee.Type,
			&attendee.UpdatedAt,
			&attendee.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, &attendee)
	}
	return attendees, nil
}

func (a *Attendee) GetByCourses(courses string) ([]*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from attendees where program=$1`

	rows, err := db.QueryContext(ctx, query, courses)
	if err != nil {
		return nil, err
	}
	var attendees []*Attendee
	for rows.Next() {
		var attendee Attendee
		err := rows.Scan(
			&attendee.ID,
			&attendee.Name,
			&attendee.Program,
			&attendee.YearLevel,
			&attendee.UpdatedAt,
			&attendee.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, &attendee)
	}
	return attendees, nil
}
//func (c *Comment) GetCommentsByPostID(id string) ([]*Comment, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
//	defer cancel()
//	query := `SELECT * FROM comments WHERE post_id = $1`
//	rows, err := db.QueryContext(ctx, query, id)
//	if err != nil {
//		return nil, err
//	}
//	var comments []*Comment
//	for rows.Next() {
//		var comment Comment
//		err := rows.Scan(
//			&comment.ID,
//			&comment.Author,
//			&comment.PostID,
//			&comment.CommentBody,
//			&comment.CreatedAt,
//			&comment.UpdatedAt,
//		)
//		if err != nil {
//			return nil, err
//		}
//		comments = append(comments, &comment)
//	}
//	return comments, nil
//}

func (a *Attendee) GetAttendeeById(id uuid.UUID) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM attendees WHERE id = $1
	`
	var attendee Attendee

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&attendee.ID,
		&attendee.Name,
		&attendee.Program,
		&attendee.YearLevel,
		&attendee.Type,
		&attendee.CreatedAt,
		&attendee.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &attendee, nil
}

// TODO: check if all fields is populated
func (a *Attendee) UpdateAttendee(id uuid.UUID, body Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE attendees
		SET
			name = $1,
			program = $2,
			year_level = $3,
			updated_at = $4
		WHERE id=$5
	`
	fmt.Println(body.YearLevel)
	_, err := db.ExecContext(
		ctx,
		query,
		body.Name,
		body.Program,
		body.YearLevel,
		time.Now(),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (a *Attendee) DeleteAttendee(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM attendees WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Attendee) CreateAttendee(attendee Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO attendees (name, program, year_level, type, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		attendee.Name,
		attendee.Program,
		attendee.YearLevel,
		attendee.Type,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &attendee, nil
}

