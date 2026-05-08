package parser

import (
	"bufio"
	"cotonetes/types"
	"cotonetes/utils"
	"log"
	"os"
	"strings"
	"testing"
)

// latexToCotonetes takes the input note, with text in markdown format,
// calls the Export_to_latex function to generate a latex file and returns
// a string slice containing the latex file lines
func LatexToCotonetes(t *testing.T, note types.Note) ([]string, string) {
	folder_path := t.TempDir()
	file_path := folder_path + "/test.tex"

	err := Export_to_latex_file(file_path, "test", 0, []types.Note{note})

	utils.FailNotEquals(t, "Failed export", nil, err)

	f, err := os.Open(file_path)

	if err != nil {
		log.Fatalf("Error opening file %s: %v", file_path, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	output_file_contents := make([]string, 0, 20)

	for scanner.Scan() {
		line := scanner.Text()

		output_file_contents = append(output_file_contents, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading %s: %v", file_path, err)
	}

	return output_file_contents, folder_path
}

func baseMarkdownParserTest(t *testing.T, test_input utils.TestInput) {
	output_file_contents, _ := LatexToCotonetes(t, test_input.Markdown)

	expected_content := []string{
		`\section{test}`,
		"",
	}

	expected_content = append(expected_content, strings.Join(utils.NoteToLatex(test_input.Latex), "\n"))

	utils.FailNotEqualsSlice(t, "Failed to process note content line", expected_content, output_file_contents)
}

////
// Tests
////

func TestBasicString(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdTextOnly)
}

func TestSpecialCharsString(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdTitleSpecialChars)
}

func TestBoldString(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteBoldText)
}

func TestUrl(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteUrl)
}

func TestNewline(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteNewline)
}

func TestItemize(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteItemize)
}

func TestEnumerate(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteEnumerate)
}

func TestVerbatim(t *testing.T) {
	baseMarkdownParserTest(t, utils.TdNoteVerbatim)
}
