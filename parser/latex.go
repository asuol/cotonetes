package parser

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"path/filepath"
	"strings"
	"regexp"
)

type Note struct {
	Title string
	Url string
	Created_date string
	Updated_date string
	Text []string
}

type FileNotes struct {
	File_path string
	Notes []Note
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

func latex_note_content_to_txt(latex_note []string) []string {
	note := make([]string, 0)

	command_re := regexp.MustCompile(`\\[^{]*{([^}]*)}`)
	newline_re := regexp.MustCompile(`\\\\`)
	special_re := regexp.MustCompile(`\\(.)`)
	item_re := regexp.MustCompile(`\\item (.*)`)
	is_itemize := false
	is_enumerate := false
	count := 0

	for _, line := range latex_note {
		// replace blocks
		// grep -rPo "begin{.*}" . | awk -F: '{print $2}' | sort -u
		if line == `\begin{enumerate}` {
			is_enumerate = true
			count = 0
			continue
		} else if line == `\begin{itemize}` {
			is_itemize = true
			continue
		} else if line == `\begin{verbatim}` {
			continue
		}

		if line == `\end{enumerate}` {
			is_enumerate = false
			continue
		} else if line == `\end{itemize}` {
			is_itemize = false
			continue
		} else if line == `\end{verbatim}` {
			continue
		}

		if line == "" {
			continue
		}

		if is_enumerate {
			count++
			line = item_re.ReplaceAllString(line, fmt.Sprintf("%d. %s", count, "$1"))
		} else if is_itemize {
			line = item_re.ReplaceAllString(line, "* $1")
		}

		// replace any \command{text} with text
		line = command_re.ReplaceAllString(line, "$1")
		// replace any \\ with \n
		line = newline_re.ReplaceAllString(line, "\n")
		// replace any \<symbol> with <symbol>
		line = special_re.ReplaceAllString(line, "$1")

		note = append(note, line)
	}

	return note
}

func process_latex_file(file_path string) []Note {
	fmt.Println("Processing " + file_path)
	f, err := os.Open(file_path)

	if err != nil {
		log.Fatalf("Error opening file %s: %v", file_path, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	is_note := false
	cur_note := make([]string, 0, 20)
	notes := make([]Note, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !is_note && strings.HasPrefix(line, "\\textbf{Title:}") {
			is_note = true
		}
		if is_note {
			if strings.HasPrefix(line, "\\hrulefill") {
				notes = append(notes, Note{
					latex_note_title_to_txt(cur_note[0]),
					latex_note_url_to_txt(cur_note[1]),
					latex_note_title_to_txt(cur_note[2]),
					latex_note_title_to_txt(cur_note[3]),
					latex_note_content_to_txt(cur_note[4:len(cur_note)-1]),
				})

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
			return Process_files(filepath.Join(folder_path, file.Name()), extension)
		} else {
			switch extension {
				case "tex":
					file_path := filepath.Join(folder_path, file.Name())
					file_notes = append(file_notes, FileNotes{file_path, process_latex_file(file_path)})
			}
		}
	}

	return file_notes
}
