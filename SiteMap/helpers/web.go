package helpers

import (
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetDocFromUrl(url string) (*html.Node,error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil , err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil , err
	}
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}
	return doc,nil

}