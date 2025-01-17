#include <stdio.h>
#include <string.h>

__global__ void kernel()
{
    printf("hello from GPU\n");
}

struct Error {
    int code;
    char *message;
};

extern "C"
#ifdef _WIN32
__declspec(dllexport)
#else
__attribute__((visibility("default")))
#endif
Error cuda_call()
{
    kernel<<<1,1>>>();
    auto err = cudaGetLastError();
    if (err != cudaSuccess)
        return {err, strdup(cudaGetErrorString(err))};
    err = cudaDeviceSynchronize();
    return {err, strdup(cudaGetErrorString(err))};
}
