![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)

#### Results

Go: 1.16 Machine: Macbook Pro 2,6 GHz Intel Core i7 16 GB 2133 MHz LPDDR3

```
BenchmarkGoAvroDecode-8      	  298760	      3770 ns/op	     418 B/op	      27 allocs/op
BenchmarkGoAvroEncode-8      	  248575	      4611 ns/op	     893 B/op	      63 allocs/op
BenchmarkGoGenAvroDecode-8   	  411472	      2722 ns/op	     728 B/op	      45 allocs/op
BenchmarkGoGenAvroEncode-8   	 1220017	       973.0 ns/op	     256 B/op	       3 allocs/op
BenchmarkHambaDecode-8       	 1936148	       607.3 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-8       	 2679397	       453.7 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-8    	  558864	      2216 ns/op	    1688 B/op	      35 allocs/op
BenchmarkLinkedinEncode-8    	 1536460	       770.4 ns/op	     248 B/op	       5 allocs/op
```