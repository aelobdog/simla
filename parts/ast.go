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

// import "fmt"

type Node interface {
	NodeName() string
}

type Program struct {
	Statements []Node
}

type Identifier struct {
    Token Token
    Value string
}

type IntLit struct {
    Token Token
    Value int32
}

type RealLit struct {
    Token Token
    Value float32
}

type CharLit struct {
    Token Token
    Value byte
}

type BoolLit struct {
    Token Token
    Value bool
}

type StringLit struct {
    Token Token
}

// TODO: make ast nodes for the different statements and expressions
// statements:
//      declaration
//      initialization/function defn/struct defn/ all that
//      conditional
//      loops

// making all the structs belong to interface:Node

func (o *Program) NodeName() string {
    return "Program"
}

func (o *Identifier) NodeName() string {
    return "Id: " + o.Token.Lexeme
}

func (o *IntLit) NodeName() string {
    return "Int: " + o.Token.Lexeme
}

func (o *RealLit) NodeName() string {
    return "Real: " + o.Token.Lexeme
}

func (o *CharLit) NodeName() string {
    return "Char: " + o.Token.Lexeme
}

func (o *BoolLit) NodeName() string {
    return "Bool: " + o.Token.Lexeme
}

func (o *StringLit) NodeName() string {
    return "Str: " + o.Token.Lexeme
}
