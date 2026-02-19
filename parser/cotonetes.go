package parser

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"cotonetes/types"
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

func markdown_note_to_latex(markdown_note []string) []string {
	note := make([]string, 0)

	bold_re := regexp.MustCompile(`\*\*([^*]*)\*\*`)
	url_re := regexp.MustCompile(`\.*\[([^]]*)\]\(([^)]*)\)\.*`)
	newline_re := regexp.MustCompile(`^ *$`)
	item_re := regexp.MustCompile(`^ *\* +(.*)$`)
	num_re := regexp.MustCompile(`^ *[1-9]+\. +(.*)$`)
	is_itemize := false
	is_enumerate := false
	is_verbatim := false

	// Some notes have sections written as e.g. "1. something", followed by 1 or more paragraphs before a "2."
	// These are not a numeric list, so we buffer potential enumerate blocks until we confirm if it is indeed a enumerate block
	// otherwise leave the "1." tokens as is
	is_num_buf := false
	num_buffer := make([]string, 0)

	for index, line := range markdown_note {
		// replace blocks
		if !is_enumerate && num_re.MatchString(line) {
			is_enumerate = true
			is_num_buf = true
			num_buffer = append(num_buffer, `\begin{enumerate}` + "\n")
		} else if !is_itemize && item_re.MatchString(line) {
			is_itemize = true
			note = append(note, `\begin{itemize}` + "\n")
		} else if !is_verbatim && line == "```" {
			is_verbatim = true
			note = append(note, `\begin{verbatim}` + "\n")
			continue
		}

		if (is_enumerate || is_itemize) && newline_re.MatchString(line) {
			continue
		}

		if is_enumerate && !num_re.MatchString(line) {
			is_enumerate = false
			// if the "enum" block has ended during the buffering phase, then it is a fake enumerate block
			if is_num_buf {
				note = append(note, num_buffer[1])
				note = append(note, `\\` + "\n")
				num_buffer = nil
				is_num_buf = false
			} else {
				note = append(note, "\n" + `\end{enumerate}` + "\n")
			}
		} else if is_itemize && !item_re.MatchString(line) {
			is_itemize = false
			note = append(note, "\n" + `\end{itemize}` + "\n")
		} else if is_verbatim && line == "```" {
			is_verbatim = false
			note = append(note, "\n" + `\end{verbatim}` + "\n")
			if len(markdown_note) == index + 1 {
				// if this is the last line on the note, return here to avoid adding a \\ at the end of a verbatim block (which would be invalid latex syntax)
				return note
			} else {
				continue
			}
		}

		if is_enumerate {
			if !is_num_buf {
				line = num_re.ReplaceAllString(line, fmt.Sprintf("\t\\item %s", "$1"))
			}
		} else if is_itemize {
			line = item_re.ReplaceAllString(line, fmt.Sprintf("\t\\item %s", "$1"))
		} else if is_verbatim {
			note = append(note, line)
			continue
		}

		bold_matches := bold_re.FindAllString(line, -1)

		if len(bold_matches) > 0 {
			non_bold_matches := bold_re.Split(line, -1)

			line = non_bold_matches[0]
			non_match_index := 1

			for _, match := range bold_matches {
				bold_text := bold_re.FindStringSubmatch(match)[1]

				line = line + `\textbf{` + bold_text + `}` + non_bold_matches[non_match_index]

				if non_match_index < len(non_bold_matches) - 1 {
					non_match_index = non_match_index + 1
				}
			}
		}

		url_matches := url_re.FindAllString(line, -1)

		// we do not want to replace any symbols that are part of an URL
		line = escape_special_chars(line)

		if len(url_matches) > 0 {
			non_url_matches := url_re.Split(line, -1)

			line = non_url_matches[0]
			non_match_index := 1

			for _, match := range url_matches {
				url_parts := url_re.FindStringSubmatch(match)

				url_text := url_parts[1]
				url_url := url_parts[2]

				if url_text == url_url {
					line = line + `\url{` + url_url + `}` + non_url_matches[non_match_index]
				} else {
					line = line + `\url{` + url_url + `} (` + url_text + `)` + non_url_matches[non_match_index]
				}

				if non_match_index < len(non_url_matches) - 1 {
					non_match_index = non_match_index + 1
				}
			}
		}

		line = newline_re.ReplaceAllString(line, `\\` + "\n")
		line = strings.ReplaceAll(line, `\*`, `*`)

		if is_num_buf {
			num_buffer = append(num_buffer, line)
			// If the buffer has more than two lines (begin + 1. line), then it is a real enumerate block
			if len(num_buffer) > 2 {
				for _, num := range num_buffer {
					note = append(note, num_re.ReplaceAllString(num, fmt.Sprintf("\t\\item %s", "$1")))
				}
				is_num_buf = false
				num_buffer = nil

			}
		} else {
			note = append(note, line)
		}
	}

	if is_enumerate {
		// if the "enum" block has not ended during the buffering phase, then it is a real enumerate block
		if is_num_buf {
			note = append(note, num_re.ReplaceAllString(num_buffer[1], fmt.Sprintf("\t\\item %s", "$1")))
			is_num_buf = false
			num_buffer = nil
		}

		note = append(note, "\n" + `\end{enumerate}` + "\n")
	} else if is_itemize {
		note = append(note, "\n" + `\end{itemize}`)
	} 

	return note
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

	for _, note := range(note_list) {
		if _, err = writer.WriteString(`\textbf{Title:} ` + escape_special_chars(note.Title) + `\\` + "\n"); err != nil {
			return err
		}

		if _, err = writer.WriteString(`\textbf{URL:} \url{` + note.Url + `}\\` +  "\n"); err != nil {
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

		if _, err = writer.WriteString(`\hrulefill` + "\n" + `\\` + "\n\n"); err !=nil {
			return err
		}
	}

	return writer.Flush()
}
