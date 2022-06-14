//   Copyright (C) 2022 Ashwin Godbole
//
//   This file is part of simla.
//
//   simla is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   simla is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with simla. If not, see <https://www.gnu.org/licenses/>.

package parts

import (
	"fmt"
	"os"
)

// Lexer: Current state of the lexer
type Lexer struct {
	Source  string
	CurrPos int
	NextPos int
	char    byte
	line    int
}

// CreateLexer : to create a new lexer state and initialize it
func CreateLexer(source string) *Lexer {
	lexer := Lexer{
		Source:  source,
		CurrPos: 0,
		NextPos: 1,
		char:    source[0],
		line:    0,
	}
	return &lexer
}

// ReadChar : to read the next character
func (l *Lexer) ReadChar() {
	if l.NextPos >= len(l.Source) {
		l.char = 0
	} else {
		l.char = l.Source[l.NextPos]
	}
	l.CurrPos = l.NextPos
	l.NextPos++
}

// PeekChar : to peek at the next character
func (l *Lexer) PeekChar() byte {
	if l.NextPos >= len(l.Source) {
		return 0
	}
	return l.Source[l.NextPos]
}

func (l *Lexer) updateChar() {
	if ! (l.NextPos >= len(l.Source)) {
        l.char = l.Source[l.CurrPos]
    }
}

// consumeWhiteSpace : consume extra whitespace characters
func (l *Lexer) consumeWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\r' {
		l.ReadChar()
	}
}

func (l *Lexer) readNumber() Token {
	start := l.CurrPos
	end := l.CurrPos
	floating := false

	for IsDigit(l.PeekChar()) || l.PeekChar() == '.' {
		if l.char == '.' {
			if !IsDigit(l.PeekChar()) {
				return NewToken(Illegal, "", l.line)
			}
			floating = true
		}
		end++
		l.ReadChar()
	}

	numberType := Integer
	if floating == true {
		numberType = Real
	}
    l.CurrPos = end
    l.updateChar()
	return NewToken(numberType, l.Source[start:end+1], l.line)
}

func (l *Lexer) readWord() Token {
	start := l.CurrPos
	end := l.CurrPos

	for IsLetter(l.PeekChar()) {
		end++
		l.ReadChar()
	}
	word := l.Source[start : end+1]
    l.CurrPos = end
    l.updateChar()

	if IsKeyword(word) {
		return NewToken(Keywords[word], word, l.line)
	} else if word == "true" || word == "false" {
		return NewToken(Boolean, word, l.line)
	}

	return NewToken(Ident, word, l.line)
}

func (l *Lexer) readString() Token {
	start := l.CurrPos
	end := l.CurrPos

	for l.char != '"' {
		if l.char == '\\' && l.PeekChar() == '"' {
			end += 2
			l.ReadChar()
		} else {
			end++
		}
		l.ReadChar()

		if l.char == 0 {
			fmt.Printf("ERROR (line : %d) String literal may not be closed.\n", l.line)
			os.Exit(1)
		}
	}
    l.CurrPos = end
    l.updateChar()
	return NewToken(String, l.Source[start:end+1], l.line)
}

func (l *Lexer) readComment() Token {
	l.ReadChar()
	for l.char != '\n' {
		l.ReadChar()
	}
	return NewToken(EOL, "", l.line)
}

// NextToken : to ge the next token from the source
func (l *Lexer) NextToken() Token {
	var token Token
	l.consumeWhiteSpace()

	switch l.char {
	case '\n':
		token = NewToken(EOL, "", l.line)
		l.line++
	case ',':
		token = NewToken(Comma, ",", l.line)
	case '+':
		token = NewToken(Plus, "+", l.line)
	case '-':
		token = NewToken(Minus, "-", l.line)
	case '*':
		token = NewToken(Star, "*", l.line)
	case '/':
		token = NewToken(Fwrdsl, "/", l.line)
	case '%':
		token = NewToken(Percent, "%", l.line)
	case '(':
		token = NewToken(Lparen, "(", l.line)
	case ')':
		token = NewToken(Rparen, ")", l.line)
	case '[':
		token = NewToken(Lsquare, "[", l.line)
	case ']':
		token = NewToken(Rsquare, "]", l.line)
	case '=':
		if l.PeekChar() == '=' {
			token = NewToken(Equal, "==", l.line)
			l.ReadChar()
		} else {
			token = NewToken(Assign, "=", l.line)
		}
	case '!':
		if l.PeekChar() == '=' {
			token = NewToken(Nequal, "!=", l.line)
			l.ReadChar()
		} else {
			token = NewToken(Lognot, "!", l.line)
		}
	case '>':
		if l.PeekChar() == '=' {
			token = NewToken(Grequal, ">=", l.line)
			l.ReadChar()
		} else {
			token = NewToken(Grthan, ">", l.line)
		}
	case '<':
		if l.PeekChar() == '=' {
			token = NewToken(Lsequal, "<=", l.line)
			l.ReadChar()
		} else {
			token = NewToken(Lsthan, "<", l.line)
		}
	case '&':
		token = NewToken(Logand, "&", l.line)
	case '|':
		token = NewToken(Logor, "|", l.line)
	case 0:
		token = NewToken(EOF, "", l.line)
	case ':':
		token = NewToken(Colon, ":", l.line)
	default:
		if IsDigit(l.char) {
			token = l.readNumber()
		} else if IsLetter(l.char) {
			token = l.readWord()
		} else if l.char == '"' {
			token = l.readString()
		} else if l.char == ';' {
			token = l.readComment()
		} else if l.NextPos > len(l.Source) {
            token = NewToken(EOF, "", l.line)
        } else {
			token = NewToken(Illegal, string(l.char), l.line)
		}
	}

    l.ReadChar()
	if token.Type == Illegal {
		fmt.Printf("ERROR (line : %d) Found illegal token (%s)\n.", token.Line, token.Lexeme)
	}
	return token
}
