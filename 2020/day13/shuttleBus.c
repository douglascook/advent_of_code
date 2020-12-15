#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include "../lib/input.h"

void getBuses(char *start, int *buses, int *busCount);
int compareInt(const void* a, const void* b);
void checkFirstBusWeCanCatch(int *buses, int busCount, int departureTime);

int main() {
  char input[1000], *start;

  getInput(input);
  int departureTime = strtol(input, &start, 10);
  printf("Departure time is %d\n", departureTime);

  int buses[200], busCount = 0;
  getBuses(start, buses, &busCount);

  checkFirstBusWeCanCatch(buses, busCount, departureTime);
}


/* Return a sorted array containing all bus IDs */
void getBuses(char *start, int *buses, int *busCount) {
  char *next;
  int busId;
  while (*start != '\0') {
    busId = strtol(start, &next, 10);
    // if we're at an integer then add it to the array of bus IDs
    if (busId != 0) {
      buses[(*busCount)++] = busId;
      start = next;
    // otherwise increment pointer so we continue at next char
    } else {
      start = start + 1;
    }
  }
  qsort(buses, *busCount, sizeof(int), compareInt);
}

void checkFirstBusWeCanCatch(int *buses, int busCount, int departureTime) {
  int earliest = INT_MAX, busId, j, nextDeparture, earliestBusId;
  for (int b = 0; b < busCount; b++) {
    busId = buses[b];
    printf("Processing bus ID %d\n", busId);
    for (j = 1; (nextDeparture = j*busId) <= departureTime; j++) {
      ;
    }
    if (nextDeparture < earliest) {
      earliest = nextDeparture;
      earliestBusId = busId;
      printf("Found earlier bus %d, leaving at %d\n", earliestBusId, earliest);
    }
  }
  printf("The earliest bus we can catch is %d\n", earliestBusId);
  printf("and will leave %d minutes after our desired departure time\n", earliest - departureTime);
}

int compareInt(const void* a, const void* b) {
  // cast the void pointer as an int pointer and dereference to get value
  return *(int*)a - *(int*)b;
}
