package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"quizGame/helpers"
	"quizGame/models"
)




func main() {
	var items []models.Item

	file,err := os.Open("problems.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	csvReader := csv.NewReader(file)
	data,err := csvReader.ReadAll()
	if err != nil{
		log.Fatal(err)
	}

	items = helpers.MapDataToItems(data)
	
	intTimer := helpers.GetTimer()
	right,total := helpers.StartQuiz(items,intTimer)
	fmt.Printf("Total Rights : %d , Total Quiz: %d",right,total)

}
