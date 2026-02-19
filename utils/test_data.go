package utils

import (
	"cotonetes/types"
)

var TdTextOnly = TestInput{
	types.Note{
		"Sample title",
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{"Sample line"},
	},
	types.Note{
		"Sample title",
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{"Sample line"},
	},
}

var TdTitleSpecialChars = TestInput{
	types.Note{
		`\&\#\%\_\$\^`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{"Sample line"},
	},
	types.Note{
		`&#%_$^`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{"Sample line"},
	},
}

var TdNoteBoldText = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{`some text with \textbf{bold content} for test. Real \textbf{bold}`},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{`some text with **bold content** for test. Real **bold**`},
	},
}

var TdNoteUrl = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with \url{link text} for test`,
			`another line with another \url{link}`,
		},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with [link text](link text) for test`,
			`another line with another [link](link)`,
		},
	},
}

var TdNoteNewline = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text for test\\`,
			``,
			`another line with another`,
			`\\`,
			``,
		},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text for test`,
			``,
			`another line with another`,
			``,
		},
	},
}

var TdNoteItemize = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with \url{link text} for test`,
			`\begin{itemize}`,
			``,
			``,
			``,
			`\item item 1`,
			`\item item 2`,
			``,
			`\item item 3`,
			`\end{itemize}`,
		},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with [link text](link text) for test`,
			``,
			`* item 1`,
			`* item 2`,
			``,
			`* item 3`,
		},
	},
}

var TdNoteEnumerate = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with \url{link text} for test`,
			`\begin{enumerate}`,
			`\item item 1`,
			`\item item 2`,
			`\item item 3`,
			`\end{enumerate}`,
		},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with [link text](link text) for test`,
			`1. item 1`,
			`2. item 2`,
			`3. item 3`,
		},
	},
}

var TdNoteVerbatim = TestInput{
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with \url{link text} for test`,
			`\begin{verbatim}`,
			``,
			` #######&>>>>>`,
			`\\`,
			``,
			` some regular text`,
			``,
			``,
			`\end{verbatim}`,
		},
	},
	types.Note{
		`Sample title`,
		"Sample url",
		"Sample create date",
		"Sample update date",
		[]string{
			`some text with [link text](link text) for test`,
			"```",
			``,
			` #######&>>>>>`,
			``,
			` some regular text`,
			``,
			``,
			"```",
		},
	},
}
