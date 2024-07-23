

---

## Tetris Optimizer README

### Project Overview

The Tetris Optimizer is a Go-based project designed to take a set of tetrominoes and efficiently place them on the smallest possible square board. This project includes functionalities to validate tetromino shapes, process and trim unnecessary rows and columns, and assemble the tetrominoes into an optimized board.

### Project Structure

The project directory has the following structure:

```
tetris-optimizer/
├── tetromino/
     ├── array.go
│   ├── map.go
│   ├── print.go
│   ├── valid.go
├── go.mod
├── LICENSE
├── main.go
├── README.md
├── sample.txt
```

### Installation

1. **Clone the repository:**

```sh
git clone https://github.com/kevwasonga/Tetris-optimizer
cd tetris-optimizer
```

2. **Install dependencies:**

Ensure you have Go installed on your machine. If not, download and install it from [https://golang.org/dl/](https://golang.org/dl/).

```sh
go mod tidy
```

### Usage

#### Preparing Tetrominoes

1. **Create a tetromino file (e.g., `sample.txt`):**

```
....
.##.
.##.
....

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..
```

Each tetromino must be a 4x4 grid of characters with `#` representing the tetromino blocks and `.` representing empty spaces. Ensure each tetromino is separated by an empty line.

#### Running the Program

1. **Run the program without timing:**

```sh
go run main.go
```

2. **Run the program with a specific tetromino file:**

```sh
go run main.go sample.txt
```

3. **Run the program with timing:**

If you want to measure the total execution time, use:

```sh
time go run main.go sample.txt
```

### Code Explanation

#### main.go

The entry point of the program. It loads tetrominoes from the file, validates them, and attempts to assemble them onto the smallest possible square board.

#### tetromino/loadbanner.go

Loads and processes the tetrominoes from the file.

#### tetromino/print.go

Handles the assembly of tetrominoes onto the board.

### Error Handling

- If any error occurs during file reading or tetromino validation, the program prints `ERROR` and terminates immediately using `os.Exit(0)`.

### Future Improvements

- Enhance the tetromino validation to include more complex rules.
- Add a graphical interface to visualize the tetromino placement.
- Optimize the algorithm for better performance with larger sets of tetrominoes.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

### Contact

For any issues or questions, please open an issue in the repository or contact the maintainer at [kevinwasonga116@gmail.com].

---

