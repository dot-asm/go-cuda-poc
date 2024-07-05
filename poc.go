package main

// #cgo linux linux LDFLAGS: -ldl -Wl,-rpath,"$ORIGIN"
// #ifdef _WIN32
// # include <windows.h>
// # include <stdio.h>
// #else
// # include <dlfcn.h>
// # include <errno.h>
// #endif
// #include <stdlib.h>
// #include <string.h>
//
// typedef struct {
//     int code;
//     char *message;
// } Error;
//
// static Error (*cuda_call)(/* args */);
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
// static _Bool load(_GoString_ *go_err)
// {
//     if (cuda_call == NULL) {
// #ifdef _WIN32
//         HMODULE h = LoadLibraryA("poc.dll");
//         if (h == NULL && system("nvcc -shared -o poc.dll poc.cu") == 0) {
//             h = LoadLibraryA("poc.dll");
//         }
//         if (h != NULL) cuda_call = (void*)GetProcAddress(h, "cuda_call");
//         if (cuda_call == NULL) {
//             static char buf[24];
//             snprintf(buf, sizeof(buf), "WIN32 Error #0x%x", GetLastError());
//             toGoString(go_err, buf);
//         }
// #else
//         void* h = dlopen("poc.so", RTLD_NOW|RTLD_GLOBAL);
//         if (h == NULL &&
//             system("nvcc -shared -o poc.so poc.cu -cudart shared "
//                         "-Xcompiler -fPIC,-fvisibility=hidden "
//                         "-Xlinker -Bsymbolic") == 0) {
//             h = dlopen("./poc.so", RTLD_NOW|RTLD_GLOBAL);
//         }
//         if (h != NULL) cuda_call = dlsym(h, "cuda_call");
//         if (cuda_call == NULL) {
//             toGoString(go_err, dlerror());
//         }
// #endif
//     }
//
//     return cuda_call != NULL;
// }
//
// static void bridge(GoError *go_err /*, args */)
// {   toGoError(go_err, cuda_call(/* args */));   }
import "C"

//export toGoString
func toGoString(go_str *string, c_str *C.char) {
    *go_str = C.GoString(c_str)
}

func init() {
    var err string
    if !C.load(&err) {
        panic(err)
    }
}

func main() {
    var err C.GoError
    C.bridge(&err /*, args */)
    if err.code != 0 {
        panic(err.message)
    }
}
