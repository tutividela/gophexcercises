package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"quizGame/models"
	"strconv"
	"strings"
	"time"
)

func MapDataToItems(data [][]string) []models.Item {
	var items []models.Item
	var anItem models.Item
	for _, v := range data {
		anItem.Question = strings.Trim(v[0], "\r\n ")
		anItem.Solution = strings.Trim(v[1], "\r\n ")
		items = append(items, anItem)
	}
	return items
}

func StartQuiz(items []models.Item, intTimer int) (int, int) {
	var (
		right int
		total int = len(items)
		start     = make(chan bool, 1)
	)

	//fmt.Println("Press any key to start timer...")
	//bufio.NewReader(os.Stdin)

	t1 := time.NewTimer(time.Duration(intTimer) * time.Second)
	go func() {

		for i, v := range items {
			fmt.Printf("Question #%d: %s\n", i+1, v.Question)
			fmt.Println("Enter your answer: ")

			reader := bufio.NewReader(os.Stdin)
			ans, err := reader.ReadString('\n')
			ans = strings.Trim(ans, "\r\n ")

			if err != nil {
				log.Fatal(err)
			}
			if ans == v.Solution {
				right++
			}
		}
		t1.Stop()
		start <- true
	}()
	select {
	case <-t1.C:
		fmt.Println("Timer expired")
	case <-start:
		fmt.Println("Timer stopped")
	}
	return right, total
}

func GetTimer() int {
	fmt.Println("Set Timer to (press 0 if use default value): ")
	reader := bufio.NewReader(os.Stdin)
	stringTimer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	intTimer, err := strconv.Atoi(strings.Trim(stringTimer, "\r\n "))
	if err != nil {
		log.Fatal(err)
	}
	if intTimer == 0 {
		intTimer = 30
	}
	return intTimer
}