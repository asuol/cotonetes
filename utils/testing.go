package utils

import (
	"testing"
	"cotonetes/types"
	"reflect"
)

type TestInput struct{
	Latex types.Note
	Markdown types.Note
}

func NoteToLatex(note types.Note) []string {
	latex := make([]string, 0)

	latex = append(latex, `\textbf{Title:} ` + note.Title + `\\`)
	latex = append(latex, `\textbf{URL:} \url{` + note.Url + `}\\`)
	latex = append(latex, `\textbf{Created:} ` + note.Created_date + `\\`)
	latex = append(latex, `\textbf{Last Updated:} ` + note.Updated_date + `\\`)
	latex = append(latex, `\\`)

	for _, line := range note.Text {
		latex = append(latex, line)
	}

	latex = append(latex, `\hrulefill`)

	latex = append(latex, `\\`)

	latex = append(latex, ``)

	return latex
}

func FailTest(t *testing.T, message string, expected any, obtained any) {
	t.Fatalf("%s:\n expected %+v\n obtained %+v", message, expected, obtained)
}

func FailNotEquals(t *testing.T, message string, expected any, obtained any) {
	if expected != obtained {
		FailTest(t, message, expected, obtained)
	}
}

func FailNotEqualsSlice(t *testing.T, message string, expected []string, obtained []string) {
	for index, line := range obtained {
		if expected[index] != line {
			t.Fatalf("%s:\n expected %+v\n obtained %+v\n full expected slice %+v\n full obtained slice %+v", message, expected[index], line, expected, obtained)
		}
	}
}

func FailNotEqualsStruct(t *testing.T, message string, expected any, obtained any) {
	if !reflect.DeepEqual(expected, obtained) {
		FailTest(t, message, expected, obtained)
	}
}
