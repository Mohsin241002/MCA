#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define WIDTH  41
#define HEIGHT 41

int maze[WIDTH][HEIGHT];

// Directions: Right, Down, Left, Up
int dx[] = {1, 0, -1, 0};
int dy[] = {0, 1, 0, -1};

void initialize_maze() {
    for (int x = 0; x < WIDTH; x++) {
        for (int y = 0; y < HEIGHT; y++) {
            maze[x][y] = 1; // 1 for wall
        }
    }
}

void carve_passage(int x, int y) {
    int dir = rand() % 4;
    int count = 0;
    while (count < 4) {
        int nx = x + dx[dir];
        int ny = y + dy[dir];
        int nx2 = nx + dx[dir];
        int ny2 = ny + dy[dir];
        if (nx2 > 0 && ny2 > 0 && nx2 < WIDTH && ny2 < HEIGHT && maze[nx2][ny2] == 1) {
            maze[nx][ny] = 0; // 0 for path
            maze[nx2][ny2] = 0;
            carve_passage(nx2, ny2);
        }
        dir = (dir + 1) % 4;
        count++;
    }
}

void generate_maze() {
    initialize_maze();
    carve_passage(1, 1); // Start carving from the top-left corner
    maze[1][0] = 0; // Entry point
    maze[WIDTH - 2][HEIGHT - 1] = 0; // Exit point
}

void print_maze() {
    for (int y = 0; y < HEIGHT; y++) {
        for (int x = 0; x < WIDTH; x++) {
            if (maze[x][y] == 1) {
                printf("|");
            } else {
                printf(" ");
            }
        }
        printf("\n");
    }
}

int main() {
    srand(time(NULL));

    generate_maze();
    print_maze();

    return 0;
}
