#include <assert.h>
#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_COL 26

typedef struct {
    int **grid;
    int rows;
} Spreadsheet;

Spreadsheet *spreadsheetCreate(int rows)
{
    Spreadsheet *new_spreadsheet = malloc(sizeof(Spreadsheet));
    if (!new_spreadsheet) {
        fprintf(stderr, "Failed to allocate memory");
        return NULL;
    }

    int **grid = calloc(rows, sizeof(int *));
    if (!grid) {
        fprintf(stderr, "Failed to allocate memory");
        free(new_spreadsheet);
        return NULL;
    }

    for (int i = 0; i < rows; i++) {
        grid[i] = calloc(MAX_COL, sizeof(int));
        if (!grid[i]) {
            for (int j = 0; j < i; j++) {
                free(grid[j]);
            }
            free(grid);
            free(new_spreadsheet);
            return NULL;
        }
    }

    new_spreadsheet->rows = rows;
    new_spreadsheet->grid = grid;
    return new_spreadsheet;
}

int *get_cell(Spreadsheet *obj, char *cell)
{
    int j = cell[0] - 'A';
    cell++;
    int i = atoi(cell) - 1;
    return &obj->grid[i][j];
}

void spreadsheetSetCell(Spreadsheet *obj, char *cell, int value)
{
    *(get_cell(obj, cell)) = value;
}

void spreadsheetResetCell(Spreadsheet *obj, char *cell)
{
    *(get_cell(obj, cell)) = 0;
}

int spreadsheetGetValue(Spreadsheet *obj, char *formula)
{
    size_t i;
    for (i = 0; i < strlen(formula); i++) {
        if (formula[i] == '+') {
            i++;
            break;
        }
    }

    int x = isalpha(formula[1]) ? *(get_cell(obj, formula + 1))
                                : atoi(formula + 1);

    int y = isalpha(formula[i]) ? *(get_cell(obj, formula + i))
                                : atoi(formula + i);
    return x + y;
}

void spreadsheetFree(Spreadsheet *obj)
{
    if (!obj) {
        return;
    }
    if (obj->grid) {
        for (int i = 0; i < obj->rows; i++) {
            free(obj->grid[i]);
        }
        free(obj->grid);
    }

    free(obj);
}

#define assert_equals(expected, actual)                                        \
    do {                                                                       \
        if (actual != expected) {                                              \
            printf("Expected %d, but got %d\n", expected, actual);             \
            status = 1;                                                        \
            goto defer;                                                        \
        }                                                                      \
    } while (0)

int main(void)
{

    int status = 0;

    Spreadsheet *s = spreadsheetCreate(458);
    assert(s);

    int actual = spreadsheetGetValue(s, "=O126+10272");
    assert_equals(10272, actual);
    actual = spreadsheetGetValue(s, "=5+7");
    assert_equals(12, actual);
    spreadsheetSetCell(s, "A1", 10);
    actual = spreadsheetGetValue(s, "=A1");
    assert_equals(10, actual);
    actual = spreadsheetGetValue(s, "=A1+6");
    assert_equals(16, actual);
    spreadsheetSetCell(s, "B2", 15);
    actual = spreadsheetGetValue(s, "=B2");
    assert_equals(15, actual);
    actual = spreadsheetGetValue(s, "=A1+B2");
    assert_equals(25, actual);
    spreadsheetResetCell(s, "A1");
    actual = spreadsheetGetValue(s, "=A1");
    assert_equals(0, actual);
    actual = spreadsheetGetValue(s, "=A1+B2");
    assert_equals(15, spreadsheetGetValue(s, "=A1+B2"));
    printf("Correct\n");

defer:
    spreadsheetFree(s);
    return status;
}
