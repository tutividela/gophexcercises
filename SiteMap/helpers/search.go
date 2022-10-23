package helpers

import (
	"log"
	"net/url"
	"strings"

	"github.com/tutividela/gophexcercises/HtmlParser/models"
)

//Parseo a un  objeto URL y obtengo el host.
//Si host es vacio entonces se agrega.
//Si no es vacio y es diferente a 'url' entonces no lo agrego.
func GetBelongingLinksToMap(domain *url.URL , links []models.Link , m map[string]string) {
	for _,v := range links {
		storeNonCyclicalUrls(v,m,domain)
	}
}

func checkIfLinkIsRelativePath(link models.Link) bool {
	u,err := url.Parse(link.Href)
	if err != nil {
		log.Fatal("CheckIfLinkIsRelativePath: ",err)
		return false
	}
	return u.Host == "" && u.Path != ""
}

func checkIfLinkIsAbsoultePath(link models.Link,domain *url.URL) bool {
	u,err := url.Parse(link.Href)
	if err != nil {
		log.Fatal("CheckIfLinkIsRelativePath: ",err)
		return false
	}
	return u.Host == domain.Host
}

func storeNonCyclicalUrls(link models.Link , mapLinks map[string]string,domain *url.URL) {
	if _,ok := mapLinks[link.Href]; !ok {
		if checkIfLinkIsRelativePath(link) {
			mapLinks[link.Href] = strings.Join([]string{strings.TrimRight(domain.String(),"/"),link.Href},"")
		}
		if checkIfLinkIsAbsoultePath(link,domain){
			u,err := url.Parse(link.Href)
			if err != nil {
				log.Fatal("StoreNonCyclicalUrls: ",err)
				return
			}
			mapLinks[u.Path] = link.Href
		}
	}else{
		return
	}

}

func CheckIfUrlExistsInHistoricMap(relativePath string,historicMap map[string]string) bool {
	_,ok := historicMap[relativePath]
	return ok
}

func CopyUniqueValuesToMap(m1 ,m2 map[string]string) {
	for i,v := range m1 {
		if _,ok := m2[i];!ok {
			m2[i]=v
		}
	}
}
func RemoveRepeatedValues(layerI1 ,historicMap map[string]string) {
	for i := range layerI1 {
		if _,ok := historicMap[i]; ok {
			delete(layerI1,i)
		}
	}
}