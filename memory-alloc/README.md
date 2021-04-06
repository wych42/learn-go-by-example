```bash
go test -bench . -benchmem
go test -run TestCopyIt -trace=copy_trace.out
go tool trace copy_trace.out
```