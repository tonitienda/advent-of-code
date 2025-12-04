#include <stdio.h>
#include <stdlib.h>

int main(void) {
    size_t len = 0;
    char *data = read_file("day01/input.txt", &len);
    if (!data) {
        fprintf(stderr, "Error reading input\n");
        return 1;
    }

    // Solve here…

    free(data);
    return 0;
}
