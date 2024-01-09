package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (e *Event) GetAllEvents() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from events`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var events []*Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.StartTime,
			&event.EndTime,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (e *Event) GetEventById(id uuid.UUID) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM events WHERE id = $1
	`
	var event Event

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&event.ID,
		&event.StartTime,
		&event.EndTime,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// TODO: check if all fields is populated
func (e *Event) UpdateEvent(id uuid.UUID, body Event) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE events
		SET
			name = $1,
			start_time = $2,
			end_time = $3,
			updated_at = $4
		WHERE id=$5
	`

	_, err := db.ExecContext(
		ctx,
		query,
		body.Name,
		body.StartTime,
		body.EndTime,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (e *Event) DeleteEvent(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM events WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) CreateEvent(event Event) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO events (name, start_time, end_time, created_at, updated_at)
		values ($1, $2, $3, $4, $5) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		event.Name,
		event.StartTime,
		event.EndTime,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
