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
	upperBody rune = (Z_byteCode-A_byteCode)
	lowerBody rune = (z_byteCode-a_byteCode)
)

func caesarCipher(s string, k int32) string {
    // Write your code here
	var (
		cipher string
		c rune
	)
	for _,p := range s {
		if isSpecialCharacter(p) {
			cipher = strings.Join([]string{cipher,string(p)},"")
		}else {
			if p >= a_byteCode && p <= z_byteCode {
				c =cipherCharacter(a_byteCode,z_byteCode,p,k)
				cipher =strings.Join([]string{cipher,string(c)},"")
			}
			if p >= A_byteCode && p <= Z_byteCode{
				c= cipherCharacter(A_byteCode,Z_byteCode,p,k)
				cipher =strings.Join([]string{cipher,string(c)},"")
			}
		}
	}
	return cipher
}

func cipherCharacter(firstCharacterRange , lastCharacterRange ,plainCharacter rune,k int32) rune{
	var (
		entero rune
		resto rune
		c rune
	)
	entero = k / (lastCharacterRange-firstCharacterRange) 
	resto = k % (lastCharacterRange-firstCharacterRange)
	c = plainCharacter - entero + resto
	if c > lastCharacterRange {
		dif := c - lastCharacterRange
		c = firstCharacterRange  + dif -1
	}
		
	return c
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
