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

import "fmt"

type Parser struct {
    lexer        *Lexer
    currentToken Token
    peekToken    Token
}

func CreateParser(lexer *Lexer) *Parser {
    p := &Parser{lexer: lexer}
    p.advanceToken()
    p.advanceToken()
    return p
}

func (p *Parser) advanceToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
    // fmt.Println(">>", p.currentToken.Lexeme, "\n>>", p.peekToken.Lexeme, "\n")
}

func (p *Parser) ParseProgram() *Decl {
    // a program is basically just a bunch of decls

    var decl *Decl
    cp := decl

    for p.currentToken.Type != EOF {
        cp = p.parseDecl();
        cp = cp.Next
    }

    return decl

    // token := p.currentToken
    // for token.Type != EOF {
    //     fmt.Println(token.Line, "|", token.Type.StringFor(), "|", token.Lexeme)
    //     p.advanceToken()
    //     token = p.currentToken
    // }
}

func isType(t TokenType) bool {
    switch t {
    case Integer:   fallthrough
    case Real:      fallthrough
    case String:    fallthrough
    case Character: fallthrough
    case Boolean:   fallthrough
    case Array:     fallthrough
    case Function:  fallthrough
    case Struct:
        return true
    }

    return false
}

func (p *Parser) parseDecl() *Decl {
    d := &Decl{}

    p.expectCurrentTokenOrError(Ident)
    d.Name = p.currentToken.Lexeme

    p.expectPeekTokenOrError(Colon)
    p.advanceToken()
    p.advanceToken()

    if !isType(p.currentToken.Type) {
        fmt.Printf("Error: on line %d, expected a type but got `%s`",
                    p.currentToken.Line, 
                    p.currentToken.Type.StringFor())
    }

    return d
}

func (p *Parser) expectCurrentTokenOrError(t TokenType) {
    if p.currentToken.Type != t {
        fmt.Printf("Error: on line %d, expected token of type `%s` but got `%s`",
                    p.currentToken.Line, 
                    t.StringFor(),
                    p.currentToken.Type.StringFor())
    }
}

func (p *Parser) expectPeekTokenOrError(t TokenType) {
    if p.peekToken.Type != t {
        fmt.Printf("Error: on line %d, expected token of type `%s` but got `%s`",
                    p.peekToken.Line, 
                    t.StringFor(),
                    p.peekToken.Type.StringFor())
    }
}
