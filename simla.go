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

package main

import (
	tools "simla/parts"
)

func main() {
    number := "variable : int\nrealvar : real"
	l := tools.CreateLexer(number)
	p := tools.CreateParser(l)
    
    // for {
    //     t := l.NextToken()
    //     if t.Lexeme == "\n" {
    //         t.Lexeme = "NEWLINE"
    //     }
    //     fmt.Printf("%3d \t %10s \t %10s\n", t.Line, t.Lexeme, t.Type.StringFor())
    //     
    //     if t.Type == tools.EOF {
    //         break
    //     }
    // }

    print("\n\n\n")
	p.ParseProgram()
}
