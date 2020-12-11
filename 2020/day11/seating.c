#include <stdio.h>
#include <string.h>
#include <math.h>
#include "../lib/input.h"

#define MAX_DIM 100

char FLOOR = '.';
char EMPTY = 'L';
char OCCUPIED = '#';

void loadSeats(char seats[MAX_DIM][MAX_DIM], int *rowCount, int *colCount);
int arraysEqual(char a[MAX_DIM][MAX_DIM], char b[MAX_DIM][MAX_DIM], int rowCount, int colCount);
void updateSeats(char seats[MAX_DIM][MAX_DIM], char next[MAX_DIM][MAX_DIM], int rowCount, int colCount);
char getNextState(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, int col, int row);
int countState(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, char state);
void printSeats(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount);
void updateSeatsAgain(char seats[MAX_DIM][MAX_DIM], char next[MAX_DIM][MAX_DIM], int rowCount, int colCount);
char getNextStateAgain(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, int row, int col);


int main() {
  char seats[MAX_DIM][MAX_DIM], originalSeats[MAX_DIM][MAX_DIM], next[MAX_DIM][MAX_DIM];
  int rowCount = 0, colCount = 0;

  printf("Running with original update method\n");
  loadSeats(seats, &rowCount, &colCount);
  printSeats(seats, rowCount, colCount);
  memcpy(&originalSeats, &seats, sizeof(originalSeats));

  int iteration = 0;
  while (1) {
    printf("Iteration %d\n", iteration++);
    updateSeats(seats, next, rowCount, colCount);
    printSeats(next, rowCount, colCount);

    if (arraysEqual(seats, next, rowCount, colCount)) {
      break;
    }
    memcpy(&seats, &next, sizeof(seats));
  }
  printf("Final state\n");
  printSeats(next, rowCount, colCount);
  printf("Occupied seats = %d\n", countState(next, rowCount, colCount, OCCUPIED));

  printf("Running with new update method\n");
  memcpy(&seats, &originalSeats, sizeof(originalSeats));
  printSeats(seats, rowCount, colCount);

  iteration = 0;
  while (1) {
    printf("Iteration %d\n", iteration++);
    updateSeatsAgain(seats, next, rowCount, colCount);
    printSeats(next, rowCount, colCount);

    if (arraysEqual(seats, next, rowCount, colCount)) {
      break;
    }
    memcpy(&seats, &next, sizeof(seats));
  }
  printf("Final state\n");
  printSeats(next, rowCount, colCount);
  printf("Occupied seats = %d\n", countState(next, rowCount, colCount, OCCUPIED));
}

void loadSeats(char seats[MAX_DIM][MAX_DIM], int *rowCount, int *colCount) {
  int r = 0, length = 0;
  char row[MAX_DIM];

  while ((length = getLine(row)) != 0) {
    // don't include the newline
    *colCount = length - 1;
    for (int c = 0; c < *colCount; c++) {
      seats[r][c] = row[c];
    }
    r++;
  }
  *rowCount = r;
  printf("Seats have %d rows and %d columns\n", *rowCount, *colCount);
}


void updateSeats(char seats[MAX_DIM][MAX_DIM], char next[MAX_DIM][MAX_DIM], int rowCount, int colCount) {
  int nextState;
  for (int r = 0; r < rowCount; r++) {
    for (int c = 0; c < colCount; c++) {
      nextState = getNextState(seats, rowCount, colCount, r, c);
      next[r][c] = nextState;
    }
  }
}

char getNextState(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, int row, int col) {
  char current = seats[row][col];

  // no need to calculate anything for floor
  if (current == FLOOR) {
    return FLOOR;
  }

  // loop over all neighbours and count occupied
  int neighbours = 0;
  for (int r = fmax(row-1, 0); r < fmin(row+2, rowCount); r++) {
    for (int c = fmax(col-1, 0); c < fmin(col+2, colCount); c++) {
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

void updateSeatsAgain(char seats[MAX_DIM][MAX_DIM], char next[MAX_DIM][MAX_DIM], int rowCount, int colCount) {
  int nextState;
  for (int r = 0; r < rowCount; r++) {
    for (int c = 0; c < colCount; c++) {
      nextState = getNextStateAgain(seats, rowCount, colCount, r, c);
      next[r][c] = nextState;
    }
  }
}

char getNextStateAgain(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, int row, int col) {
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
  while (--r >= 0 && ++c < colCount) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // east
  r = row;
  c = col;
  while (++c < colCount) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south east
  r = row;
  c = col;
  while (++r < rowCount && ++c < colCount) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south
  r = row;
  c = col;
  while (++r < rowCount) {
    if (seats[r][c] != FLOOR) {
      neighbours += seats[r][c] == OCCUPIED;
      break;
    }
  }
  // south west
  r = row;
  c = col;
  while (++r < rowCount && --c >= 0) {
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

int arraysEqual(char a[MAX_DIM][MAX_DIM], char b[MAX_DIM][MAX_DIM], int rowCount, int colCount) {
  for (int r = 0; r < rowCount; r++) {
    for (int c = 0; c < colCount; c++) {
      if (a[r][c] != b[r][c]) {
        return 0;
      }
    }
  }
  return 1;
}

int countState(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount, char state) {
  int count = 0;
  for (int r = 0; r < rowCount; r++) {
    for (int c = 0; c < colCount; c++) {
      if (seats[r][c] == state) {
        count++;
      }
    }
  }
  return count;
}

void printSeats(char seats[MAX_DIM][MAX_DIM], int rowCount, int colCount) {
  int count = 0;
  for (int r = 0; r < rowCount; r++) {
    for (int c = 0; c < colCount; c++) {
      putchar(seats[r][c]);
    }
    putchar('\n');
  }
  putchar('\n');
}
