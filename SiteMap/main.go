package main

import (
	"encoding/xml"
	"log"
	"os"

	htmlparser "github.com/tutividela/gophexcercises/HtmlParser/helpers"
	"github.com/tutividela/gophexcercises/SiteMap/helpers"
	sitemap "github.com/tutividela/gophexcercises/SiteMap/helpers"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Expected 2 arguments , got %d",len(os.Args))
		return
	}
	domain := os.Args[1]

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
				sitemap.CopyUniqueValuesToMap(layerI1,historicMap)
				sitemap.RemoveRepeatedValues(layerI1,historicMap)
			}
		}
		return layerI1
	}

	nextLayer := ProcessLayer(resultMap)
	for len(nextLayer) > 0 {
		nextLayer = ProcessLayer(nextLayer) 
	}
	err := sitemap.CreateXMLFile("urlset.xml")
	if err != nil {
		log.Fatal(err)
	}
	urlSet := helpers.MapToUrlSet(historicMap)
	bs,err := xml.Marshal(urlSet)
	if err != nil {
		log.Fatal(err)
	}
	err = sitemap.WriteToXMLFile("urlset.xml",bs)
	if err != nil {
		log.Fatal(err)
	}
}

