![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro](https://github.com/linkedin/goavro)

#### Results

Go: 1.12 Machine: Macbook Pro 2,6 GHz Intel Core i7 16 GB 2133 MHz LPDDR3

```
BenchmarkGoAvroDecode-8     	  500000	      3448 ns/op	     442 B/op	      27 allocs/op
BenchmarkGoAvroEncode-8     	  300000	      4123 ns/op	     841 B/op	      63 allocs/op
BenchmarkHambaDecode-8      	 3000000	       531 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-8      	 3000000	       432 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-8   	 1000000	      2065 ns/op	    1776 B/op	      40 allocs/op
BenchmarkLinkedinEncode-8   	 2000000	       830 ns/op	     288 B/op	      10 allocs/op
```