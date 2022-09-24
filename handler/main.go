package main

import (
	"HandlerMod/functions"
	"fmt"
	"net/http"
)

func main() {

	mux := defaultMux()
	

	/* pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	} */
	mapHandler := functions.MapHandler(functions.GetPathUrlFromDB(), mux)

	/*  yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := functions.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	} */
	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
 }

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}