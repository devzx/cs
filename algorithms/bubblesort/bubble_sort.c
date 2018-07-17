#include <stdio.h>

int bubbleSort(int arr[], int size) {
    int i, j, swap;
    for (i = 0; i < size; i++){
        for (j = 0; j < size - 1; j++) {
            if (arr[j] > arr[j+1]) {
                swap = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = swap;
            }
        }
    }
    for (i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }

    return 0;
}

int main() {
    int unsorted[] = {9, 2, 8, 3, 1, 10, 6, 7, 4, 5,};
    size_t size = sizeof(unsorted) / sizeof(unsorted[0]);
    bubbleSort(unsorted, size);

    return 0;
}
