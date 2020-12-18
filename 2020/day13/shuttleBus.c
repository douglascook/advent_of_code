#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include "../lib/input.h"

void getBuses(char *start, int *buses, int *busCount);
void getBusesAndIndexes(char *start, int *buses, int *indexes);
int compareInt(const void* a, const void* b);
void checkFirstBusToCatch(int *buses, int busCount, int departureTime);
void checkAllSubsequentDepartures(int *buses, int *indexes, int busCount);

int main() {
  char input[1000], *start;

  getInput(input);
  int departureTime = strtol(input, &start, 10);
  printf("Departure time is %d\n", departureTime);

  int buses[200], indexes[200], busCount = 0;
  getBuses(start, buses, &busCount);
  checkFirstBusToCatch(buses, busCount, departureTime);

  getBusesAndIndexes(start, buses, indexes);
  int i, j;
  for (i = 0; i < busCount; i++) {
    printf("Bus %d is at index %d\n", buses[i], indexes[i]);
  }
  checkAllSubsequentDepartures(buses, indexes, busCount);
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

void getBusesAndIndexes(char *start, int *buses, int *indexes) {
  char *next;
  int busId, index = 0, busCount = 0;
  while (*start != '\0') {
    busId = strtol(start, &next, 10);
    if (*start == ',') {
      index++;
    }
    // if we're at an integer then add it to the array of bus IDs
    if (busId != 0) {
      buses[busCount] = busId;
      indexes[busCount] = index;
      busCount++;
      start = next;
    // otherwise increment pointer so we continue at next char
    } else {
      start = start + 1;
    }
  }

}

void checkFirstBusToCatch(int *buses, int busCount, int departureTime) {
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

void checkAllSubsequentDepartures(int *buses, int *indexes, int busCount) {
  // for each multiple of the bus with highest ID
  //    check that multiple % earlier bus = highest bus index - bus being checked index
  //    check that multiple % later bus = later bus - (later bus - bus being checked)
  //
  long long time, i;
  int j;
  int good = 0;
  for (i = 1; !good; i++) {
    time = buses[0] * i;
    good = 1;

    if (i % 100000 == 0)
      printf("Checking time %lld\n", time);

    for (j = 1; j < busCount; j++) {
      if ((time + indexes[j]) % buses[j] != 0) {
        good = 0;
        break;
      }
    }
  }
  printf("Done! Timestamp is %lld\n", time);
}

int compareInt(const void* a, const void* b) {
  // cast the void pointer as an int pointer and dereference to get value
  return *(int*)a - *(int*)b;
}
