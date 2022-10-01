package main

import (
	"log"
	"strings"
	"testing"

	"github.com/tutividela/gophexcercises/HtmlParser/helpers"
	"golang.org/x/net/html"
)


func TestEx1(t *testing.T) {
	tag := "a"
	ex1 :=`<a href="/dog"><span>Something in a span</span>Text not in a span<b>Bold text!</b></a>`
	docEx1, err := html.Parse(strings.NewReader(ex1))
	if err != nil {
		log.Fatal(err)
	}
	ex1Link := "/dog"
	ex1Text := "Something in a span Text not in a span Bold text!"
	links := helpers.SearchHtmlLinks(docEx1,tag)
	if len(links) > 1 {
		t.Errorf("Expected lenght 1 , got length %d",len(links))
	}
	if ex1Link != links[0].Href {
		t.Errorf("Expected href %s , got href %s",ex1Link,links[0].Href)
	}
	if ex1Text != links[0].Text {
		t.Errorf("Expected href text %s , got href text %s",ex1Link,links[0].Text)
	}
}