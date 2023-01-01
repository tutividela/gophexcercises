package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'alternate' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func alternate(s string) int32 {
    // Write your code here
    characters := getCharactersSet(s) //O(n)
    
    var (
        max int32
        str string
        filter []string
    )
    
    for i:= int32(0);i<= int32(len(characters)) -1;i++ {
        for j:= int32(i+1); j<= int32(len(characters)) -1;j++ { //O(n^2)
            filter = getRestOfCharacters(i,j,characters) //O(1)
            str = s
            for _,f := range filter {
                str = strings.Join(strings.Split(str, f), "")
            }
            ok := checkAlternate(str) //O(n)
            if ok && (len(str) > int(max)) {
                max = int32(len(str))
            }
        }
    }
    return max
}

func getRestOfCharacters(i,j int32,characters []string) []string {
    var filter []string
    if i == 0 && j == 1{
        filter = append(filter, characters[j+1:]...)
        return filter
    }
    if i == int32(len(characters))-2 && j == int32(len(characters)-1) {
        filter = append(filter, characters[0:i]...)
        return filter
    }
    if j - i > 1 {
        filter = append(filter, characters[0:i]...)
        filter = append(filter, characters[i+1:j]...)
        if j < int32(len(characters)) -1 {
            filter = append(filter, characters[j+1:]...)
        }
        return filter
    }else{
        filter = append(filter, characters[0:i]...)
        filter = append(filter, characters[j+1:]...)
    }
    return filter
}

func getCharactersSet(s string) []string {
    m := make(map[string]bool)
    var (
        set []string
    )
    for _,v := range s {
        _,ok := m[string(v)]
        if !ok {
            m[string(v)] = true
        }
    }
    for i:= range m {
        set = append(set, i)
    }
    return set
}

func checkAlternate(s string) bool {
    
    first := int32(s[0])
    second := int32(s[1])
    dif := first - second
    pattern := math.Abs(float64(dif))
    for i := 1 ; i< len(s) -1 ; i++ {
        first= int32(s[i])
        second= int32(s[i+1])
        dif = first - second
        p := math.Abs(float64(dif))
        if p != pattern {
            return false
        }
    }
    return true
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    l := int32(lTemp)
    fmt.Println(l)

    s := readLine(reader)

    result := alternate(s)

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
