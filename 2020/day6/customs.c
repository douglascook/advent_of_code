#include <stdio.h>
#include "../lib/input.h"

#define ALPHABET_SIZE 26

int countAnyAnswered(int questions[ALPHABET_SIZE]);
int countAllAnswered(int questions[ALPHABET_SIZE], int groupSize);
void resetQuestions(int questions[ALPHABET_SIZE]);

int main() {
  int questions[ALPHABET_SIZE];
  resetQuestions(questions);

  char c, line[100];
  int length, anyAnswered = 0, allAnswered = 0, groupSize = 0;
  while ((length = getLine(line)) != 0) {
    if (length == 1) {
      // end of one group of passengers, count total questions answered by group
      anyAnswered += countAnyAnswered(questions);
      allAnswered += countAllAnswered(questions, groupSize);
      resetQuestions(questions);
      groupSize = 0;
    } else {
      // record which questions were answered by this passenger
      for (int i = 0; (c = line[i]) != '\n'; i++) {
        questions[c - 'a'] += 1;
      }
      groupSize++;
    }
  }
  anyAnswered += countAnyAnswered(questions);
  allAnswered += countAllAnswered(questions, groupSize);

  printf("Total any answered questions = %d\n", anyAnswered);
  printf("Total all answered questions = %d\n", allAnswered);
}

/* Return number of questions with a count of at least 1 */
int countAnyAnswered(int questions[ALPHABET_SIZE]) {
  int count = 0;
  for (int i = 0; i < ALPHABET_SIZE; i++) {
    count += questions[i] > 0;
  }
  return count;
}

/* Return number of questions with count matching size of the group */
int countAllAnswered(int questions[ALPHABET_SIZE], int groupSize) {
  int count = 0;
  for (int i = 0; i < ALPHABET_SIZE; i++) {
    count += questions[i] == groupSize;
  }
  return count;
}

void resetQuestions(int questions[ALPHABET_SIZE]) {
  for (int i = 0; i < ALPHABET_SIZE; i++) {
    questions[i] = 0;
  }
}
