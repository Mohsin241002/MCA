#include <stdio.h>

#define WIDTH  21
#define HEIGHT 21

extern int maze[HEIGHT][WIDTH];
int solution[HEIGHT][WIDTH];

int dfs(int x, int y) {
    if (x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT || maze[y][x] == 1 || solution[y][x] == 1) {
        return 0;
    }
    solution[y][x] = 1;

    if (x == WIDTH - 2 && y == HEIGHT - 2) {
        return 1;
    }

    if (dfs(x + 1, y) || dfs(x, y + 1) || dfs(x - 1, y) || dfs(x, y - 1)) {
        return 1;
    }

    solution[y][x] = 0;
    return 0;
}

void print_solution() {
    for (int y = 0; y < HEIGHT; y++) {
        for (int x = 0; x < WIDTH; x++) {
            if (solution[y][x] == 1) {
                printf(".");
            } else {
                printf("%c", maze[y][x] == 1 ? '#' : ' ');
            }
        }
        printf("\n");
    }
}

int main() {
    // Assuming maze is already generated
    if (dfs(1, 1)) {
        printf("Maze solved:\n");
        print_solution();
    } else {
        printf("No solution found.\n");
    }
    return 0;
}
