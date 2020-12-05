/* 7 bit row number -> max = 1111111 == 127
 * 3 bit col number -> max = 111 = 7
 */
#include <stdio.h>
#include "../lib/input.h"

int parseRowNumber(const char* line);
int parseColNumber(const char* line);
int parseBinary(const char* line, int magnitude, char setChar);

int main() {
  char line[20];
  int row, col, seatId, highestId = 0, lowestId = 1024;
  int occupiedSeats[1024];

  while (getLine(line) != 0) {
    row = parseBinary(line, 7, 'B');
    col = parseBinary(&line[7], 3, 'R');

    seatId = row*8 + col;
    occupiedSeats[seatId] = 1;

    if (seatId > highestId) {
      highestId = seatId;
    }
    if (seatId < lowestId) {
      lowestId = seatId;
    }
  }
  printf("Highest seat ID = %d\n", highestId);

  for (int i = lowestId; i < highestId; i++) {
    if (occupiedSeats[i] != 1) {
      printf("My seat is %d\n", i);
    }
  }
}

/* Parse the given string into binary where bits are set with the isSet value */
int parseBinary(const char* string, int magnitude, char isSet) {
  int result = 0;
  for (int i = 0; i < magnitude; i++) {
    if (string[i] == isSet) {
      // bit shifting 1 << x gives us 2**x
      result |= 1 << (magnitude - i - 1);
    }
  }
  return result;
}
