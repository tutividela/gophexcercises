package main

import (
	"fmt"
	"log"

	htmlparser "github.com/tutividela/gophexcercises/HtmlParser/helpers"
	sitemap "github.com/tutividela/gophexcercises/SiteMap/helpers"
)

func main() {
	domain := "https://www.calhoun.io/"

	resultMap := make(map[string]string)
	historicMap := make(map[string]string)

	resultMap["/"]=domain

	var ProcessLayer = func(resultMap map[string]string) map[string]string {
		
		layerI1 := make(map[string]string)
		for i,v := range resultMap {
			if !sitemap.CheckIfUrlExistsInHistoricMap(i,historicMap) {
				dom,err := sitemap.GetUrlObjectFromString(v)
				if err != nil {
					log.Fatal("GetUrlObjectFromString: ",err)
					break
				}
				doc,err := sitemap.GetDocFromUrl(v)
				if err != nil {
					log.Fatal("GetDocFromUrl: ",err)
				}
				links := htmlparser.SearchHtmlLinks(doc,"a")
				sitemap.GetBelongingLinksToMap(dom,links,layerI1)
				sitemap.CopyMapToMap(layerI1,historicMap)
				CleanMap(layerI1,historicMap)
			}
		}
		return layerI1
	}

	nextLayer := ProcessLayer(resultMap)
	for len(nextLayer) > 0 {
		nextLayer = ProcessLayer(nextLayer) 
	}
	for i,v := range historicMap {
		fmt.Printf("Path: %s , Url : %s \n",i,v)
	}
	
}

func CleanMap(layerI1 ,historicMap map[string]string) {
	for i := range layerI1 {
		if _,ok := historicMap[i]; ok {
			delete(layerI1,i)
		}
	}
}