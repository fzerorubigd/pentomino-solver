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
  -color
    	Use color output (default true)
  -count int
    	The count of the solution to show before exit, -1 to show all (default -1)
  -height int
    	Width of the puzzle (default 6)
  -width int
    	Width of the puzzle (default 10)
```

**Example:**

To solve a 6x10 puzzle and see the first solution in color:

```bash
./bin/pentomino -width 10 -height 6 -count 1 -color
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
  -color
    	Use color output (default true)
  -count int
    	The count of the solution to show before exit, -1 to show all (default -1)
  -day int
    	The day of the month, 1 to 31 (default 1)
  -jalali
    	Use jalali calendar
  -month int
    	The month, 1 to 12 (default 1)
  -output-dir string
    	Output directory for SVG files
  -svg
    	Output SVG files (1.svg, 2.svg, ...)
  -today
    	Output today's calendar, ignore all other date related flags
  -weekday int
    	The weekday, 1 for the first day, 7 for 7th day. (Shanbe is first for Persian, Monday for Gregorian) (default 1)
  -year int
    	The year of the calendar, for Persian 1=1404 and for Gregorian 1=2025, max 10 (default 1)
```

**Example:**

To solve the puzzle for the first day of the first month of the first year (Gregorian) and see the first solution:

```bash
./bin/pcalendar -weekday 1 -day 1 -month 1 -year 1 -count 1
```

To solve for today's date using the Jalali calendar and output the first 5 solutions as SVGs to the `jalali` directory:

```bash
./bin/pcalendar -today -jalali -count 5 -svg -output-dir jalali
```

## Features

- **Text & Color Output**: Supports both plain text and colored ANSI output for terminal viewing.
- **SVG Export**: Can export solutions as SVG images (via `pcalendar -svg`).
- **Support for Calendars**: Supports both Gregorian and Jalali (Persian) calendars.
- **Daily Puzzle**: Use GitHub Actions to generate and send daily puzzles via Telegram.

## Daily Puzzle (GitHub Action)

This repository includes a GitHub Action (`.github/workflows/daily_puzzle.yml`) that runs daily at 6:00 AM. It:

1.  Generates 5 solutions for the current day's puzzle for both Gregorian and Jalali calendars.
2.  Converts the SVG output to PNG.
3.  Sends the images to a Telegram chat.

**Required Secrets:**

- `TELEGRAM_TO`: The Telegram chat ID to send messages to.
- `TELEGRAM_TOKEN`: The user bot token.

## Output Example

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
*(Colors are rendered in the terminal)*

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
*(Colors are rendered in the terminal)*
