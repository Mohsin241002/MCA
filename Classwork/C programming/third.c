#include <stdio.h>

int main() {
  int boxWidth = 20;
  int boxHeight = 10;
  int circle1X = 6;
  int circle1Y = 4;
  int circle2X = 14;
  int circle2Y = 4;
  int radius = 3;

  for (int y = 0; y < boxHeight; y++) {
    for (int x = 0; x < boxWidth; x++) {
      if (x == 0 || x == boxWidth - 1 || y == 0 || y == boxHeight - 1) {
        printf("*");
      } else if ((x - circle1X) * (x - circle1X) + (y - circle1Y) * (y - circle1Y) <= radius * radius &&
                 (x - circle2X) * (x - circle2X) + (y - circle2Y) * (y - circle2Y) <= radius * radius) {
        printf("X");
      } else if ((x - circle1X) * (x - circle1X) + (y - circle1Y) * (y - circle1Y) <= radius * radius) {
        printf("O");
      } else if ((x - circle2X) * (x - circle2X) + (y - circle2Y) * (y - circle2Y) <= radius * radius) {
        printf("O");
      } else {
        printf(" ");
      }
    }
    printf("\n");
  }

  return 0;
}