#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../lib/input.h"

int compareInt(const void* a, const void* b);
long long countRoutesToZero(int *adapters, int index, long long *routes);

int main() {
  char input[10000];
  const int count = getInput(input);

  int i, n, adapters[count];
  char *start = input, *next;
  for (i = 0; (n = strtol(start, &next, 10)) != 0; i++) {
    adapters[i] = n;
    start = next;
  }

  // Adapters can only take inputs between 1 and 3 jolts lower than their own
  // joltage, so we must go through them in ascending order
  qsort(adapters, count, sizeof(int), compareInt);

  int differences[3] = {0, 0, 0};
  int previous = 0;
  for (i = 0; i < count; i++) {
    differences[adapters[i] - previous - 1] += 1;
    previous = adapters[i];
  }
  // Built in adapter always has joltage 3 higher than max, so add one to 3 count
  differences[2] += 1;

  for (i = 0; i < 3; i++) {
    printf("Adapters with jolt difference of %d = %d\n", i+1, differences[i]);
  }
  printf("Product of 1 and 3 jolt difference counts = %d\n", differences[0] * differences[2]);

  long long routes[1000] = {0};
  printf(
    "Total number of routes through all adapters = %lld\n",
    countRoutesToZero(adapters, count - 1, routes)
  );
}

/* Return the number of valid routes to zero from the adapter at given index.
 *
 * The total for an adapter is the sum of the totals of the adapters with joltage
 * at most 3 jolts below its own joltage, take a dynamic programming approach to
 * avoid repeated computation.
 * */
long long countRoutesToZero(int *adapters, int index, long long *routes) {
  // if we have a cached value use that
  if (routes[index] != 0) {
    return routes[index];
  }
  // if the current adapter is directly within range of zero, count that route
  long long total = adapters[index] <= 3 ? 1 : 0;

  // recurse through all routes below and add to total
  for (int i = 1; i <= 3 && (index - i) >= 0; i++) {
    if (adapters[index] - adapters[index - i] <= 3) {
      total += countRoutesToZero(adapters, index - i, routes);
    }
  }
  // cache the result
  routes[index] = total;
  printf("Adapter %d has %lld routes\n", adapters[index], total);

  return total;
}

int compareInt(const void* a, const void* b) {
  // cast the void pointer as an int pointer and dereference to get value
  return *(int*)a - *(int*)b;
}
