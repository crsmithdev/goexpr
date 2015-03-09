# goexpr

[![Travis](https://travis-ci.org/crsmithdev/goexpr.svg?branch=master)](https://travis-ci.org/crsmithdev/goexpr)
[![GoDoc](https://godoc.org/github.com/crsmithdev/goexpr?status.svg)](https://godoc.org/github.com/crsmithdev/goexpr)

## What?

Goexpr is a Go package that evaluates mathematical expressions.  Values can be constants, or variables whose values are provided at evaluation time.  Goexpr is inspired by [MathJS](http://mathjs.org/).

## Why?

Being able to evaluate basic mathematical expressions is useful in analytics, but Go is not a scripting language with `eval`.  A simple, pluggable math library abstracts out the mechanics of evaluation from the input data.

## Features

- Supports basic algebraic operations.
- Not regex-based, parses with Go's `ast`.
- Safe, restricted to a subset of language-defined syntax.
- No external dependencies.

## Quick start

Install:


```
go get github.com/crsmithdev/goxpr
```

Evaluate `(a + b) / 2`, where `a=3` and `b=1`:

```
import "github.com/crsmithdev/goexpr"

parsed, _ := goexpr.Parse("(a + b) / 2")
result, _ := goexpr.Evaluate(parsed, map[string]float64{
    "a": 3,
    "b": 1,
})

fmt.Println("result: %d", result)
```
```
result: 2
```

## Documentation

API documentation on [Godoc](https://godoc.org/github.com/crsmithdev/goexpr)

## Recommended tools

- [godep](https://github.com/tools/godep) to lock to a specific commit or tag
- [goenv](https://github.com/crsmithdev/goenv)  to isolate dependencies between projects

## Development

After cloning, build:
```
make
```

Run tests:
```
make test
```

Run tests automatically on write:
```
make test-auto
```

Run tests with coverage:
```
make test-cov
```

View HTML coverage report:
```
make html-cov
```

## Next

- Additional operators and functions
- Additional examples
