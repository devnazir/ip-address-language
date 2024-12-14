# IP ADDRESS Programming Language

In this language, all keywords and expressions are represented using ASCII values. The syntax is inspired by languages like Go, but with ASCII code representations for all keywords and expressions.

## Installation

Install via curl:

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/devnazir/ip-address-language/refs/heads/master/install.sh)"
```

Install via wget:

```bash
bash -c "$(wget -O- https://raw.githubusercontent.com/devnazir/ip-address-language/refs/heads/master/install.sh)"
```

## Usage
```bash
ipl <filename>.n
```

## Features

- [x] Variable Assignment

  - [x] int
  - [x] string
  - [x] bool
  - [x] float64
  - [x] array
  - [x] object
  - [x] function
  - [x] capture shell output

- [x] Arithmetic Operations

  - [x] Addition
  - [x] Subtraction
  - [x] Multiplication
  - [x] Division
  - [x] Modulus

- [x] Comparison Operations

  - [x] Equals
  - [x] Not Equals
  - [x] Greater Than
  - [x] Less Than
  - [x] Greater Than or Equal To
  - [x] Less Than or Equal To

- [x] Logical Operations

  - [x] And
  - [x] Or
  - [x] Not

- [x] Conditional Statements
  - [x] If
  - [x] Else
  - [x] Else If
  - [ ] Switch
  - [ ] Case
- [ ] Loops

  - [ ] For
  - [ ] While
  - [ ] Do While

- [x] Functions

  - [x] Declaration
  - [x] Return
  - [x] Arguments
  - [x] Anonymous Functions

- [ ] Built-in Functions

  - [x] echo
  - [ ] array
    - [ ] append
    - [ ] delete
    - [ ] len
  - [ ] string
    - [ ] split
    - [ ] join
    - [ ] replace
    - [ ] toUpper
    - [ ] toLower

- [ ] Error Handling

- [x] Comments

  - [x] Single Line
  - [x] Multi Line

- [ ] Static Typing

- [x] Import & Export
  - [x] Variables
  - [x] Functions
  - [x] Aliases

## Syntax

When writing code in IP Address Language, you will use ASCII values to represent keywords and expressions. For example, the keyword `var` is represented by the ASCII value `118 97 114`. But you have to separate each value with a dot `.`. to represent the keyword `var` in IP Address Language, you will write `118.97.114`.

> **Note:** not all keyword or expression should be separated by a dot. (see example below or check the [examples](examples) directory

Here some example keyword or exression that is should not be separated by a dot:

#### String

when creating a string, you should wrap the ascii values with double quotes `34`. So, to create a string `John Doe`, you will write:

```
34 74.111.104.110 68.111.101 34 // "John Doe"
```

`34` is the ASCII value for double quotes `"`

#### Function

`(` and `)` after the function name should not be separated by a dot. So, to call a function like `sayHello()`, you will write:

```
115.97.121.72.101.108.108.111 40 41

// 115.97.121.72.101.108.108.111 = sayHello
// 40 = (
// 41 = )
```

Some example of syntax that you can use in IP Address Programming Language:

### Variable Assignment

```
118.97.114 120 61 52.50 // var x = 42
```

### Arithmetic Operations

```
118.97.114 120 61 52.50 43 52.50 // var x = 42 + 42
```

## Print Output

```
101.99.104.111 36.120 // echo $x
```

## Conditional Statements

```
105.102 120 60 53 123
    101.99.104.111 34 120 105.115 103.114.101.97.116.101.114 116.104.97.110 53 34
125 101.108.115.101 105.102 120 61.61 53 123
    101.99.104.111 34 120 105.115 101.113.117.97.108 116.111 53 34
125 101.108.115.101 123
    101.99.104.111 34 120 105.115 108.101.115.115 116.104.97.110 53 34
125
```

is equivalent to:

```go
if x < 5 {
    echo "x is greater than 5"
} else if x == 5 {
    echo "x is equal to 5"
} else {
    echo "x is less than 5"
}
```

## Functions

```
118.97.114 110.97.109.101 61 34 74.111.104.110 68.111.101 34
101.99.104.111 34 72.101.108.108.111.44 36.110.97.109.101.33 34

102.117.110.99 115.97.121.72.101.108.108.111 40 41 123
  118.97.114 110.97.109.101 61 34 85.108.117.109 34
  101.99.104.111 34 72.101.108.108.111.44 36.110.97.109.101.33 34

  102.117.110.99 115.97.121.72.101.108.108.111.50 40 41 123
    118.97.114 110.97.109.101 61 34 68.111.101 34
    118.97.114 97.103.101 61 50.48
    101.99.104.111 34 72.101.108.108.111.44 36.110.97.109.101.33 34

    102.117.110.99 115.97.121.72.101.108.108.111.51 40 41 123
      118.97.114 110.97.109.101 61 34 74.111.104.110 34
      101.99.104.111 34 72.101.108.108.111.44 36.110.97.109.101.33 34
      101.99.104.111 34 65.103.101 58 36.97.103.101 34
    125

    115.97.121.72.101.108.108.111.51 40 41
  125

  115.97.121.72.101.108.108.111.50 40 41
125

115.97.121.72.101.108.108.111 40 41
```

is equivalent to:

```go
var name = "John Doe"
echo "Hello, $name!"

func sayHello() {
  var name = "Ulum"
  echo "Hello, $name!"

  func sayHello2() {
    var name = "Doe"
    var age = 20
    echo "Hello, $name!"

    func sayHello3() {
      var name = "John"
      echo "Hello, $name!"
      echo "Age: $age"
    }

    sayHello3()
  }

  sayHello2()
}

sayHello()
```

To see more examples, check the [examples](examples) directory.

> FYI, initially, this project was called [gosh](https://github.com/devnazir/gosh), but I decided to modify the syntax and name it IP Address Programming Language.

> This is just a fun project.

And if you like this project, please give it a star. Thank you!
