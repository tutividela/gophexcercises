package helpers

import (
	"fmt"

	"golang.org/x/net/html"
)

func SearchHtmlLinks(doc *html.Node, tag string) /* []models.Link */ {
	//var links []models.Link
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("<a> href: ",a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fmt.Printf("Data: %s\n",c.Data)
			f(c)
		}
	}
	f(doc)
}