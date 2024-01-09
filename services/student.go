package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Program string `json:"program"`
	YearLevel string `json:"year_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Student) GetAllStudents() ([]*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from comments`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var students []*Student
	for rows.Next() {
		var student Student
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Program,
			&student.YearLevel,
			&student.UpdatedAt,
			&student.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (s *Student) GetByCourses(courses string) ([]*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from students where program=$1`

	rows, err := db.QueryContext(ctx, query, courses)
	if err != nil {
		return nil, err
	}
	var students []*Student
	for rows.Next() {
		var student Student
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Program,
			&student.YearLevel,
			&student.UpdatedAt,
			&student.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
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

func (s *Student) GetStudentById(id uuid.UUID) (*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT * FROM students WHERE id = $1
	`
	var student Student

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Program,
		&student.YearLevel,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// TODO: check if all fields is populated
func (s *Student) UpdateStudent(id uuid.UUID, body Student) (*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE students
		SET
			name = $1,
			program = $2,
			year_level = $3,
			updated_at = $4
		WHERE id=$5
	`

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

func (s *Student) DeleteStudent(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM students WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) CreateComment(student Student) (*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		INSERT INTO students (name, program, year_level, created_at, updated_at)
		values ($1, $2, $3, $4, $5) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		student.Name,
		student.Program,
		student.YearLevel,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

