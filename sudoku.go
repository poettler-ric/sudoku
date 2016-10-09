// solves sudoku puzzles

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "unicode"
)

// ReadSudoku reads a sudoku file and returns the parsed matrix
func ReadSudoku(f *os.File) *[9][9]int {
    scanner := bufio.NewScanner(f)
    s := [9][9]int{}
    line, column := 0, 0

    for scanner.Scan() {
        text := scanner.Text()
        if len(strings.TrimSpace(text)) == 0 {
            // ignore empty lines
            continue
        }

        column = 0
        for _, c := range (text) {
            if unicode.IsSpace(c) {
                // ignore spaces
                continue
            }
            i, err := strconv.Atoi(string(c))
            if err == nil {
                s[line][column] = i
            }
            column++
        }
        line++
    }
    return &s
}

// PrintSudoku prints the sudoku matrix to stdout
func PrintSudoku(s *[9][9]int) {
    line, column := 1, 1
    for _, i := range s {
        column = 1
        for _, j := range i {
            fmt.Print(j)
            if column % 3 == 0 {
                fmt.Print(" ")
            }
            column++
        }
        fmt.Println()
        if line % 3 == 0 {
            fmt.Println()
        }
        line++
    }
}

// Candidates calculates possibilities for one specific field
// TODO: return a pointer?
func Candidates(s *[9][9]int, row, column int) []int {
    seen := [10]bool{false}
    // iterate over the row
    for i := 0; i < 9; i++ {
        seen[s[row][i]] = true
    }
    // iterate over the column
    for i := 0; i < 9; i++ {
        seen[s[i][column]] = true
    }
    // iterate over the square
    squareRow := row / 3
    squareColumn := column / 3
    for i := squareRow * 3; i < (squareRow + 1) * 3; i++ {
        for j := squareColumn * 3; j < (squareColumn + 1) * 3; j++ {
            seen[s[i][j]] = true
        }
    }

    out := make([]int, 0, 9)
    for i, present := range seen {
        if !present {
            out = append(out, i)
        }
    }
    return out
}

// SolveSudoku iterates over and updates the sudoku matrix
func SolveSudoku(s *[9][9]int) {
    for changed := true; changed; {
        changed = false
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                if s[i][j] != 0 {
                    // ignore already solved numbers
                    continue
                }

                candidates := Candidates(s, i, j)
                if len(candidates) == 1 {
                    s[i][j] = candidates[0]
                    changed = true
                }
            }
        }
    }
}

func main()  {
    if len(os.Args) < 2 {
        log.Fatalf("usage: %s <file>", os.Args[0])
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    s := ReadSudoku(f)
    PrintSudoku(s)
    SolveSudoku(s)
    fmt.Println("===")
    PrintSudoku(s)
}
