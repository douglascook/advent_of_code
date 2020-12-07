#include <stdio.h>

/* Read the next line from input, store contents and return its length */
int getLine(char* line) {
  char c;
  int length = 0;
  while ((c = getchar()) != EOF) {
    line[length++] = c;
    if (c == '\n') {
      break;
    }
  }
  line[length] = '\0';
  return length;
}

/* Read the entire input into a character array and return number of lines */
int getInput(char* input) {
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

