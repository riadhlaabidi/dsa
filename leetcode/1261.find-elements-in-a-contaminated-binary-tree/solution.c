#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

#define ELEMENTS_SIZE 1000001

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

typedef struct {
    int *elements;
} FindElements;

static void recover_bt(struct TreeNode *root, int x, FindElements *fe)
{
    if (root == NULL) {
        return;
    }

    if (x < ELEMENTS_SIZE) {
        fe->elements[x] = 1;
    }

    recover_bt(root->left, 2 * x + 1, fe);
    recover_bt(root->right, 2 * x + 2, fe);
}

FindElements *create(struct TreeNode *root)
{
    FindElements *fe = (FindElements *)malloc(sizeof(FindElements));
    if (fe == NULL) {
        fprintf(stderr, "Error allocating memory");
        return NULL;
    }

    int e[ELEMENTS_SIZE] = {0};
    fe->elements = e;
    recover_bt(root, 0, fe);
    return fe;
}

int find(FindElements *fe, int target) { return fe->elements[target]; }

void free_fe(FindElements *fe)
{
    free(fe);
    fe = NULL;
}

int main(void) { return 1; }
