# ANTLR

## Grammar development

To test the grammar during its development, use the `antlr4-parse` utility. You can use a dockerized version by going to the `antlr-tools` directory at the same level of this README, and then:

1. make build 
2. Create a folder `folder` to be used as a container volume
3. make sh
4. Assuming a `Latex.g4` grammar file, run as `antlr4-parse Latex.g4 latex -trace`
5. Type a sample text and then type `<CTRL>+D` to see the trace of the parsing of the input text using the provided grammar

Some relevant references to create grammars:

* https://tomassetti.me/best-practices-for-antlr-parsers/
* https://tomassetti.me/antlr-mega-tutorial/
* Existing grammars, such as https://github.com/antlr/codebuff/blob/master/grammars/org/antlr/codebuff/Java.g4
* https://yalin.dev/blog/introduction-to-dsls-using-antlr/

## Generating parser

With the grammar file ready, generate the parser using the following docker command:

1. git clone https://github.com/antlr/antlr4
2. Build the docker image in the `docker` directory, as per the README.md in that folder

Then, on the folder where you have the grammar file run the following to generate the parser (on a `latex_parser` folder/golang package:

`docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work antlr/antlr4 -Dlanguage=Go -o latex_parser -package latex_parser Latex.g4`

