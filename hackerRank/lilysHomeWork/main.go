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
 * Complete the 'lilysHomework' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
 
 func countSwapsAsc(arr []int32) int32 {
     var (
         index int32
         swap int32 =0
         aux int32
         min int32
        )
     arr_cpy:= make([]int32,0)
     arr_cpy = append(arr_cpy, arr...)
     for i := 0 ; i < len(arr_cpy) ; i++{
        index=int32(i)
        min= arr_cpy[i]
        for j := i+1 ; j < len(arr_cpy);j++ {
            if arr_cpy[j] < arr_cpy[i] && arr_cpy[j] < min {
                index = int32(j)
                min = arr_cpy[j]
            }
        }
        if index != int32(i) {
            
            swap++
            aux = arr_cpy[i]
            arr_cpy[i]=arr_cpy[index]
            arr_cpy[index]=aux
        }
     }
     return swap
 }
 
 func countSwapsDesc(arr []int32) int32 {
     var (
         index int32
         swap int32 =0
         aux int32
         max int32
        )
     arr_cpy:= make([]int32,0)
     arr_cpy = append(arr_cpy, arr...)
     for i := 0 ; i < len(arr_cpy) ; i++{
        index=int32(i)
        max = arr_cpy[i]
        for j := i+1 ; j < len(arr_cpy);j++ {
            if arr_cpy[j] > arr_cpy[i] && arr_cpy[j] > max {
                index = int32(j)
                max = arr_cpy[j]
            }
        }
        if index != int32(i) {
            swap++
            aux = arr_cpy[i]
            arr_cpy[i]=arr_cpy[index]
            arr_cpy[index]=aux
        }
     }
     return swap
 }

func lilysHomework(arr []int32) int32 {
    // Write your code here
    

    asc := countSwapsAsc(arr)
    desc := countSwapsDesc(arr)
    if asc < desc {
        return asc
    }else{
        return desc
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := lilysHomework(arr)

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
