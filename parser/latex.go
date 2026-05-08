package parser

import (
	"bufio"
	"cotonetes/latex_parser"
	"cotonetes/types"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileNotes struct {
	File_path string
	Notes     []types.Note
}

func escape_special_chars_to_markdown(line string) string {
	special_re := regexp.MustCompile(`\\(.)`)

	return special_re.ReplaceAllString(line, "$1")
}

func latex_note_title_to_txt(line string) string {
	re := regexp.MustCompile(`\.*{.*} (.*[^\\])\\`)
	match := re.FindStringSubmatch(line)

	if len(match) == 0 {
		log.Fatalf("Error processing line: %s", line)
	}

	return match[1]
}

func latex_note_url_to_txt(line string) string {
	line_net_content := latex_note_title_to_txt(line)
	re := regexp.MustCompile(`\\url{(.*)}`)
	return re.FindStringSubmatch(line_net_content)[1]
}

type LatexListener struct {
	*latex_parser.BaseLatexListener

	Title                  string
	Url                    string
	Created                string
	Updated                string
	Note                   []string
	block_stack            []string
	word_stack             []string
	text_stack             []string
	verbatim_content_stack []string
	is_verbatim_block      bool
}

func (s *LatexListener) getTagValues(tag []antlr.Token) string {
	values := make([]string, 0)
	for _, tag := range tag {
		values = append(values, tag.GetText())
	}

	return strings.Join(values, "")
}

func (s *LatexListener) getWord() string {
	word := strings.Join(s.word_stack, "")

	s.word_stack = nil

	return word
}

// a line break and an empty line both mean a single empty line, even if combined or there are multiple empty lines
func (s *LatexListener) addParagraph() {
	if s.is_verbatim_block == false && (s.Note == nil || s.Note[len(s.Note)-1] != "") {
		s.Note = append(s.Note, "")
	}
}

func (s *LatexListener) EnterTag(ctx *latex_parser.TagContext) {
	if s.word_stack != nil {
		s.text_stack = append(s.text_stack, s.getWord())
	}
}

func (s *LatexListener) ExitTag(ctx *latex_parser.TagContext) {
	tagVal := s.getWord()
	switch ctx.GetName().GetText() {
	case "textbf":
		s.text_stack = append(s.text_stack, fmt.Sprintf("**%s**", tagVal))
	case "url":
		s.text_stack = append(s.text_stack, fmt.Sprintf("[%s](%s)", tagVal, tagVal))
	default:
		log.Fatalf("Unknown tag: %s", ctx.GetText())
	}
}

func (s *LatexListener) ExitNote_title(ctx *latex_parser.Note_titleContext) {
	s.Title = s.getWord()
}

func (s *LatexListener) ExitNote_url(ctx *latex_parser.Note_urlContext) {
	s.Url = s.getWord()
}

func (s *LatexListener) ExitNote_created(ctx *latex_parser.Note_createdContext) {
	s.Created = s.getWord()
}

func (s *LatexListener) ExitNote_updated(ctx *latex_parser.Note_updatedContext) {
	s.Updated = s.getWord()
}

func (s *LatexListener) ExitText(ctx *latex_parser.TextContext) {
	if s.word_stack != nil {
		s.text_stack = append(s.text_stack, s.getWord())
	}
	s.Note = append(s.Note, strings.Join(s.text_stack, ""))

	s.text_stack = nil
}

func (s *LatexListener) ExitEscaped_word(ctx *latex_parser.Escaped_wordContext) {
	content := ctx.GetContent().GetText()
	space := ""
	if ctx.GetSpace() != nil {
		space = ctx.GetSpace().GetText()
	}
	s.word_stack = append(s.word_stack, content+space)
}

func (s *LatexListener) ExitLetter(ctx *latex_parser.LetterContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *LatexListener) ExitPunctuation(ctx *latex_parser.PunctuationContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *LatexListener) ExitNumber(ctx *latex_parser.NumberContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *LatexListener) ExitWs(ctx *latex_parser.WsContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *LatexListener) ExitLine_break(ctx *latex_parser.Line_breakContext) {
	s.addParagraph()
}

func (s *LatexListener) ExitEmpty_line(ctx *latex_parser.Empty_lineContext) {
	s.addParagraph()
}

func (s *LatexListener) ExitBlock_line(ctx *latex_parser.Block_lineContext) {
	s.block_stack = append(s.block_stack, s.getWord())
}

func (s *LatexListener) ExitItemize(ctx *latex_parser.ItemizeContext) {
	for _, item := range s.block_stack {
		s.Note = append(s.Note, fmt.Sprintf("* %s", item))
	}

	s.block_stack = nil
}

func (s *LatexListener) ExitEnumerate(ctx *latex_parser.EnumerateContext) {
	for i, item := range s.block_stack {
		s.Note = append(s.Note, fmt.Sprintf("%d. %s", i+1, item))
	}

	s.block_stack = nil
}

func (s *LatexListener) ExitVerbatim_word(ctx *latex_parser.Verbatim_wordContext) {
	s.verbatim_content_stack = append(s.verbatim_content_stack, s.getWord())
}

func (s *LatexListener) ExitVerbatim_symbol(ctx *latex_parser.Verbatim_symbolContext) {
	s.verbatim_content_stack = append(s.verbatim_content_stack, ctx.GetText())
}

func (s *LatexListener) ExitVerbatim_linebreak(ctx *latex_parser.Verbatim_linebreakContext) {
	s.verbatim_content_stack = append(s.verbatim_content_stack, ctx.GetText())
}

func (s *LatexListener) ExitVerbatim_line(ctx *latex_parser.Verbatim_lineContext) {
	s.block_stack = append(s.block_stack, strings.Join(s.verbatim_content_stack, ""))

	s.verbatim_content_stack = nil
}

func (s *LatexListener) EnterVerbatim(ctx *latex_parser.VerbatimContext) {
	s.is_verbatim_block = true
}

func (s *LatexListener) ExitVerbatim(ctx *latex_parser.VerbatimContext) {
	s.Note = append(s.Note, "```")

	for _, item := range s.block_stack {
		s.Note = append(s.Note, item)
	}

	s.Note = append(s.Note, "```")

	s.block_stack = nil

	s.is_verbatim_block = false
}

func latex_to_note(latex_note []string) types.Note {
	// Setup the input, replicating a text file (lines ending with newline)
	is := antlr.NewInputStream(strings.Join(latex_note, "\n"))

	// Create the Lexer
	lexer := latex_parser.NewLatexLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := latex_parser.NewLatexParser(stream)

	var listener LatexListener

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Latex())

	return types.Note{
		listener.Title,
		listener.Url,
		listener.Created,
		listener.Updated,
		listener.Note,
	}
}

