#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define WIDTH  21
#define HEIGHT 25

int maze[HEIGHT][WIDTH];

void initialize_maze() {
    for (int y = 0; y < HEIGHT; y++) {
        for (int x = 0; x < WIDTH; x++) {
            maze[y][x] = 1;
        }
    }
}

void generate_maze(int x, int y) {
    int directions[] = { 0, 1, 2, 3 };
    for (int i = 0; i < 4; i++) {
        int r = rand() % 4;
        int temp = directions[r];
        directions[r] = directions[i];
        directions[i] = temp;
    }

    for (int i = 0; i < 4; i++) {
        int dx = 0, dy = 0;
        switch (directions[i]) {
            case 0: dx = 1; break;
            case 1: dy = 1; break;
            case 2: dx = -1; break;
            case 3: dy = -1; break;
        }

        int nx = x + dx * 2;
        int ny = y + dy * 2;

        if (nx >= 0 && nx < WIDTH && ny >= 0 && ny < HEIGHT && maze[ny][nx] == 1) {
            maze[ny - dy][nx - dx] = 0;
            maze[ny][nx] = 0;
            generate_maze(nx, ny);
        }
    }
}

void print_maze() {
    for (int y = 0; y < HEIGHT; y++) {
        for (int x = 0; x < WIDTH; x++) {
            if (maze[y][x] == 1) {
                printf("#");
            } else {
                printf(" ");
            }
        }
        printf("\n");
    }
}

int main() {
    srand(time(NULL));
    initialize_maze();
    maze[1][1] = 0; // Start point
    generate_maze(1, 1);
    print_maze();
    return 0;
}
