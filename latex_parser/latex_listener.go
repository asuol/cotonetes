// Code generated from Latex.g4 by ANTLR 4.13.2. DO NOT EDIT.

package latex_parser // Latex
import "github.com/antlr4-go/antlr/v4"


// LatexListener is a complete listener for a parse tree produced by LatexParser.
type LatexListener interface {
	antlr.ParseTreeListener

	// EnterLatex is called when entering the latex production.
	EnterLatex(c *LatexContext)

	// EnterNote_title is called when entering the note_title production.
	EnterNote_title(c *Note_titleContext)

	// EnterNote_url is called when entering the note_url production.
	EnterNote_url(c *Note_urlContext)

	// EnterNote_created is called when entering the note_created production.
	EnterNote_created(c *Note_createdContext)

	// EnterNote_updated is called when entering the note_updated production.
	EnterNote_updated(c *Note_updatedContext)

	// EnterNote_text is called when entering the note_text production.
	EnterNote_text(c *Note_textContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterLine_break is called when entering the line_break production.
	EnterLine_break(c *Line_breakContext)

	// EnterEmpty_line is called when entering the empty_line production.
	EnterEmpty_line(c *Empty_lineContext)

	// EnterEscaped_word is called when entering the escaped_word production.
	EnterEscaped_word(c *Escaped_wordContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterEscaped is called when entering the escaped production.
	EnterEscaped(c *EscapedContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterPunctuation is called when entering the punctuation production.
	EnterPunctuation(c *PunctuationContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterWs is called when entering the ws production.
	EnterWs(c *WsContext)

	// EnterVerbatim_word is called when entering the verbatim_word production.
	EnterVerbatim_word(c *Verbatim_wordContext)

	// EnterVerbatim_symbol is called when entering the verbatim_symbol production.
	EnterVerbatim_symbol(c *Verbatim_symbolContext)

	// EnterVerbatim_linebreak is called when entering the verbatim_linebreak production.
	EnterVerbatim_linebreak(c *Verbatim_linebreakContext)

	// EnterVerbatim_line is called when entering the verbatim_line production.
	EnterVerbatim_line(c *Verbatim_lineContext)

	// EnterBlock_line is called when entering the block_line production.
	EnterBlock_line(c *Block_lineContext)

	// EnterBlock_item is called when entering the block_item production.
	EnterBlock_item(c *Block_itemContext)

	// EnterItemize is called when entering the itemize production.
	EnterItemize(c *ItemizeContext)

	// EnterEnumerate is called when entering the enumerate production.
	EnterEnumerate(c *EnumerateContext)

	// EnterVerbatim is called when entering the verbatim production.
	EnterVerbatim(c *VerbatimContext)

	// ExitLatex is called when exiting the latex production.
	ExitLatex(c *LatexContext)

	// ExitNote_title is called when exiting the note_title production.
	ExitNote_title(c *Note_titleContext)

	// ExitNote_url is called when exiting the note_url production.
	ExitNote_url(c *Note_urlContext)

	// ExitNote_created is called when exiting the note_created production.
	ExitNote_created(c *Note_createdContext)

	// ExitNote_updated is called when exiting the note_updated production.
	ExitNote_updated(c *Note_updatedContext)

	// ExitNote_text is called when exiting the note_text production.
	ExitNote_text(c *Note_textContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitLine_break is called when exiting the line_break production.
	ExitLine_break(c *Line_breakContext)

	// ExitEmpty_line is called when exiting the empty_line production.
	ExitEmpty_line(c *Empty_lineContext)

	// ExitEscaped_word is called when exiting the escaped_word production.
	ExitEscaped_word(c *Escaped_wordContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitEscaped is called when exiting the escaped production.
	ExitEscaped(c *EscapedContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitPunctuation is called when exiting the punctuation production.
	ExitPunctuation(c *PunctuationContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitWs is called when exiting the ws production.
	ExitWs(c *WsContext)

	// ExitVerbatim_word is called when exiting the verbatim_word production.
	ExitVerbatim_word(c *Verbatim_wordContext)

	// ExitVerbatim_symbol is called when exiting the verbatim_symbol production.
	ExitVerbatim_symbol(c *Verbatim_symbolContext)

	// ExitVerbatim_linebreak is called when exiting the verbatim_linebreak production.
	ExitVerbatim_linebreak(c *Verbatim_linebreakContext)

	// ExitVerbatim_line is called when exiting the verbatim_line production.
	ExitVerbatim_line(c *Verbatim_lineContext)

	// ExitBlock_line is called when exiting the block_line production.
	ExitBlock_line(c *Block_lineContext)

	// ExitBlock_item is called when exiting the block_item production.
	ExitBlock_item(c *Block_itemContext)

	// ExitItemize is called when exiting the itemize production.
	ExitItemize(c *ItemizeContext)

	// ExitEnumerate is called when exiting the enumerate production.
	ExitEnumerate(c *EnumerateContext)

	// ExitVerbatim is called when exiting the verbatim production.
	ExitVerbatim(c *VerbatimContext)
}
