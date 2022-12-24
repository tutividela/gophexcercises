package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//export OUTPUT_PATH="words.txt"

func camelcase(s string) int {
    // Write your code here
	if len(s) == 0 {
		return 0
	}
	count := 1
	for _,c := range s {
		if c >= 65 && c <= 90 {
			count++
		}
	}
	return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := camelcase(s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}