![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)
* [https://github.com/heetch/avro](https://github.com/heetch/avro)

#### Results

Go: 1.20.5 Machine: Macbook Pro M1

```
BenchmarkGoAvroDecode-10       	  759916	      1547 ns/op	     418 B/op	      27 allocs/op
BenchmarkGoAvroEncode-10       	  643017	      1874 ns/op	     800 B/op	      63 allocs/op
BenchmarkGoGenAvroDecode-10    	 1345588	       897.6 ns/op	     320 B/op	      11 allocs/op
BenchmarkGoGenAvroEncode-10    	 3037244	       396.3 ns/op	     240 B/op	       3 allocs/op
BenchmarkHambaDecode-10        	 5046981	       235.8 ns/op	      47 B/op	       0 allocs/op
BenchmarkHambaEncode-10        	 5976478	       201.0 ns/op	     112 B/op	       1 allocs/op
BenchmarkHeetchDecode-10       	 1314978	       908.5 ns/op	     640 B/op	      10 allocs/op
BenchmarkHeetchEncode-10       	 2938995	       408.2 ns/op	     512 B/op	       5 allocs/op
BenchmarkLinkedinDecode-10     	 1204290	      1004 ns/op	    1688 B/op	      35 allocs/op
BenchmarkLinkedinEncode-10     	 3165674	       380.6 ns/op	     248 B/op	       5 allocs/op
```