String substitution benchmark
=============================

```
BenchmarkReplacer-4       	 1000000	      1188 ns/op	    2488 B/op	       8 allocs/op
BenchmarkReplacerOnce-4   	 5000000	       282 ns/op	      88 B/op	       4 allocs/op
BenchmarkStrings-4        	 5000000	       250 ns/op	      64 B/op	       2 allocs/op
BenchmarkRegex-4          	 1000000	      1147 ns/op	     136 B/op	       7 allocs/op
```