grammar Markdown;

markdown
    : (text | block | empty_line)* EOF
    ;

text
    : (word | bold | url)+ NEWLINE?
    ;

empty_line
    : NEWLINE+
    ;

bold
    : '**' word+ '**'
    ;

url
    : '[' url_title+ '](' url_value+ ')'
    ;

word
    : SYMBOL          #symbol
    | LETTER          #letter
    | PUNCTUATION     #punctuation
    | BACKSLASH       #backslash
    | SQUARE_PARENS   #square
    | ROUND_PARENS    #round
    | NUMBER          #number
    | WS              #ws
    ;

url_title
    : (
          SYMBOL
        | LETTER
        | PUNCTUATION
        | ROUND_PARENS
        | NUMBER
        | WS
        )+
    ;

url_value
    : (
          SYMBOL
        | LETTER
        | PUNCTUATION
        | SQUARE_PARENS
        | NUMBER
        | WS
        )+
    ;

item_line
    : '*' WS* word+ NEWLINE*
    ;

number_line
    : NUMBER_ITEM WS* word+ NEWLINE*
    ;

verbatim_line
    : word+ NEWLINE
    | NEWLINE
    ;

block
    : item_line+                                    #itemize
    | number_line+                                  #enumerate
    | '```' NEWLINE? verbatim_line+ '```' NEWLINE?  #verbatim
    ;

LETTER
    : [\p{L}]+
    ;

NUMBER_ITEM
    : NUMBER '.'
    ;

PUNCTUATION
    // all the punctuation that does not require escaping in latex
    // except the parens, which interfere with the markdown URL syntax
    : (
          '.'
        | ','
        | ';'
        | ':'
        | '!'
        | '?'
        | '-'
        | '\''
        | '"'
        | '/'
        | '@'
        | '*'
        | '+'
        | '='
        )+
    ;

BACKSLASH
    : '\\'
    ;

SQUARE_PARENS
    : '['
    | ']'
    ;

ROUND_PARENS
    : '('
    | ')'
    ;

SYMBOL
    : '#'
    | '$'
    | '%'
    | '&'
    | '_'
    | '^'
    | '>'
    | '<'
    ;

NUMBER
    : '-'? INT ('.' [0-9]+)?
    ;

fragment INT
    // integer part forbids leading 0s (e.g. `01`)
    : '0'
    | [1-9] [0-9]*
    ;

// Required to detect end of blocks (bullet list or enumerations), as well as to preserve lines
NEWLINE
    : '\n'
    ;

WS
    : [ \t]+
    ;

CR
    : '\r' -> skip
    ;
