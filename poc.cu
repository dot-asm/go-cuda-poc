#include <stdio.h>

__global__ void kernel()
{
    printf("hello from GPU\n");
}

extern "C"
#ifdef _WIN32
__declspec(dllexport)
#else
__attribute__((visibility("default")))
#endif
void cuda_call()
{
    kernel<<<1,1>>>();
    cudaDeviceSynchronize();
}
