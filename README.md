FastReflect
-----------

([English](./README.md) / [中文](./README-zh-CN.md))

A faster method to get elements from an interface (Struct or Slice type) for Go.


# Usage
 
```go
import (
    "github.com/karminski/fastreflect"
)

// Access an interface (struct type) property by name
TheFieldYouWant := fastreflect.StructFieldByName(YourStructInterface, "FieldName")

// Get all elements from an interface (slice type)
AllElementsYouWant := fastreflect.SliceAllElements(YourSliceInterface)

```

For more details please see: [FastReflect_test.go](./FastReflect_test.go)


# Performance

- Access an interface (struct type) property by name, **3.58x** fast.
- Get all elements from an interface (slice type), **4.3x** fast.

```
root# go test -bench=.
goos: linux
goarch: amd64
BenchmarkStructFieldByName-8        14354080            78.6 ns/op
BenchmarkInternalFieldByName-8       4009170           297 ns/op
BenchmarkSliceAllElements-8         10825730           124 ns/op
BenchmarkInternalIndex-8             2514120           489 ns/op
PASS
ok      _/fastreflect/src   5.883s
```
