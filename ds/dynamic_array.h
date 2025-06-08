#ifndef DS_DYNAMIC_ARRAY_H
#define DS_DYNAMIC_ARRAY_H

#include "stdlib.h"

#define DARRAY_DEFAULT_CAPACITY 32

#define darray_realloc(darray, new_capacity)                                   \
    do {                                                                       \
        if ((new_capacity) > (darray)->capacity) {                             \
            if ((darray)->capacity == 0) {                                     \
                (darray)->capacity = DARRAY_DEFAULT_CAPACITY;                  \
            }                                                                  \
            while ((new_capacity) > (darray)->capacity) {                      \
                (darray)->capacity *= 2;                                       \
            }                                                                  \
            (darray)->elements =                                               \
                realloc((darray)->elements,                                    \
                        (darray)->capacity * sizeof(*(darray)->elements));     \
        }                                                                      \
    } while (0)

#define darray_append(darray, item)                                            \
    do {                                                                       \
        darray_realloc((darray), (darray)->count + 1);                         \
        (darray)->elements[(darray)->count++] = (item);                        \
    } while (0)

#define darray_free(darray)                                                    \
    do {                                                                       \
        free((darray).elements);                                               \
    } while (0)

#endif /* end of include guard: DS_DYNAMIC_ARRAY_H */
