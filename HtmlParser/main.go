package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tutividela/gophexcercises/HtmlParser/helpers"

	"golang.org/x/net/html"
)

func main() {
	file ,err := os.Open("pages/ex5.html")
	if err != nil {
		log.Fatal("Open: ",err)
	}
	defer file.Close()
	byteValue ,err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ReadAll: ",err)
	}
	doc, err := html.Parse(strings.NewReader(string(byteValue)))
	if err != nil {
		log.Fatal(err)
	}
	links:=helpers.SearchHtmlLinks(doc,"a")
	for _,v := range links {
		fmt.Printf("Link: %s , Text: %s \n",v.Href,v.Text)
	}
}