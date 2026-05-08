package parser

import (
	"cotonetes/types"
	"cotonetes/utils"
	"log"
	"os"
	"testing"
)

func createLatexFile(folder_path string, file_name string, note types.Note) {
	file_path := folder_path + "/" + file_name

	f, err := os.OpenFile(file_path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Error opening file %s: %v", file_path, err)
	}

	defer f.Close()

	for _, line := range utils.NoteToLatex(note) {
		f.WriteString(line + "\n")
	}
}

func setupTest(t *testing.T, note types.Note) (string, string) {
	folder_path := t.TempDir()
	file_name := "sample.tex"

	createLatexFile(folder_path, file_name, note)

	return folder_path, file_name
}

func LatexParserTest(t *testing.T, folder_path string, expected_markdown types.Note) {
	processed_notes := Process_files(folder_path, "tex")

	utils.FailNotEquals(t, "Failed to process expected number of files", 1, len(processed_notes))

	utils.FailNotEquals(t, "Failed to process expected number of notes", 1, len(processed_notes[0].Notes))

	processed_note := processed_notes[0].Notes[0]

	utils.FailNotEquals(t, "Failed to process note title", expected_markdown.Title, processed_note.Title)
	utils.FailNotEquals(t, "Failed to process note url", expected_markdown.Url, processed_note.Url)
	utils.FailNotEquals(t, "Failed to process note created date", expected_markdown.Created_date, processed_note.Created_date)
	utils.FailNotEquals(t, "Failed to process note updated date", expected_markdown.Updated_date, processed_note.Updated_date)

	emptyline, note_content := processed_note.Text[0], processed_note.Text[1:]

	utils.FailNotEquals(t, "Failed to process note empty line preamble", "", emptyline)

	utils.FailNotEqualsSlice(t, "Failed to process note content line", expected_markdown.Text, note_content)
}

func baseLatexParserTest(t *testing.T, test_input utils.TestInput) {
	folder_path, _ := setupTest(t, test_input.Latex)

	LatexParserTest(t, folder_path, test_input.Markdown)
}

////
// Tests
////

func TestTextOnly(t *testing.T) {
	baseLatexParserTest(t, utils.TdTextOnly)
}

func TestTitleSpecialChars(t *testing.T) {
	baseLatexParserTest(t, utils.TdTitleSpecialChars)
}

func TestNoteBoldText(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteBoldText)
}

func TestNoteUrl(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteUrl)
}

func TestNoteNewline(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteNewline)
}

func TestNoteItemize(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteItemize)
}

func TestNoteEnumerate(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteEnumerate)
}

func TestNoteVerbatim(t *testing.T) {
	baseLatexParserTest(t, utils.TdNoteVerbatim)
}
