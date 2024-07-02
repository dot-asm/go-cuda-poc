package main

// #cgo LDFLAGS: -ldl -Wl,-rpath,"$ORIGIN"
// #include <dlfcn.h>
// #include <stdlib.h>
//
// static void (*cuda_call)();
//
// static void bridge()
// {
//     if (cuda_call == NULL) {
//         void* h = dlopen("poc.so", RTLD_NOW|RTLD_GLOBAL);
//         if (h == NULL &&
//             system("nvcc -shared -o poc.so poc.cu -cudart shared "
//                         "-Xcompiler -fPIC,-fvisibility=hidden "
//                         "-Xlinker -Bsymbolic") == 0) {
//             h = dlopen("./poc.so", RTLD_NOW|RTLD_GLOBAL);
//         }
//         if (h != NULL) cuda_call = dlsym(h, "cuda_call");
//     }
//
//     if (cuda_call) cuda_call();
// }
import "C"

func main() {
    C.bridge()
}
