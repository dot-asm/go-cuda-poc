package main

// #cgo LDFLAGS: -Wl,-rpath,"$ORIGIN"
//
// extern __attribute__((weak)) void cuda_call();
//
// static void bridge()
// {
//     if (cuda_call) cuda_call();
// }
import "C"

func main() {
    C.bridge()
}
