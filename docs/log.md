# Simla Development Log

### Day 1
- modified: docs/lang.md, parts/tokens.go, parts/lexer.go
- notes:
    - copied over coltok.go from the colon-lang repo
    - copied over collex.go from the colon-lang repo
    - removed the `string = string + string` pattern
    - replaced said patter with `string = string[start:end]`
    - edited the tokens and lexer file to fit language description

### Day 2
- modified: parts/tokens.go
- notes:
    - added `print`, `true`, `false` to the token list
    - added code to lexer to support these new tokens
