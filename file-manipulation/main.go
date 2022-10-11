package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	/*
		1. Open Output file
		2. Open source file to copy
		3. Get line count of output and take difference of 25M - line count
		4. See line count of source file, see if line count fits in difference
		5. If so, copy rows then close source file
		6. Else, copy what you can and remember location in source file for next output file then close output file
	*/

	//const MAX_ROWS int32 = 250000000 // The real max rows value
	const MAX_ROWS int32 = 5 // For testing purposes

	//inputList := [4]string {"test1.csv", "test2.csv", "test3.csv", "test4.csv"}
	inputList := [2]string {"test1.csv", "test2.csv"}

	var source *os.File
	outputCtr := 1
	outputFilename := "./output" + strconv.Itoa(outputCtr) + ".csv"
	output := writeFile(outputFilename)

	var filePosition int64 = 0
	var middleOfFile bool = false
	i := 0
	for i < len(inputList) {
		fmt.Printf("input: %s\n", inputList[i])
		source = readFile(inputList[i])

		numRowsOutput := getNumRows(outputFilename)
		fmt.Printf("numRowsOutput: %d\n", numRowsOutput)

		var linesRemaining int32 = MAX_ROWS - numRowsOutput
		fmt.Printf("linesRemaining: %d\n", linesRemaining)

		writeableLines := linesRemaining - 1 // Reserve 1 line for ending new line
		fmt.Printf("writeableLines: %d\n", writeableLines)

		numRowsSource := getNumRows(inputList[i])
		fmt.Printf("numRowsSource: %d\n", numRowsSource)

		if numRowsSource <= writeableLines && !middleOfFile{
			n, err := io.Copy(output, source)
			if err != nil {
				fmt.Printf("Failed to append %s to %s: %s", inputList[i], outputFilename, err)
			}
			fmt.Printf("wrote %d bytes\n", n)
			source.Close()
			i++
		} else {
			// WIP This section of code is buggy
			// copy a certain amount
			offset := copyNRows(source, output, writeableLines, filePosition)
			var err error
			filePosition, err = source.Seek(offset, 1)
			if err != nil {
				fmt.Printf("Error seeking: %s", err)
			}
			middleOfFile = true
			numRowsOutput = getNumRows(outputFilename)
			linesRemaining = MAX_ROWS - numRowsOutput
			writeableLines = linesRemaining - 1
			if writeableLines <= 0 {
				output.Close()
			}

			// create new outputfile, then retry to copy source file
			outputCtr++
			outputFilename = "./output" + strconv.Itoa(outputCtr) + ".csv"
			output = writeFile(outputFilename)
		}
	}
}

func readFile(fileName string) (*os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening %s: %s", fileName, err)
		return nil
	}
	return file
}

func writeFile(fileName string) (*os.File) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("Error opening output file: %s", err)
		return nil
	}
	return file
}

func getNumRows(fileName string) (int32) {
	cmd, err := exec.Command("wc", "-l", fileName).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	cmdOutput := string(cmd[:])
	lineCount, err := strconv.Atoi(strings.Fields(cmdOutput)[0]) // TODO can Atoi return int32? what if cmdOutput is more than an int?
	if err != nil {
		fmt.Printf("%s", err)
	}

	return int32(lineCount)
}

func copyNRows(source io.Reader, destination io.Writer, n int32, currentPosition int64) (int64) {
	csvReader := csv.NewReader(source)
	for i := int32(0); i < n; i++ { 
		row, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("EOF reached. Stopping")
			break
		}
		if err != nil {
			fmt.Printf("Error reading: %s", err)
		}

		csv := strings.Join(row[:], ",") + "\n"
		reader := strings.NewReader(csv)
		n, err := io.Copy(destination, reader)
		if err != nil {
			fmt.Printf("Failed to append: %s", err)
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	offset := csvReader.InputOffset() // Gives us beginning of next row
	return offset
}
