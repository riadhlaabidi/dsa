#include <stdio.h>
#include <assert.h>

void sort (int *arr, int n) 
{
  for (int i = 0; i < n - 1; ++i) {
    int min = i;
    for (int j = i + 1; j < n; ++j) {
      if (arr[j] < arr[min]) 
	min = j;
    }

    arr[i] = arr[i] ^ arr[min];
    arr[min] = arr[i] ^ arr[min];
    arr[i] = arr[i] ^ arr[min];
  }
}

int is_sorted(int *arr, int n) 
{
  for (int i = 0; i < n - 1; ++i) { 
    if (arr[i] > arr[i + 1])
      return 0;
  }

  return 1;
}

int main(void)
{
  int n = 10;
  int arr[] = {5, 2, 63, 0, 114, 86, 99, 4, 2, 76};
  int i;

  sort(arr, n);
  assert(is_sorted(arr, n));

  for (i = 0; i < n; ++i) 
    printf("%d ", arr[i]);

  return 0;
}
