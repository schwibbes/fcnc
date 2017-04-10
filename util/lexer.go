package util

import (
	"log"
)

type Lexer struct {
	Input  string
	Tokens []string

	opening      rune
	lastRune     rune
	currentToken string
}

const (
	SINGLE_QUOTE = '\''
	DOUBLE_QUOTE = '"'
	WS           = ' '
)

const (
	UNDEF = -1
)

func NewLexer(input string) (res Lexer) {
	res = Lexer{}
	res.Input = input
	res.Tokens = make([]string, 0)

	res.currentToken = ""
	res.lastRune = UNDEF
	res.opening = UNDEF
	return
}

func (self *Lexer) Tokenize() {

	for _, x := range self.Input {
		log.Println("line: ", string(x))
		switch {
		case isQuote(x) && (self.opening == UNDEF):
			log.Println("open-quote")
			self.opening = x
		case isQuote(x) && isQuote(self.opening):
			log.Println("close-quote and push")
			self.opening = UNDEF
			self.push()
		case (x == WS) && self.lastRune == WS && (self.opening == UNDEF):
			log.Println("skip multiple-ws")
		case (x == WS) && self.lastRune != WS && (self.opening == UNDEF) && (len(self.currentToken) == 0):
			log.Println("first-ws -> nothing-to-push: ")
		case (x == WS) && self.lastRune != WS && (self.opening == UNDEF):
			log.Println("first-ws -> push: ")
			self.push()
		default:
			log.Println("consume: ", string(x))
			self.currentToken += string(x)

		}
		self.lastRune = x
	}

	if len(self.currentToken) > 0 {
		self.push()
	}

	return
}

func (self *Lexer) push() {
	self.Tokens = append(self.Tokens, self.currentToken)
	self.currentToken = ""

}

func isQuote(r rune) bool {
	return r == SINGLE_QUOTE || r == DOUBLE_QUOTE
}
