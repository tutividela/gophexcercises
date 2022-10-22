package helpers

import (
	"log"
	"net/url"
)

func GetUrlObjectFromString(u string) (*url.URL,error) {
	resource,err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return resource,nil
}