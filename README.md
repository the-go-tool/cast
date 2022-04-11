[![GoDoc](https://godoc.org/github.com/the-go-tool/cast?status.svg)](https://godoc.org/github.com/the-go-tool/cast)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/the-go-tool/cast)
![Coverage](https://img.shields.io/badge/Coverage-91.6%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/the-go-tool/cast)](https://goreportcard.com/report/github.com/the-go-tool/cast)
[![codebeat badge](https://codebeat.co/badges/9c37fc37-7990-4498-bf70-a213abf6fbfe)](https://codebeat.co/projects/github-com-the-go-tool-cast-init)
![example workflow](https://github.com/the-go-tool/cast/actions/workflows/.github/workflows/ci.yaml/badge.svg)

# :sparkles: Cast
> Generics based GoLang library to allow type to type casting.
> No extra dependencies.

> :construction: The module under construction! :construction:

## :package: Install
`go get -u github.com/the-go-tool/cast`

## 🚀 Usage
As long as the module uses generics [Go 1.18](https://go.dev)
or above must be installed. Learn more
[about generics](https://go.dev/blog/why-generics) if you haven't.

Key feature that we have only several methods with type inference to make casts. Let's look at the example:

```go
str := "42"
i := cast.To[int](str)
// Result: i is int(42) now
```

It's possible to use any of primitive types like all types of
**int**, **uint**, **bool**, **string** and their derivatives also.
Another example with custom type based on primitive value:

```go
package main

import (
    "fmt"
    "github.com/the-go-tool/cast"
)

type Custom string

func main() {
    custom := cast.To[Custom](256)
    fmt.Println(custom) // string "256"
}
```

### :hash: Error Handling
We learned about most common function **cast.To**. But what if our program can't cast one type to other by some reason? For example, we're trying
to cast string "non-int" to int. In this case we'll get int(0) and
won't know about error via **cast.To**.

Another one used function is **cast.WithError** which returns casted value
and error if it happens:

```go
i, err := cast.WithError[int]("non-int")
fmt.Println(i) // int(0) so it's default value for int
fmt.Println(err) // "strconv.ParseInt: parsing "non-int": invalid syntax"
```

Also we can check casting possibility like so:

```go
if cast.Possible[int]("non-int") {
    fmt.Println("casting possible")
} else {
    fmt.Println("mission impossible") // will print
}
```

And very rare case when we need to panic in case of impossible casting:

```go
i := cast.MustTo[int]("non-int") // panics
```

### :hash: Casting Extension
This module has flexible API. So, you can register your own type for casting:

```go
type Custom int

func init() {
    MustRegister(func(in Custom) (out int, err error) {
		return int(in), nil
	})
}
```

Now we can use _Custom -> int_ casting like so:

```go
func main() {
    fmt.Println(cast.To[int](Custom(5))) // int 5
}
```

It's going to be tedious to write casting function for each types.
But this module provides casting proxies. Look:

```go
MustRegisterProxy[Custom, string, int]() // Input, Output, Proxy
fmt.Println(cast.To[string](Custom(5))) // string 5
```

Because we already registered _Custom -> int_ and we have default
_int -> string_ casters just register chain _Custom -> int -> string_.

## :arrows_counterclockwise: Built-in Casters

### Related Cross Casting (RelatedXCast)
> Casting between related types but different sizes.

```go
To[int](int32(5)) //> int(5)
To[uint8](uint16(500)) //> uint8(244), so 500-256=244
To[float32](float64(6.9)) //> float32(6.9)
To[byte](uint(5)) //> uint8(5), so byte is alias for uint8
To[rune](int(50)) //> int32(50), so rune is alias for int32
```

| Type X    | X from/to X        |
|-----------|--------------------|
| ints      | :white_check_mark: |
| uints     | :white_check_mark: |
| floats    | :white_check_mark: |
| complexes | :white_check_mark: |

### Number Cross Casting (NumberXCast)
> Casting between number types (with possible different sizes).

```go
To[int](56.6) //> int(56), only integer part
To[float32](56) //> float32(56)
To[uint8](int8(-40)) //> uint8(216), so 256-40=216
To[uint64](int8(-1)) //> uint64(MAX_VALUE)
To[float32](complex(5, -1)) //> float32(5)
```

| Type N1 | Type N2    | N1 from/to N2      |
|---------|------------|--------------------|
| ints    | uints      | :white_check_mark: |
| ints    | floats     | :white_check_mark: |
| ints    | complexes  | :white_check_mark: |
| uints   | floats     | :white_check_mark: |
| uints   | complexes  | :white_check_mark: |
| floats  | complexes  | :white_check_mark: |

### String Cross Casting (StringXCast)
> Casting between string and anything basic type.
> Usually, it's close to `strconv` standard behavior except
> boolean.
> 
> Casting from string to boolean extended and accepts any
> of `1`, `+`, `t`, `true`, `y`, `yes` as `true` and any of
> `0`, `-`, `f`, `false`, `n`, `no` as `false`.
> It's space trimmed and case insensitive.
> Any other string values will produce an error.

```go
To[bool]("y") //> true
To[bool]("  1 ") //> true
To[bool](" n  ") //> false
To[bool](" dsa ") //> false, and error for WithError
To[string](true) //> "true"
To[string](complex(-1, -2)) //> "(-1-2i)"
To[int]("-5") //> int(-5)
To[int]("+5") //> int(5)
To[int]("5.6") //> int(0), and error for WithError
To[float32]("5.6") //> float32(5.6)
```

| Type S | Type A    | S from/to A        |
|--------|-----------|--------------------|
| string | ints      | :white_check_mark: |
| string | uints     | :white_check_mark: |
| string | floats    | :white_check_mark: |
| string | complexes | :white_check_mark: |
| string | bool      | :white_check_mark: |

### Boolean to Number Types Casting (BoolNumCast)
> Casting between boolean and number types.
> `true` alway will be `1` if cast to number or `1+0i` for complexes.
> Number will `true` if value isn't equal `0`. For complexes only
> real part is matter.

```go
To[int](true) //> int(1)
To[int](false) //> int(0)
To[bool](int(0)) //> false
To[bool](int(-1)) //> true
To[bool](uint(1)) //> true
To[bool](complex(0, 5)) //> false, so only real part matter
To[bool](complex(5, 0)) //> true
``` 

| Type B | Cast N    | B from/to N        |
|--------|-----------|--------------------|
| bool   | ints      | :white_check_mark: |
| bool   | uints     | :white_check_mark: |
| bool   | floats    | :white_check_mark: |
| bool   | complexes | :white_check_mark: |

## :link: Similar Projects
Please, star this repository if you find it helpful :star:  
Also, [make an issue](https://github.com/the-go-tool/cast/issues)
if you found a bug or would like for some improvements.

If this module doesn't fit here is links to similar projects:
- :link: https://github.com/spf13/cast