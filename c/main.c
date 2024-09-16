#include <stdio.h>

void main() {
    int numbers[5] = {11, 22, 33, 44, 55};  
    printf("number of array items: %ld,\n", sizeof(numbers)/sizeof(int));

    for (int i=0; i< 5; ++i) {
         printf("index: %d: %d,\n", i, numbers[i]);
    }
}