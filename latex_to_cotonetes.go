package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"cotonetes/parser"
	"cotonetes/utils"
	"regexp"
)

func main() {
	db_path_ptr := flag.String("db", "cotonetes.db", "Path to new database file")
	latex_notes_path_ptr := flag.String("notes", "/tmp/notes", "Path to folder containing notes in latex format")

	flag.Parse()

	if  _, error := os.Stat(*latex_notes_path_ptr); error != nil {
		log.Fatal(fmt.Sprintf("Provided note folder does not exist!: %s", *latex_notes_path_ptr))
	}

	if  _, error := os.Stat(*db_path_ptr); error == nil {
		if utils.User_confirmation(fmt.Sprintf("File %s already exists. Delete?", *db_path_ptr)) {
			if e := os.Remove(*db_path_ptr); e != nil {
				log.Fatal(e)
			}
		} else {
			fmt.Println("Please re-run the program with another database path to continue")
			os.Exit(0)
		}
	}

	db_manager := utils.DatabaseManager{*db_path_ptr}

	db := db_manager.OpenDatabase()
	defer db.Close()

	tx := db_manager.BeginTransaction(db)

	db_manager.CreateDatabase(tx)

	file_notes := parser.Process_files(*latex_notes_path_ptr, "tex")

	category_re := regexp.MustCompile(regexp.QuoteMeta(*latex_notes_path_ptr) + string(os.PathSeparator) + `(.*)` + string(os.PathSeparator) + `.*$`)

	for _, f := range file_notes {

		file_path := f.File_path

		category := category_re.ReplaceAllString(file_path, "$1")

		cat_id := db_manager.CreateCategory(tx, category)

		for _, note := range f.Notes {
			note_id := db_manager.AddNote(tx, note)

			db_manager.AddNoteCategory(tx, note_id, cat_id)
		}
	}

	db_manager.CommitTransaction(tx)
}
