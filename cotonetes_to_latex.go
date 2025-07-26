package main

import (
	"fmt"
	"flag"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"cotonetes/types"
	"cotonetes/parser"
	"strings"
	"regexp"
	"path/filepath"
)

type category struct {
	Id int
	Category string
}

func main() {
	db_path_ptr := flag.String("db", "cotonetes.db", "Path to database file")
	export_notes_path_ptr := flag.String("notes", "/tmp/export", "Path to folder to store exported notes in latex format")

	flag.Parse()

	if  _, error := os.Stat(*export_notes_path_ptr); error != nil {
		log.Fatal(fmt.Sprintf("Provided note folder does not exist!: %s", *export_notes_path_ptr))
	}

	if  _, error := os.Stat(*db_path_ptr); error != nil {
		log.Fatal(fmt.Sprintf("Provided database file does not exist!: %s", *db_path_ptr))
	}

	var err error
	var db *sql.DB

	if db, err = sql.Open("sqlite3", *db_path_ptr); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	select_categories_stmt := `SELECT * FROM categories`
	
	select_notes_stmt := `SELECT notes.title, notes.url, notes.created, notes.last_updated, notes.note FROM notes INNER JOIN note_categories ON notes.id = note_categories.note_id WHERE note_categories.category_id = $1;`

	var rows *sql.Rows

	if rows, err = db.Query(select_categories_stmt); err != nil {
		log.Fatalf("%q: %s\n", err, select_categories_stmt)
	}

	defer rows.Close()

	for rows.Next() {
		var cat category

		if err = rows.Scan(&cat.Id, &cat.Category); err != nil {
			log.Fatalf("%q\n", err)
		}

		category_folder_path := filepath.Join(*export_notes_path_ptr, cat.Category)

		if err := os.MkdirAll(category_folder_path, 755); err != nil {
			log.Fatalf("%q\n", err)
		}

		file_name_re := regexp.MustCompile(`^.*` + string(os.PathSeparator) + `(.*)$`)

		file_name := file_name_re.ReplaceAllString(cat.Category, "$1")

		sub_section_index := len(strings.Split(cat.Category, string(os.PathSeparator))) - 1

		file_name_path := filepath.Join(category_folder_path, file_name + ".tex")

		var cat_rows *sql.Rows
		
		if cat_rows, err = db.Query(select_notes_stmt, cat.Id); err != nil {
			log.Fatalf("%q: %s\n", err, select_notes_stmt)
		}

		note_list := make([]types.Note, 0)

		for cat_rows.Next() {
			var note types.Note
			var text string

			if err := cat_rows.Scan(&note.Title, &note.Url, &note.Created_date, &note.Updated_date, &text); err != nil {
				log.Fatalf("%q\n", err)
			}

			note.Text = strings.Split(text,"\n")

			note_list = append(note_list, note)
		}
		
		parser.Export_to_latex_file(file_name_path, file_name, sub_section_index, note_list)
	}
}
