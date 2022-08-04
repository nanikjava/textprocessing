package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"rockt/repository/model"
	"rockt/repository/repository"
	"strings"
	"time"
)

const createDB string = `
  CREATE TABLE IF NOT EXISTS userdata (
  date DATETIME NOT NULL,
  email TEXT  NOT NULL,
  sessionid TEXT  NOT NULL
  );`

const queryDB string = `SELECT * FROM userdata WHERE date BETWEEN ? and ? ORDER BY date ASC;`

type repo struct {
	db *sql.DB
}

func (r repo) Query(from string, to string) []model.Datarecord {
	var records []model.Datarecord
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, queryDB, from, to)
	if err != nil {
		fmt.Println("Error querying :", err)
		return records
	}

	for rows.Next() {
		r := model.Datarecord{}
		err := rows.Scan(&r.DateISO8601, &r.EmailAddress, &r.SessionID)

		if err != nil {
			fmt.Println("could not scan row: %v", err)
		}

		records = append(records, r)
	}

	defer rows.Close()

	return records
}

func NewRepository() (repository.Repository, error) {
	//db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	db, err := sql.Open("sqlite3", "/home/nanik/GolandProjects/rockt/repository/data.db")
	if err != nil {
		return nil, err
	}

	return repo{db}, nil
}

func (r repo) Create() error {
	statement, err := r.db.Prepare(createDB)

	if err != nil {
		log.Println("Problem creating table ")
	}
	_, err = statement.Exec()

	return err
}

func (r repo) Close() {
	r.db.Close()
}

func (r repo) BulkInsert(records []model.Datarecord) error {
	params := []string{}
	args := []interface{}{}

	for _, l := range records {
		params = append(params, "(?, ?, ?)")
		args = append(args, l.DateISO8601)
		args = append(args, l.EmailAddress)
		args = append(args, l.SessionID)
	}

	sInsert := "INSERT INTO userdata (date, email, sessionid) VALUES %s"
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
