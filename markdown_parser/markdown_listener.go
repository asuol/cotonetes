// Code generated from Markdown.g4 by ANTLR 4.13.2. DO NOT EDIT.

package markdown_parser // Markdown
import "github.com/antlr4-go/antlr/v4"


// MarkdownListener is a complete listener for a parse tree produced by MarkdownParser.
type MarkdownListener interface {
	antlr.ParseTreeListener

	// EnterMarkdown is called when entering the markdown production.
	EnterMarkdown(c *MarkdownContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterEmpty_line is called when entering the empty_line production.
	EnterEmpty_line(c *Empty_lineContext)

	// EnterBold is called when entering the bold production.
	EnterBold(c *BoldContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterPunctuation is called when entering the punctuation production.
	EnterPunctuation(c *PunctuationContext)

	// EnterBackslash is called when entering the backslash production.
	EnterBackslash(c *BackslashContext)

	// EnterSquare is called when entering the square production.
	EnterSquare(c *SquareContext)

	// EnterRound is called when entering the round production.
	EnterRound(c *RoundContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterWs is called when entering the ws production.
	EnterWs(c *WsContext)

	// EnterUrl_title is called when entering the url_title production.
	EnterUrl_title(c *Url_titleContext)

	// EnterUrl_value is called when entering the url_value production.
	EnterUrl_value(c *Url_valueContext)

	// EnterItem_line is called when entering the item_line production.
	EnterItem_line(c *Item_lineContext)

	// EnterNumber_line is called when entering the number_line production.
	EnterNumber_line(c *Number_lineContext)

	// EnterVerbatim_line is called when entering the verbatim_line production.
	EnterVerbatim_line(c *Verbatim_lineContext)

	// EnterItemize is called when entering the itemize production.
	EnterItemize(c *ItemizeContext)

	// EnterEnumerate is called when entering the enumerate production.
	EnterEnumerate(c *EnumerateContext)

	// EnterVerbatim is called when entering the verbatim production.
	EnterVerbatim(c *VerbatimContext)

	// ExitMarkdown is called when exiting the markdown production.
	ExitMarkdown(c *MarkdownContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitEmpty_line is called when exiting the empty_line production.
	ExitEmpty_line(c *Empty_lineContext)

	// ExitBold is called when exiting the bold production.
	ExitBold(c *BoldContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitPunctuation is called when exiting the punctuation production.
	ExitPunctuation(c *PunctuationContext)

	// ExitBackslash is called when exiting the backslash production.
	ExitBackslash(c *BackslashContext)

	// ExitSquare is called when exiting the square production.
	ExitSquare(c *SquareContext)

	// ExitRound is called when exiting the round production.
	ExitRound(c *RoundContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitWs is called when exiting the ws production.
	ExitWs(c *WsContext)

	// ExitUrl_title is called when exiting the url_title production.
	ExitUrl_title(c *Url_titleContext)

	// ExitUrl_value is called when exiting the url_value production.
	ExitUrl_value(c *Url_valueContext)

	// ExitItem_line is called when exiting the item_line production.
	ExitItem_line(c *Item_lineContext)

	// ExitNumber_line is called when exiting the number_line production.
	ExitNumber_line(c *Number_lineContext)

	// ExitVerbatim_line is called when exiting the verbatim_line production.
	ExitVerbatim_line(c *Verbatim_lineContext)

	// ExitItemize is called when exiting the itemize production.
	ExitItemize(c *ItemizeContext)

	// ExitEnumerate is called when exiting the enumerate production.
	ExitEnumerate(c *EnumerateContext)

	// ExitVerbatim is called when exiting the verbatim production.
	ExitVerbatim(c *VerbatimContext)
}
