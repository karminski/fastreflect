FastReflect
-----------

([English](./README.md) / [中文](./README-zh-CN.md))

一个更快的方法, 用来获取 interface (Struct) 的属性, 或 interface (Slice)  的元素.


# 使用方法

```go
import (
    "github.com/karminski/fastreflect"
)

// 如果 interface 是 Struct, 根据属性名称获取它的属性
TheFieldYouWant := fastreflect.StructFieldByName(YourStructInterface, "FieldName")

// 如果 interface 是 Slice, 获取它所有的元素
AllElementsYouWant := fastreflect.SliceAllElements(YourSliceInterface)

```

详细使用方法见: [FastReflect_test.go](./FastReflect_test.go)


# Performance

- 如果 interface 是 Struct, 根据属性名称获取它的属性, 性能是原生 reflect 的 **3.58x** 倍.
- 如果 interface 是 Slice, 获取它所有的元素, 性能是原生 reflect 的 **4.3x** 倍.

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
