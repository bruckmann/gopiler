# Gopiler

An interpreter for Summer language written in Golang. 

## The Summer language:

### Syntax

```summer
 let five = 5;
 let ten = 10;

 let add = fn(a, b) {
   a + b;
  };

 let result = add(five, ten);
```

### Features

- First Class functions
- Higher order functions
- Variable bindings 
- Integer and booleans
- Closures
- String data structure
- Array data structure
- A hash data structure
- Built-in functions
- arithmetic expressions

### To use the interactive mode just run the command:

```shell 
  go run main.go
```

The language and the interpreter was done based on the book "Writing An Interpreter in Go" from Thorsten Ball.


