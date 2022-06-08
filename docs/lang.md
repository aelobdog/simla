# Language Features

## Features

### Values
```
12, 12.3, 'c', "hello", true
```

### Variables
```
name : int
name : char = 'c'
```

### Arrays
```
name : array [number] type
name : array [number] type = {values}
```

### Conditionals
```
if CONDITION then
    code
else
    code
end
```

### Loops
```
loop CONDITION

end
```

### Functions
```
name : func RET_TYPE ( PARAMS ) = {

}
```

### Structs
```
name : struct = {
    sub-components
}
```

### Comments
```
; comment text
```

## Design Notes

References used to design simla
    - _Introduction to Compilers and Language Design_, Douglas Thain, https://www3.nd.edu/~dthain/compilerbook/
    - _Writing An Interpreter In Go_, Thorsten Ball, https://interpreterbook.com/
    - _Writing A Compiler in Go_, Thorsten Ball, https://compilerbook.com/
