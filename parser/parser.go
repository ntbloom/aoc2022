package parser

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ntbloom/aoc2022/errors"
)

const InputsDirectory = "inputs"
const TestInputsDirectory = "test_inputs"

// ParseFlags gets the command-line flags for the day of the week and the filename
func ParseFlags() (int, int, *os.File) {

	day := flag.Int("day", -1, "Which day of the month! (1-24)")
	puzzle := flag.Int("puzzle", 1, "which puzzle to use")
	flag.Parse()

	if *day == -1 {
		fmt.Println("Must provide day!")
		os.Exit(errors.DayNotFound)
	}

	filename, err := GetFileName(*day, *puzzle, InputsDirectory)
	if err != nil {
		panic(err)
	}
	fd := GetFileDescriptor(filename)

	return *day, *puzzle, fd

}

// GetFileDescriptor checks for the file and opens it
func GetFileDescriptor(filename string) *os.File {
	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("File %s does not exist\n", filename)
		os.Exit(errors.FileNotExists)
	}

	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", filename, err)
		os.Exit(errors.FileOpeningError)
	}
	return fd
}

// GetFileName gets a filename for a puzzle and string
func GetFileName(day, puzzle int, directory string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	name := fmt.Sprintf("%d-%d", day, puzzle)
	return filepath.Join(path, directory, name), nil
}
