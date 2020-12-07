#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../lib/input.h"

#define MAX_LINES 1000
#define MAX_LINE_LENGTH 200
#define MAX_COLOUR 100

int getBagId(const char bags[MAX_LINES][MAX_COLOUR], const char* bagColour, int bagCount);
void getWordsInRange(char* line, int start, int end, char* output);
void setToZero(int* array, int length);
void findBag(
    const char bagNames[MAX_LINES][MAX_COLOUR],
    const char bagContents[MAX_LINES][MAX_LINE_LENGTH],
    int bagCount,
    int bagToFind,
    int* foundIn
);
int countNestedBags(
    char bagNames[MAX_LINES][MAX_COLOUR],
    char bagContents[MAX_LINES][MAX_LINE_LENGTH],
    int bagCount,
    int bagId
);

int main() {
  char bagContents[MAX_LINES][MAX_LINE_LENGTH];
  char bagNames[MAX_LINES][MAX_COLOUR];

  int i, bagCount = 0;
  // create an array containing all bag colours
  for (i = 0; (getLine(bagContents[i])) != 0; i++) {
    getWordsInRange(bagContents[i], 0, 2, bagNames[i]);
  }
  bagCount = i;
  printf("Total number of bags described is %d\n", bagCount);

  // First we need to find all bags that contain a shiny gold bag, directly or indirectly
  int shinyGold = getBagId(bagNames, "shiny gold", bagCount);
  int foundIn[bagCount];
  // haven't found it in any bags yet
  setToZero(foundIn, bagCount);

  findBag(bagNames, bagContents, bagCount, shinyGold, foundIn);
  int total = 0;
  for (i = 0; i < bagCount; i++) {
    total += foundIn[i];
  }
  printf("Total number of bags containing a shiny gold bag is %d\n", total);

  // Now we're interested in how many bags are contained within the shiny gold bag
  total = countNestedBags(bagNames, bagContents, bagCount, shinyGold);
  printf("Total number of bags a shiny gold bag contains is %d\n", total - 1);
}

/* Recursively search for all bags containing the bag with ID bagToFind */
void findBag(
  const char bagNames[MAX_LINES][MAX_COLOUR],
  const char bagContents[MAX_LINES][MAX_LINE_LENGTH],
  int bagCount,
  int bagToFind,
  int* foundIn
) {
  char* match;
  for (int bagId = 0; bagId < bagCount; bagId++) {
    match = strstr(bagContents[bagId], bagNames[bagToFind]);
    // if match is the start of the line then it is the description of the bag itself
    if (match != NULL && match != bagContents[bagId]) {
      foundIn[bagId] = 1;
      // recurse to find all bags containing the bag we just found
      findBag(bagNames, bagContents, bagCount, bagId, foundIn);
    }
  }
}

/* Recursively count the total number of bags a single bag with given bagId must contain */
int countNestedBags(
  char bagNames[MAX_LINES][MAX_COLOUR],
  char bagContents[MAX_LINES][MAX_LINE_LENGTH],
  int totalBags,
  int bagId
) {
  // count the current bag first
  int total = 1;

  char* contents = bagContents[bagId];
  if (strstr(contents, "no other") != NULL) {
    // bag has no children
    return total;
  }

  int nestedBagCount;
  char *next, nestedBagName[MAX_COLOUR];
  char *start = strpbrk(contents, "0123456789");
  // parse out the count of the next nested bag
  while (start != NULL && ((nestedBagCount = strtol(start, &next, 10)) != 0)) {
    // parse out the name of this nested bag
    // FIXME need the +1 because the function doesn't handle whitespace well...
    getWordsInRange(next + 1, 0, 2, nestedBagName);

    total += nestedBagCount * countNestedBags(
      bagNames, bagContents, totalBags, getBagId(bagNames, nestedBagName, totalBags)
    );
    start = strpbrk(next, "0123456789");
  }
  return total;
}

/* Return the ID of the bag with the given colour */
int getBagId(const char bags[MAX_LINES][MAX_COLOUR], const char* bagColour, int bagCount) {
  for (int i = 0; i < bagCount; i++) {
    if (strcmp(bagColour, bags[i]) == 0) {
      return i;
    }
  }
  return -1;
}

// TODO use strtok for this
/* Return a string containing words from <line> in the token range defined by
 * <start> and <end>.
 *
 * Tokens are assumed to be separated by single space characters.
 */
void getWordsInRange(char* line, int start, int end, char* output) {
  int spaceCount = 0, startChar = -1;
  char c;
  for (int i = 0; (c = line[i]) != '\0'; i++) {
    if (spaceCount >= start) {
      if (startChar == -1) {
        // first character in range so set character offset
        startChar = i;
      }
      output[i - startChar] = c;
    }
    if (c == ' ') {
      if (++spaceCount == end) {
        // reached the end of the desired range so remove trailing space and exit
        output[i - startChar] = '\0';
        break;
      }
    }
  }
  // TODO handle case where spaceCount < start ie no words found
}

// TODO move to helpers
void setToZero(int* array, int length) {
  for (int i = 0; i < length; i++) {
    array[i] = 0;
  }
}
