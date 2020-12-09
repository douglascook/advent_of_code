#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#include "../lib/input.h"

int validateNumber(long long *preceding, long long number);

int main() {
  char input[1000000];
  getInput(input);

  long long n, target, numbers[1000];

  int i, j, lineCount = 0;
  char *next, *start = input;
  for (lineCount = 0; (n = strtoll(start, &next, 10)) != 0; lineCount++) {
    numbers[lineCount] = n;
    start = next;
  }

  // Find the first number that is not the sum of two of the preceding 25 numbers
  for (i = 25; i < lineCount; i++) {
    target = numbers[i];
    if (validateNumber(&numbers[i - 25], target) == 0) {
      printf("First invalid number is %lld\n", target);
      break;
    }
  }

  // Find a contiguous block of numbers summing to the number found above
  int startI, endI;
  for (i = 0; i < lineCount; i++) {
    n = numbers[i];
    for (j = i + 1; n < target; j++) {
      n += numbers[j];
      if (n == target) {
        printf("Found numbers summing to %lld in positions %d to %d\n", target, i, j);
        startI = i;
        endI = j;
      }
    }
  }

  // Find the min and max in the contigous block (numbers is not ordered)
  long long min = LLONG_MAX, max = 0;
  for (i = startI; i <= endI; i++) {
    n = numbers[i];
    if (n > max) {
      max = n;
    }
    if (n < min) {
      min = n;
    }
  }
  printf("Min and max numbers in the block are %lld and %lld\n", min, max);
  printf("Their sum is %lld\n", min + max);
}

int validateNumber(long long *preceding, long long number) {
  int i, j;
  for (i = 0; i < 25; i++) {
    for (j = i+1; j < 25; j++) {
      // if the number is the sum of two preceding numbers it is valid
      if (preceding[i] + preceding[j] == number) {
        return 1;
      }
    }
  }
  return 0;
}
