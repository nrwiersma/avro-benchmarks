![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)

#### Results

Go: 1.20 Machine: Macbook Pro M1

```
BenchmarkGoAvroDecode-8      	  732993	      1570 ns/op	     418 B/op	      27 allocs/op
BenchmarkGoAvroEncode-8      	  589030	      1953 ns/op	     819 B/op	      63 allocs/op
BenchmarkGoGenAvroDecode-8   	 1338111	       896.1 ns/op	     320 B/op	      11 allocs/op
BenchmarkGoGenAvroEncode-8   	 2870955	       416.9 ns/op	     240 B/op	       3 allocs/op
BenchmarkHambaDecode-8       	 4168944	       287.9 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-8       	 5691699	       215.1 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-8    	 1000000	      1026 ns/op	    1688 B/op	      35 allocs/op
BenchmarkLinkedinEncode-8    	 3149337	       382.6 ns/op	     248 B/op	       5 allocs/op
```