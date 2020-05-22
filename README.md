![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/hamba/avro](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)

#### Results

Go: 1.14 Machine: Macbook Pro 2,6 GHz Intel Core i7 16 GB 2133 MHz LPDDR3

```
BenchmarkGoAvroDecode-8     	  296842	      3868 ns/op	     442 B/op	      27 allocs/op
BenchmarkGoAvroEncode-8     	  257130	      4798 ns/op	     883 B/op	      63 allocs/op
BenchmarkHambaDecode-8      	 1885256	       621 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-8      	 2444282	       495 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-8   	  512622	      2344 ns/op	    1776 B/op	      40 allocs/op
BenchmarkLinkedinEncode-8   	 1275567	       975 ns/op	     288 B/op	      10 allocs/op
```