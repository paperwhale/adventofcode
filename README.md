# Advent of Code 🎄

Getting into the holiday spirit with [Advent of Code](https://adventofcode.com). 
As Linus would say,
> "That's what Christmas is all about, Charlie Brown."

## Dependencies

To run these solutions, you'll need:

*   **Go**: The programming language used for the solutions.
*   **curl**: Used by the `install.sh` script to download puzzle inputs.

## Getting Started

The `install.sh` script automates setting up the directory structure and downloading the puzzle input for a given day.

To run the script for a specific day and year, use the following command:

```bash
./install.sh -d <day> -y <year>
```

The script requires a session cookie from the Advent of Code website to download your personalized puzzle input. You must set the `AOC_SESSION` environment variable with your cookie value.

```bash
export AOC_SESSION="your_session_cookie_here"
```

If you don't specify a day or year, the script will default to the current day and year.
