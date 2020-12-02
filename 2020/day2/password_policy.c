#include <stdio.h>
#include <stdlib.h>
#include "../lib/input.h"

int validateMinMax(const char line[], int length, int min, int max);
int validatePositional(const char line[], int length, int min, int max);

int main() {
  char line[1000];
  int minMaxValidCount = 0;
  int positionalValidCount = 0;

  int length, min, max;
  char *next;

  while ((length = getLine(line)) > 0) {
    printf("Line = %s\n", line);

    min = strtol(&line[0], &next, 10);
    // strtol parses the hyphen as a minus, so we need to negate it
    max = -strtol(next, &next, 10);
    printf("Min = %d, Max = %d\n", min, max);

    minMaxValidCount += validateMinMax(line, length, min, max);
    positionalValidCount += validatePositional(line, length, min, max);
  }
  printf("Total number of min-max valid passwords = %d\n", minMaxValidCount);
  printf("Total number of positionally valid passwords = %d\n", positionalValidCount);
}

/* Validate password given min and max criteria */
int validateMinMax(const char line[], int length, int min, int max) {
  char letter = '\0';
  int letterCount = 0;

  // Now find the letter and validate the password
  for (int i = 0; i < length; i++) {
    if (line[i] == ':') {
      letter = line[i-1];
      printf("Letter is %c\n\n", letter);
    }
    // check letter is set first so we know we are in the password section
    if (letter != '\0' && line[i] == letter) {
      letterCount++;
    }
  }
  return letterCount >= min && letterCount <= max;
}

/* Validate password given positional criteria */
int validatePositional(const char line[], int length, int pos1, int pos2) {
  char letter = '\0';
  int matchCount = 0;
  int passwordStart = 0;

  // Now find the letter and validate the password
  for (int i = 0; i < length; i++) {
    if (line[i] == ':') {
      letter = line[i-1];
      passwordStart = i+1;
    }
    // if we're in the password and on expected letter
    if (passwordStart && line[i] == letter) {
      if (i - passwordStart == pos1 || i - passwordStart == pos2) {
        matchCount++;
      }
    }
  }
  return matchCount == 1;
}
