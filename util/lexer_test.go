package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestTokenize(t *testing.T) {

	lex := NewLexer(`  foo 123  " abc" 'def ' `)
	lex.Tokenize()

	log.Println(lex.Tokens)
	require.Equal(t, 4, len(lex.Tokens))
	assert.Equal(t, "foo", lex.Tokens[0])
	assert.Equal(t, "123", lex.Tokens[1])
	assert.Equal(t, " abc", lex.Tokens[2])
	assert.Equal(t, "def ", lex.Tokens[3])
}

func TestTokenize2(t *testing.T) {

	lex := NewLexer(` state 'STARTUP' `)
	lex.Tokenize()

	log.Println(lex.Tokens)
	require.Equal(t, 2, len(lex.Tokens))
	assert.Equal(t, "state", lex.Tokens[0])
	assert.Equal(t, "STARTUP", lex.Tokens[1])
}
