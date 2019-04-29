include github.com/hamba/make/golang

# Run benchmarks
bench:
	@go test -bench=. .
.PHONY: bench
