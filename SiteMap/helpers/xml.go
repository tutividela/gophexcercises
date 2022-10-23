package helpers

import (
	"log"
	"os"

	"github.com/tutividela/gophexcercises/SiteMap/models"
)

func CreateXMLFile(name string) error{
	xmlFilePointer, err := os.Create(name)
	defer xmlFilePointer.Close()
	if err != nil {
		log.Fatalf("os.Create: %v \n",err)
		return err
	}
	return nil
}

func WriteToXMLFile(name string,bytes []byte) error {
	/* xmlFilePointer,err := os.Open(name)
	defer xmlFilePointer.Close()
	if err != nil {
		log.Fatalf("os.Open: %v \n",err)
		return err
	} */
	err := os.WriteFile(name,bytes,0644)
	if err != nil {
		log.Fatalf("*os.File.Write(): %v \n",err)
		return err
	}
	return nil
}

func MapToUrlSet(m map[string]string) models.UrlSet {
	var (
		locs []models.Url
		u models.Url
	)
	for _,v:= range m {
		u = models.Url{
			U: v,
		}
		locs = append(locs, u)
	}
	return models.UrlSet{
		Locations: locs,
	}
}
