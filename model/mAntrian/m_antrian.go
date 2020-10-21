package mAntrian

import (
	"database/sql"
	"fmt"
	"time"
)

type Antrian struct {
	Id        int64
	Nomor     int64  // `json:"nomor"`
	Status    string //["waiting","done","expired"]
	Pelanggan string
	Teller    string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
}

const dateFormat = `2006-01-02 15:04:05`

func Getlast(db *sql.DB) (nomor int64, err error) {
	rows, err := db.Query(`SELECT nomor FROM antrian WHERE status="waiting" ORDER BY nomor ASC LIMIT 0,1;`)
	if err != nil {
		return 0, nil
	}
	s := Antrian{}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&s.Nomor)
	}
	return s.Nomor, nil
}

func Create(db *sql.DB) (antrian Antrian, err error) {
	rows, err := db.Query(`SELECT nomor FROM antrian ORDER BY nomor DESC LIMIT 0,1;`)
	s := Antrian{}
	if err != nil {
		s.Nomor = 0
	} else {
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&s.Nomor)
		}
	}

	var dataantrian Antrian
	now := time.Now()
	res, err := db.Exec(`INSERT INTO antrian(nomor,status,created_at,updated_at,expired_at) VALUES(?,?,?,?,?)`,
		s.Nomor+1,
		"waiting",
		now, now, now)
	if err != nil {
		return
	}
	dataantrian.Nomor = s.Nomor + 1
	dataantrian.Status = "waiting"
	dataantrian.CreatedAt = now
	dataantrian.UpdatedAt = now
	dataantrian.ExpiredAt = now
	fmt.Println(res)
	return dataantrian, nil
}

func Update(db *sql.DB, m *Antrian) (affectedRows int64, err error) {
	now := time.Now()
	res, err := db.Exec(`UPDATE antrian SET status=?, pelanggan=?, teller=?, updated_at=? WHERE id=?`,
		m.Status,
		m.Pelanggan,
		m.Teller,
		now,
		m.Id)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

// s := Antrian{}
// rows, err := db.Query(`SELECT * FROM antrian WHERE id=?`, m.Id)
// if err != nil {
// 	return
// }
// defer rows.Close()
// for rows.Next() {
// 	err = rows.Scan(
// 		&s.Id,
// 		&s.Nomor,
// 		&s.Status,
// 		&s.Pelanggan,
// 		&s.Teller,
// 		&s.CreatedAt,
// 		&s.UpdatedAt,
// 		&s.ExpiredAt,)
// 	if err != nil {
// 		return
// 	}
// }

// func SelectAll(db *sql.DB) (students []Student, err error) {
// 	rows, err := db.Query(`SELECT * FROM students ORDER BY id DESC`)
// 	students = []Student{}
// 	if err != nil {
// 		return students, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		s := Student{}
// 		createdAt, updatedAt := ``, ``
// 		err = rows.Scan(
// 			&s.Id,
// 			&s.Nim,
// 			&s.Name,
// 			&s.Semester,
// 			&createdAt,
// 			&updatedAt)
// 		if err != nil {
// 			return
// 		}
// 		s.CreatedAt, _ = time.Parse(dateFormat, createdAt)
// 		s.UpdatedAt, _ = time.Parse(dateFormat, updatedAt)
// 		students = append(students, s)
// 	}
// 	return
// }

// func Insert(db *sql.DB, m *Student) (err error) {
// 	now := time.Now()
// 	res, err := db.Exec(`INSERT INTO students(name,nim,semester,created_at,updated_at)
// VALUES(?,?,?,?,?)`,
// 		m.Name,
// 		m.Nim,
// 		m.Semester,
// 		now,
// 		now)
// 	if err != nil {
// 		return err
// 	}
// 	m.Id, err = res.LastInsertId()
// 	if err == nil {
// 		m.CreatedAt = now
// 		m.UpdatedAt = now
// 	}
// 	return err
// }

// func Update(db *sql.DB, m *Student) (affectedRows int64, err error) {
// 	s := Student{}
// 	rows, err := db.Query(`SELECT * FROM students WHERE id=?`, m.Id)
// 	if err != nil {
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		createdAt, updatedAt := ``, ``
// 		err = rows.Scan(
// 			&s.Id,
// 			&s.Nim,
// 			&s.Name,
// 			&s.Semester,
// 			&createdAt,
// 			&updatedAt)
// 		if err != nil {
// 			return
// 		}
// 	}
// 	if m.Name == "" {
// 		m.Name = s.Name
// 	}
// 	if m.Nim == "" {
// 		m.Nim = s.Nim
// 	}
// 	if m.Semester == 0 {
// 		m.Semester = s.Semester
// 	}
// 	now := time.Now()
// 	res, err := db.Exec(`UPDATE students SET name=?, nim=?, semester=?, updated_at=? WHERE id=?`,
// 		m.Name,
// 		m.Nim,
// 		m.Semester,
// 		now,
// 		m.Id)
// 	if err != nil {
// 		return
// 	}
// 	return res.RowsAffected()
// }

// func Delete(db *sql.DB, m *Student) (deletedRecord Student, success bool) {
// 	s := Student{}
// 	rows, err := db.Query(`SELECT * FROM students WHERE id=?`, m.Id)
// 	if err != nil {
// 		return s, false
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		createdAt, updatedAt := ``, ``
// 		err = rows.Scan(
// 			&s.Id,
// 			&s.Nim,
// 			&s.Name,
// 			&s.Semester,
// 			&createdAt,
// 			&updatedAt)
// 		if err != nil {
// 			return
// 		}
// 		s.CreatedAt, _ = time.Parse(dateFormat, createdAt)
// 		s.UpdatedAt, _ = time.Parse(dateFormat, updatedAt)
// 	}
// 	res, err := db.Exec(`DELETE FROM students WHERE id=?`, m.Id)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("%s", res)
// 	return s, true
// }
