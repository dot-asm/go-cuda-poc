# go-cuda-poc

```
$ go run poc.go
hello from GPU
$ 
```

Recall that on Windows [`cgo`](https://pkg.go.dev/cmd/cgo) depends on MinGW GCC being available on your %PATH%. In addition to which this PoC requires [`nvcc`](https://developer.nvidia.com/cuda-downloads?target_os=Windows) and even [`cl`](https://learn.microsoft.com/en-us/visualstudio/ide/reference/command-prompt-powershell).
