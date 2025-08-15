# miniid

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

`miniid` generates ID encoded in base62.

## Features

* Simple API.
* Clean and tested code.
* Dependency-free.

## Install

Go version 1.20+

```
go get github.com/cristalhq/miniid
```

## Example

```go
gen := miniid.New(1200)

for i := 0; i < 10; i++ {
	fmt.Print(gen.Next(), " ")
}
fmt.Println()

// Output:
// JN JO JP JQ JR JS JT JU JV JW
```

Also see examples: [example_test.go](https://github.com/cristalhq/miniid/blob/main/example_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/miniid/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/miniid/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/miniid
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/miniid
[reportcard-img]: https://goreportcard.com/badge/cristalhq/miniid
[reportcard-url]: https://goreportcard.com/report/cristalhq/miniid
[coverage-img]: https://codecov.io/gh/cristalhq/miniid/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/miniid
[version-img]: https://img.shields.io/github/v/release/cristalhq/miniid
[version-url]: https://github.com/cristalhq/miniid/releases
