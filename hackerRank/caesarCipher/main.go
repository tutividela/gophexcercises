package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */
var (
	A_byteCode rune = 65
	Z_byteCode rune = 90
	a_byteCode rune = 97
	z_byteCode rune = 122
)

func caesarCipher(s string, k int32) string {
    // Write your code here
	var cipher string
	for _,p := range s {
		if isSpecialCharacter(p) {
			cipher = strings.Join([]string{cipher,string(p)},"")
		}else {
			if p >= a_byteCode && p <= z_byteCode {
				entero := k / (z_byteCode-a_byteCode)
				resto := k % (z_byteCode-a_byteCode)
				c := p - entero + resto
				if c > z_byteCode {
					dif := c - z_byteCode
					c = a_byteCode  + dif -1
				}
				cipher =strings.Join([]string{cipher,string(c)},"")
			}
			if p >= A_byteCode && p <= Z_byteCode{
				entero := k / (Z_byteCode-A_byteCode)
				resto := k % (Z_byteCode-A_byteCode)
				c := p - entero + resto
				if c > Z_byteCode {
					dif := c -Z_byteCode
					c = A_byteCode  + dif -1
				}
				cipher =strings.Join([]string{cipher,string(c)},"")
			}
		}
	}
	return cipher
}

func isSpecialCharacter(p rune) bool {
	return !((p >= a_byteCode && p <= z_byteCode) || (p >= A_byteCode && p <= Z_byteCode))
}

//export OUTPUT_PATH="cipher.txt"

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
	fmt.Println(n)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)
    fmt.Fprintf(writer, "%s\n", result)

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
