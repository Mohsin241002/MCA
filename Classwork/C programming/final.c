#include <stdio.h>
#include <math.h>

int main() {
  int boxWidth = 20;
  int boxHeight = 10;

  int circle1X = 6;
  int circle1Y = 4;
  int circle2X = 11;
  int circle2Y = 4;

  int radius = 3;

  for (int y = 0; y < boxHeight; y++) {
    for (int x = 0; x < boxWidth; x++) {
      if (x == 0 || x == boxWidth - 1 || y == 0 || y == boxHeight - 1) {
        printf("*");
      } else {
        double distance1 = sqrt(pow(x - circle1X, 2) + pow(y - circle1Y, 2));
        double distance2 = sqrt(pow(x - circle2X, 2) + pow(y - circle2Y, 2));
        if (distance1 <= radius && distance2 <= radius) {
          printf("X"); // Intersection point
        } else if (distance1 <= radius) {
          printf("O"); // Circle 1
        } else if (distance2 <= radius) {
          printf("O"); // Circle 2
        } else {
          printf(" ");
        }
      }
    }
    printf("\n");
  }

  return 0;
}