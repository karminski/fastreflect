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

- 如果 interface 是 Struct, 根据属性名称获取它的属性, 性能是原生 reflect 的 **4.35x** 倍.
- 如果 interface 是 Slice, 获取它所有的元素, 性能是原生 reflect 的 **3.1x** 倍.

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
