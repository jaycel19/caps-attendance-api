package services

import (
	"context"
	"log"

	"github.com/jaycel19/capstone-api/util"
)


type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func (a *Admin) AdminLogin(adminPayload Admin) (*Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select * from admin where username = $1`

	var admin Admin 

	row := db.QueryRowContext(ctx, query, adminPayload.Username)
	err := row.Scan(
		&admin.Username,
		&admin.Password,
	)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (a *Admin) GetAllAdmins() ([]*Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from admin`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var admins []*Admin
	for rows.Next() {
		var admin Admin
		err := rows.Scan(
			&admin.Username,
			&admin.Password,
		)
		if err != nil {
			return nil, err
		}
		admins = append(admins, &admin)
	}
	return admins, nil
}

// TODO: check if all fields is populated
func (a *Admin) UpdateAdmin(username string, body Admin) (*Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE admin
		SET
			username = $1,
			password = $2,
		WHERE username=$1
	`
	_, err := db.ExecContext(
		ctx,
		query,
		body.Username,
		body.Password,
		username,
	)

	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (a *Attendee) DeleteAdmin(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM admin WHERE username=$1`
	_, err := db.ExecContext(ctx, query, username)
	if err != nil {
		return err
	}
	return nil
}

func (a *Admin) CreateAdmin(admin Admin) (*Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO admin (username, password)
		values ($1, $2 ) returning *
	`
	
	hashedPassword, err := util.HashPassword(admin.Password)
	if err != nil {
		log.Print("Error Hashing password admin")
		return nil, err
	}

	_, err = db.ExecContext(
		ctx,
		query,
		admin.Username,
		hashedPassword,
	)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

