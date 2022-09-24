package functions

import (
	"HandlerMod/models"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		var uri string = r.RequestURI
		endpointsArray := strings.Split(uri, "/")
		lastEndpoint := endpointsArray[len(endpointsArray)-1]
		v,ok := pathsToUrls["/"+lastEndpoint]
		if ok {
			http.Redirect(w,r,v,308)
		}else{
			fallback.ServeHTTP(w,r)
		}
	}
	return h
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var yamlValues []models.T
	

	err :=yaml.Unmarshal(yml,&yamlValues)	
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	//fmt.Printf("--- urlMap:\n%v\n\n",yamlValues)
	pathsToUrls:=MapYAML(yamlValues)
	return MapHandler(pathsToUrls,fallback), nil
}

func MapYAML(values []models.T) map[string]string {
	m := make(map[string]string)
	for _,v := range values {
		m[v.Path] = v.Url
	}
	return m
}