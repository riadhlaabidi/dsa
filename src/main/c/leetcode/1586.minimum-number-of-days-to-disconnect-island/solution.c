#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#include <assert.h>

void dfs(int** grid, int row, int col, int n, int* m, int** visited) 
{
  if (row < 0 || row >= n || col < 0 || col >= m[row] || grid[row][col] == 0 || visited[row][col])
    return; 
   
  visited[row][col] = 1;

  int dirs[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
  for (int i = 0; i < 4; ++i)
    dfs(grid, row + dirs[i][0], col + dirs[i][1], n, m, visited);
}

int count_islands(int** grid, int gridSize, int* gridColSize) 
{
  int** visited = (int**) malloc(gridSize * sizeof(int*));

  if (visited == NULL) {
    fprintf(stderr, "Memory allocation failed\n");
    exit(EXIT_FAILURE);
  }

  for (int i = 0; i < gridSize; ++i) { 
    visited[i] = (int*) calloc(gridColSize[i], sizeof(int));

    if (visited[i] == NULL) {
      fprintf(stderr, "Memory allocation failed\n");
      exit(EXIT_FAILURE);
    }
  }

  int island_count = 0;

  for (int i = 0; i < gridSize; ++i) {
    for (int j = 0; j < gridColSize[i]; ++j) {
      if (grid[i][j] && !visited[i][j]) {
	dfs(grid, i, j, gridSize, gridColSize, visited);
	island_count++;
      }
    }
  } 

  free(visited);
  return island_count;
}
 
int min_days(int** grid, int gridSize, int* gridColSize)
{
  if (count_islands(grid, gridSize, gridColSize) != 1) {
    return 0;
  }   

  for (int i = 0; i < gridSize; ++i) {
    for (int j = 0; j < gridColSize[i]; ++j) {
      if (grid[i][j] == 0) continue;

      grid[i][j] = 0;
      
      if (count_islands(grid, gridSize, gridColSize) != 1)
	return 1;

      grid[i][j] = 1;
    }
  }
  return 2;
}

int main(void) 
{
  int grid_col_size[3] = {4, 4, 4};
  int c1[4] = {0,1,1,0};
  int c2[4] = {0,1,1,0};
  int c3[4] = {0,0,0,0};
  int* grid[3] = {c1, c2, c3};

  int expected = 2;
  int ans = min_days(grid, 3, grid_col_size);

  printf("Expected: %d\n", expected); 
  printf("Result: %d\n===> ", ans); 
  if (ans == expected) 
    printf("Correct");
  else 
    printf("Mismatch");

  return 0;
}
