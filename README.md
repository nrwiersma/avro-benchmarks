![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)

#### Results

Go: 1.17 Machine: Macbook Pro 2,6 GHz Intel Core i7 16 GB 2133 MHz LPDDR3

```
BenchmarkGoAvroDecode-8      	  333607	      3182 ns/op	     418 B/op	      27 allocs/op
BenchmarkGoAvroEncode-8      	  303241	      3975 ns/op	     839 B/op	      63 allocs/op
BenchmarkGoGenAvroDecode-8   	  467928	      2603 ns/op	     728 B/op	      45 allocs/op
BenchmarkGoGenAvroEncode-8   	 1500506	       878.0 ns/op	     256 B/op	       3 allocs/op
BenchmarkHambaDecode-8       	 1898246	       558.9 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-8       	 3031268	       399.7 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-8    	  578799	      2084 ns/op	    1688 B/op	      35 allocs/op
BenchmarkLinkedinEncode-8    	 1506548	       704.0 ns/op	     248 B/op	       5 allocs/op
```