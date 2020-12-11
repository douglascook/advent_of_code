#include <stdio.h>
#include <string.h>
#include <math.h>
#include "../lib/input.h"

#define MAX_DIM 100

char FLOOR = '.';
char EMPTY = 'L';
char OCCUPIED = '#';
char ORIGINAL_SEATS[MAX_DIM][MAX_DIM];
int ROW_COUNT;
int COL_COUNT;

void loadSeats();
void run(char (*updateMethod)());
char getNextState(char seats[MAX_DIM][MAX_DIM], int col, int row);
char getNextState2(char seats[MAX_DIM][MAX_DIM], int row, int col);
int arraysEqual(char a[MAX_DIM][MAX_DIM], char b[MAX_DIM][MAX_DIM]);
int countState(char seats[MAX_DIM][MAX_DIM], char state);
void printSeats(char seats[MAX_DIM][MAX_DIM]);


int main() {
  loadSeats();
  printf("Running with first update method\n");
  run(getNextState);
  printf("\nRunning with second update method\n");
  run(getNextState2);
}

void run(char (*updateMethod)()) {
  char seats[MAX_DIM][MAX_DIM], next[MAX_DIM][MAX_DIM];
  memcpy(&seats, &ORIGINAL_SEATS, sizeof(seats));
  printf("Initial state\n");
  printSeats(seats);

  int iteration = 0, nextState;
  while (1) {
    /* printf("Iteration %d\n", iteration++); */
    for (int r = 0; r < ROW_COUNT; r++) {
      for (int c = 0; c < COL_COUNT; c++) {
        nextState = updateMethod(seats, r, c);
        next[r][c] = nextState;
      }
    }
    /* printSeats(next); */

    if (arraysEqual(seats, next)) {
      break;
    }
    memcpy(&seats, &next, sizeof(seats));
  }
  printf("Final state\n");
  printSeats(next);
  printf("Occupied seats = %d\n", countState(next, OCCUPIED));
}

void loadSeats() {
  int r = 0, length = 0;
  char row[MAX_DIM];

  while ((length = getLine(row)) != 0) {
    // don't include the newline
    COL_COUNT = length - 1;
    for (int c = 0; c < COL_COUNT; c++) {
      ORIGINAL_SEATS[r][c] = row[c];
    }
    r++;
  }
  ROW_COUNT = r;
  printf("There are %d rows and %d columns of seats\n", ROW_COUNT, COL_COUNT);
}


/* Consider immediate neighbours when looking for occupied seats */
char getNextState(char seats[MAX_DIM][MAX_DIM], int row, int col) {
  char current = seats[row][col];
  // no need to calculate anything for floor
  if (current == FLOOR) {
    return FLOOR;
  }

  // loop over all neighbours and count occupied
  int neighbours = 0;
  for (int r = fmax(row-1, 0); r < fmin(row+2, ROW_COUNT); r++) {
    for (int c = fmax(col-1, 0); c < fmin(col+2, COL_COUNT); c++) {
      // don't consider the seat itself
      if (r == row && c == col) {
        continue;
      }
      neighbours += seats[r][c] == OCCUPIED;
    }
  }
  // empty seat with no occupied neighbours becomes occupied
  if (current == EMPTY && neighbours == 0) {
    return OCCUPIED;
  }
  // occupied seat with at least 4 neighbours becomes empty
  if (current == OCCUPIED && neighbours >= 4) {
    return EMPTY;
  }
  // otherwise nothing changes
  return current;
}

/* Consider the first non-floor cell when looking for occupied seats, if the
 * immediate neigbour is a floor cell then look at the next in that direction */
char getNextState2(char seats[MAX_DIM][MAX_DIM], int row, int col) {
  char current = seats[row][col];
  // no need to calculate anything for floor
  if (current == FLOOR) {
    return FLOOR;
  }

  int r, c, neighbours = 0;
  // north west
  r = row;
  c = col;
  while (--r >= 0 && --c >= 0) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // north
  r = row;
  c = col;
  while (--r >= 0) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // north east
  r = row;
  c = col;
  while (--r >= 0 && ++c < COL_COUNT) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // east
  r = row;
  c = col;
  while (++c < COL_COUNT) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south east
  r = row;
  c = col;
  while (++r < ROW_COUNT && ++c < COL_COUNT) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south
  r = row;
  c = col;
  while (++r < ROW_COUNT) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south west
  r = row;
  c = col;
  while (++r < ROW_COUNT && --c >= 0) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // west
  r = row;
  c = col;
  while (--c >= 0) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // empty seat with no occupied neighbours becomes occupied
  if (current == EMPTY && neighbours == 0) {
    return OCCUPIED;
  }
  // occupied seat with at least 4 neighbours becomes empty
  if (current == OCCUPIED && neighbours >= 5) {
    return EMPTY;
  }
  // otherwise nothing changes
  return current;
}

int arraysEqual(char a[MAX_DIM][MAX_DIM], char b[MAX_DIM][MAX_DIM]) {
  for (int r = 0; r < ROW_COUNT; r++) {
    for (int c = 0; c < COL_COUNT; c++) {
      if (a[r][c] != b[r][c]) {
        return 0;
      }
    }
  }
  return 1;
}

int countState(char seats[MAX_DIM][MAX_DIM], char state) {
  int count = 0;
  for (int r = 0; r < ROW_COUNT; r++) {
    for (int c = 0; c < COL_COUNT; c++) {
      if (seats[r][c] == state) {
        count++;
      }
    }
  }
  return count;
}

void printSeats(char seats[MAX_DIM][MAX_DIM]) {
  for (int r = 0; r < ROW_COUNT; r++) {
    for (int c = 0; c < COL_COUNT; c++) {
      putchar(seats[r][c]);
    }
    putchar('\n');
  }
  putchar('\n');
}
