package main

// #cgo unix LDFLAGS: -ldl -Wl,-rpath,"$ORIGIN"
// #ifdef _WIN32
// # include <windows.h>
// #else
// # include <dlfcn.h>
// # include <stdlib.h>
// #endif
//
// static void (*cuda_call)();
//
// static void bridge()
// {
//     if (cuda_call == NULL) {
// #ifdef _WIN32
//         HMODULE h = LoadLibraryA("poc.dll");
//         if (h == NULL && system("nvcc -shared -o poc.dll poc.cu") == 0) {
//             h = LoadLibraryA("poc.dll");
//         }
//         if (h != NULL) cuda_call = (void*)GetProcAddress(h, "cuda_call");
// #else
//         void* h = dlopen("poc.so", RTLD_NOW|RTLD_GLOBAL);
//         if (h == NULL &&
//             system("nvcc -shared -o poc.so poc.cu -cudart shared "
//                         "-Xcompiler -fPIC,-fvisibility=hidden "
//                         "-Xlinker -Bsymbolic") == 0) {
//             h = dlopen("./poc.so", RTLD_NOW|RTLD_GLOBAL);
//         }
//         if (h != NULL) cuda_call = dlsym(h, "cuda_call");
// #endif
//     }
//
//     if (cuda_call) cuda_call();
// }
import "C"

func main() {
    C.bridge()
}
