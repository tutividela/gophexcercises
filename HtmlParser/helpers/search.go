package helpers

import (
	"strings"

	"github.com/tutividela/gophexcercises/HtmlParser/models"

	"golang.org/x/net/html"
)

func SearchHtmlLinks(doc *html.Node, tag string)  []models.Link  {
	var links []models.Link
	
	var f func(*html.Node)
	f = func(n *html.Node) {
		var hrefText string
		var hrefLink string
		if n.Type == html.ElementNode && n.Data == tag {
			for _, a := range n.Attr {
				if a.Key == "href" {
					hrefLink = a.Val
					break
				}
			}
			GetTextContenFromNode(n,&hrefText)
			links = append(links, MapNodeToLink(hrefLink,hrefText))
		}
		
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}

func MapNodeToLink(href,text string) models.Link {
	return models.Link{
		Href: href,
		Text:text,
	}
}

func GetTextContenFromNode(n *html.Node,textNode *string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			(*textNode) = strings.Join([]string{*textNode,c.Data}," ")
		}
		GetTextContenFromNode(c,textNode)
	}
}