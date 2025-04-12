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
- [Arrays](#arrays)
- [Slices](#slices)
- [Maps](#maps)
- [Structs](#structs)
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

    nameSchema.Validate("Malma")
    nameSchema.Validate(77)
}
```

Creating a struct schema

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

    personSchema.Validate(Person{
        Name: "Malma",
        Age: 19,
    })
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

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.String()
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

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.Int()
    gosch.Int().MinValue(-100)
    gosch.Int().MaxValue(100)
    gosch.Int().
        MinValue(-100).
        MaxValue(100)

    gosch.Uint()
    gosch.Uint().MinValue(0)
    gosch.Uint().MaxValue(100)
    gosch.Uint().
        MinValue(0).
        MaxValue(100)

    gosch.Float32()
    gosch.Float32().MinValue(-77.7)
    gosch.Float32().MaxValue(77.7)
    gosch.Float32().
        MinValue(-77.7).
        MaxValue(77.7)
}
```

## Arrays

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.Array(3, gosch.Int())
}
```

## Slices

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.Slice(gosch.Int())
    gosch.Slice(gosch.Int()).MinLength(1)
    gosch.Slice(gosch.Int()).MaxLength(10)
    gosch.Slice(gosch.Int()).
        MinLength(1).
        MaxLength(10)
}
```

## Maps

```go
package main

import "github.com/ItsMalma/gosch"

func main() {
    gosch.Map(gosch.String(), gosch.Int())
    gosch.Map(gosch.String(), gosch.Int()).MinLength(1)
    gosch.Map(gosch.String(), gosch.Int()).MaxLength(10)
    gosch.Map(gosch.String(), gosch.Int()).
        MinLength(1).
        MaxLength(10)
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
    - [ ] Value
    - [ ] Not Value
    - [ ] Starts With
    - [ ] Not Starts With
    - [ ] Ends With
    - [ ] Not Ends With
    - [ ] Includes
    - [ ] Not Includes
    - [ ] Excludes
    - [ ] Not Excludes
    - [ ] Pattern
        - [ ] Email
        - [ ] ISO Date
        - [ ] Phone Number
    - [ ] Not Pattern
        - [ ] Email
        - [ ] ISO Date
        - [ ] Phone Number
    - [ ] Enum
    - [ ] Not Enum
- [ ] Int
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
    - [ ] Multiple Of
    - [ ] Not Multiple Of
    - [ ] Enum
    - [ ] Not Enum
- [ ] Uint
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
    - [ ] Multiple Of
    - [ ] Not Multiple Of
    - [ ] Enum
    - [ ] Not Enum
- [ ] Float
    - [x] Data Type
    - [x] Nil
    - [x] Min Value
    - [x] Max Value
    - [ ] Not Value
    - [ ] Multiple Of
    - [ ] Not Multiple Of
    - [ ] Enum
    - [ ] Not Enum
- [x] Struct
    - [x] Data Type
    - [x] Nil
    - [x] Field
    - [ ] Not Field
- [x] Array
    - [x] Data Type
    - [x] Nil
    - [x] Element
    - [x] Length
    - [ ] Contains Element
    - [ ] Not Contains Element
- [x] Slice
    - [x] Data Type
    - [x] Nil
    - [x] Element
    - [x] Min Length
    - [x] Max Length
    - [ ] Contains Element
    - [ ] Not Contains Element
- [x] Map
    - [x] Data Type
    - [x] Nil
    - [x] Key
    - [x] Element
    - [x] Min Length
    - [x] Max Length
    - [ ] Contains Key
    - [ ] Not Contains Key
    - [ ] Contains Element
    - [ ] Not Contains Element
- [ ] Custom
    - [ ] Error Message
    - [ ] Rule
    - [ ] Schema
- [ ] Unit Tests