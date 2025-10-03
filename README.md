# Pentomino Solver

This project is a solver for the classic [Pentomino](https://en.wikipedia.org/wiki/Pentomino) puzzle, written in Go. It also includes a special solver for a calendar-based pentomino puzzle.

## What is a Pentomino?

A pentomino is a polyomino of order 5, meaning it is a geometric shape composed of five congruent squares, connected edge-to-edge. There are 12 free pentominoes, and they are used in various tiling puzzles. The total area of the 12 pentominoes is 60 squares.

## Solvers

This project provides two command-line solvers:

### 1. `pentomino`

This is a classic pentomino solver that attempts to tile a rectangular board with the 12 pentomino pieces. The area of the board must be 60 squares.

#### Usage

```bash
./bin/pentomino [flags]
```

To get the help message:

```bash
./bin/pentomino -h
```

```
Usage of ./bin/pentomino:
  -count int
    	The count of the solution to show before exit, -1 to show all (default -1)
  -height int
    	Width of the puzzle (default 6)
  -width int
    	Width of the puzzle (default 10)
```

**Example:**

To solve a 6x10 puzzle and see the first solution:

```bash
./bin/pentomino -width 10 -height 6 -count 1
```

### 2. `pcalendar`

This solver works on a special 10x7 board that represents a calendar. You can specify a date, and the solver will block out the corresponding cells for the weekday, day, month, and year, and then solve the puzzle with the remaining pentomino pieces. This solver is based on the [pentomino-calendar](https://github.com/fzerorubigd/pentomino-calendar) project.

#### Usage

```bash
./bin/pcalendar [flags]
```

To get the help message:

```bash
./bin/pcalendar -h
```

```
Usage of ./bin/pcalendar:
  -count int
    	The count of the solution to show before exit, -1 to show all (default -1)
  -day int
    	The day of the month, 1 to 31 (default 1)
  -month int
    	The month, 1 to 12 (default 1)
  -weekday int
    	The weekday, 1 for the first day, 7 for 7th day. (Shanbe is first for Persian, Monday for Gregorian) (default 1)
  -year int
    	The year of the calendar, for Persian 1=1404 and for Gregorian 1=2025, max 10 (default 1)
```

**Example:**

To solve the puzzle for the first day of the first month of the first year and see the first solution:

```bash
./bin/pcalendar -weekday 1 -day 1 -month 1 -year 1 -count 1
```

## Output Example

The solver will print the solutions to the console.

### `pentomino` example

Here is an example of a solved 6x10 puzzle:

```
1 ===>
VIIIIILLLL
VYYYYWZZXL
VVVYWWZXXX
UUFWWZZTXP
UFFNNTTTPP
UUFFNNNTPP
```

### `pcalendar` example

Here is an example of a solved calendar puzzle. The `O` characters represent the blocked-out date pieces (weekday, day, month, year) and are not part of the pentomino solution.

```
1 ===>
OWOIIIIIOT
LWWYYYYTTT
LFWWYNNNXT
LFFFNNVXXX
LLFZZOVUXU
PPPZVVVUUU
PPZZOOOOOO
```
