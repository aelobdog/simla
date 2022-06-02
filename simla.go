package main

import (
    "fmt"
    tools "simla/parts"
)

func main() {
    fmt.Println("hello, world")
    number := "hello 123.45"
    l := tools.CreateLexerState(number)

    token := l.NewToken()
    fmt.Println(token.Lexeme)

    token = l.NewToken()
    fmt.Println(token.Lexeme)
}
