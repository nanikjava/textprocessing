package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"rockt/model"
	"rockt/repo"
	"strings"
)

const createDB string = `
  CREATE TABLE IF NOT EXISTS userdata (      
  date DATETIME NOT NULL,
  email TEXT  NOT NULL,
  sessionid TEXT  NOT NULL,
  filename TEXT NOT NULL
  );`

const queryDB string = `SELECT date,email,sessionid FROM userdata WHERE filename = ?  AND date BETWEEN ? and ?  ORDER BY date ASC;`

type Repo struct {
	db *sql.DB
}

func (r Repo) Query(from string, to string, filename string) []model.Datarecord {
	var records []model.Datarecord

	rows, err := r.db.Query(queryDB, filename, from, to)
	if err != nil {
		fmt.Println("Error querying :", err)
		return records
	}

	for rows.Next() {
		r := model.Datarecord{}
		err := rows.Scan(&r.DateISO8601, &r.EmailAddress, &r.SessionID)

		if err != nil {
			log.Println("could not scan row: ", err.Error())
		}

		records = append(records, r)
	}

	defer rows.Close()

	// return empty data if none is found
	if records == nil {
		return []model.Datarecord{}
	}
	return records
}

func NewRepository() (repo.Repository, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	//db, err := sql.Open("sqlite3", "/home/nanik/GolandProjects/rockt/repository/data.db")
	if err != nil {
		return nil, err
	}

	return Repo{db}, nil
}

func (r Repo) Create() error {
	statement, err := r.db.Prepare(createDB)

	if err != nil {
		log.Println("Problem creating table ")
	}
	_, err = statement.Exec()

	return err
}

func (r Repo) Close() {
	r.db.Close()
}

func (r Repo) BulkInsert(records []model.Datarecord) error {
	params := []string{}
	args := []interface{}{}

	for _, l := range records {
		params = append(params, "(?, ?, ?, ?)")
		args = append(args, l.DateISO8601)
		args = append(args, l.EmailAddress)
		args = append(args, l.SessionID)
		args = append(args, l.FileName)
	}

	sInsert := "INSERT INTO userdata (date, email, sessionid, filename) VALUES %s"
	sStmt := fmt.Sprintf(sInsert, strings.Join(params, ","))

	//Begin: Transaction
	txn, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		_ = txn.Rollback()
	}()

	rcount, err := txn.ExecContext(context.Background(), sStmt, args...)
	if err != nil {
		//txn.Rollback()
		fmt.Println("Error executing insert : ", err)
		return err
	}

	cnt, err := rcount.RowsAffected()
	fmt.Println("Total number of records inserted  : ", cnt)

	return txn.Commit()
}
