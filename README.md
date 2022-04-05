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
Learn more about [available types here](cast.go).
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

### Error handling
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