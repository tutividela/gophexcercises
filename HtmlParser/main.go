package main

import (
	"HtmlParserMod/helpers"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)



func main() {
	file ,err := os.Open("pages/ex4.html")
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
	helpers.SearchHtmlLinks(doc,"a")
}