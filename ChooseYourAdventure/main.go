package main

import (
	"CYAMod/helpers"
	"net/http"
)

func main() {
    var handler helpers.MyHandler
    http.ListenAndServe(":8080", &handler)
}