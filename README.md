<h1>GoSch</h1>
<p>Go schema validation</p>

## Table of contents

- [Table of contents](#table-of-contents)
- [Introduction](#introduction)
- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Types](#types)
- [Strings](#strings)
- [Numbers](#numbers)
- [Todos](#todos)

## Introduction

GoSch is a Go schema validation library.

Highlights:

- Zero dependencies

## Installation

```sh
go get github.com/ItsMalma/gosch
```

## Basic Usage

Creating a simple string schema

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    nameSchema := gosch.String()

    println(nameSchema.Validate("Malma"))
    println(nameSchema.Validate(77))
}
```

Creating a complex schema (struct, array)

```go
package main

import "github.com/ItsMalma/gosch"

type Person struct {
    Name string
    Age  int
}

func main() {
    personSchema := gosch.Struct().
        Field("Name", gosch.String()).
        Field("Age", gosch.Int())
    
    personsSchema := gosch.Array().
        Element(personSchema).
        Length(3)
    
    println(personSchema.Validate(Person{
        Name: "Malma",
        Age: 19,
    }))
    println(personsSchema.Validate([3]Person{
        {
            Name: "Malma",
            Age: 19,
        },
        {
            Name: "John Doe",
            Age: 25,
        },
        {
            Name: "Bob",
            Age: 22,
        },
    }))
}
```

## Types

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    // Primitive data types
    gosch.String()
    gosch.Int()
    gosch.Int8()
    gosch.Int16()
    gosch.Int32()
    gosch.Int64()
    gosch.Uint()
    gosch.Uint8()
    gosch.Uint16()
    gosch.Uint32()
    gosch.Uint64()
    gosch.Float32()
    gosch.Float64()

    // Nilable (pointer) primitive data types
    gosch.String().Nil()
    gosch.Int().Nil()
    gosch.Int8().Nil()
    gosch.Int16().Nil()
    gosch.Int32().Nil()
    gosch.Int64().Nil()
    gosch.Uint().Nil()
    gosch.Uint8().Nil()
    gosch.Uint16().Nil()
    gosch.Uint32().Nil()
    gosch.Uint64().Nil()
    gosch.Float32().Nil()
    gosch.Float64().Nil()
}
```

## Strings

Gosch includes additional string-specific rules.

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.String().NotEmpty()
    gosch.String().MinLength(3)
    gosch.String().MaxLength(3)

    gosch.String().
        NotEmpty().
        MinLength(8).
        MaxLength(100)
}
```

## Numbers

Gosch includes additional number-specific (int, uint and float) rules.

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.Int().MinValue(-100)
    gosch.Int().MaxValue(100)

    gosch.Uint().MinValue(0)
    gosch.Uint().MaxValue(100)

    gosch.Float32().MinValue(-77.7)
    gosch.Float32().MaxValue(77.7)
}
```

## Todos

- [ ] String
    - [x] Data Type
    - [x] Nil
    - [x] Not Empty
    - [x] Min Length
    - [x] Max Length
    - [ ] Not Length
    - [ ] Starts With
    - [ ] Ends With
    - [ ] Includes
    - [ ] Excludes
    - [ ] Pattern
        - [ ] Email
        - [ ] ISO Date
        - [ ] Phone Number
- [ ] Int
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
- [ ] Uint
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
- [ ] Float
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
- [x] Struct
    - [x] Data Type
    - [x] Nil
    - [x] Field
- [x] Array
    - [x] Data Type
    - [x] Nil
    - [x] Element
    - [x] Length
- [x] Slice
    - [x] Data Type
    - [x] Nil
    - [x] Element
    - [x] Min Length
    - [x] Max Length
- [ ] Map
    - [ ] Data Type
    - [ ] Nil
    - [ ] Key
    - [ ] Min Length
    - [ ] Max Length
- [ ] Custom
    - [ ] Error Message
    - [ ] Rule
    - [ ] Schema
- [ ] Unit Tests