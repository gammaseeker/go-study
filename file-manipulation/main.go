package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func readFile(fileName string) (*os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening %s: %s", fileName, err)
		return nil
	}
	return file
}

func writeFile(fileName string) (*os.File) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("Error opening output file: %s", err)
		return nil
	}
	return file
}

func getNumRows(fileName string) (int32) {
	// Gets line count of output file
	cmd, err := exec.Command("wc", "-l", fileName).Output()
	if err != nil {
		fmt.Println("%s", err)
	}
	cmdOutput := string(cmd[:])
	lineCount, err := strconv.Atoi(strings.Fields(cmdOutput)[0]) // can Atoi return int32? what if cmdOutput is more than an int?
	if err != nil {
		fmt.Println("%s", err)
	}
	fmt.Println(lineCount)
	return int32(lineCount)
}

func main() {
	/*
		1. Open Output file
		2. Open source file to copy
		3. Get line count of output and take difference of 25M - line count
		4. See line count of source file, see if line count fits in difference
		5. If so, copy rows then close source file
		6. Else, copy what you can and remember location in source file for next output file then close output file
	*/

	output := writeFile("./output.csv")
	defer output.Close()

	source := readFile("./test1.csv")
	defer source.Close()

	numRowsOutput := getNumRows("./output.csv")
	const MAX_ROWS int32 = 25000000
	var linesRemaining int32 = MAX_ROWS - numRowsOutput
	fmt.Println(linesRemaining)

	numRowsSource := getNumRows("./test1.csv")
	if numRowsSource < linesRemaining {
		n, err := io.Copy(output, source)
		if err != nil {
			fmt.Println("Failed to append test.csv to output: %s", err)
		}
		fmt.Println("wrote %d bytes", n)
	}

	source = readFile("./test2.csv")

	numRowsOutput = getNumRows("./output.csv")
	linesRemaining = MAX_ROWS - numRowsOutput
	fmt.Println(linesRemaining)

	numRowsSource = getNumRows("./test2.csv")
	if numRowsSource < linesRemaining {
		n, err := io.Copy(output, source)
		if err != nil {
			fmt.Println("Failed to append test.csv to output: %s", err)
		}
		fmt.Println("wrote %d bytes", n)
	}


	/*


	n, err := io.Copy(out, file1)
	if err != nil {
		fmt.Println("Failed to append test.csv to output: %s", err)
	}
	fmt.Println("wrote %d bytes", n)

	_, err = out.WriteString("\n")
	if err != nil {
		fmt.Println("Failed to write new line: %s", err)
	}
	n, err = io.Copy(out, file2)
	if err != nil {
		fmt.Println("Failed to append test2.csv to output: %s", err)
	}
	fmt.Println("wrote %d bytes", n)

	file1.Close()
	file2.Close()
	out.Close()
	*/

}

/*
func readFile() {
	filePath := "./test.csv"
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error with opening file: %s", err)
	}

	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("EOF reached. Stopping")
			break
		}
		if err != nil {
			fmt.Println("Error reading: %s", err)
		}
		fmt.Println("%s", record)
	}
}
*/