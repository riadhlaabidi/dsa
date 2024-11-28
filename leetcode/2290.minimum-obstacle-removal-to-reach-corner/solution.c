#include <limits.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#include "../../ds/deque/deque.h"

typedef struct {
    Node node;
    int obstacles;
    int r;
    int c;
} Cell;

/**
 * create a new cell with embedded Deque node
 */
Cell *create_cell(int obs, int r, int c)
{
    Cell *new_cell = malloc(sizeof(Cell));

    if (!new_cell) {
        fprintf(stderr, "Failed to allocate memory");
        return NULL;
    }

    new_cell->obstacles = obs;
    new_cell->r = r;
    new_cell->c = c;

    return new_cell;
}

int is_inbound(int m, int n, int r, int c)
{
    return r >= 0 && r < m && c >= 0 && c < n;
}

/**
 * m: number of rows
 * n: number of columns
 */
int minimum_obstacles(int **grid, int m, int n)
{
    int dirs[][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

    int min_obstacles[m][n];

    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            min_obstacles[i][j] = INT_MAX;
        }
    }

    min_obstacles[0][0] = 0;

    Deque d;
    DEQUE_INIT(&d);

    Cell *c = create_cell(0, 0, 0);
    DEQUE_ADD_FRONT(&d, &c->node);

    while (!DEQUE_IS_EMPTY(&d)) {
        Cell *curr = DEQUE_DATA(DEQUE_GET_FRONT(&d), Cell, node);
        DEQUE_POP_FRONT(&d);

        int obs = curr->obstacles;
        int r = curr->r;
        int c = curr->c;

        for (int i = 0; i < 4; ++i) {
            int new_r = r + dirs[i][0];
            int new_c = c + dirs[i][1];

            if (is_inbound(m, n, new_r, new_c) &&
                min_obstacles[new_r][new_c] == INT_MAX) {

                if (grid[new_r][new_c] == 1) {
                    min_obstacles[new_r][new_c] = obs + 1;
                    Cell *new_cell = create_cell(obs + 1, new_r, new_c);
                    DEQUE_ADD_BACK(&d, &new_cell->node);
                } else {
                    min_obstacles[new_r][new_c] = obs;
                    Cell *new_cell = create_cell(obs, new_r, new_c);
                    DEQUE_ADD_FRONT(&d, &new_cell->node);
                }
            }
        }
    }

    // cleanup
    Node *tmp = d.head;
    while (tmp != NULL) {
        Node *next = tmp->next;
        Cell *c = DEQUE_DATA(tmp, Cell, node);
        DEQUE_POP_FRONT(&d);
        free(c);
        tmp = next;
    }

    return min_obstacles[m - 1][n - 1];
}

int main(void)
{
    int m = 3;
    int n = 3;
    int *grid[] = {(int[]){0, 1, 1}, (int[]){1, 1, 0}, (int[]){1, 1, 0}};

    int expected = 2;
    int actual = minimum_obstacles(grid, m, n);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);
    return EXIT_SUCCESS;
}
