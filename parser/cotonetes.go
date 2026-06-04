package parser

import (
	"bufio"
	"cotonetes/markdown_parser"
	"cotonetes/types"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os"
	"regexp"
	"strings"
)

func escape_special_chars(line string) string {
	// $ and ^ chars are replaced via strings instead of regex
	special_re := regexp.MustCompile(`(&|#|%|_)`)

	line = special_re.ReplaceAllString(line, `\$1`)
	line = strings.ReplaceAll(line, `$`, `\$`)
	line = strings.ReplaceAll(line, `^`, `\^`)

	return line
}

type MarkdownListener struct {
	*markdown_parser.BaseMarkdownListener

	Note                   []string
	block_stack            []string
	word_stack             []string
	text_stack             []string
	url_title              string
	url_value              string
	verbatim_content_stack []string
	is_verbatim_block      bool
}

func (s *MarkdownListener) getWord() string {
	word := strings.Join(s.word_stack, "")

	s.word_stack = nil

	return word
}

func (s *MarkdownListener) EnterUrl(ctx *markdown_parser.UrlContext) {
	if s.word_stack != nil {
		s.text_stack = append(s.text_stack, s.getWord())
	}
}

func (s *MarkdownListener) ExitUrl_title(ctx *markdown_parser.Url_titleContext) {
	s.url_title = ctx.GetText()
}

func (s *MarkdownListener) ExitUrl_value(ctx *markdown_parser.Url_valueContext) {
	s.url_value = ctx.GetText()
}

func (s *MarkdownListener) ExitUrl(ctx *markdown_parser.UrlContext) {
	if s.url_value == s.url_title {
		s.text_stack = append(s.text_stack, `\url{`+s.url_title+`}`)

	} else {
		s.text_stack = append(s.text_stack, `\href{`+s.url_title+`}{`+s.url_value+`}`)
	}
}

func (s *MarkdownListener) EnterBold(ctx *markdown_parser.BoldContext) {
	if s.word_stack != nil {
		s.text_stack = append(s.text_stack, s.getWord())
	}
}

func (s *MarkdownListener) ExitBold(ctx *markdown_parser.BoldContext) {
	s.text_stack = append(s.text_stack, `\textbf{`+s.getWord()+`}`)
}

func (s *MarkdownListener) ExitText(ctx *markdown_parser.TextContext) {
	if s.word_stack != nil {
		s.text_stack = append(s.text_stack, s.getWord())
	}
	s.Note = append(s.Note, strings.Join(s.text_stack, ""))

	s.text_stack = nil
}

func (s *MarkdownListener) ExitSymbol(ctx *markdown_parser.SymbolContext) {
	if s.is_verbatim_block {
		s.word_stack = append(s.word_stack, ctx.GetText())
	} else {
		s.word_stack = append(s.word_stack, `\`+ctx.GetText())
	}
}

func (s *MarkdownListener) ExitLetter(ctx *markdown_parser.LetterContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitPunctuation(ctx *markdown_parser.PunctuationContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitBackslash(ctx *markdown_parser.BackslashContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitSquare(ctx *markdown_parser.SquareContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitRound(ctx *markdown_parser.RoundContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitNumber(ctx *markdown_parser.NumberContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitWs(ctx *markdown_parser.WsContext) {
	s.word_stack = append(s.word_stack, ctx.GetText())
}

func (s *MarkdownListener) ExitEmpty_line(ctx *markdown_parser.Empty_lineContext) {
	note_count := len(s.Note)

	if note_count > 0 {
		s.Note[note_count-1] += `\\`
	}

	s.Note = append(s.Note, "")
}

func (s *MarkdownListener) ExitItem_line(ctx *markdown_parser.Item_lineContext) {
	s.block_stack = append(s.block_stack, s.getWord())
}

func (s *MarkdownListener) ExitNumber_line(ctx *markdown_parser.Number_lineContext) {
	s.block_stack = append(s.block_stack, s.getWord())
}

func (s *MarkdownListener) ExitVerbatim_line(ctx *markdown_parser.Verbatim_lineContext) {
	s.block_stack = append(s.block_stack, s.getWord())
}

func (s *MarkdownListener) ExitItemize(ctx *markdown_parser.ItemizeContext) {
	s.Note = append(s.Note, fmt.Sprintf(`\begin{itemize}`))
	for _, item := range s.block_stack {
		s.Note = append(s.Note, fmt.Sprintf("\\item %s", item))
	}
	s.Note = append(s.Note, fmt.Sprintf(`\end{itemize}`))

	s.block_stack = nil
}

func (s *MarkdownListener) ExitEnumerate(ctx *markdown_parser.EnumerateContext) {
	s.Note = append(s.Note, fmt.Sprintf(`\begin{enumerate}`))
	for _, item := range s.block_stack {
		s.Note = append(s.Note, fmt.Sprintf("\\item %s", item))
	}
	s.Note = append(s.Note, fmt.Sprintf(`\end{enumerate}`))

	s.block_stack = nil
}

func (s *MarkdownListener) EnterVerbatim(ctx *markdown_parser.VerbatimContext) {
	s.is_verbatim_block = true
}

func (s *MarkdownListener) ExitVerbatim(ctx *markdown_parser.VerbatimContext) {
	s.Note = append(s.Note, fmt.Sprintf(`\begin{verbatim}`))
	for _, line := range s.block_stack {
		s.Note = append(s.Note, line)
	}
	s.Note = append(s.Note, fmt.Sprintf(`\end{verbatim}`))

	s.block_stack = nil

	s.is_verbatim_block = false
}

func markdown_note_to_latex(markdown_note []string) []string {
	// Setup the input, replicating a text file (lines ending with newline)
	is := antlr.NewInputStream(strings.Join(markdown_note, "\n"))

	// Create the Lexer
	lexer := markdown_parser.NewMarkdownLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := markdown_parser.NewMarkdownParser(stream)

	var listener MarkdownListener

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Markdown())

	return listener.Note
}

func Export_to_latex_file(file_path string, section_name string, sub_section_index int, note_list []types.Note) error {
	fmt.Println("Processing " + file_path)

	var f *os.File
	var err error

	if f, err = os.Create(file_path); err != nil {
		log.Fatalf("Error creating file %s: %v", file_path, err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	if _, err = writer.WriteString(`\` + strings.Repeat("sub", sub_section_index) + `section{` + escape_special_chars(section_name) + `}` + "\n\n"); err != nil {
		return err
	}

	for _, note := range note_list {
		if _, err = writer.WriteString(`\textbf{Title:} ` + escape_special_chars(note.Title) + `\\` + "\n"); err != nil {
			return err
		}

		if _, err = writer.WriteString(`\textbf{URL:} \url{` + note.Url + `}\\` + "\n"); err != nil {
			return err
		}

		if _, err = writer.WriteString(`\textbf{Created:} ` + note.Created_date + `\\` + "\n"); err != nil {
			return err
		}

		if _, err = writer.WriteString(`\textbf{Last Updated:} ` + note.Updated_date + `\\` + "\n" + `\\` + "\n"); err != nil {
			return err
		}

		for _, content := range markdown_note_to_latex(note.Text) {
			if _, err = writer.WriteString(content + "\n"); err != nil {
				return err
			}
		}

		if _, err = writer.WriteString(`\hrulefill` + "\n" + `\\` + "\n\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}
