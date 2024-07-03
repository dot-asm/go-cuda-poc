package main

// #cgo unix LDFLAGS: -ldl -Wl,-rpath,"$ORIGIN"
// #ifdef _WIN32
// # include <windows.h>
// #else
// # include <dlfcn.h>
// # include <stdlib.h>
// #endif
//
// typedef struct {
//     int code;
//     char *message;
// } Error;
//
// static Error (*cuda_call)();
//
// typedef struct {
//     int code;
//     _GoString_ message;
// } GoError;
//
// void toGoString(_GoString_ *, char *);
//
// static void toGoError(GoError *go_err, Error c_err)
// {
//     go_err->code = c_err.code;
//     if (c_err.message != NULL) {
//         toGoString(&go_err->message, c_err.message);
//         free(c_err.message);
//         c_err.message = NULL;
//     }
// }
//
// static void bridge(GoError *go_err)
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
//     if (cuda_call) toGoError(go_err, cuda_call());
// }
import "C"

//export toGoString
func toGoString(go_str *string, c_str *C.char) {
    *go_str = C.GoString(c_str)
}

func main() {
    var err C.GoError
    C.bridge(&err)
    if err.code != 0 {
        panic(err.message)
    }
}
