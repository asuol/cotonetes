grammar Latex;

latex
    : note_title note_url note_created note_updated line_break NEWLINE* note_text EOF
    ;

note_title
    : '\\textbf{Title:}' WS* word+ '\\\\' NEWLINE?
    ;

note_url
    // See comment to url_tag label below
    : '\\textbf{URL:} \\url{' verbatim_content+ '}' '\\\\' NEWLINE?
    ;

note_created
    : '\\textbf{Created:}' WS* word+ '\\\\' NEWLINE?
    ;

note_updated
    : '\\textbf{Last Updated:}' WS* word+ '\\\\' NEWLINE?
    ;

note_text
    : (text | block | line_break | empty_line)* 
    ;

text
    : (tag | word)+ NEWLINE?
    ;

line_break
    : '\\\\' WS*
    ;

empty_line
    : NEWLINE+
    ;

escaped_word
    : content=LETTER+ space=WS*
    | content=(SYMBOL | '{' | '}') space=WS*
    ;

composite_tag_separator
    : '}{'
    ;

tag
    // url tag needs to account for SYMBOL in addition to regular words, which cannot be escaped in the latex sense with backslash or dedicated symbol command, especially because 
    // 1. Latex does not really need to escape the characters: at worst it ignores them and does not produce them
    // 2. Symbols such as _ and % are valid URL characters, so url tag content is verbatim content
    : '\\url{' verbatim_content+ '}'                                  #url_tag
    | '\\' name=LETTER+ '{' word+ '}'                                 #simple_tag
    | '\\' name=LETTER+ '{' word+ composite_tag_separator word+ '}'   #composite_tag
    ;

word
    // content may be empty, this way it matches the backslash char
    : '\\' escaped_word?    #escaped
    | LETTER          #letter
    | (PUNCTUATION | EXTENDED_UNICODE)    #punctuation
    | NUMBER          #number
    | WS              #ws
    ;

verbatim_content
    : word            #verbatim_word
    | SYMBOL          #verbatim_symbol
    | SPECIAL         #verbatim_special
    | line_break      #verbatim_linebreak
    ;

curly_parens
    : ( '{' | '}' )
    ;

verbatim_block_content
    : verbatim_content
    | curly_parens
    ;

verbatim_line
    : verbatim_block_content* NEWLINE
    ;

block_content
    : tag
    | word
    | curly_parens
    ;

block_line
    : block_content+ NEWLINE+
    ; 

block_item
    : WS* '\\item' WS* block_line
    ;

block
    : '\\begin{itemize}' NEWLINE* block_item* '\\end{itemize}' NEWLINE?       #itemize
    | '\\begin{enumerate}' NEWLINE* block_item* '\\end{enumerate}' NEWLINE?   #enumerate
    | '\\begin{verbatim}' NEWLINE? verbatim_line* '\\end{verbatim}' NEWLINE?  #verbatim
    ;

LETTER
    : [\p{L}]+
    ;

SPECIAL
    // all chars that require a dedicated latex command
    : '~'
    | '^'
    | '|'
    ;

SYMBOL
    // all chars that can be directly escaped in latex
    : '#'
    | '$'
    | '%'
    | '&'
    | '>'
    | '<'
    | '€'
    | '_'
    ;

NUMBER
    : '-'? INT ('.' [0-9]+)?
    ;

fragment INT
    // integer part forbids leading 0s (e.g. `01`)
    : '0'
    | [1-9] [0-9]*
    ;

// Required to detect empty lines, which counts as a paragraph in latex
NEWLINE
    : '\n'
    ;

WS
    : [ \t]+
    ;

CR
    : '\r' -> skip
    ;

PUNCTUATION
    // all the punctuation that does not require escaping in latex
    // NOTE: this rule intends to match all chars that are not mentioned in the above lexer rules, including the string literals in the parser rules (antlr creates lexer rules for those inderectly) (order of precedence)
    // Punctuation except curly braces. These are to be handled as literals in the parser rules
    : [\p{Pc}]
    | [\p{Pd}]
    | [\p{Pi}]
    | [\p{Pf}]
    | [\p{Po}]
    | ( '(' | ')' | '[' | ']' )
    | [\p{S}]
    | [\p{Z}]
    ;

EXTENDED_UNICODE
    : [\u0001-\uFFFF]
    ;