func process_latex_file(file_path string) []types.Note {
	fmt.Println("Processing " + file_path)
	f, err := os.Open(file_path)

	if err != nil {
		log.Fatalf("Error opening file %s: %v", file_path, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	is_note := false
	cur_note := make([]string, 0, 20)
	notes := make([]types.Note, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !is_note && strings.HasPrefix(line, "\\textbf{Title:}") {
			is_note = true
		}
		if is_note {
			if strings.HasPrefix(line, "\\hrulefill") {
				notes = append(notes, latex_to_note(cur_note))

				cur_note = nil
				is_note = false
				continue
			}
			cur_note = append(cur_note, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading %s: %v", file_path, err)
	}

	return notes
}

func Process_files(folder_path string, extension string) []FileNotes {
	files, err := os.ReadDir(folder_path)
	if err != nil {
		log.Fatal(err)
	}

	file_notes := make([]FileNotes, 0)

	for _, file := range files {
		if file.IsDir() {
			file_notes = append(file_notes, Process_files(filepath.Join(folder_path, file.Name()), extension)...)
		} else {
			switch extension {
			case "tex":
				file_path := filepath.Join(folder_path, file.Name())
				if filepath.Ext(file_path) == "."+extension {
					file_notes = append(file_notes, FileNotes{file_path, process_latex_file(file_path)})
				}
				break
			}
		}
	}

	return file_notes
}
