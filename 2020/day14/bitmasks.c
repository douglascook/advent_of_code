#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include "../lib/input.h"

void updateMasks(char *line, long long *trueMask, long long *falseMask);
void updateMemory(long long *memory, char *line, long long trueMask, long long falseMask);

int main() {
  long long trueMask, falseMask;
  long long memory[100000] = {0};

  char line[100];
  while (getLine(line) != 0) {
    if (strstr(line, "mask = ") == line) {
      updateMasks(line, &trueMask, &falseMask);
    } else {
      updateMemory(memory, line, trueMask, falseMask);
    }
  }
  long long total = 0;
  for (int i = 0; i < 100000; i++) {
    total += memory[i];
  }
  printf("Total of all values remaining in memory is %lld\n", total);
}

/* Parse out new values for true and false masks from the given line */
void updateMasks(char *line, long long *trueMask, long long *falseMask) {
  printf("Updating masks\n");
  // reset both to zero first
  *trueMask = *falseMask = 0;

  int exp, charIdx;
  for (int i = 0; i < 36; i++) {
    exp = 35 - i;
    charIdx = 7 + i;
    if (line[charIdx] == '1') {
      *trueMask |= (long long)pow(2, exp);
    } else if (line[charIdx] == '0') {
      *falseMask |= (long long)pow(2, exp);
    }
  }
  printf("True mask integer value is %lld\n", *trueMask);
  printf("False mask integer value is %lld\n", *falseMask);
}

/* Update memory, assigning the masks modified value to the the specified address */
void updateMemory(long long *memory, char *line, long long trueMask, long long falseMask) {
  char *next;
  int address = strtol(&line[4], &next, 10);

  // parse out the original value
  long long value = strtol(strstr(line, "=") + 1, &next, 10);

  // OR to set all matching bits
  value |= trueMask;

  // OR then SUBTRACT to clear all matching bits
  value |= falseMask;
  value -= falseMask;

  memory[address] = value;
}
