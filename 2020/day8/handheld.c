#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../lib/input.h"

enum operation{ACC, JMP, NOP};

int runInstructions(char instructions[1000][10], int instructionCount, int *accumulator);

int main() {
  char instructions[1000][10];

  int lineCount = 0;
  while ((getLine(instructions[lineCount++]) != 0)) {
    ;
  }

  int accumulator = 0;
  runInstructions(instructions, lineCount, &accumulator);
  printf("Value of accumulator at loop = %d\n", accumulator);

  // Work through the instructions backwards, updating NOP or JMP and checking
  // if we terminate
  int i;
  for (i = lineCount - 1; i > 0; i--) {
    accumulator = 0;

    // change NOP to JMP
    if (instructions[i][0] == 'n') {
      instructions[i][0] = 'j';
      instructions[i][1] = 'm';
      if (runInstructions(instructions, lineCount, &accumulator) == 0) {
        break;
      }
      instructions[i][0] = 'n';
      instructions[i][1] = 'o';
    // change JMP to NOP
    } else if (instructions[i][0] == 'j') {
      instructions[i][0] = 'n';
      instructions[i][1] = 'o';
      if (runInstructions(instructions, lineCount, &accumulator) == 0) {
        break;
      }
      instructions[i][0] = 'j';
      instructions[i][1] = 'm';
    }
  }
  printf("Changed instruction number %d - %s", i, instructions[i]);
  printf("DONE! Value of accumulator at end = %d\n", accumulator);
}

/* Run the instructions and update accumulator, returning 0 if we terminate at
 * the instruction after the last otherwise 1 if we hit a loop */
int runInstructions(char instructions[1000][10], int instructionCount, int *accumulator) {
  // partially initialised array -> remainder filled with type-appropriate zero
  int visited[1000] = {0};
  int i = 0;
  char *instruction, *next;
  // Run until we hit a loop...
  while (visited[i] == 0) {
    // ...or reach the instruction after the last
    if (i == instructionCount) {
      return 0;
    }
    instruction = instructions[i];

    if (instruction[0] == 'a') {
      visited[i] = ACC;
      *accumulator += strtol(&instruction[4], &next, 10);
      i++;
    } else if (instruction[0] == 'j') {
      visited[i] = JMP;
      i += strtol(&instruction[4], &next, 10);
    } else {
      visited[i] = NOP;
      i++;
    }
  }
  return 1;
}
