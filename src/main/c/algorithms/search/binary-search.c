#include <assert.h>
#include <stdio.h>

int search(int *arr, int n, int needle)
{
    int low = 0;
    int high = n;

    while (low < high) {
	int mid = (low + high) / 2;

	if (arr[mid] == needle) {
	    // needle found, return it's index 
	    return mid;
	} else if (arr[mid] < needle) {
	    low = mid + 1; 
	} else {
	    high = mid;
	}
    } 

    // negative if needle is not found
    return -1;
}

int main(void) 
{
    int arr[7] = {1, 3, 4, 9, 9, 190, 322}; 

    int needle1 = 1;
    int needle2 = 9;
    int needle3 = 322;
    int needle4 = 455;

    int i1 = search(arr, 7, needle1); 
    int i2 = search(arr, 7, needle2); 
    int i3 = search(arr, 7, needle3); 
    int i4 = search(arr, 7, needle4); 

    assert(i1 == 0 && arr[i1] == needle1);
    assert(i2 == 3 && arr[i2] == needle2);
    assert(i3 == 6 && arr[i3] == needle3);
    assert(i4 == -1);

    for (int i = 0; i < 7; ++i) {
	printf("%d ", arr[i]);
    } 

    printf("\nNeedle %d @ index %d\n", needle1, i1); 
    printf("Needle %d @ index %d\n", needle2, i2); 
    printf("Needle %d @ index %d\n", needle3, i3); 
    printf("Needle %d @ index %d\n", needle4, i4); 
}
