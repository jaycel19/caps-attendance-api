package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Personnel struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
}

func (p *Personnel) GetAllPersonnel() ([]*Personnel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from personnels`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var personnels []*Personnel
	for rows.Next() {
		var personnel Personnel
		err := rows.Scan(
			&personnel.ID,
			&personnel.Name,
			&personnel.UpdatedAt,
			&personnel.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		personnels = append(personnels, &personnel)
	}
	return personnels, nil
}


func (p *Personnel) GetPersonnelById(id uuid.UUID) (*Personnel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM personnels WHERE id = $1
	`
	var personnel Personnel

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&personnel.ID,
		&personnel.Name,
		&personnel.CreatedAt,
		&personnel.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &personnel, nil
}

// TODO: check if all fields is populated
func (p *Personnel) UpdatePersonnel(id uuid.UUID, body Personnel) (*Personnel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE personnels
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

func (p *Personnel) DeletePersonnel(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM personnels WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) CreatePersonnel(personnel Personnel) (*Personnel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO personnels (name, created_at, updated_at)
		values ($1, $2, $3) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		personnel.Name,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &personnel, nil
}
