package main

import (
	"log"

	htmlparser "github.com/tutividela/gophexcercises/HtmlParser/helpers"
	sitemap "github.com/tutividela/gophexcercises/SiteMap/helpers"
)

func main() {
	domain := "https://www.calhoun.io/"

	resultMap := make(map[string]string)
	layerI1 := make(map[string]string)
	historicMap := make(map[string]string)

	dom,err := sitemap.GetUrlObjectFromString(domain)
		if err != nil {
			log.Fatal("GetUrlObjectFromString: ",err)
		} 
		doc,err := sitemap.GetDocFromUrl(domain)
		if err != nil {
			log.Fatal("GetDocFromUrl: ",err)
		}

		links := htmlparser.SearchHtmlLinks(doc,"a")
		//Implementar un recorrido de arbol como BFS
		sitemap.GetBelongingLinksToMap(dom,links,resultMap)
		sitemap.CopyMapToMap(resultMap,historicMap)
		for i,v := range resultMap {
			if !sitemap.CheckIfUrlExistsInHistoricMap(i,historicMap){
				layerI1[i]=v
			}
		}

	var f = func (resultMap map[string]string)  {
		layerI1 = make(map[string]string)
		dom,err := sitemap.GetUrlObjectFromString(domain)
		if err != nil {
			log.Fatal("GetUrlObjectFromString: ",err)
		} 
		doc,err := sitemap.GetDocFromUrl(domain)
		if err != nil {
			log.Fatal("GetDocFromUrl: ",err)
		}

		links := htmlparser.SearchHtmlLinks(doc,"a")
		//Implementar un recorrido de arbol como BFS
		sitemap.GetBelongingLinksToMap(dom,links,resultMap)
		sitemap.CopyMapToMap(resultMap,historicMap)
		
		for i,v := range resultMap {
			if !sitemap.CheckIfUrlExistsInHistoricMap(i,historicMap){
				
				layerI1[i] =v
			}
		}
	}
	f(layerI1)
}