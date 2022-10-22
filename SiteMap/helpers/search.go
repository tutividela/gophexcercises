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
		StoreNonCyclicalUrls(v,m,domain)
	}
}

func CheckIfLinkIsRelativePath(link models.Link) bool {
	u,err := url.Parse(link.Href)
	if err != nil {
		log.Fatal("CheckIfLinkIsRelativePath: ",err)
		return false
	}
	return u.Host == "" && u.Path != ""
}

func CheckIfLinkIsAbsoultePath(link models.Link,domain *url.URL) bool {
	u,err := url.Parse(link.Href)
	if err != nil {
		log.Fatal("CheckIfLinkIsRelativePath: ",err)
		return false
	}
	return u.Host == domain.Host
}

func StoreNonCyclicalUrls(link models.Link , mapLinks map[string]string,domain *url.URL) {
	if _,ok := mapLinks[link.Href]; !ok {
		if CheckIfLinkIsRelativePath(link) {
			mapLinks[link.Href] = strings.Join([]string{strings.TrimRight(domain.String(),"/"),link.Href},"")
		}
		if CheckIfLinkIsAbsoultePath(link,domain){
			u,err := url.Parse(link.Href)
			if err != nil {
				log.Fatal("StoreNonCyclicalUrls: ",err)
				return
			}
			mapLinks[u.Path] = link.Href
		}
	}else{
		log.Println(link.Href," Already exists in mapLinks")
		return
	}

}

func CheckIfUrlExistsInHistoricMap(relativePath string,historicMap map[string]string) bool {
	_,ok := historicMap[relativePath]
	return ok
}

func CopyMapToMap(m1 ,m2 map[string]string) {
	for i,v := range m1 {
		if _,ok := m2[i];!ok {
			m2[i]=v
		}
	}
}