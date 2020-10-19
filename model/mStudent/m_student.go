package mStudent

import (
	"database/sql"
	"time"
)

type Student struct {
	Id int64 
	Nim string
	Name string
	Semester int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const dateFormat = `2006-01-02 15:04:05`

func SelectAll(db *sql.DB) (students []Student, err error) {
	rows, err := db.Query(`SELECT * FROM students ORDER BY id DESC`)
	students = []Student{}
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		s := Student{}
		createdAt, updatedAt := ``, ``
		err = rows.Scan(
			&s.Id,
			&s.Nim,
			&s.Name,
			&s.Semester,
			&createdAt,
			&updatedAt)
		if err != nil {
			return 
		}
		s.CreatedAt, _ = time.Parse(dateFormat, createdAt)
		s.UpdatedAt, _ = time.Parse(dateFormat, updatedAt)
		students = append(students, s)
	}
	return
}
