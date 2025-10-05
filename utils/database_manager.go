package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
	"cotonetes/types"
)

type DatabaseManager struct{
	Db_path string
}

func (d *DatabaseManager) OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", d.Db_path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (d *DatabaseManager) BeginTransaction(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	return tx
}

func (d *DatabaseManager) CommitTransaction(tx *sql.Tx) {
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (d *DatabaseManager) CreateDatabase(tx *sql.Tx) {
	create_tables_stmt := `
	create table categories (id INTEGER PRIMARY KEY, category TEXT UNIQUE NOT NULL);
	create table notes (id INTEGER PRIMARY KEY, title TEXT NOT NULL, url TEXT NOT NULL, created INTEGER NOT NULL, last_updated INTEGER NOT NULL, note TEXT NOT NULL);
	create table note_categories (note_id INTEGER, category_id INTEGER, FOREIGN KEY(note_id) REFERENCES notes(id), FOREIGN KEY(category_id) REFERENCES categories(id), PRIMARY KEY(note_id, category_id));
	`

	if _, err := tx.Exec(create_tables_stmt); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("%q: unable to rollback: %q\n", err, rollbackErr)
		}
		log.Fatalf("%q: %s\n", err, create_tables_stmt)
	}
}

func (d *DatabaseManager) CreateCategory(tx *sql.Tx, category string) int64 {
	var err error
	var res sql.Result

	insert_category_stmt := `insert into categories (category) values ($1);`

	if res, err = tx.Exec(insert_category_stmt, category); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("%q: unable to rollback: %q\n", err, rollbackErr)
		}
		log.Fatal("%q: %s\n", err, insert_category_stmt)
	}

	var cat_id int64

	if cat_id, err = res.LastInsertId(); err != nil {
		log.Fatal(err)
	}

	return cat_id
}

func (d *DatabaseManager) AddNote(tx *sql.Tx, note types.Note) int64 {
	var err error
	var res sql.Result

	insert_note_stmt := `insert into notes (title, url, created, last_updated, note) values ($1, $2, $3, $4, $5);`

	if res, err = tx.Exec(insert_note_stmt, note.Title, note.Url, note.Created_date, note.Updated_date, strings.Join(note.Text, "\n")); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("%q: unable to rollback: %q\n", err, rollbackErr)
		}
		log.Fatal("%q: %s\n", err, insert_note_stmt)
	}

	var note_id int64

	if note_id, err = res.LastInsertId(); err != nil {
		log.Fatal(err)
	}

	return note_id
}

func (d *DatabaseManager) AddNoteCategory(tx *sql.Tx, note_id int64, cat_id int64) {
	insert_note_category_stmt := `insert into note_categories (note_id, category_id) values ($1, $2);`

	if _, err := tx.Exec(insert_note_category_stmt, note_id, cat_id); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("%q: unable to rollback: %q\n", err, rollbackErr)
		}
		log.Fatal("%q: %s\n", err, insert_note_category_stmt)
	}
}
