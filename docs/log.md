# Simla Development Log

### Day 1
Created the syntax for the language. 
Wrote parts/tokens.go and parts/lexer.go
Borrowed a lot of code from my Colon-lang lexer, but modified some behavior. Most changes are basically removing the constant _string = string + string_ operations with a simpler substring _string[start:end]_ operations.

### Day 2
Updated the parts/tokens.go file to add "print" to the list of keywords
