#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "../lib/input.h"

enum direction{NORTH, EAST, SOUTH, WEST};

void directInstructions(char instructions[1000][10], int count);
void waypointInstructions(char instructions[1000][10], int count);
int modulo(int a, int b);

int main() {
  char instructions[1000][10];
  int count;
  for (count = 0; getLine(instructions[count]) > 0; count++) {
    ;
  }
  directInstructions(instructions, count);
  waypointInstructions(instructions, count);
}

/* Instructions are given as direct movements to make with the ship.
 *
 * Eg. N3 means move 3 north, R90 means turn 90 degrees right, F10 means move
 * 10 in current heading.
 */
void directInstructions(char instructions[1000][10], int count) {
  int moved[4] = {0};
  int angle = 90;
  enum direction heading = EAST;

  char *next, action;
  int value;
  for (int i = 0; i < count; i++) {
    action = instructions[i][0];
    value = strtol(&instructions[i][1], &next, 10);

    switch(action) {
      case 'N':
        moved[NORTH] += value;
        break;
      case 'S':
        moved[SOUTH] += value;
        break;
      case 'E':
        moved[EAST] += value;
        break;
      case 'W':
        moved[WEST] += value;
        break;
      case 'F':
        moved[heading] += value;
        break;
      case 'L':
        angle = modulo(angle - value, 360);
        heading = angle/90;
        break;
      case 'R':
        angle = modulo(angle + value, 360);
        heading = angle/90;
        break;
      default:
        break;
    }
  }
  int x = moved[NORTH] - moved[SOUTH];
  int y = moved[EAST] - moved[WEST];
  printf("Final coordinates = (%d, %d)\n", x, y);
  printf("Manhattan distance from origin is %d\n", abs(x) + abs(y));
}

/* Instructions are given as changes to a waypoint relative to the ship which
 * is then used to direct the movement of the ship.
 *
 * Eg. N3 means move the waypoint 3 north, R90 means rotate the waypoint 90
 * degrees right, F10 means move 10 times in vector defined by waypoint.
 */
void waypointInstructions(char instructions[1000][10], int count) {
  // starting position of waypoint relative to the ship
  int x = 0, y = 0, wx = 10, wy = 1, temp, i, j;

  char *next, action;
  int value;
  for (i = 0; i < count; i++) {
    action = instructions[i][0];
    value = strtol(&instructions[i][1], &next, 10);

    switch(action) {
      case 'N':
        wy += value;
        break;
      case 'S':
        wy -= value;
        break;
      case 'E':
        wx += value;
        break;
      case 'W':
        wx -= value;
        break;
      case 'F':
        // move value times in waypoint vector
        x += wx * value;
        y += wy * value;
        break;
      case 'L':
        // rotate waypoint around the ship
        for (j = 0; j < value/90; j++) {
          temp = wx;
          wx = -wy;
          wy = temp;
        }
        break;
      case 'R':
        // rotate waypoint around the ship
        for (j = 0; j < value/90; j++) {
          temp = wy;
          wy = -wx;
          wx = temp;
        }
        break;
      default:
        break;
    }
  }
  printf("Final coordinates = (%d, %d)\n", x, y);
  printf("Manhattan distance from origin is %d\n", abs(x) + abs(y));
}

/* C modulo does not handle negative numbers like mathematical modulo and will
 * return negative results if a is negative. This will always return a positive
 * result */
int modulo(int a, int b) {
  return ((a % b) + b) % b;
}
