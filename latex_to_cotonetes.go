package main

import (
	"fmt"
	"flag"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"cotonetes/parser"
	"regexp"
	"strings"
)

func user_confirmation(msg string) bool {
	var i string
	fmt.Println(msg, " [y/n]")
	fmt.Scan(&i)
	if i != "y" && i != "n" {
		fmt.Println("Please answer with \"y\" or \"n\"")
		return user_confirmation(msg)
	}

	return i == "y"
}

func main() {
	db_path_ptr := flag.String("db", "cotonetes.db", "Path to new database file")
	latex_notes_path_ptr := flag.String("notes", "/tmp/notes", "Path to folder containing notes in latex format")

	flag.Parse()

	if  _, error := os.Stat(*latex_notes_path_ptr); error != nil {
		log.Fatal(fmt.Sprintf("Provided note folder does not exist!: %s", *latex_notes_path_ptr))
	}

	if  _, error := os.Stat(*db_path_ptr); error == nil {
		if user_confirmation(fmt.Sprintf("File %s already exists. Delete?", *db_path_ptr)) {
			if e := os.Remove(*db_path_ptr); e != nil {
				log.Fatal(e)
			}
		} else {
			fmt.Println("Please re-run the program with another database path to continue")
			os.Exit(0)
		}
	}


	db, err := sql.Open("sqlite3", *db_path_ptr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	create_tables_stmt := `
	create table categories (id INTEGER PRIMARY KEY, category TEXT UNIQUE NOT NULL);
	create table notes (id INTEGER PRIMARY KEY, title TEXT NOT NULL, url TEXT NOT NULL, created INTEGER NOT NULL, last_updated INTEGER NOT NULL, note TEXT NOT NULL);
	create table note_categories (note_id INTEGER, category_id INTEGER, FOREIGN KEY(note_id) REFERENCES notes(id), FOREIGN KEY(category_id) REFERENCES categories(id), PRIMARY KEY(note_id, category_id));
	`
	_, err = db.Exec(create_tables_stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, create_tables_stmt)
		return
	}
	file_notes := parser.Process_files(*latex_notes_path_ptr, "tex")

	insert_category_stmt := `insert into categories (category) values ($1);`
	insert_note_stmt := `insert into notes (title, url, created, last_updated, note) values ($1, $2, $3, $4, $5);`
	insert_note_category_stmt := `insert into note_categories (note_id, category_id) values ($1, $2);`
	
	category_re := regexp.MustCompile(regexp.QuoteMeta(*latex_notes_path_ptr) + string(os.PathSeparator) + `(.*)` + string(os.PathSeparator) + `.*$`)

	for _, f := range file_notes {

		file_path := f.File_path

		category := category_re.ReplaceAllString(file_path, "$1")

		var err error
		var res sql.Result

		if res, err = db.Exec(insert_category_stmt, category); err != nil {
			log.Printf("%q: %s\n", err, insert_category_stmt)
			return
		}

		var cat_id int64

		if cat_id, err = res.LastInsertId(); err != nil {
			log.Printf("%q\n", err)
			return
		}

		for _, note := range f.Notes {
			if res, err = db.Exec(insert_note_stmt, note.Title, note.Url, note.Created_date, note.Updated_date, strings.Join(note.Text, "\n")); err != nil {
				log.Printf("%q: %s\n", err, insert_note_stmt)
				return
			}

			var note_id int64

			if note_id, err = res.LastInsertId(); err != nil {
				log.Printf("%q\n", err)
				return
			}

			if res, err = db.Exec(insert_note_category_stmt, note_id, cat_id); err != nil {
				log.Printf("%q: %s\n", err, insert_note_category_stmt)
				return
			}
		}
	}
}
