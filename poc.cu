#include <stdio.h>

__global__ void kernel()
{
    printf("hello from GPU\n");
}

extern "C" __attribute__((visibility("default")))
void cuda_call()
{
    kernel<<<1,1>>>();
    cudaDeviceSynchronize();
}
