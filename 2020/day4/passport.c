#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../lib/input.h"

#define FIELD_COUNT 8

const char *FIELDS[FIELD_COUNT] = {
  "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"
};

int checkAndResetFields(int fields[FIELD_COUNT]);
void resetFields(int fields[]);
int validateFieldData(const char* fieldName, const char* input, int index, int* validFields);
int checkNumberInRange(const char *data, int min, int max);
int checkValidHeight(const char *data);
int checkValidHairColour(const char *data);
int checkValidEyeColour(const char *data);
int checkValidPassportId(const char *data);

int main() {
  char input[100000];
  getInput(input);

  int validNumFields = 0, validFieldData = 0;
  int fields[FIELD_COUNT] = {0, 0, 0, 0, 0, 0, 0, 0};
  int validFields[FIELD_COUNT] = {0, 0, 0, 0, 0, 0, 0, 0};

  char c;
  for (int i = 0; (c = input[i]) != '\0'; i++) {
    // got to the end of one entry - validate and reset fields
    if (c == '\n' && input[i+1] == '\n') {
      validNumFields += checkAndResetFields(fields);
      validFieldData += checkAndResetFields(validFields);

    // found a field - check its type type and validate the contents
    } else if (c == ':') {
      char *fieldName = &input[i-3];
      for (int f = 0; f < FIELD_COUNT; f++) {
        if (strncmp(fieldName, FIELDS[f], 3) == 0) {
          fields[f] = 1;
          validFields[f] = validateFieldData(FIELDS[f], input, i, validFields);
        }
      }
    }
  }
  // TODO nicer way to check final entry in loop?
  validNumFields += checkAndResetFields(fields);
  validFieldData += checkAndResetFields(validFields);

  printf("Total number of entries with correct fields = %d\n", validNumFields);
  printf("Total number of entries with correct and valid fields = %d\n", validFieldData);
}

/* Validate that data is correct, based on criteria for that field type */
int validateFieldData(const char* fieldName, const char* input, int index, int* validFields) {
  char data[100];
  int i;
  for (i = 1; input[index+i] != ' ' && input[index+i] != '\n'; i++) {
    data[i-1] = input[index+i];
  }
  data[i-1] = '\0';

  if (strcmp(fieldName, "byr") == 0) {
    return checkNumberInRange(data, 1920, 2002);
  }
  if (strcmp(fieldName, "iyr") == 0) {
    return checkNumberInRange(data, 2010, 2020);
  }
  if (strcmp(fieldName, "eyr") == 0) {
    return checkNumberInRange(data, 2020, 2030);
  }
  if (strcmp(fieldName, "hgt") == 0) {
    return checkValidHeight(data);
  }
  if (strcmp(fieldName, "hcl") == 0) {
    return checkValidHairColour(data);
  }
  if (strcmp(fieldName, "ecl") == 0) {
    return checkValidEyeColour(data);
  }
  if (strcmp(fieldName, "pid") == 0) {
    return checkValidPassportId(data);
  }
  if (strcmp(fieldName, "cid") == 0) {
    return 1;
  }
  return 0;
}

/* Check that data is a number within the provided interval */
int checkNumberInRange(const char *data, int min, int max) {
  char *next;
  int number = strtol(data, &next, 10);
  return min <= number && number <= max;
}

/* Valid height is an integer given in either centimetres or inches */
int checkValidHeight(const char *data) {
  char *next;
  int height = strtol(data, &next, 10);
  if (strcmp(next, "cm") == 0)
    return 150 <= height && height <= 193;
  if (strcmp(next, "in") == 0) {
    return 59 <= height && height <= 76;
  }
  return 0;
}

/* A valid hair colour ID is a hash followed by precisely 6 alphanumerical characters */
int checkValidHairColour(const char *data) {
  if (data[0] != '#') {
    return 0;
  }
  char c;
  int i;
  for (i = 1; (c = data[i]) != '\0'; i++) {
    if (i > 6 || c < '0' || c > 'z' || (c > '9' && c < 'a')) {
      return 0;
    }
  }
  return i == 7;
}

/* Eye colour must be in some pre-defined set */
int checkValidEyeColour(const char *data) {
  return (
    strcmp(data, "amb") == 0 ||
    strcmp(data, "blu") == 0 ||
    strcmp(data, "brn") == 0 ||
    strcmp(data, "gry") == 0 ||
    strcmp(data, "grn") == 0 ||
    strcmp(data, "hzl") == 0 ||
    strcmp(data, "oth") == 0
  );
}

/* A valid passport ID contains precisely 9 digits */
int checkValidPassportId(const char *data) {
  char c;
  int i;
  for (i = 0; (c = data[i]) != '\0'; i++) {
    if (c < '0' || c > '9') {
      return 0;
    }
  }
  return i == 9;
}

/* Check that all fields apart from Country ID are present */
int checkAndResetFields(int fields[]) {
  int result = 1;
  // final entry is country ID which can be ignored
  for (int i = 0; i < FIELD_COUNT-1; i++) {
    result *= fields[i];
  }
  resetFields(fields);
  return result;
}

/* Reset all fields to zero for next entry */
void resetFields(int fields[]) {
  for (int i = 0; i < FIELD_COUNT; i++) {
    fields[i] = 0;
  }
}
