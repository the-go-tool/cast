[![GoDoc](https://godoc.org/github.com/the-go-tool/cast?status.svg)](https://godoc.org/github.com/the-go-tool/cast)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/the-go-tool/cast)
![Coverage](https://img.shields.io/badge/Coverage-64.8%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/the-go-tool/cast)](https://goreportcard.com/report/github.com/the-go-tool/cast)
![example workflow](https://github.com/the-go-tool/cast/actions/workflows/.github/workflows/ci.yaml/badge.svg)

# :sparkles: Cast
> Generics based GoLang library to allow type to type casting.
> No extra dependencies.

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

## :link: Similar Projects
Please, star this repository if you find it helpful :star:  
Also, [make an issue](https://github.com/the-go-tool/cast/issues)
if you found a bug or would like for some improvements.

If this module doesn't fit here is links to similar projects:
- :link: https://github.com/spf13/cast