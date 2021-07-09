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

- Access an interface (struct type) property by name, **4.35x** fast.
- Get all elements from an interface (slice type), **3.1x** fast.

```
goos: linux
goarch: amd64
pkg: fastreflect
cpu: Intel(R) Xeon(R) CPU E3-1246 v3 @ 3.50GHz
BenchmarkStructFieldByName-8        14697340            88.99 ns/op        0 B/op          0 allocs/op
BenchmarkInternalFieldByName-8       3378478           353.4 ns/op        32 B/op          4 allocs/op
BenchmarkSliceAllElements-8          9508502           128.9 ns/op       144 B/op          1 allocs/op
BenchmarkInternalIndex-8             3015903           398.2 ns/op       216 B/op          9 allocs/op
PASS
ok      fastreflect 5.907s
```
