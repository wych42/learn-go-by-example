run bench

```bash
go test -bench .
# or
go test -benchmem -run=^$ -bench ^(BenchmarkLoadOrStoreBalanced)$
```

#+RESULT
```plaintext
goos: darwin
goarch: amd64
pkg: github.com/wych42/learn-go-by-example/concurrent-map/safemap
BenchmarkLoadOrStoreBalanced/*safemap.WrapConcurrentMap-4         	10342911	       159 ns/op	      24 B/op	       2 allocs/op
BenchmarkLoadOrStoreBalanced/*safemap.SafeMap-4                   	 3915858	       398 ns/op	      24 B/op	       2 allocs/op
BenchmarkLoadOrStoreBalanced/*sync.Map-4                          	 6016010	       200 ns/op	      32 B/op	       2 allocs/op
PASS
ok  	github.com/wych42/learn-go-by-example/concurrent-map/safemap	5.049s
```
