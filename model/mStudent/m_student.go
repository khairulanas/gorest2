package mStudent

import (
	"database/sql"
	"fmt"
	"time"
)

type Student struct {
	Id        int64
	Nim       string // `json:"nim"`
	Name      string
	Semester  int
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

func Insert(db *sql.DB, m *Student) (err error) {
	now := time.Now()
	res, err := db.Exec(`INSERT INTO students(name,nim,semester,created_at,updated_at)
VALUES(?,?,?,?,?)`,
		m.Name,
		m.Nim,
		m.Semester,
		now,
		now)
	if err != nil {
		return err
	}
	m.Id, err = res.LastInsertId()
	if err == nil {
		m.CreatedAt = now
		m.UpdatedAt = now
	}
	return err
}
func Update(db *sql.DB, m *Student) (affectedRows int64, err error) {
	s := Student{}
	rows, err := db.Query(`SELECT * FROM students WHERE id=?`, m.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
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
	}
	if m.Name == "" {
		m.Name = s.Name
	}
	if m.Nim == "" {
		m.Nim = s.Nim
	}
	if m.Semester == 0 {
		m.Semester = s.Semester
	}
	now := time.Now()
	res, err := db.Exec(`UPDATE students SET name=?, nim=?, semester=?, updated_at=? WHERE id=?`,
		m.Name,
		m.Nim,
		m.Semester,
		now,
		m.Id)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

func Delete(db *sql.DB, m *Student) (deletedRecord Student, success bool) {
	s := Student{}
	rows, err := db.Query(`SELECT * FROM students WHERE id=?`, m.Id)
	if err != nil {
		return s, false
	}
	defer rows.Close()
	for rows.Next() {
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
	}

	res, err := db.Exec(`DELETE FROM students WHERE id=?`, m.Id)
	if err != nil {
		return
	}
	fmt.Printf("%s", res)
	return s, true
}
