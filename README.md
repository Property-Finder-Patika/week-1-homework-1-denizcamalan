# week-1-homework-1

## Weekly Contents

- Go Code Structure and Analysis
- Global and Local Variables
- Constants
- Variables, Basic Data Types
- Scope
- Type Conversions

# Chapter1

## Main Properties

- Automatic memory management via garbage collection
- First-class functions
- Lexical (static) scope
- Standard library and APIs for basic structures such as collections (containers), networking, I/O, text processing, etc.
- Go and all of its libraries and tools are open-source.
- It is a compiled language

## Object-Oriented Programming (OOP) in Go
1. Struct
- Go does not support custom types through classes but structs.
2. Encapsulation
- It means hiding sensitive data from users.
> **Note:** if function/struct is public as it begins with a capital letter, else it begins with a lower case.
3. Inheritance
- subclass/child class are the terms used for the class which acquire properties in golang with struct.
4. Interfaces
- The interface is a custom type that is used to specify a set of one or more method signatures.By treating objects of different types in a consistent way, as long as they stick to one interface, Golang implements polymorphism.

## Functional Programming in Go
Golang is not a functional language but has a lot of features that enable to turning our code more elegant, concise, maintainable, easier to understand and test.
- Functions are first-class citizens.
- Supports anonymous functions or closures.

## CommandLine Arguments

- os.Args is a slice of strings.

## Lecture 1 Codes

[Lecture Codes](./code/ch1/greeting)

# Exercises
- Exercise 1.1
    [Code Example and Output text file](./code/ch1/echo)
- Exercise 1.2
    [Code Example and Output text file](./code/ch1/echo2)
- Exercise 1.3
    [Code Example and Output text file](./code/ch1/echo3)   &   [Output Picture](./code/pics/echo3.png) 
- Excercise 1.4
    [Code Example and Output text file](./code/ch1/dup2)
- Excercise 1.5
    [Code Example and Output text file](./code/ch1/lissajous)   &    [Output Picture](./code/pics/lissajouss_green.png)
- Excercise 1.6
    [Code Example and Output text file](./code/ch1/lissajous)   &    [Output Picture](./code/pics/changing_img.png)
- Excercise 1.7 & Excercise 1.8 & Excercise 1.9
    [Code Example](./code/ch1/fetch)
- Ecercise 1.10
    [Code Example and Output text file](./code/ch1/fetch)
- Ecercise 1.11
    [Code Example and Output text file](./code/ch1/fetchall)
- Ecercise 1.12
    [Code Example](./code/ch1/fetch)    &   [Output Picture](./code/pics/cycles=20.png)
# Chapter2

## Names
- Go programmers use "camel case" when forming names by combining words; that is, interior capital letters are preferred over interior underscores.
## Declarations
- Four major kinds of declarations: var, const, type, and func.

## Variables
- A pointer value is the address of a variable.
- Pointers are key to the flag package, which uses a program's command-line arguments to set
the values of cer tain variables distributed throughout the program.
- The expression new(T) creates an unnamed variable of typ e T.

## Assignments
- Tuple assignment
## Type Declarations
    type Celsius float64
    type Fahrenheit float64

- Two types, Celsius and Fahrenheit, for the two units of temperature. Even though both have the same underlying typ e, float64, they are not the same type, so they cannot be compared or combined in arithmetic expressions.

## Scope

- If a name is declared in both an outer block and an inner block, the inner declaration will be found first.

# Exercises
- Exercise 2.1
    [Code Example and Output text file](./code/ch2/cf)
- Exercise 2.2
    [Code Example and Output text file](./code/ch2/meterconv)
- Exercise 2.3
    [Code Example and Output text file](./code/ch2/popcount)   &   [Output Picture](./code/pics/popcountVSlooppopcount.png)
- Exercise 2.3
    [Code Example and Output text file](./code/ch2/popcount)   &   [Output Picture](./code/pics/popcount2.4.png)
- Exercise 2.3
    [Code Example and Output text file](./code/ch2/popcount)   &   [Output Picture](./code/pics/popcount2.5.png)

# Chapter3

## Integers
        & bitwise AND
        | bitwise OR
        ^ bitwise XOR
        &^ bit cle ar (AND NOT)
        << lef t shif t
        >> right shif t
## Floating-Point Numbers
## Complex Numbers
## Booleans
## Strings
## Constants

# Exercises
- Exercise 3.1
    [Code Example and Output text file](./code/ch3/surface) &   [Output Picture](./code/ch3/surface/output1.png)
- Exercise 3.2
    [Output Picture](./code/ch3/surface/output2.png)
- Exercise 3.3
    [Output Picture](./code/ch3/surface/output3.png)
- Exercise 3.4
    [Code Example and Output text file](./code/ch3/surface) &   [[Output Picture](./code/pics/surface_web.png)
- Exercise 3.5
    [Output Picture](./code/pics/web_mandelbrot.png)
- Exercise 3.6
    [Output Picture](./code/pics/web_colored.png)
- Exercise 3.10
    [Code Example and Output text file](./code/ch3/printints)
- Exercise 3.11
    [Code Example and Output text file](./code/ch3/printints)   &   [Output Picture](./code/pics/floating_point.png)
- Exercise 3.12
    [Code Example and Output text file](./code/ch3/printints)

## Documentation

[Go Doc](https://go.dev/doc/)
[GoPL book](https://drive.google.com/file/d/1kvsEfCuOYecBrfy12tTI1kDDbC4e4AVy/view?usp=sharing)
[GeeksForGreeks](www.geeksforgeeks.org/object-oriented-programming-in-golang/)