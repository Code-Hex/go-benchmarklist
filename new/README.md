Struct new benchmark
=============================

```
BenchmarkNew-4         	200000000	         6.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewNormal-4   	200000000	         6.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewSimple-4   	200000000	         6.67 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Code-Hex/go-benchmarklist/new	6.000s
```

多分ビルド時の最適化によって同じコードに修正されているのだろうと予想