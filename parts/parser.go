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

	// fmt.Println("current : ", p.currentToken.Lexeme, "\npeeked : ", p.peekToken.Lexeme)
}

func (p *Parser) ParseProgram() *Program {
	// a program is basically just a bunch of decls

	program := &Program{}
    var cp **Decl

	for p.currentToken.Type != EOF {
		// fmt.Println("LOG: Current decl is for", p.currentToken.Lexeme)
        if program.Decls == nil {
            program.Decls = p.parseDecl()
            cp = &program.Decls
        } else {
            *cp = p.parseDecl()
        }

        // if cp != nil {
        //     println(cp.Name)
        // }

		fmt.Println("LOG: Parsed a decl")
		// fmt.Printf("LOG: (%s, %d)\n", cp.Name, cp.Type.TypeKind)
		cp = &(*cp).Next
	}

	return program
}

func (p *Parser) parseDecl() *Decl {
	d := &Decl{}

	p.expectCurrentTokenOrError(Ident)
	d.Name = p.currentToken.Lexeme
	// print(d.Name, ") ")
	p.advanceToken()

	p.expectCurrentTokenOrError(Colon)
	// print(d.Name, ") ")
	p.advanceToken()

	// if !isType(p.currentToken.Type) {
	//	fmt.Printf("Error: on line %d, expected a type but got `%s`\n",
	//		p.currentToken.Line,
	// 		p.currentToken.Type.StringFor())
	// }

	t := p.currentToken.Type
	dtype := &Type{}

	switch t {
	case Integer:
		dtype.TypeKind = TInt

	case Real:
		dtype.TypeKind = TReal

	case String:
		dtype.TypeKind = TString

	case Character:
		dtype.TypeKind = TChar

	case Boolean:
		dtype.TypeKind = TBool

	case Array:
		dtype.TypeKind = TArray
		fmt.Println("OUTPUT WRONG FROM THIS POINT ON")
		panic("Unimplemented")

	case Function:
		dtype.TypeKind = TFunc
		fmt.Println("OUTPUT WRONG FROM THIS POINT ON")
		panic("Unimplemented")

	case Struct:
		dtype.TypeKind = TStruct
		fmt.Println("OUTPUT WRONG FROM THIS POINT ON")
		panic("Unimplemented")

	default:
		fmt.Printf("Error: on line %d, expected a type but got `%s`\n",
			p.currentToken.Line,
			p.currentToken.Type.StringFor())

	}

	// print(d.Name, ") ")
	p.advanceToken()
	if p.currentToken.Type == EOL || p.currentToken.Type == EOF {
		d.Type = dtype
		// print(d.Name, ") ")
		p.advanceToken()
		return d
	}

	return d
}

func (p *Parser) expectCurrentTokenOrError(t TokenType) {
	if p.currentToken.Type != t {
		fmt.Printf("Error: on line %d, expected token of type `%s` but got `%s`\n",
			p.currentToken.Line,
			t.StringFor(),
			p.currentToken.Type.StringFor())
	}
}

func (p *Parser) expectPeekTokenOrError(t TokenType) {
	if p.peekToken.Type != t {
		fmt.Printf("Error: on line %d, expected token of type `%s` but got `%s`\n",
			p.peekToken.Line,
			t.StringFor(),
			p.peekToken.Type.StringFor())
	}
}
