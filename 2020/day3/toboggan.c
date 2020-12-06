#include <stdio.h>
#include "../lib/input.h"

#define TREE '#'
#define OPEN '.'

#define ROWS 1000
#define COLS 100

int checkRoute(char route[ROWS][COLS], int length, int width, int right, int down);

int main() {
  char line[100];
  char route[ROWS][COLS];

  int width, w, i;
  int length = 0;

  // TODO work out how to return/update height and width and refactor into function
  while((w = getLine(line)) > 0) {
    // TODO better way to store width without using a temp variable?
    // ignore newline char
    width = w - 1;
    for (i = 0; i < width; i++) {
      route[length][i] = line[i];
    }
    route[length][i] = '\0';
    length++;
  }

  int treeProduct = 1;
  treeProduct = treeProduct * checkRoute(route, length, width, 1, 1);
  treeProduct = treeProduct * checkRoute(route, length, width, 3, 1);
  treeProduct = treeProduct * checkRoute(route, length, width, 5, 1);
  treeProduct = treeProduct * checkRoute(route, length, width, 7, 1);
  treeProduct = treeProduct * checkRoute(route, length, width, 1, 2);
  printf("Product of trees encountered on all routes = %d\n", treeProduct);
}

int checkRoute(char route[ROWS][COLS], int length, int width, int right, int down) {
  int x = 0, y = 0;
  int treeCount = 0;
  while (x <= length) {
    y = (y + right) % width;
    x = x + down;
    if (route[x][y] == TREE) {
      treeCount++;
      /* printf("Tree at: x = %d, y = %d\t\t %s\n", x, y, route[x]); */
    }
  }
  printf("Right = %d, Down = %d -> Tree Count = %d\n", right, down, treeCount);
  return treeCount;
}
