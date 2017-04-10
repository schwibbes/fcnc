package task

import (
	"bufio"
	"bytes"
	"github.com/schwibbes/fcnc/action"
	"github.com/schwibbes/fcnc/query"
	"github.com/schwibbes/fcnc/util"
	"log"
	"regexp"
	"strings"
)

type Parser struct {
	// in
	asByes []byte

	// out
	Result  []Task
	mode    parserMode
	taskNum int
	linNum  int
}

type parserMode int

const (
	TASK parserMode = 1 + iota
	VAR
	COND
	ACT
	ASSERT
	ERR
)

func (self *Parser) parse() (result []Task, err error) {
	buf := bytes.NewBuffer(self.asByes)
	r := bufio.NewReader(buf)

	for {
		line, prefix, err := r.ReadLine()
		if err != nil || prefix {
			break
		}
		self.parseLine(string(line))
	}
	return
}

func (self *Parser) parseLine(line string) {

	for _, x := range line {
		switch {
		case regexp.MustCompile(`\s`).MatchString(string(x)):
			continue
		case x == '#':
			self.mode = TASK
			self.Result = append(self.Result, *NewTask(line))
			self.taskNum++
			log.Printf("#Task %v", self.taskNum)
			return
		case x == '%':
			self.mode = VAR
			return
		case x == '?':
			self.mode = COND
			return
		case x == '!':
			self.mode = ACT
			return
		case x == '+':
			self.mode = ASSERT
			return
		case x == '-':
			self.addToCurrent(line)
			return
		default:
			self.mode = ERR
			return
		}
	}

}

func (self *Parser) addToCurrent(line string) {
	current := &self.Result[len(self.Result)-1]
	line = removePrefix(line)
	log.Printf("processing task(%v) %v      with line:%v", len(self.Result)-1, current, line)
	switch self.mode {
	case VAR:
		log.Println("new-var", line)
		p := parseProperty(line)
		current.Vars[p[0]] = p[1]
	case COND:
		line = cleanup(line)
		log.Println("new-cond", line)
		lexer := util.NewLexer(line)
		lexer.Tokenize()
		log.Println("yy " + strings.Join(lexer.Tokens, ","))
		current.Conditions = append(current.Conditions, query.NewQuery(lexer.Tokens))
	case ACT:
		line = cleanup(line)
		log.Println("new-act", line)
		lexer := util.NewLexer(line)
		lexer.Tokenize()
		log.Println(strings.Join(lexer.Tokens, ","))
		current.Actions = append(current.Actions, action.NewAction(lexer.Tokens))
	case ASSERT:
		line = cleanup(line)
		log.Println("new-assert", line)
		lexer := util.NewLexer(line)
		lexer.Tokenize()
		log.Println(strings.Join(lexer.Tokens, ","))
		current.Asserts = append(current.Asserts, query.NewQuery(lexer.Tokens))
	}
}

func parseProperty(expr string) (out []string) {
	out = strings.Split(expr, "=")
	if 2 != len(out) {
		panic("cannot parse property from line: " + expr)
	}
	out[0] = cleanup(out[0])
	out[1] = cleanup(out[1])
	return
}

func cleanup(in string) string {
	return strings.Trim(in, " \t'\"")
}

func removePrefix(input string) string {
	return regexp.MustCompile(`\s\-`).ReplaceAllString(input, "")
}
