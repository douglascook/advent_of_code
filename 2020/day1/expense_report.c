#include <stdio.h>
#include <stdlib.h>

#define TARGET 2020

int readInput(char input[]);
void parseInput(char input[], int parsed[]);
int solveTwoEntries(const int parsed[], int numberEntries);
int solveThreeEntries(const int parsed[], int numberEntries);

int main() {
  char input[10000];
  int parsed[1000];

  int lineCount = readInput(input);
  parseInput(input, parsed);

  printf("SOLVING FOR TWO ENTRIES\n");
  printf("Product is %d\n", solveTwoEntries(parsed, lineCount));
  printf("SOLVING FOR THREE ENTRIES\n");
  printf("Product is %d\n", solveThreeEntries(parsed, lineCount));
}

/* Read the input into a character array and return number of lines */
int readInput(char input[]) {
  char c;
  int i;
  int lineCount = 0;
  for (i = 0; (c = getchar()) != EOF; i++) {
    input[i] = c;
    if (c == '\n') {
      lineCount++;
    }
  }
  input[i] = '\0';
  return lineCount;
}

/* Parse the character array into an array of integers */
void parseInput(char input[], int parsed[]) {
  char *start = &input[0];
  char *next;
  int value;
  int index = 0;
  // strtol parses the first integer and sets pointer to the next character, so
  // no need to handle newlines ourselves
  while ((value = strtol(start, &next, 10)) != '\0') {
    parsed[index++] = value;
    start = next;
  }
}

/* Brute force solution to find the two entries that sum to 2000 and return their product */
int solveTwoEntries(const int values[], int length) {
  int x, y;
  for (int i = 0; i < length; i++) {
    x = values[i];

    for (int j = i+1; j < length; j++) {
      y = values[j];

      if (x + y == TARGET) {
        printf("Values are %d and %d\n", x, y);
        return x * y;
      }
    }
  }
  return 0;
};

/* Brute force solution to find the three entries that sum to 2000 and return their product */
int solveThreeEntries(const int values[], int length) {
  int x, y, z;
  for (int i = 0; i < length; i++) {
    x = values[i];

    for (int j = i+1; j < length; j++) {
      y = values[j];

      for (int k = j+1; k < length; k++) {
        z = values[k];

        if (x + y + z == TARGET) {
        printf("Values are %d, %d and %d\n", x, y, z);
          return x * y * z;
        }
      }
    }
  }
  return 0;
};
